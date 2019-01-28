package distance

import (
	"MapApi/util/request/url"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type distanceMatrix struct {
	DestinationAddresses []string `json:"destination_addresses"`
	OriginAddresses      []string `json:"origin_addresses"`
	Rows                 []row    `json:"rows"`
	Status               string   `json:"status"`
}

type row struct {
	Elements []element `json:"elements"`
}

type element struct {
	Distance distance `json:"distance"`
	Duration duration `json:"duration"`
	Status   string   `json:"status"`
}

type distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type duration struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

func GetDurationSecond(from string, to string) int {
	resp, _ := http.Get(url.GetDistanceUrl(from, to))
	html, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return parseDurationSecond(html)
}

func parseDurationSecond(html []byte) int {
	matrix := &distanceMatrix{}
	json.Unmarshal(html, &matrix)
	return matrix.Rows[0].Elements[0].Duration.Value
}
