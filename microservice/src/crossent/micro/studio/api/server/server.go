package server

import (
	"code.cloudfoundry.org/lager"
	"crossent/micro/studio/db"
	"crossent/micro/studio/client"
	"time"
	"math/rand"
)

type Server struct {
	logger        lager.Logger
	repositoryFactory db.RepositoryFactory
	uaa *client.UAA
	//tempFactory   cixpertDb.TempFactory
	//teamFactory db.TeamFactory
	//oauthConfig   *uaaxpert.GenericOAuthConfig
	//authTokenGenerator auth.AuthTokenGenerator
	//csrfTokenGenerator auth.CSRFTokenGenerator
	//expire        time.Duration
	//rejector      auth.Rejector
}

func NewServer(
	logger lager.Logger,
	repositoryFactory db.RepositoryFactory,
	uaa *client.UAA,
	//tempFactory cixpertDb.TempFactory,
	//dbTeamFactory db.TeamFactory,
	//oauthConfig *uaaxpert.GenericOAuthConfig,
	//authTokenGenerator auth.AuthTokenGenerator,
	//csrfTokenGenerator auth.CSRFTokenGenerator,
	//expire time.Duration,
) *Server {
	return &Server{
		logger:      logger,
		repositoryFactory:   repositoryFactory,
		uaa: uaa,
		//tempFactory: tempFactory,
		//teamFactory: dbTeamFactory,
		//oauthConfig: oauthConfig,
		//authTokenGenerator:  authTokenGenerator,
		//csrfTokenGenerator:  csrfTokenGenerator,
		//expire:      expire,
		//rejector:    auth.UnauthorizedRejector{},
	}
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
