package aliyun

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type WebhookBodyPushData struct {
	Digest   string `json:"digest"`
	PushedAt string `json:"pushed_at"`
	Tag      string `json:"tag"`
}

type WebhookBodyRepository struct {
	DateCreated            string `json:"date_created"`
	Name                   string `json:"name"`
	Namespace              string `json:"namespace"`
	Region                 string `json:"region"`
	RepoAuthenticationType string `json:"repo_authentication_type"`
	RepoFullName           string `json:"repo_full_name"`
	RepoOriginType         string `json:"repo_origin_type"`
	RepoType               string `json:"repo_type"`
}

type WebhookBody struct {
	PushData   WebhookBodyPushData   `json:"push_data"`
	Repository WebhookBodyRepository `json:"repository"`
}

func Parse(r *http.Request) (*WebhookBody, error) {
	// parse request
	//fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	//println("json:", string(body))

	// Unmarshal json
	var data WebhookBody
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	//fmt.Printf("%+v", data)
	return &data, nil
}
