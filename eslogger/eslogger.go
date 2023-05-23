package eslogger

import (
	"context"
	"encoding/json"
	"errors"
	"loggerStruct"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var es *elasticsearch.Client

func MakeConnection() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://cc37831c624e432d9b3c7943eb8fd6bb.us-central1.gcp.cloud.es.io:443",
		},
		Username: "elastic",
		Password: "x7AOv9fXKzUumAcE5wHVBvEq",
	}

	elasticClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, errors.New("can't connect to elastic cloud, please check your internet connectivity or firewall settings for any errors")
	}
	es = elasticClient
	return es, nil
}

func LogError(jsonString loggerStruct.Logger) error {
	body, err := json.Marshal(jsonString)
	if err != nil {
		return err
	}
	// Prepare the Elasticsearch request
	req := esapi.IndexRequest{
		Index:      "search-logger",
		DocumentID: "", // Elasticsearch will generate a unique document ID
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return errors.New("[" + strconv.Itoa(res.StatusCode) + "] " + res.String())
	} else {
		return nil
	}
}
