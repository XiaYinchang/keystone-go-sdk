package keystone

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetProjectById(projectid string) (*ResProjectBody, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/v3/projects/" + projectid,
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

func (c *Client) GetProjectByName(projectName string) (*ResProject, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/v3/projects?name=" + projectName,
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
	if len(resProjectsBody.Projects) <= 0 {
		return nil, err
	}
	return &(resProjectsBody.Projects[0]), nil
}

func (c *Client) CheckProjectExist(projectName string) (bool, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/v3/projects?name=" + projectName,
		Method:       http.MethodGet,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return true, err
	}
	var resProjectsBody ResProjectsBody
	err = json.Unmarshal(resp.Body, &resProjectsBody)

	if err != nil {
		return true, err
	}
	if len(resProjectsBody.Projects) <= 0 {
		return false, nil
	}
	return true, nil
}

func (c *Client) CreateProject(userName, projectName string) error {
	reqProjectBodyArray, err := json.Marshal(&ReqProject{
		ProjectInfo: Project{
			Name: projectName,
		},
	})
	if err != nil {
		return err
	}
	_, err = c.DoRequest(KeyRequest{
		URL:          "/v3/projects",
		Method:       http.MethodPost,
		OkStatusCode: http.StatusCreated,
		Body:         reqProjectBodyArray,
	})
	if err != nil {
		return err
	}
	err = c.AssignRoleToUserOnProject("admin", userName, projectName)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteProject(userName, projectName string) error {
	err := c.UnassignRoleToUserOnProject("admin", userName, projectName)
	if err != nil {
		return err
	}
	projectInfo, err := c.GetProjectByName(projectName)
	if err != nil {
		return err
	}
	_, err = c.DoRequest(KeyRequest{
		URL:          "/v3/projects/" + projectInfo.Id,
		Method:       http.MethodDelete,
		OkStatusCode: http.StatusNoContent,
	})
	if err != nil {
		return err
	}
	return nil
}
