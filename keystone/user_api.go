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

func (c *Client) GetUserByName(userName string) (*ResUser, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/v3/users?name=" + userName,
		Method:       http.MethodGet,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resUsersBody ResUsersBody
	err = json.Unmarshal(resp.Body, &resUsersBody)

	if err != nil {
		return nil, err
	}
	if len(resUsersBody.Users) <= 0 {
		return nil, err
	}
	return &(resUsersBody.Users[0]), nil
}

func (c *Client) ListUsers() ([]ResUser, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/v3/users",
		Method:       http.MethodGet,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resUsersBody ResUsersBody
	err = json.Unmarshal(resp.Body, &resUsersBody)

	if err != nil {
		return nil, err
	}
	return resUsersBody.Users, nil
}

func (c *Client) CreateUser(name, password string) error {
	bodyByteArray, err := json.Marshal(&ReqUser{
		UserInfo: User{
			Name:     name,
			Password: password,
		},
	})
	if err != nil {
		return err
	}
	_, err = c.DoRequest(KeyRequest{
		URL:          "/v3/users",
		Method:       http.MethodPost,
		OkStatusCode: http.StatusCreated,
		Body:         bodyByteArray,
	})

	if err != nil {
		return err
	}
	c.CreateProject(name, name+"origin")
	return nil
}
