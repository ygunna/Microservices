package client

// /cloudfoundry/cf-networking-release/src/lib/json_client/json_client.go reference

import (
	"net/http"
	"fmt"
	//"os"
	"io/ioutil"
	"bytes"
	//"encoding/base64"
	"log"
	"encoding/json"
	"net/url"
	"errors"
	"crypto/tls"
	"io"

	"github.com/cloudfoundry-community/go-cfclient"
	"encoding/base64"
	"crossent/micro/studio/domain"
)


//var (
//	userName       = os.Getenv("UAA_USER_NAME") //"admin"
//	userPassword   = os.Getenv("UAA_USER_PASSWORD") //"admin"
//	clientName     = os.Getenv("UAA_CLIENT_NAME") //"portal-id"
//	clientSecret   = os.Getenv("UAA_CLIENT_SECRET") //"portal-secret"
//	uaaEndpoint    = os.Getenv("UAA_ENDPOINT") //"https://uaa.bosh-lite.com"
//	apiEndpoint    = os.Getenv("API_ENDPOINT") //"https://api.bosh-lite.com"
//)

type UAA struct {
	UaaURL       string
	ApiURL       string
	Username     string
	Password     string
	ClientID     string
	ClientSecret string
	TraefikApiUrl string
	TraefikPort  uint16
	TraefikUser  string
	TraefikPassword string
	ExternalURL  string
	GrafanaUrl   string
	GrafanaPort  uint16
	GrafanaAdminPassword string
}

type CF_TOKEN struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//func DefaultClient() *UAA {
//	return &UAA{
//		UaaURL: uaaEndpoint,
//		ApiURL: apiEndpoint,
//		Username: userName,
//		Password: userPassword,
//		ClientID: clientName,
//		ClientSecret: clientSecret,
//	}
//}

type V2Info struct {
	AuthorizationEndpoint    string `json:"authorization_endpoint"`
	TokenEndpoint            string `json:"token_endpoint"`
	DopplerLoggingEndpoint   string `json:"doppler_logging_endpoint"`
	AppSSHEndpoint           string `json:"app_ssh_endpoint"`
	AppSSHHostKeyFingerprint string `json:"app_ssh_host_key_fingerprint"`
	AppSSHOauthCLient        string `json:"app_ssh_oauth_client"`
}

func NewClient(apiEndpoint string, uaaEndpoint string, name string, password string, clientName string, clientSecret string,
	       traefikApiUrl string, traefikPort uint16, traefikUser string, traefikPassword string, externalUrl string,
	       grafanaUrl string, grafanaPort uint16, grafanaAdminPassword string) *UAA {
	return &UAA{
		UaaURL: uaaEndpoint,
		ApiURL: apiEndpoint,
		Username: name,
		Password: password,
		ClientID: clientName,
		ClientSecret: clientSecret,
		TraefikApiUrl: traefikApiUrl,
		TraefikPort: traefikPort,
		TraefikUser: traefikUser,
		TraefikPassword: traefikPassword,
		ExternalURL: externalUrl,
		GrafanaUrl: grafanaUrl,
		GrafanaPort: grafanaPort,
		GrafanaAdminPassword: grafanaAdminPassword,
	}
}

func (u *UAA) GetAuthToken() (*CF_TOKEN, error) {
	uaaURL := fmt.Sprintf("%s/oauth/token", u.UaaURL)
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", u.Username)
	data.Set("password", u.Password)
	//data.Set("client_id", u.ClientID)
	//data.Set("client_secret", u.ClientSecret)

	//data.Set("response_type", "token")
	//data.Set("scope", "")

	//fmt.Println(u.Username, u.Password, u.ClientID, u.ClientSecret)

	um := &CF_TOKEN{}

	r, err := http.NewRequest("POST", uaaURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return um, err
	}
	//basicAuthToken := base64.StdEncoding.EncodeToString([]byte(u.ClientID + ":" + u.ClientSecret))
	//r.Header.Set("Authorization", "Basic "+basicAuthToken)
	r.Header.Set("Authorization", "Basic Y2Y6")
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	config := &tls.Config{InsecureSkipVerify: true}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	//resp, err := http.DefaultClient.Do(r)
	resp, err := client.Do(r)
	if err != nil {
		return um, err
	}
	if resp.StatusCode != http.StatusOK {
		errout, _ := ioutil.ReadAll(resp.Body)
		log.Printf("StatusCode: %d\n Error: %s\n", resp.StatusCode, errout)
		if errout != nil {
			return um, errors.New(fmt.Sprintf("Response %v: %v", resp.StatusCode, string(errout)))
		} else {
			return um, errors.New(fmt.Sprintf("Response %v", resp.StatusCode))
		}
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return um, err
	}

	//um := &struct {
	//	AccessToken string `json:"access_token"`
	//	RefreshToken string `json:"refresh_token"`
	//}{}
	//um := &CF_TOKEN{}
	json.Unmarshal(content, um)

	return um, nil
}

