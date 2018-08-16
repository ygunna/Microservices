package broker

import (
	"code.cloudfoundry.org/lager"
	"context"
	"errors"
	//"github.com/18F/concourse-broker/cf"
	//"github.com/18F/concourse-broker/concourse"
	//"github.com/18F/concourse-broker/config"
	"github.com/pivotal-cf/brokerapi"
	"crossent/micro/broker/config"
	//"github.com/concourse/go-concourse/concourse"
	"fmt"
	"crossent/micro/broker/microservicebroker"
	//"code.cloudfoundry.org/localip"
	//"io/ioutil"
	//"os"
	"bytes"
	"strings"
	"net/url"
	"encoding/json"
)

// New returns a new concourse service broker instance.
func New(credentialsGenerator microservicebroker.CredentialsGenerator, logger lager.Logger, config config.Config) brokerapi.ServiceBroker {
	return &microserviceBroker{credentialsGenerator: credentialsGenerator, logger: logger, config: config}
}

type InstanceBinder struct {
	username string
	password string
}

type microserviceBroker struct {
	credentialsGenerator microservicebroker.CredentialsGenerator
	logger   lager.Logger
	config      config.Config
	InstanceBinders  map[string]InstanceBinder
}

func (c *microserviceBroker) Services(context context.Context) []brokerapi.Service {
	c.logger.Info("Accessing service catalog")
	brokerapiservices := []brokerapi.Service{}

	for _, serviceConfig := range c.config.Service {
		if serviceConfig.ServiceName != "micro-gateway-server" {
			service := brokerapi.Service{
				ID:          serviceConfig.ServiceID,
				Name:        serviceConfig.ServiceName,
				Description: serviceConfig.Description,
				Bindable:    true,
				Tags:        serviceConfig.Tags, //[]string{"configuration", "spring-cloud"},
				Plans: []brokerapi.ServicePlan{
					brokerapi.ServicePlan{
						ID:          serviceConfig.PlanID,
						Name:        "standard",
						Description: serviceConfig.Description,
						Metadata: &brokerapi.ServicePlanMetadata{
							DisplayName: "Standard",
							Bullets:     []string{"configuration", "Multi-tenant"},
							Costs: []brokerapi.ServicePlanCost{
								brokerapi.ServicePlanCost{
									Amount: map[string]float64{"usd": 0.0},
									Unit:   "MONTHLY",
								},
							},
						},
					},
				},
				Metadata: &brokerapi.ServiceMetadata{
					DisplayName:         serviceConfig.DisplayName,
					ImageUrl:            fmt.Sprintf("data:image/png;base64,%s", serviceConfig.IconImage),
					LongDescription:     serviceConfig.LongDescription,
					ProviderDisplayName: serviceConfig.ProviderDisplayName,
					DocumentationUrl:    serviceConfig.DocumentationURL,
					SupportUrl:          serviceConfig.SupportURL,
				},
			}

			brokerapiservices = append(brokerapiservices, service)
		}
	}
	return brokerapiservices
}
func (c *microserviceBroker) Provision(context context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {

	var param microservicebroker.Parameters
	if details.RawParameters != nil {
		p_err := json.Unmarshal(details.RawParameters, &param)
		if p_err != nil {
			return brokerapi.ProvisionedServiceSpec{}, p_err
		}
	}

	microserviceClient := microservicebroker.NewClient(c.config, c.logger, details)
	credentials := c.credentialsGenerator.Generate(instanceID)
	err := microserviceClient.CreateService(instanceID, credentials, param)
	if err != nil {
		return brokerapi.ProvisionedServiceSpec{}, err
	}

	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (c *microserviceBroker) Deprovision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {

	return brokerapi.DeprovisionServiceSpec{}, nil
}

func (c *microserviceBroker) Bind(context context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	//// rabbitmq broker_binding.go reference
	fmt.Println(">>",instanceID, bindingID, details.ServiceID)

	var buffer bytes.Buffer
	var appName string

	for _, l := range c.config.Service {
		if l.ServiceID == details.ServiceID {

			buffer.WriteString(l.DisplayName)
			buffer.WriteString(instanceID)
			appName = buffer.String()
		}
	}

	creds := map[string]interface{}{}

	if appName != "" {
		credentials := c.credentialsGenerator.Generate(instanceID)
		scheme, domain := getDomain(c.config.Api, c.config.SkipCertCheck)
		creds["uri"] = fmt.Sprintf("%s://%s:%s@%s.%s", scheme, credentials.Username, credentials.Password, appName, domain)
		creds[microservicebroker.BASIC_USER] = credentials.Username
		creds[microservicebroker.BASIC_SECRET] = credentials.Password
	}



	return brokerapi.Binding{Credentials: creds }, nil
}
func (c *microserviceBroker) Unbind(context context.Context, instanceID, bindingID string,	details brokerapi.UnbindDetails) error {
	return nil //errors.New("service does not support bind")
}

func (c *microserviceBroker) Update(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, errors.New("service does not support update")
}

func (c *microserviceBroker) LastOperation(context context.Context, instanceID, operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, errors.New("service does not support LastOperation")
}

func getDomain(surl string, skipcertcheck bool) (string, string) {
	u, err := url.Parse(surl)
	if err != nil {
		fmt.Println(err)
		return "http", "bosh-lite.com"
	}
	parts := strings.Split(u.Hostname(), ".")

	domain := "";
	if len(parts) > 3 {
		domain = parts[len(parts)-3] + "." + parts[len(parts)-2] + "." + parts[len(parts)-1]
	} else {
		domain = parts[len(parts)-2] + "." + parts[len(parts)-1]
	}

	scheme := u.Scheme
	if skipcertcheck {
		scheme = "http"
	}

	return scheme, domain
}