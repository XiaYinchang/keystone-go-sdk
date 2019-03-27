package keystone

import (
	"net/http"
)

const X_SUBJECT_TOKEN_HEADER = "X-Subject-Token"
const X_AUTH_TOKEN = "X-Auth-Token"

type KeystoneAuth struct {
	AuthURL     string
	APIVersion  string
	DomainName  string
	ProjectName string
	UserName    string
	UserId      string
	Password    string
	Token       string
}

type Domain struct {
	Name string `json:"name"`
}

type IdentityUser struct {
	User User `json:"user"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Domain   Domain `json:"domain"`
}

type Identity struct {
	Methods  []string     `json:"methods"`
	Password IdentityUser `json:"password"`
}

type Auth struct {
	Identity Identity `json:"identity"`
}

type SingleAuth struct {
	Auth Auth `json:"auth"`
}

type SystemScopedAuth struct {
	Auth SysAuth `json:"auth"`
}

type SysAuth struct {
	Identity Identity `json:"identity"`
	Scope Scope `json:"scope"`
}

type Scope struct {
	System System `json:"system"`
}

type System struct {
	All bool `json:"all"`
}

type ResTokenBody struct {
	Token ResToken `json:"token"`
}

type ResToken struct {
	User ResUser `json:"user"`
}

type ResUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ReqUser struct {
	UserInfo User `json:"user"`
}

type KeyRequest struct {
	URL          string
	Method       string
	Headers      map[string]string
	Body         []byte
	OkStatusCode int
}

type KeyResponse struct {
	Body       []byte
	StatusCode int
	Headers    http.Header
}

type ResProjectsBody struct {
	Projects []ResProject `json:"projects"`
	Links    ResLinks     `json:"links"`
}

type ResProjectBody struct {
	Project ResProject `json:"project"`
}

type ResProject struct {
	Description string   `json:"description"`
	DomainId    string   `json:"domain_id"`
	Enabled     bool     `json:"enabled"`
	Id          string   `json:"id"`
	Links       ResLinks `json:"links"`
	Name        string   `json:"name"`
	ParentId    string   `json:"parent_id"`
}

type Project struct {
	Name string `json:"name"`
}

type ReqProject struct {
	ProjectInfo Project `json:"project"`
}

type ResUsersBody struct {
	Users []ResUser `json:"users"`
	Links ResLinks  `json:"links"`
}

type ResUserBody struct {
	User ResUser `json:"user"`
}

type ResRole struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ResRolesBody struct {
	Roles []ResRole `json:"roles"`
	Links ResLinks  `json:"links"`
}

type ResLinks map[string]string
