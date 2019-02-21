# go-keystone :key:

Go library for Keystone v3.0 API

### Installation

Install using `go get github.com/XiaYinchang/go-keystone`.


### Usage

```go
// create new client
authInfo := KeystoneAuth {
    AuthURL:     "http://192.168.56.101:5000",
	APIVersion:  "v3",
	DomainName:  "Default",
	ProjectName: "admin",
	UserName:    "admin",
	Password:    "test",
}

client, err := keystone.NewClient(authInfo)
if err != nil {
    log.Fatal(err)
}

// get token
token := client.AuthInfo.Token

// get projects of user with userid
userProjects := client.UserProjects(client.AuthInfo.UserId)
```
