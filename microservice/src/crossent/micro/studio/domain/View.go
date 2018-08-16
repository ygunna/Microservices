package domain

type View struct {
	ID          int   `json:"id"`
	Name        string   `json:"name"`
	OrgGuid     string   `json:"orgGuid"`
	OrgName     string   `json:"orgName"`
	SpaceGuid   string   `json:"spaceGuid"`
	SpaceName   string   `json:"spaceName"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	App         int      `json:"app"`
	AppGuid     string   `json:"appGuid"`
	Service     int      `json:"service"`
	ServiceGuid string   `json:"serviceGuid"`
	Visible     string   `json:"visible"`
	Status      string   `json:"status"`
	Url         string   `json:"url"`
	ServiceName string   `json:"serviceName"`
	ServiceInstanceName string `json:"serviceInstanceName"`
	Plan        string   `json:"plan"`
	AppName     string   `json:"appName"`
	Source      string   `json:"source"`
	Target      string   `json:"target"`
	Port        int      `json:"port"`
	Swagger	    string   `json:"swagger"`
}

type MicroDetail struct {
	App []View `json:"app"`
	Service []View `json:"service"`
	Policy []View `json:"policy"`
	ServiceApp []View `json:"serviceApp"`
}
