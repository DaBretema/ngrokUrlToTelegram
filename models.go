package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//
// Helpers
//

func doGetReq(uri string) *http.Response {
	response, err := http.Get(uri)
	if err != nil {
		log.Fatalf("Request to URI [%v] failed:\n[Reason]-> %v\n", uri, err)
	}
	return response
}

func uriToOBJ(uri string, obj interface{}) {
	bytes, err := ioutil.ReadAll(doGetReq(uri).Body)
	if err != nil {
		log.Fatalf("Fail reading URI response [%v]\n[Reason]->%v\n", uri, err)
	}
	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		log.Fatalf("Err unmarshaling JSON: %v\n", err)
	}
}

//

// Reply of --> "https://api.telegram.org/bot<TOKEN>/getUpdates"

type tgUpdatesReply struct {
	Ok     bool `json:"ok"`
	Result []struct {
		UpdateID int `json:"update_id"`
		Message  struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Chat struct {
				ID        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date     int    `json:"date"`
			Text     string `json:"text"`
			Entities []struct {
				Offset int    `json:"offset"`
				Length int    `json:"length"`
				Type   string `json:"type"`
			} `json:"entities"`
		} `json:"message"`
	} `json:"result"`
}

func (o *tgUpdatesReply) FillFromURI(uri string) {
	uriToOBJ(uri, o)
	if !o.Ok {
		log.Fatalf("Telegram request error. Maybe BAD token?")
	}
}

//

// Reply of --> "http://127.0.0.1:4040/api/tunnels"

type ngrokReply struct {
	Tunnels []struct {
		Name      string `json:"name"`
		URI       string `json:"uri"`
		PublicURL string `json:"public_url"`
		Proto     string `json:"proto"`
		Config    struct {
			Addr    string `json:"addr"`
			Inspect bool   `json:"inspect"`
		} `json:"config"`
		Metrics struct {
			Conns struct {
				Count  float64 `json:"count"`
				Gauge  float64 `json:"gauge"`
				Rate1  float64 `json:"rate1"`
				Rate5  float64 `json:"rate5"`
				Rate15 float64 `json:"rate15"`
				P50    float64 `json:"p50"`
				P90    float64 `json:"p90"`
				P95    float64 `json:"p95"`
				P99    float64 `json:"p99"`
			} `json:"conns"`
			HTTP struct {
				Count  float64 `json:"count"`
				Rate1  float64 `json:"rate1"`
				Rate5  float64 `json:"rate5"`
				Rate15 float64 `json:"rate15"`
				P50    float64 `json:"p50"`
				P90    float64 `json:"p90"`
				P95    float64 `json:"p95"`
				P99    float64 `json:"p99"`
			} `json:"http"`
		} `json:"metrics"`
	} `json:"tunnels"`
	URI string `json:"uri"`
}

func (o *ngrokReply) FillFromURI(uri string) {
	uriToOBJ(uri, o)
}

//
