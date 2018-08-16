package handlers

// reference : https://github.com/cloudfoundry-incubator/stratos/blob/master/components/cloud-foundry/backend/cf_websocket_streams.go

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	//log "github.com/Sirupsen/logrus"
	//"github.com/cloudfoundry/noaa"
	"github.com/cloudfoundry/noaa/consumer"
	noaa_errors "github.com/cloudfoundry/noaa/errors"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/engine/standard"
	"code.cloudfoundry.org/lager"
	"crossent/micro/studio/client"
	"strings"
	"github.com/cloudfoundry/noaa"
)

const (
	// Time allowed to read the next pong message from the peer
	pongWait = 30 * time.Second

	// Send ping messages to peer with this period (must be less than pongWait)
	pingPeriod = (pongWait * 9) / 10

	// Time allowed to write a ping message
	pingWriteTimeout = 10 * time.Second
)

type ParamPage struct {
	Id  string
}

// Allow connections from any Origin
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type LogHandler struct {
	logger        lager.Logger
	uaa *client.UAA
}

func NewLogHandler(
	logger lager.Logger,
	uaa *client.UAA,
) *LogHandler {




	return &LogHandler{
		logger:      logger,
		uaa: uaa,
	}
}


//func (c CloudFoundrySpecification) appStream(echoContext echo.Context) error {
//	return c.commonStreamHandler(echoContext, appStreamHandler)
//}