func (u *UAA) GetResourcesFromToken(method, route string, reqData, respData interface{}, token string) error {
	var reader io.Reader
	if method != "GET" {
		bodyBytes, err := json.Marshal(reqData)
		if err != nil {
			return fmt.Errorf("json marshal request body: %s", err)
		}
		reader = bytes.NewReader(bodyBytes)
	}

	reqURL := u.ApiURL + route
	request, err := http.NewRequest(method, reqURL, reader)
	if err != nil {
		return fmt.Errorf("http new request: %s", err)
	}

	request.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	config := &tls.Config{InsecureSkipVerify: true}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	//client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		//panic(err)
		return err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//panic(err)
		return err
	}
	//str := string(bytes)
	//fmt.Println(str)

	if response.StatusCode > 299 {
		var errDescription string

		var cfErr domain.CloudFoundryErr
		err = json.Unmarshal(bytes, &cfErr)
		if err != nil {
			errDescription = string(bytes)
		} else {
			errDescription = cfErr.Description
		}

		return fmt.Errorf(`{"httpStatusCode": %v, "cfErrorCode": "%s", "Message": "%s"}`, response.StatusCode, cfErr.ErrorCode, errDescription)
	}


	//logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	//logger.Println(string(bytes))

	if respData != nil {
		err = json.Unmarshal(bytes, respData)
		if err != nil {
			return fmt.Errorf("json unmarshal: %s", err)
		}
	}

	return nil
}

func (u *UAA) GetResources(method, route string, reqData, respData interface{}) error {
	token, err := u.GetAuthToken()
	if err != nil {
		//panic(err)
		return err
	}
	return u.GetResourcesFromToken(method, route , reqData, respData , token.AccessToken)
}

func (u *UAA) CfClient() (client *cfclient.Client, err error) {
	//token, err := u.GetAuthToken()
	//if err != nil {
	//	//panic(err)
	//	return nil, err
	//}

	config := &tls.Config{InsecureSkipVerify: true}
	tr := &http.Transport{TLSClientConfig: config}
	hc := &http.Client{Transport: tr}

	c := &cfclient.Config{
		ApiAddress:   u.ApiURL,
		Username:     u.Username,
		Password:     u.Password,
		SkipSslValidation: true,
		HttpClient: hc,
		//Token:        token.AccessToken,
	}

	client, err = cfclient.NewClient(c)

	if err != nil {
		return nil, err
	}
	return client, nil
}


func (u *UAA) Info() (V2Info, error) {
	var reader io.Reader
	var v2Response V2Info

	reqURL := u.ApiURL + "/v2/info"
	//request, err := http.Get(reqURL)
	request, err := http.NewRequest("GET", reqURL, reader)

	if err != nil {
		return v2Response, fmt.Errorf("http new request: %s", err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	config := &tls.Config{InsecureSkipVerify: true}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	//client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return v2Response, err
	}
	if response.StatusCode != http.StatusOK {
		errout, _ := ioutil.ReadAll(response.Body)
		log.Printf("StatusCode: %d\n Error: %s\n", response.StatusCode, errout)
		if errout != nil {
			return v2Response, errors.New("response not 200: " + string(errout))
		} else {
			return v2Response, errors.New("response not 200")
		}
	}

	defer response.Body.Close()

	content := json.NewDecoder(response.Body)
	if err = content.Decode(&v2Response); err != nil {
		return v2Response, err
	}

	return v2Response, nil
}

func (u *UAA) GetInfo() (V2Info, error) {
	var v2Response V2Info
	err := u.GetResources("GET", fmt.Sprintf("/v2/info"), nil, &v2Response)
	if err != nil {
		return v2Response, errors.New(("Error getting /v2/info:") + err.Error())
	}
	return v2Response, nil
}

func (u *UAA) GetUaaToken(data url.Values) (*CF_TOKEN, error) {
	uaaURL := fmt.Sprintf("%s/oauth/token", u.UaaURL)

	data.Set("client_id", u.ClientID)
	data.Set("client_secret", u.ClientSecret)
	um := &CF_TOKEN{}
	r, err := http.NewRequest("POST", uaaURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return um, err
	}
	basicAuthToken := base64.StdEncoding.EncodeToString([]byte(u.ClientID + ":" + u.ClientSecret))
	r.Header.Set("Authorization", "Basic "+basicAuthToken)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Accept", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Do(r)
	if err != nil {
		return um, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errout, _ := ioutil.ReadAll(res.Body)
		log.Printf("StatusCode: %d\n Error: %s\n", res.StatusCode, errout)
		if errout != nil {
			return um, errors.New(fmt.Sprintf("Response %v: %v", res.StatusCode, string(errout)))
		} else {
			return um, errors.New(fmt.Sprintf("Response %v", res.StatusCode))
		}
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return um, err
	}
	json.Unmarshal(content, um)

	return um, nil
}

// Required UAA scopes: uaa.admin or token.revoke
func (u *UAA) RevokeUserUaaToken(userId string, token string) error {
	uaaURL := fmt.Sprintf("%s/oauth/token/revoke/user/%s", u.UaaURL, userId)

	r, err := http.NewRequest("GET", uaaURL, nil)
	if err != nil {
		return err
	}
	r.Header.Set("Authorization", "Bearer "+token)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		errout, _ := ioutil.ReadAll(res.Body)
		log.Printf("StatusCode: %d\n Error: %s\n", res.StatusCode, errout)
		if errout != nil {
			return errors.New(fmt.Sprintf("Response %v: %v", res.StatusCode, string(errout)))
		} else {
			return errors.New(fmt.Sprintf("Response %v", res.StatusCode))
		}
	}
	return nil
}