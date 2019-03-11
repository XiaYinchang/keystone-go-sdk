package keystone

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetRoleByName(roleName string) (*ResRole, error) {
	resp, err := c.DoRequest(KeyRequest{
		URL:          "/v3/roles?name="+roleName,
		Method:       http.MethodGet,
		OkStatusCode: http.StatusOK,
	})
	if err != nil {
		return nil, err
	}
	var resRolesBody ResRolesBody
	err = json.Unmarshal(resp.Body, &resRolesBody)

	if err != nil {
		return nil, err
	}
	if len(resRolesBody.Roles)<=0 {
		return nil, err
	}
	return &(resRolesBody.Roles[0]), nil
}

func (c *Client) AssignRoleToUserOnProject(roleName, userName, projectName string) ( error) {
	roleInfo, err := c.GetRoleByName(roleName)
		if err != nil {
		return err
	}
	userInfo, err := c.GetUserByName(userName)
		if err != nil {
		return err
	}
	projectInfo, err := c.GetProjectByName(projectName)
	if err != nil {
		return err
	}
	_, err = c.DoRequest(KeyRequest{
		URL:          "/v3/projects/" +projectInfo.Id + "/users/"+ userInfo.Id + "/roles/" +roleInfo.Id,
		Method:       http.MethodPut,
		OkStatusCode: http.StatusNoContent,
	})
	if err != nil {
		return err
	}
	return nil
}