func (s *LogHandler) Stream(w http.ResponseWriter, r *http.Request) {

	//template := template.Must(template.New("homePage").Parse(homePageTemplate))
	//id := r.FormValue("id")
	//err := template.Execute(w, ParamPage{
	//	Id:  id,
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}

	if err :=  s.commonStreamHandler(w, r, appStreamHandler); err != nil {
		s.logger.Error("failed Stream", err)
		//w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (s *LogHandler) Firehose(w http.ResponseWriter, r *http.Request) {

	//stream := r.FormValue("stream")

	if err :=  s.commonStreamHandler(w, r, firehoseStreamHandler); err != nil {
		s.logger.Error("failed Firehose", err)
		//w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (s *LogHandler) commonStreamHandler(w http.ResponseWriter, r *http.Request, bespokeStreamHandler func(http.ResponseWriter, *http.Request, *AuthorizedConsumer, *websocket.Conn, lager.Logger) error) error {
	ac, err := s.openNoaaConsumer(w, r)
	if err != nil {
		return err
	}
	defer ac.consumer.Close()

	clientWebSocket, pingTicker, err := upgradeToWebSocket(w, r, s.logger)
	if err != nil {
		return err
	}
	defer clientWebSocket.Close()
	defer pingTicker.Stop()

	if err := bespokeStreamHandler(w, r, ac, clientWebSocket, s.logger); err != nil {
		return err
	}

	// This blocks until the WebSocket is closed
	drainClientMessages(clientWebSocket)
	return nil
}

type AuthorizedConsumer struct {
	consumer     *consumer.Consumer
	authToken    string
	refreshToken func() error
}

// Refresh the Authorization token if needed and create a new Noaa consumer
func (s *LogHandler) openNoaaConsumer(w http.ResponseWriter, r *http.Request) (*AuthorizedConsumer, error) {

	ac := &AuthorizedConsumer{}

	// Get the CNSI and app IDs from route parameters
	//cnsiGUID := echoContext.Param("cnsiGuid")
	//userGUID := echoContext.Get("user_id").(string)

	//id := r.FormValue(":id")

	//// Extract the Doppler endpoint from the CNSI record
	//cnsiRecord, err := c.portalProxy.GetCNSIRecord(cnsiGUID)
	//if err != nil {
	//	return nil, fmt.Errorf("Failed to get record for CNSI %s: [%v]", cnsiGUID, err)
	//}
	//
	ac.refreshToken = func() error {
		t, err := s.uaa.GetAuthToken()
		if err != nil {
			//panic(err)
			return err
		}
		ac.authToken = "bearer " + t.RefreshToken
		return nil
	//	newTokenRecord, err := c.portalProxy.RefreshToken(cnsiRecord.SkipSSLValidation, cnsiGUID, userGUID, "", "", cnsiRecord.TokenEndpoint)
	//	if err != nil {
	//		msg := fmt.Sprintf("Error refreshing token for CNSI %s : [%v]", cnsiGUID, err)
	//		return echo.NewHTTPError(http.StatusUnauthorized, msg)
	//	}
	//	ac.authToken = "bearer " + newTokenRecord.AuthToken
	//	return nil
	}

	t, err := s.uaa.GetAuthToken()
	if err != nil {
		//panic(err)
		return ac, err
	}
	ac.authToken = "bearer " + t.AccessToken

	info, err := s.uaa.Info()
	if err != nil {
		//panic(err)
		return ac, err
	}
	dopplerAddress := info.DopplerLoggingEndpoint

	//dopplerAddress := cnsiRecord.DopplerLoggingEndpoint
	//log.Debugf("CNSI record Obtained! Using Doppler Logging Endpoint: %s", dopplerAddress)
	//
	//// Get the auth token for the CNSI from the DB, refresh it if it's expired
	//if tokenRecord, ok := c.portalProxy.GetCNSITokenRecord(cnsiGUID, userGUID); ok && !tokenRecord.Disconnected {
	//	ac.authToken = "bearer " + tokenRecord.AuthToken
	//	expTime := time.Unix(tokenRecord.TokenExpiry, 0)
	//	if expTime.Before(time.Now()) {
	//		log.Debug("Token obtained has expired, refreshing!")
	//		if err = ac.refreshToken(); err != nil {
	//			return nil, err
	//		}
	//	}
	//} else {
	//	return nil, fmt.Errorf("Error getting token for user %s on CNSI %s", userGUID, cnsiGUID)
	//}
	//
	//// Open a Noaa consumer to the doppler endpoint
	//log.Debugf("Creating Noaa consumer for Doppler endpoint %s", dopplerAddress)
	ac.consumer = consumer.New(dopplerAddress, &tls.Config{InsecureSkipVerify: true}, http.ProxyFromEnvironment)

	return ac, nil
}

// Upgrade the HTTP connection to a WebSocket with a Ping ticker
func upgradeToWebSocket(w http.ResponseWriter, r *http.Request, logger lager.Logger) (*websocket.Conn, *time.Ticker, error) {

	// Adapt echo.Context to Gorilla handler
	responseWriter := w
	request := r

	// We're now ok talking to CF, time to upgrade the request to a WebSocket connection
	logger.Debug("Upgrading request to the WebSocket protocol...")
	clientWebSocket, err := upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		logger.Error("Upgrading connection to a WebSocket failed", err)
		return nil, nil, err
	}
	logger.Debug("Successfully upgraded to a WebSocket connection")

	// HSC-1276 - handle pong messages and reset the read deadline
	clientWebSocket.SetReadDeadline(time.Now().Add(pongWait))
	clientWebSocket.SetPongHandler(func(string) error {
		clientWebSocket.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// HSC-1276 - send regular Pings to prevent the WebSocket being closed on us
	ticker := time.NewTicker(pingPeriod)
	go func() {
		for range ticker.C {
			clientWebSocket.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(pingWriteTimeout))
		}
	}()

	return clientWebSocket, ticker, nil
}

// Attempts to get the recent logs, if we get an unauthorized error we will refresh the auth token and retry once
func getRecentLogs(ac *AuthorizedConsumer, cnsiGUID, appGUID string) ([]*events.LogMessage, error) {
	//fmt.Println("getRecentLogs", ac.authToken)
	messages, err := ac.consumer.RecentLogs(appGUID, ac.authToken)
	//fmt.Println(">>>>>>>>>>>>>>", err)
	if err != nil {
		errorPattern := "Failed to get recent messages for App %s on CNSI %s [%v]"
		if _, ok := err.(*noaa_errors.UnauthorizedError); ok {
			// If unauthorized, we may need to refresh our Auth token
			// Note: annoyingly, older versions of CF also send back "401 - Unauthorized" when the app doesn't exist...
			// This means we sometimes end up here even when our token is legit
			err := ac.refreshToken();
			if  err != nil {
				return nil, fmt.Errorf(errorPattern, appGUID, cnsiGUID, err)
			}
			messages, err = ac.consumer.RecentLogs(appGUID, ac.authToken)
			if err != nil {
				msg := fmt.Sprintf(errorPattern, appGUID, cnsiGUID, err)
				return nil, echo.NewHTTPError(http.StatusUnauthorized, msg)
			}
		} else {
			return nil, fmt.Errorf(errorPattern, appGUID, cnsiGUID, err)
		}
	}
	return messages, nil
}

func drainErrors(errorChan <-chan error) {
	for err := range errorChan {
		// Note: we receive a nil error before the channel is closed so check here...
		if err != nil {
			fmt.Errorf("Received error from Doppler %v", err.Error())
		}
	}
}

func drainLogMessages(msgChan <-chan *events.LogMessage, callback func(msg *events.LogMessage)) {
	for msg := range msgChan {
		callback(msg)
	}
}

func drainFirehoseEvents(eventChan <-chan *events.Envelope, callback func(msg *events.Envelope)) {
	for event := range eventChan {
		callback(event)
	}
}

// Drain and discard incoming messages from the WebSocket client, effectively making our WebSocket read-only
func drainClientMessages(clientWebSocket *websocket.Conn) {
	for {
		_, _, err := clientWebSocket.ReadMessage()
		if err != nil {
			// We get here when the client (browser) disconnects
			break
		}
	}
}

func appStreamHandler(w http.ResponseWriter, r *http.Request, ac *AuthorizedConsumer, clientWebSocket *websocket.Conn, log lager.Logger) error {
	logger := log.Session("cf_websocket_handler")
	logger.Debug("stream")

	// Get the CNSI and app IDs from route parameters
	//cnsiGUID := echoContext.Param("cnsiGuid")
	//appGUID := echoContext.Param("appGuid")
	appGUID := r.FormValue("id")

	logger.Info("Received request for log stream for App ID:", lager.Data{"appGUID": appGUID})

	messages, err := getRecentLogs(ac, "msxpert", appGUID)
	if err != nil {
		return err
	}
	// Reusable closure to pump messages from Noaa to the client WebSocket
	// N.B. We convert protobuf messages to JSON for ease of use in the frontend
	relayLogMsg := func(msg *events.LogMessage) {
		if jsonMsg, err := json.Marshal(msg); err != nil {
			logger.Error("Received unparsable message from Doppler", err, lager.Data{"jsonMsg": jsonMsg})
		} else {
			err := clientWebSocket.WriteMessage(websocket.TextMessage, jsonMsg)
			if err != nil {
				logger.Error("Error writing data to WebSocket", err)
			}
		}
	}

	// Send the recent messages, sorted in Chronological order
	for _, msg := range noaa.SortRecent(messages) {
		relayLogMsg(msg)
	}

	msgChan, errorChan := ac.consumer.TailingLogs(appGUID, ac.authToken)

	// Process the app stream
	go drainErrors(errorChan)
	go drainLogMessages(msgChan, relayLogMsg)

	logger.Info("Now streaming log for App ID:", lager.Data{"appGUID": appGUID})
	return nil
}

func firehoseStreamHandler(w http.ResponseWriter, r *http.Request, ac *AuthorizedConsumer, clientWebSocket *websocket.Conn, log lager.Logger) error {
	logger := log.Session("cf_websocket_handler")
	logger.Debug("firehose")

	// Get the CNSI and app IDs from route parameters
	//cnsiGUID := echoContext.Param("cnsiGuid")
	guid := r.FormValue("id")

	logger.Debug("Received request for Firehose stream for ID: ", lager.Data{"appGUID":guid})

	//userGUID := echoContext.Get("user_id").(string)
	//userGUID := "test"
	firehoseSubscriptionId := "msxpert" + "@" + strconv.FormatInt(time.Now().UnixNano(), 10)
	logger.Debug("Connecting the Firehose with subscription ID: ", lager.Data{"firehoseSubscriptionId":firehoseSubscriptionId})

	eventChan, errorChan := ac.consumer.Firehose(firehoseSubscriptionId, ac.authToken)

	// Process the app stream
	go drainErrors(errorChan)
	go drainFirehoseEvents(eventChan, func(msg *events.Envelope) {
		filterMsg := msgFilter(msg, guid)
		if filterMsg != nil {
			if jsonMsg, err := json.Marshal(msg); err != nil {
				logger.Error("Received unparsable message from Doppler", err, lager.Data{"jsonMsg":jsonMsg})
			} else {
				err := clientWebSocket.WriteMessage(websocket.TextMessage, jsonMsg)
				if err != nil {
					logger.Error("Error writing data to WebSocket", err)
				}
			}
		}
	})

	logger.Debug("Firehose connected and streaming: subscription ID: ", lager.Data{"firehoseSubscriptionId":firehoseSubscriptionId})
	return nil
}

func msgFilter(msg *events.Envelope, guid string) *events.Envelope {

	m := msg.GetLogMessage()
	if m != nil && guid != "" {
		guids := strings.Split(guid, ":")
		for _, g := range guids {
			if g != "" && m.GetAppId() == g {
				return msg
			}
		}

	}
	return nil
}

