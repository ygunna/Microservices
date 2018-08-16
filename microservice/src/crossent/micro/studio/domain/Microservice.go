package domain

import (
	"github.com/alexedwards/scs"
)

const (
	UAA_TOKEN_NAME = "access_token"
	APP_STATE_STOPPED = "STOPPED"
	APP_STATE_STARTED = "STARTED"
	USER_ID = "user_id"
	SERVICE_REGISTRY_SERVER = "registry-server"
	SERVICE_CONFIG_SERVER = "config-server"
	SERVICE_GATEWAY_SERVER = "gateway-server"
	MSA_REGISTRY_APP = "registryapp"
	MSA_CONFIG_APP = "configapp"
	MSA_GATEWAY_APP = "gatewayapp-micro"
	SAMPLE_APP_FRONT = "front"
	SAMPLE_APP_BACK = "back"

	BASIC_USER = "basic-user"
	BASIC_SECRET = "basic-secret"
)
var SessionManager scs.Manager

type TokenRequest struct {
	GrantType string `json:"grantType"`
	ResponseType string `json:"responseType"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	UserId string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	UaaToken TokenJSON `json:"tokenJson"`
}

type TokenJSON struct {
	User_id string `json:"user_id"`
	User_name string `json:"user_name"`
	Email string `json:"email"`
	Exp int32 `json:"exp"`
	Scope []string `json:"scope"`
}

type CloudFoundryErr struct {
	Description string `json:"description"`
	ErrorCode string `json:"error_code"`
}

type CloudFoundryErrBody struct {
	HttpStatusCode int `json:"httpStatusCode"`
	CfErrorCode string `json:"cfErrorCode"`
	Message string `json:"Message"`
}
