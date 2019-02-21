package keystone

import (
	"encoding/json"
	"net/http"
)

func (c *Client) UserProjects(userid string) (*ResProjectBody, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/v3/users/" + userid + "/projects",
		Method:       http.MethodGet,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resProjectBody ResProjectBody
	err = json.Unmarshal(resp.Body, &resProjectBody)

	if err != nil {
		return nil, err
	}
	return &resProjectBody, nil
}
