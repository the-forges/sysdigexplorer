package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.ibm.com/ZaaS/sysdigexplorer/model"
)

type Dashboard struct {
	model.Dashboard
}

const (
	dashboardsEndPoint        = "/api/v2/dashboards"
	defaultDashboardsEndPoint = "/api/v2/defaultDashboards"
)

// NewDashboard instantiates a dashboard with a client or returns an error
func (c *Client) NewDashboard() *Dashboard {
	return &Dashboard{}
}

// GetDashboard fetches a dashboard by ID under a given user account
func (c *Client) GetDashboard(id string) (Dashboard, error) {
	var (
		resp  []byte
		err   error
		value map[string]Dashboard
		d     Dashboard
		ok    bool
	)

	resp, err = c.Request(http.MethodGet, dashboardsEndPoint+"/"+id, nil)
	if err != nil {
		return d, err
	}
	if err = json.Unmarshal(resp, &value); err != nil {
		return d, err
	}

	if d, ok = value["dashboard"]; !ok {
		return d, errors.New(ErrMsgRecordsNotFound)
	}

	return d, nil
}

// ListDashboards fetch all the dashboards under a given user account
func (c *Client) ListDashboards() ([]Dashboard, error) {
	var (
		resp   []byte
		err    error
		values map[string][]Dashboard
		d      []Dashboard
		ok     bool
	)

	resp, err = c.Request(http.MethodGet, dashboardsEndPoint, nil)
	if err != nil {
		return d, err
	}
	if err = json.Unmarshal(resp, &values); err != nil {
		return d, err
	}
	d, ok = values["dashboards"]
	if !ok {
		return d, errors.New(ErrMsgRecordsNotFound)
	}

	return d, nil
}
