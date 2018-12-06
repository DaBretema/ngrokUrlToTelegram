package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//
// Helpers
//

func hGet(uri string) *http.Response {
	response, err := http.Get(uri)
	if err != nil {
		errxit("FAIL GETing @ URI: " + uri)
	}
	return response
}

func uriToOBJ(uri string, obj interface{}) {
	if bytes, err := ioutil.ReadAll(hGet(uri).Body); err != nil {
		errxit("FAIL READing @ URI: " + uri)
	} else if err = json.Unmarshal(bytes, &obj); err != nil {
		errxit("FAIL UNMARSHALing @ URI: " + uri)
	}
}

//

// Reply of --> "https://api.telegram.org/bot<TOKEN>/getUpdates"

type tgModel struct {
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

func (o *tgModel) FillFromURI(uri string) {
	uriToOBJ(uri, o)
	if !o.Ok {
		errxit(_TgBadToken)
	}
}

//

// Reply of --> "http://127.0.0.1:4040/api/tunnels"

type ngModel struct {
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

func (o *ngModel) FillFromURI(uri string) {
	defer recov(_NgrokDown)
	uriToOBJ(uri, o)
}

//
