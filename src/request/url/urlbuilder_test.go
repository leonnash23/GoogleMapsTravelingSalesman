package url

import (
	"MapApi/src/domain"
	"testing"
)

func TestGetUrlWithMarkersMap(t *testing.T) {
	list := domain.MarkerList{}
	list.AddMarker("5 просека 99Б")
	markersMap := GetUrlWithMarkersMap(list)
	if markersMap != "https://maps.googleapis.com/maps/api/staticmap?&size=1900x1080&scale=2&maptype=roadmap&markers=label:0|5%20%D0%BF%D1%80%D0%BE%D1%81%D0%B5%D0%BA%D0%B0%2099%D0%91&key=AIzaSyBdQ-zn9G0VsgWMZCsshlAtd6g13STrQ90" {
		t.Error("Error link")
		t.Log(markersMap)
	}
}

func TestApiKey(t *testing.T) {
	if apiKey == "" {
		t.Error("Empty api key")
	}
}
