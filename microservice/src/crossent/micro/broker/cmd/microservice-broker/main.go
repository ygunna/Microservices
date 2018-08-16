package main

import (
	"net/http"
	"crossent/micro/broker"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-cf/brokerapi/auth"
	"code.cloudfoundry.org/lager"
	"fmt"
	"flag"
	"crossent/micro/broker/config"
	"crossent/micro/broker/microservicebroker"
	"github.com/gorilla/mux"
)

// reference: https://github.com/18F/concourse-broker

var configPath string
var appPath string

func init() {
	flag.StringVar(&configPath, "configPath", "", "Config file location")
	flag.StringVar(&appPath, "appPath", "", "App file location")
}


func main() {
	flag.Parse()

	logger := lager.NewLogger("microservice-broker")

	env, err := config.ParseConfig(configPath)
	env.AppPath = appPath

	if err != nil {
		logger.Fatal("Loading config file", err, lager.Data{
			"broker-config-path": configPath,
		})
	}

	credentials := brokerapi.BrokerCredentials{
		Username: env.BrokerUsername,
		Password: env.BrokerPassword,
	}

	serviceBroker := broker.New(microservicebroker.RandomCredentialsGenerator{}, logger, env)  // rabbitmq broker main.do reference.
	brokerAPI := newBroker(serviceBroker, logger, credentials, env)

	http.Handle("/", brokerAPI)
	http.ListenAndServe(fmt.Sprintf(":%s", env.Port), nil)
}

// reference: brokerapi.New()
func newBroker(serviceBroker brokerapi.ServiceBroker, logger lager.Logger, brokerCredentials brokerapi.BrokerCredentials, env config.Config) http.Handler {
	router := mux.NewRouter()
	brokerapi.AttachRoutes(router, serviceBroker, logger)

	subConfig := router.PathPrefix("/preapp").Subrouter()
	configClient := microservicebroker.NewClient(env, logger, brokerCredentials)
	configClient.DoPre(subConfig)

	return auth.NewWrapper(brokerCredentials.Username, brokerCredentials.Password).Wrap(router)
}

