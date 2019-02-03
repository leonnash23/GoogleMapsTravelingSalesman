package request

import (
	"MapApi/src/request/url"
	"testing"
)

func TestMapRequest_GetDistances(t *testing.T) {
	mapRequest := BuildMapRequest4()
	distances := mapRequest.GetDistances()

	if _, ok := distances["0->1"]; !ok {
		t.Error("Should contains \"mins\"")
	}
}

func TestMapRequest_GetUrl(t *testing.T) {
	mapRequest := BuildMapRequest4()
	estURL := url.GetUrlWithMarkersMap(mapRequest.Markers)
	er := "https://maps.googleapis.com/maps/api/staticmap?&size=1900x1080&scale=2&maptype=roadmap&markers=label:0|5%20%D0%BF%D1%80%D0%BE%D1%81%D0%B5%D0%BA%D0%B0%2099%D0%91&markers=label:1|%D0%9D%D0%BE%D0%B2%D0%BE-%D1%81%D0%B0%D0%B4%D0%BE%D0%B2%D0%B0%D1%8F%20333&markers=label:2|%D0%93%D0%B0%D1%81%D1%82%D0%B5%D0%BB%D0%BB%D0%BE%2043%D0%B0&markers=label:3|%D0%94%D1%8B%D0%B1%D0%B5%D0%BD%D0%BA%D0%BE%2030&key=AIzaSyBdQ-zn9G0VsgWMZCsshlAtd6g13STrQ90"
	if estURL != er {
		t.Error("ER:" + er + ", AR:" + estURL)
	}
}
