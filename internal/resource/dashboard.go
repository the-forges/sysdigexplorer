package resource

import (
	"encoding/json"
	"net/http"

	"github.ibm.com/ZaaS/sysdigexplorer/internal/common"
	"github.ibm.com/ZaaS/sysdigexplorer/internal/errors"
)

type Dashboard struct {
	values map[string]interface{}
	client common.Client
}

const (
	dashboardsEndPoint        = "/api/v2/dashboards"
	defaultDashboardsEndPoint = "/api/v2/defaultDashboards"
)

// NewDashboard instantiates a dashboard with a client or returns an error
func NewDashboard(c common.Client) (*Dashboard, error) {
	if c == nil || !c.Valid() {
		return nil, errors.New(errInvalidClient)
	}
	return &Dashboard{client: c}, nil
}

// Value takes a key, returning the value as an interface which should be then
// type hinted in order to be used
func (d *Dashboard) Value(string) interface{}

// Set takes a key and value to be set for later use
func (d *Dashboard) Set(string, interface{})

// String will return a well formatted url used for any requests
func (d *Dashboard) String() string

// Get fetches a dashboard by ID under a given user account
func (d *Dashboard) Get() (common.Resource, error)

// List fetchs all the dashboards under a given user account
func (d *Dashboard) List() ([]common.Resource, error) {
	var (
		resp       []byte
		err        error
		values     []map[string]interface{}
		dashboards []common.Resource
	)
	resp, err = d.client.Request(http.MethodGet, dashboardsEndPoint, nil)
	if err != nil {
		return dashboards, err
	}
	if err = json.Unmarshal(resp, &values); err != nil {
		return dashboards, err
	}
	for _, v := range values {
		r, err := NewDashboard(d.client)
		if err != nil {
			return dashboards, err
		}
		r.values = v
		dashboards = append(dashboards, r)
	}
	return dashboards, nil
}
