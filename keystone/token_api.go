package keystone

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) ValidateToken(token string) (*ResTokenBody, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/v3/auth/tokens",
		Method:       http.MethodGet,
		OkStatusCode: http.StatusOK,
		Headers: map[string]string{
			X_SUBJECT_TOKEN_HEADER: token,
		},
	})
	if err != nil {
		return nil, err
	}
	var resTokenBody ResTokenBody
	err = json.Unmarshal(resp.Body, &resTokenBody)
	if err != nil {
		return nil, err
	}
	return &resTokenBody, nil
}

//Tokens issue keystone token
func (c *Client) Tokens(auth interface{}) (string, string, error) {
	jsonStr, err := json.Marshal(auth)
	if err != nil {
		return "", "", fmt.Errorf("invalid auth request: %s", err)
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
