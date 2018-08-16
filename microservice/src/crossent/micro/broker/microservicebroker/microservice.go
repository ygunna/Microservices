package microservicebroker

import (
	//"errors"
	"fmt"

	"code.cloudfoundry.org/lager"
	"crossent/micro/broker/config"
	"github.com/pivotal-cf/brokerapi"
	"github.com/concourse/cf-resource/out"
	"path/filepath"
	"os"
	"encoding/json"
	"github.com/concourse/cf-resource"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"io/ioutil"
)

const BASIC_USER = "basic-user"
const BASIC_SECRET = "basic-secret"

// Client defines the capabilities that any concourse client should be able to do.
type Client interface {
	CreateService(instanceID string, credentials Credentials, param Parameters) error
	DoPre(router *mux.Router)
}

type microserviceClient struct {
	config    config.Config
	logger lager.Logger
	details interface{}
}
type Parameters struct {
	TargetOrg   string `json:"target_org,omitempty"`
	TargetSpace string `json:"target_space,omitempty"`
}

// NewClient returns a client that can be used to interface with a deployed Concourse CI instance.
func NewClient(config config.Config, logger lager.Logger, details interface{}) Client {

	return &microserviceClient{
		config:    config,
		logger: logger.Session("microservice-client"),
		details: details,
	}
}

func (c *microserviceClient) CreateService(instanceID string, credentials Credentials, param Parameters) error {
	details := c.details.(brokerapi.ProvisionDetails)

	targetOrg := param.TargetOrg
	if targetOrg == "" {
		targetOrg = c.config.Organization
	}
	targetSpace := param.TargetSpace
	if targetSpace == "" {
		targetSpace = c.config.Space
	}

	source := resource.Source{API: c.config.Api, Username: c.config.Username, Password: c.config.Password, Organization: targetOrg, Space: targetSpace, SkipCertCheck: c.config.SkipCertCheck}

	env := map[string]string{}
	env[BASIC_USER] = credentials.Username
	env[BASIC_SECRET] = credentials.Password

 	params := out.Params{ManifestPath: "", Path: "", CurrentAppName: "", EnvironmentVariables: map[string]string{}}

	var request out.Request

	request = out.Request{Source: source, Params: params}

	cloudFoundry := NewCloudFoundry(instanceID, env)
	command := out.NewCommand(cloudFoundry)



	for _, l := range c.config.Service {

		if l.ServiceID == details.ServiceID {
			go pushApp(command, request, c.config, l.DisplayName)
		}
	}


	return nil

}

func pushApp(command *out.Command, request out.Request, config config.Config, displayName string) {
	request.Params.ManifestPath = filepath.Join(displayName, "manifest.yml")
	request.Params.CurrentAppName = displayName

	// make it an absolute path
	request.Params.ManifestPath = filepath.Join(config.AppPath, request.Params.ManifestPath)

	manifestFiles, err := filepath.Glob(request.Params.ManifestPath)
	if err != nil {
		fatal("searching for manifest files", err)
		return
	}

	if len(manifestFiles) != 1 {
		fatal("invalid manifest path", fmt.Errorf("found %d files instead of 1 at path: %s", len(manifestFiles), request.Params.ManifestPath))
		return
	}

	request.Params.ManifestPath = manifestFiles[0]

	if request.Params.Path != "" {
		request.Params.Path = filepath.Join(config.AppPath, request.Params.Path)
		pathFiles, err := filepath.Glob(request.Params.Path)
		if err != nil {
			fatal("searching for path", err)
			return
		}

		if len(pathFiles) != 1 {
			fatal("invalid path", fmt.Errorf("found %d files instead of 1 at path: %s", len(pathFiles), request.Params.Path))
			return
		}

		request.Params.Path = pathFiles[0]
	}

	response, err := command.Run(request)
	if err != nil {
		fatal("running command", err)
		return
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		fatal("writing response to stdout", err)
		return
	}

}

