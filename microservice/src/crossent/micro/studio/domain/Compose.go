package domain

const (
	STATUS_INITIAL = "INITIAL"
	STATUS_RUNNING = "RUNNING"
	STATUS_STOPED = "STOPED"
	STATUS_ERROR = "ERROR"
)

type Compose struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type ComposeRequest struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	OrgGuid     string   `json:"orgGuid,omitempty"`
	OrgName     string   `json:"orgName,omitempty"`
	SpaceGuid   string   `json:"spaceGuid,omitempty"`
	SpaceName   string   `json:"spaceName,omitempty"`
	Version     string   `json:"version,omitempty"`
	Description string   `json:"description,omitempty"`
	Apps        Apps     `json:"apps,omitempty"`
	Services    ServiceInstances `json:"services,omitempty"`
	Visible     bool     `json:"visible,omitempty"`
	Status      string   `json:"status,omitempty"`
	Composition MicroserviceComposition `json:"composition,omitempty"`
}

type MicroserviceComposition struct {
	Apps        Apps     `json:"apps,omitempty"`
	Services    ServiceInstances `json:"services,omitempty"`
	ServiceBindings    ServiceBindings `json:"serviceBindings,omitempty"`
	Policies    []Policy `json:"policies"`
	Routes      []map[string]interface{} `json:"routes"`
	Configs     []map[string]interface{} `json:"configs"`
}

type MicroserviceService struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	MicroID     int    `json:"microId,omitempty"`
	ServiceGuid string `json:"serviceGuid,omitempty"`
	ServicePlanGuid string `json:"servicePlanGuid,omitempty"`
}

type MicroserviceApp struct {
	ID          int    `json:"id,omitempty"`
	MicroID     int    `json:"microId,omitempty"`
	AppGuid     string `json:"appGuid,omitempty"`
	SourceGuid  string `json:"sourceGuid,omitempty"`
}


type Meta struct {
	Guid      string `json:"guid"`
	Url       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type ServiceInstances struct {
	Count     int                       `json:"total_results"`
	Pages     int                       `json:"total_pages"`
	NextUrl   string                    `json:"next_url"`
	Resources []ServiceInstanceResource `json:"resources"`
}
type ServiceInstanceResource struct {
	Meta   Meta            `json:"metadata"`
	Entity ServiceInstance `json:"entity"`
}
type ServiceInstance struct {
	Name               string                 `json:"name"`
	Credentials        map[string]interface{} `json:"credentials"`
	ServicePlanGuid    string                 `json:"service_plan_guid"`
	Parameters         map[string]interface{} `json:"parameters,omitempty"`
	SpaceGuid          string                 `json:"space_guid"`
	Type               string                 `json:"type"`
	ServiceGuid        string                 `json:"service_guid"`
}
type Services struct {
	Count     int                `json:"total_results"`
	Pages     int                `json:"total_pages"`
	NextUrl   string             `json:"next_url"`
	Resources []ServicesResource `json:"resources"`
}

type ServicesResource struct {
	Meta   Meta    `json:"metadata"`
	Entity Service `json:"entity"`
}

type Service struct {
	Guid              string   `json:"guid"`
	Label             string   `json:"label"`
	Description       string   `json:"description"`
	Active            bool     `json:"active"`
	Bindable          bool     `json:"bindable"`
	ServiceBrokerGuid string   `json:"service_broker_guid"`
	PlanUpdateable    bool     `json:"plan_updateable"`
	Tags              []string `json:"tags"`
	ServicePlans      []string `json:"service_plans"`
}
type ServiceBindings struct {
	Count     int                       `json:"total_results"`
	Pages     int                       `json:"total_pages"`
	NextUrl   string                    `json:"next_url"`
	Resources []ServiceBindingResource  `json:"resources"`
}
type ServiceBindingResource struct {
	Meta   Meta            `json:"metadata"`
	Entity ServiceBinding  `json:"entity"`
}
type ServiceBinding struct {
	AppGuid             string                 `json:"app_guid"`
	ServiceInstanceGuid string                 `json:"service_instance_guid"`
	Credentials         map[string]interface{} `json:"credentials"`
}
type Apps struct {
	Count     int           `json:"total_results"`
	Pages     int           `json:"total_pages"`
	NextUrl   string        `json:"next_url"`
	Resources []AppResource `json:"resources"`
}
type AppResource struct {
	Meta   Meta `json:"metadata"`
	Entity App  `json:"entity"`
}
type App struct {
	Name                     string                 `json:"name"`
	Memory                   int                    `json:"memory"`
	Instances                int                    `json:"instances"`
	DiskQuota                int                    `json:"disk_quota"`
	State                    string                 `json:"state"`
	SpaceGuid                string                 `json:"space_guid"`
	Buildpack                string                 `json:"buildpack"`
	Environment              map[string]interface{} `json:"environment_json"`
}

type Policies struct {
	TotalPolicies int `json:"total_policies"`
	Policies      []Policy `json:"policies"`
}

type Policy struct {
	Source struct {
		ID string `json:"id"`
	} `json:"source"`
	Destination struct {
		ID string `json:"id"`
		Port struct {
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"ports"`
		Protocol string `json:"protocol"`
	} `json:"destination"`
}

type Results struct {
	Count     int           `json:"total_results"`
	Pages     int           `json:"total_pages"`
	NextUrl   string        `json:"next_url"`
	Resources []Resource `json:"resources"`
}
type Resource struct {
	Meta   Meta  `json:"metadata"`
	Entity Entity  `json:"entity"`
}
type Entity struct {
	Name  string  `json:"name"`
}

type RoutesResponse struct {
	Count     int              `json:"total_results"`
	Pages     int              `json:"total_pages"`
	NextUrl   string           `json:"next_url"`
	Resources []RoutesResource `json:"resources"`
}
type RoutesResource struct {
	Meta   Meta  `json:"metadata"`
	Entity Route `json:"entity"`
}
type Route struct {
	Guid                string `json:"guid"`
	Host                string `json:"host"`
	Path                string `json:"path"`
	DomainGuid          string `json:"domain_guid"`
	SpaceGuid           string `json:"space_guid"`
	ServiceInstanceGuid string `json:"service_instance_guid"`
	Port                int    `json:"port"`
}