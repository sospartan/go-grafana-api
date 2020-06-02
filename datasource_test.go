package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	createdDataSourceJSON = `
{
	"id": 1,
	"message": "Datasource added",
	"name": "test_datasource"
}
`
)

func TestNewDataSource(t *testing.T) {
	server, client := gapiTestTools(200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "foo",
		Type:      "cloudwatch",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JSONData{
			AssumeRoleArn:           "arn:aws:iam::123:role/some-role",
			AuthType:                "keys",
			CustomMetricsNamespaces: "SomeNamespace",
			DefaultRegion:           "us-east-1",
			TlsSkipVerify:           true,
		},
		SecureJSONData: SecureJSONData{
			AccessKey: "123",
			SecretKey: "456",
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewPrometheusDataSource(t *testing.T) {
	server, client := gapiTestTools(200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "foo_prometheus",
		Type:      "prometheus",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JSONData{
			HttpMethod:   "POST",
			QueryTimeout: "60s",
			TimeInterval: "1m",
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}
