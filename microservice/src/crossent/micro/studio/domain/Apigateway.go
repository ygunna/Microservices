package domain

import (
	"time"
)

type MicroApi struct {
	ID          	int   	 `json:"id"`
	Name        	string   `json:"name"`
	Part     	string   `json:"part"`
	Host     	string   `json:"host"`
	Path   		string   `json:"path"`
	Version     	string   `json:"version"`
	Description 	string   `json:"description"`
	Image 		string   `json:"image"`
	Restapi     	string   `json:"restapi"`
	Active      	string   `json:"active"`
	Userid      	string   `json:"userid"`
	Updated   	time.Time   `json:"updated"`
	Rule        	string   `json:"rule,omitempty"`
	Method 		string   `json:"method,omitempty"`
	PathStrip	string   `json:"pathStrip,omitempty"`
	WhiteList	string   `json:"whitelist,omitempty"`
	Headers		[]HeaderKeyValue    `json:"headers,omitempty"`
	Period		string   `json:"period,omitempty"`
	Average		string   `json:"average,omitempty"`
	Burst		string   `json:"burst,omitempty"`
	MaxConn		string   `json:"maxconn,omitempty"`
	MicroId	        int      `json:"microId,omitempty"`
	Favorite	int	 `json:"favorite,omitempty"`
	Username        string   `json:"username,omitempty"`
	Userpassword    string   `json:"userpassword,omitempty"`
	OrgGuid     	string   `json:"orgguid,omitempty"`
}

type MicroApiRule struct {
	ID          	int   	 `json:"id"`
	Rule        	string   `json:"rule"`
	Active     	string   `json:"active"`
}

type HeaderKeyValue struct {
	Key        	string   `json:"key"`
	Value     	string   `json:"value"`
}