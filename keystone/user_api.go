package keystone

import (
	"encoding/json"
	"net/http"
)

func (c *Client) UserProjects(userid string) (*ResProjectsBody, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/v3/users/" + userid + "/projects",
		Method:       http.MethodGet,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resProjectsBody ResProjectsBody
	err = json.Unmarshal(resp.Body, &resProjectsBody)

	if err != nil {
		return nil, err
	}
	return &resProjectsBody, nil
}
