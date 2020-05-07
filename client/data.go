package client

import (
	"bytes"
	"encoding/json"

	"github.com/the-forges/sysdigexplorer/model"
)

var (
	DataEndpoint = "/api/data"
)

type DataOptions struct {
	Last           int         `json:"last,omitempty"`
	Sampling       int         `json:"sampling,omitempty"`
	Filter         interface{} `json:"filter,omitempty"`
	Metrics        []DataOptionsMetric `json:"metrics,omitempty"`
	DataSourceType string      `json:"dataSourceType,omitempty"`
	Paging         Paging      `json:"paging,omitempty"`
}

type Aggregations map[string]string

type DataOptionsMetric struct {
	ID           string       `json:"id,omitempty"`
	Aggregations Aggregations `json:"aggregations,omitempty"`
}
type Paging struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type Data struct {
	Data []model.Datum `json:"data"`
	Start int `json:"start"`
	End int `json:"end"`
}

// GetData uses the data API to get metrics and other data
// https://docs.sysdig.com/en/working-with-the-data-api.html
func (c *Client) GetData (options DataOptions) (Data, error) {
	var (
		resp []byte
		err error
		d Data
		reqOptions = make(map[string]interface{})
		jsonBody []byte
	)
		

	jsonBody, err = json.Marshal(options)
	if err != nil {
		return d, err
	}
	
	reqOptions["body"] = bytes.NewReader(jsonBody)

	resp, err = c.Request("POST", DataEndpoint, reqOptions)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(resp, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}