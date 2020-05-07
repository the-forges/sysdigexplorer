package client

// var APIKey = ""
// var client, _ = NewClient("", APIKey)

// func TestGetData(t *testing.T) {
// 	options := DataOptions{
// 		Last: 600,
// 		Sampling: 600,
// 		Filter: nil,
// 		Metrics: []DataOptionsMetric{
// 			{
// 				ID: "agent.tag.lpar",
// 			},
// 		},
// 		DataSourceType: "container",
// 		Paging: Paging{
// 			To: 99,
// 			From: 0,
// 		},
// 	}
// 	r, err := client.GetData(options)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, r)
// 	fmt.Println(r.Data[0].D[0], r.Data[0].T)
// }

// func TestClientRequest (t *testing.T) {
// 	d, err := client.ListDashboards()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println(d)
// }