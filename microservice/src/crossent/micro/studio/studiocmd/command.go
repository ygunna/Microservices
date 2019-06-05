package microcmd

import (
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/ifrit/sigmon"
	"code.cloudfoundry.org/cfhttp"
	"crossent/micro/studio"
	"crossent/micro/studio/api"
	"fmt"
	"crossent/micro/studio/publichandler"
	"database/sql"
	"crossent/micro/studio/db"
	"crossent/micro/studio/db/lock"
	"database/sql/driver"
	_ "github.com/lib/pq"
	"crossent/micro/studio/client"
	"crossent/micro/studio/wrapper"
	"crossent/micro/studio/handlers"
)

type MicroCommand struct {
	HealthAddress string `long:"health-address"   default:"0.0.0.0" description:"IP address on which to health check."`
	HealthPort uint16 `long:"health-port" default:"8081"    description:"Port on which to listen for health check."`

	LogStreamAddress string `long:"logstream-address"   default:"0.0.0.0" description:"IP address on which to listen for log stream."`
	LogStreamPort uint16 `long:"logstream-port" default:"8082"    description:"Port on which to listen for log stream."`

	Logger LagerFlag
	//
	BindIP   IPFlag `long:"bind-ip"   default:"0.0.0.0" description:"IP address on which to listen for web traffic."`
	BindPort uint16 `long:"bind-port" default:"8080"    description:"Port on which to listen for HTTP traffic."`
	//
	TLSBindPort uint16   `long:"tls-bind-port" description:"Port on which to listen for HTTPS traffic."`
	ServerCertFile FileFlag `long:"server-tls-cert"      description:"File containing an Server SSL certificate."`
	CACertFile     FileFlag `long:"tls-cert"      description:"File containing an SSL certificate."`
	ServerKeyFile      FileFlag `long:"tls-key"       description:"File containing an RSA private key, used to encrypt HTTPS traffic."`

	Postgres PostgresConfig `group:"PostgreSQL Configuration" namespace:"postgres"`

	ApiUrl string `long:"api-url"   default:"https://api.bosh-lite.com" description:"CF api url"`
	UaaUrl string `long:"uaa-url"   default:"https://uaa.bosh-lite.com" description:"CF uaa url"`
	CfUsername string `long:"username"   description:"CF username"`
	CfPassword string `long:"password"   description:"CF password"`
	CfSkipCertCheck bool `long:"skip-cert-check"   description:"CF api skip_cert_check"`

	ClientID string `long:"client_id"   default:"micro" description:"Client ID"`
	ClientSecret string `long:"client_secret"   default:"micro-secret" description:"Client Secret"`

	TraefikApiURL URLFlag `long:"traefik-api-url" default:"http://127.0.0.1:8080" description:"Traefik URL."`
	TraefikPort uint16 `long:"traefik-port" default:"8089"    description:"Port on which to listen for Traefik HTTP."`
	TraefikUser string `long:"traefik-user"   default:"sudouser" description:"Traefik sudo user"`
	TraefikPassword string `long:"traefik-password"   default:"sudouser" description:"Traefik sudo password"`

	ExternalURL URLFlag `long:"external-url" default:"http://127.0.0.1:8080" description:"Studio URL."`

	GrafanaURL URLFlag `long:"grafana-url" default:"http://127.0.0.1:3003" description:"Grafana URL."`
	GrafanaPort uint16 `long:"grafana-port" default:"3003"    description:"Port on which to listen for Grafana HTTP."`
	GrafanaAdminPassword string `long:"grafana-admin-password"   default:"adminpassword" description:"Grafana admin password"`
}

type connectionRetryingDriver struct {
	driver.Driver
}