func (c *microserviceClient) preappServiceHandler(w http.ResponseWriter, r *http.Request) {

	type services struct {
		ServiceID       []string   `json:"service_id"`
	}

	var details services
	err := json.NewDecoder(r.Body).Decode(&details)
	if err != nil {
		fmt.Println("Decode err >>>", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	targetOrg := c.config.Organization
	targetSpace := c.config.Space

	source := resource.Source{API: c.config.Api, Username: c.config.Username, Password: c.config.Password, Organization: targetOrg, Space: targetSpace, SkipCertCheck: c.config.SkipCertCheck}

	env := map[string]string{}

	params := out.Params{ManifestPath: "", Path: "", CurrentAppName: "", EnvironmentVariables: map[string]string{}}

	var request out.Request

	request = out.Request{Source: source, Params: params}

	cloudFoundry := NewCloudFoundry("", env)
	command := out.NewCommand(cloudFoundry)


	for _, l := range c.config.Service {

		for _, service := range details.ServiceID {
			if l.ServiceID == service {
				request.Params.ManifestPath = filepath.Join(l.DisplayName, "manifest.yml")
				if l.ServiceName == "micro-gateway-server" {
					request.Params.CurrentAppName = l.DisplayName + "-micro"
				} else {
					request.Params.CurrentAppName = l.DisplayName
				}

				// make it an absolute path
				request.Params.ManifestPath = filepath.Join(c.config.AppPath, request.Params.ManifestPath)

				manifestFiles, err := filepath.Glob(request.Params.ManifestPath)
				if err != nil {
					fatal("searching for manifest files", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				if len(manifestFiles) != 1 {
					fatal("invalid manifest path", fmt.Errorf("found %d files instead of 1 at path: %s", len(manifestFiles), request.Params.ManifestPath))
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				request.Params.ManifestPath = manifestFiles[0]

				if request.Params.Path != "" {
					request.Params.Path = filepath.Join(c.config.AppPath, request.Params.Path)
					pathFiles, err := filepath.Glob(request.Params.Path)
					if err != nil {
						fatal("searching for path", err)
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					if len(pathFiles) != 1 {
						fatal("invalid path", fmt.Errorf("found %d files instead of 1 at path: %s", len(pathFiles), request.Params.Path))
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					request.Params.Path = pathFiles[0]
				}

				response, err := command.Run(request)
				if err != nil {
					fatal("running command", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
					fatal("writing response to stdout", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
	}

	return

}

func (c *microserviceClient) DoPre(router *mux.Router) {
	router.HandleFunc("/v1/create_service_app", c.preappServiceHandler).Methods("PUT")
}

//func (c *microserviceClient) DoConfig(router *mux.Router) {
//	router.HandleFunc("/v1/", c.configReadHandler).Methods("GET")
//	router.HandleFunc("/v1/", c.configWriteHandler).Methods("POST")
//}

// reference : https://github.com/cloudfoundry/cf-networking-release/blob/develop/src/example-apps/proxy/handlers/proxy_handler.go
// reference : https://github.com/cloudfoundry-community/vault-broker/blob/master/main.go
func (c *microserviceClient) configReadHandler(w http.ResponseWriter, req *http.Request) {
	destination := strings.TrimPrefix(req.URL.Path, "/config/")
	url := fmt.Sprintf("http://%s:%s/%s", "0.0.0.0", "8200", destination)

	token := req.Header.Get("X-Config-Token")
	if token == "" {
		fmt.Fprintf(os.Stderr, "No VAULT_TOKEN\n")
		os.Exit(1)
	}
	fmt.Println("VAULT URL:" + url)
	fmt.Println("VAULT Token:" + token)

	res, err := c.Do("GET", url, nil, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("read body failed: %s", err)))
		return
	}
	if res.StatusCode == 404 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if res.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Received %s from Vault", res.Status)))
		return
	}

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("read body failed: %s", err)))
		return
	}



	w.Write(b)
}

func (c *microserviceClient) configWriteHandler(w http.ResponseWriter, req *http.Request) {

}

func (c *microserviceClient) Do(method, url string, data interface{}, token string) (*http.Response, error) {
	req, err := c.newRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Vault-Token", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *microserviceClient) newRequest(method, url string, data interface{}) (*http.Request, error) {
	if data == nil {
		return http.NewRequest(method, url, nil)
	}
	cooked, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(method, url, strings.NewReader(string(cooked)))
}

func fatal(message string, err error) {
	fmt.Fprintf(os.Stderr, "error %s: %s\n", message, err)
	//os.Exit(1)
}

