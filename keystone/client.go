package keystone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AuthInfo *KeystoneAuth
}

func NewClient(authInfo KeystoneAuth) (*Client, error) {

	if authInfo.AuthURL == "" {
		return nil, fmt.Errorf("missing URL")
	}
	client := Client{AuthInfo: &authInfo}
	auth := NewAuth(authInfo.UserName, authInfo.Password, authInfo.DomainName)
	token, userid, err := client.Tokens(auth)
	if err != nil {
		return nil, fmt.Errorf("get token error")
	}
	client.AuthInfo.Token = token
	client.AuthInfo.UserId = userid
	return &client, nil
}

func (c *Client) DoRequest(r KeyRequest) (KeyResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest(r.Method, c.AuthInfo.AuthURL+r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return KeyResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(X_SUBJECT_TOKEN_HEADER, c.AuthInfo.Token)

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

func (c *Client) Tokens(auth Auth) (string, string, error) {
	jsonStr, err := json.Marshal(SingleAuth{Auth: auth})
	if err != nil {
		return "", "", fmt.Errorf("invalid auth request: ", err)
	}

	resp, err := c.doRequest(KeyRequest{
		URL:          fmt.Sprintf("%s/v3/auth/tokens", c.AuthInfo.AuthURL),
		Method:       http.MethodPost,
		Body:         jsonStr,
		OkStatusCode: http.StatusCreated,
	})

	if err != nil {
		return "", "", err
	}

	// note: not unmarshalling response body right now
	// since dont need anything from it yet
	token := resp.Headers.Get(X_SUBJECT_TOKEN_HEADER)
	if token == "" {
		return "", "", errors.New("No token found in response")
	}
	var keyTokenBody ResTokenBody
	err = json.Unmarshal(resp.Body, &keyTokenBody)

	if err != nil {
		return "", "", err
	}

	return token, keyTokenBody.Token.User.Id, nil
}