func (cmd *MicroCommand) Execute(args []string) error {

	SetupConnectionRetryingDriver("postgres", cmd.Postgres.ConnectionString(), "too-many-connections-retrying")

	logger, _ := cmd.constructLogger()

	uaa := client.NewClient(cmd.ApiUrl, cmd.UaaUrl, cmd.CfUsername, cmd.CfPassword, cmd.ClientID, cmd.ClientSecret,
				cmd.TraefikApiURL.String(), cmd.TraefikPort, cmd.TraefikUser, cmd.TraefikPassword,
				cmd.ExternalURL.String(), cmd.GrafanaURL.String(), cmd.GrafanaPort, cmd.GrafanaAdminPassword)

	httpHandler, err := cmd.constructHTTPHandler(logger, uaa)
	if err != nil {
		return err
	}

	members := []grouper.Member{}

	var microServer ifrit.Runner
	if cmd.ServerCertFile != "" || cmd.ServerKeyFile != "" || cmd.CACertFile != "" {
		tlsConfig, err := cfhttp.NewTLSConfig(string(cmd.ServerCertFile), string(cmd.ServerKeyFile), string(cmd.CACertFile))
		if err != nil {
			logger.Fatal("invalid-tls-config", err)
		}
		microServer = http_server.NewTLSServer(fmt.Sprintf("%s:%d", cmd.BindIP, cmd.TLSBindPort), httpHandler, tlsConfig)
	} else {
		microServer = http_server.New(fmt.Sprintf("%s:%d", cmd.BindIP, cmd.BindPort), httpHandler)
	}


	//healthcheckServer := http_server.New(fmt.Sprintf("%s:%d", cmd.HealthAddress, cmd.HealthPort), http.HandlerFunc(healthCheckHandler))
	logStreamServer := http_server.New(fmt.Sprintf("%s:%d", cmd.LogStreamAddress, cmd.LogStreamPort), cmd.logStreamHandler(logger, uaa))

	members = append(members, grouper.Members{
		{"auction-runner", microServer},
		//{"healthcheck", healthcheckServer},
		{"logstream", logStreamServer},
	}...)

	group := grouper.NewOrdered(os.Interrupt, members)

	monitor := ifrit.Invoke(sigmon.New(group))

	logger.Info("started")
	//fmt.Println("api:"+cmd.ApiUrl)
	//fmt.Println("uaa:"+cmd.UaaUrl)
	//fmt.Println("username:"+cmd.CfUsername)
	//fmt.Println("pwd:"+cmd.CfPassword)
	//fmt.Printf("skip:%v", cmd.CfSkipCertCheck)

	err = <-monitor.Wait()

	if err != nil {
		logger.Error("exited-with-failure", err)
		//os.Exit(1)
		return err
	}

	logger.Info("exited")

	return err


	//runner, err := cmd.Runner(args)
	//if err != nil {
	//	return err
	//}
	//
	//return <-ifrit.Invoke(sigmon.New(runner)).Wait()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (cmd *MicroCommand) constructHTTPHandler(
	logger lager.Logger,
	uaa *client.UAA,
) (http.Handler, error) {

	lockConn, err := cmd.constructLockConn("too-many-connections-retrying")
	if err != nil {
		return nil, err
	}

	lockFactory := lock.NewLockFactory(lockConn)

	dbConn, err := cmd.constructDBConn("postgres", logger, 32, "api", lockFactory)
	if err != nil {
		return nil, err
	}

	repositoryFactory := db.NewRepositoryFactory(dbConn, lockFactory)
	//bus := dbConn.Bus()

	webHandler, err := studio.NewHandler(logger)
	if err != nil {
		return nil, err
	}

	apiWrapper := wrappa.MultiWrappa{
		wrappa.NewAPIWrappa(),
	}

	apiHandler, err := api.NewHandler(logger, repositoryFactory, uaa, apiWrapper)
	if err != nil {
		return nil, err
	}

	publicHandler, err := publichandler.NewHandler()
	if err != nil {
		return nil, err
	}

	swaggerIndexHandler, err := publichandler.NewSwaggerGatewayHandler()
	if err != nil {
		return nil, err
	}

	swaggerEntryIndexHandler, err := publichandler.NewSwaggerEntryHandler()
	if err != nil {
		return nil, err
	}

	webMux := http.NewServeMux()
	webMux.Handle("/api/v1/", apiHandler)
	webMux.Handle("/public/", publicHandler)
	webMux.Handle("/swagger/", swaggerIndexHandler)
	webMux.Handle("/swagger/entry/", swaggerEntryIndexHandler)
	webMux.Handle("/", webHandler)

	//httpHandler := wrappa.LoggerHandler{
	//	Logger: logger,
	//
	//	Handler: wrappa.SecurityHandler{
	//		XFrameOptions: cmd.Server.XFrameOptions,
	//
	//		Handler: webMux,
	//	},
	//}

	return webMux, nil
}

