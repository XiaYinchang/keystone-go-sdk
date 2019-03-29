package keystone

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AuthInfo *KeystoneAuth
}

func NewClient(authInfo *KeystoneAuth) (*Client, error) {

	if authInfo.AuthURL == "" {
		return nil, fmt.Errorf("missing URL")
	}
	client := Client{AuthInfo: authInfo}
	var auth interface{}
	if authInfo.UserName == "admin" {
		auth = NewSystemAuth(authInfo.UserName, authInfo.Password, authInfo.DomainName)
	} else {
		auth = SingleAuth{Auth: NewAuth(authInfo.UserName, authInfo.Password, authInfo.DomainName)}
	}
	token, userid, err := client.Tokens(auth)
	if err != nil {
		return nil, err
	}
	client.AuthInfo.Token = token
	client.AuthInfo.UserId = userid
	return &client, nil
}

func NewClientWithToken(authInfo *KeystoneAuth) (*Client, error) {
	if authInfo.Token == "" {
		return nil, fmt.Errorf("missing token")
	}
	client := Client{AuthInfo: authInfo}
	token, err := client.ValidateToken(authInfo.Token)
	if err != nil {
		return nil, err
	}
	client.AuthInfo.UserId = token.Token.User.Id
	client.AuthInfo.UserName = token.Token.User.Name
	return &client, nil
}

func (c *Client) DoRequest(r KeyRequest) (KeyResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, c.AuthInfo.AuthURL+r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return KeyResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(X_AUTH_TOKEN, c.AuthInfo.Token)
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return KeyResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return KeyResponse{}, err
	}

	if resp.StatusCode != r.OkStatusCode {
		return KeyResponse{}, fmt.Errorf("Error: %s details: %s\n", resp.Status, body)
	}

	return KeyResponse{
		Body:       body,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header}, nil
}

func (c *Client) doRequest(r KeyRequest) (KeyResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return KeyResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return KeyResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return KeyResponse{}, err
	}

	if resp.StatusCode != r.OkStatusCode {
		return KeyResponse{}, fmt.Errorf("Error: %s details: %s\n", resp.Status, body)
	}

	return KeyResponse{
		Body:       body,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header}, nil
}
