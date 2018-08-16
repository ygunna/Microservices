package config

// referencd : cf-redis-broker/brokerconfig/config.go

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

type Config struct {
	Service []ServiceConfiguration `yaml:"service"`
	BrokerUsername string `yaml:"broker_username"`
	BrokerPassword string `yaml:"broker_password"`
	//AdminUsername  string `yaml:"admin_username"`
	//AdminPassword  string `yaml:"admin_password"`
	//ConcourseURL   string `yaml:"concourse_url"`
	//CFURL          string `yaml:"cf_url"`
	//TokenURL       string `yaml:"token_url"`
	//AuthURL        string `yaml:"auth_url"`
	//ClientID       string `yaml:"client_id"`
	//ClientSecret   string `yaml:"client_secret"`
	Port           string `yaml:"port"`
	//PipelinePath   string `yaml:"pipeline_path`

	AppPath   string `yaml:"app_path`

	Api 	       string `yaml:"api"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	Organization   string `yaml:"organization"`
	Space          string `yaml:"space"`
	SkipCertCheck  bool   `yaml:"skip_cert_check"`
}

type ServiceConfiguration struct {
	ServiceName                 string    `yaml:"service_name"`
	ServiceID                   string    `yaml:"service_id"`
	Description                 string    `yaml:"description"`
	PlanID	                    string    `yaml:"plan_id"`
	DisplayName                 string    `yaml:"display_name"`
	LongDescription             string    `yaml:"long_description"`
	ProviderDisplayName         string    `yaml:"provider_display_name"`
	DocumentationURL            string    `yaml:"documentation_url"`
	SupportURL                  string    `yaml:"support_url"`
	IconImage                   string    `yaml:"icon_image"`
	Tags			    []string  `yaml:"tags"`
}


func ParseConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	if err := candiedyaml.NewDecoder(file).Decode(&config); err != nil {
		return Config{}, err
	}

	return config, ValidateConfig(config)
}

func ValidateConfig(config Config) error {
	if config.BrokerUsername == "" {
		return fmt.Errorf("broker_username is not set")
	}
	if config.BrokerPassword == "" {
		return fmt.Errorf("broker_password is not set")
	}
	//if config.AdminUsername == "" {
	//	return fmt.Errorf("admin_username is not set")
	//}
	//if config.AdminPassword == "" {
	//	return fmt.Errorf("admin_password is not set")
	//}
	//if config.ConcourseURL == "" {
	//	return fmt.Errorf("concourse_url is not set")
	//}
	//if config.CFURL == "" {
	//	return fmt.Errorf("cf_url is not set")
	//}
	//if config.TokenURL == "" {
	//	return fmt.Errorf("token_url is not set")
	//}
	//if config.AuthURL == "" {
	//	return fmt.Errorf("auth_url is not set")
	//}
	//if config.ClientID == "" {
	//	return fmt.Errorf("client_id is not set")
	//}
	//if config.ClientSecret == "" {
	//	return fmt.Errorf("client_secret is not set")
	//}
	if config.Port == "" {
		return fmt.Errorf("port is not set")
	}

	return nil
}