func (cmd *MicroCommand) constructDBConn(
	driverName string,
	logger lager.Logger,
	//newKey *db.EncryptionKey,
	//oldKey *db.EncryptionKey,
	maxConn int,
	connectionName string,
	lockFactory lock.LockFactory,
) (db.Conn, error) {
	dbConn, err := db.Open(logger.Session("db"), driverName, cmd.Postgres.ConnectionString() /*newKey, oldKey*/, connectionName, lockFactory)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err)
	}

	// Prepare
	dbConn.SetMaxOpenConns(maxConn)

	return dbConn, nil
}

func (cmd *MicroCommand) constructLockConn(driverName string) (*sql.DB, error) {
	dbConn, err := sql.Open(driverName, cmd.Postgres.ConnectionString())
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxOpenConns(1)
	dbConn.SetMaxIdleConns(1)
	dbConn.SetConnMaxLifetime(0)

	return dbConn, nil
}

func logWrap(loggable func(http.ResponseWriter, *http.Request, lager.Logger), logger lager.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestLog := logger.Session("request", lager.Data{
			"method":  r.Method,
			"request": r.URL.String(),
		})

		requestLog.Info("serving")
		loggable(w, r, requestLog)
		requestLog.Info("done")
	}
}

func (cmd *MicroCommand) constructLogger() (lager.Logger, *lager.ReconfigurableSink) {
	logger, reconfigurableSink := cmd.Logger.Logger("micro")

	//if cmd.Metrics.YellerAPIKey != "" {
	//	yellerSink := zest.NewYellerSink(cmd.Metrics.YellerAPIKey, cmd.Metrics.YellerEnvironment)
	//	logger.RegisterSink(yellerSink)
	//}

	return logger, reconfigurableSink
}

func SetupConnectionRetryingDriver(delegateDriverName, sqlDataSource, newDriverName string) {
	delegateDBConn, err := sql.Open(delegateDriverName, sqlDataSource)
	if err == nil {
		// ignoring any connection errors since we only need this to access the driver struct
		_ = delegateDBConn.Close()
	}
	//fmt.Println(err)
	//fmt.Println(delegateDriverName)
	//fmt.Println(sqlDataSource)
	//fmt.Println(delegateDBConn)
	connectionRetryingDriver := &connectionRetryingDriver{delegateDBConn.Driver()}
	sql.Register(newDriverName, connectionRetryingDriver)
}

func (cmd *MicroCommand) logStreamHandler(logger lager.Logger, uaa *client.UAA) (http.Handler) {
	l := handlers.NewLogHandler(logger, uaa)
	//uaa := client.NewClient(cmd.ApiUrl, cmd.UaaUrl, cmd.CfUsername, cmd.CfPassword, "", "")
	//return  http.Handle()erFunc(handlers.NewLogHandler(logger, uaa).Firehose)
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//
	//
	//})
	webMux := http.NewServeMux()
	webMux.HandleFunc("/stream", l.Stream)
	webMux.HandleFunc("/firehose", l.Firehose)
	return webMux
}