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

type ResProjectBody struct {
	Projects []ResProject `json:"projects"`
	Links    ResLinks     `json:links`
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

type ResLinks []map[string]string
