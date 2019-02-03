package url

import (
	"MapApi/src/domain"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	apiKey      = ""
	distanceUrl = "https://maps.googleapis.com/maps/api/distancematrix/json?" +
		"origins=%s" +
		"&destinations=%s" +
		"&mode=walking" +
		"&key=" + apiKey
	mapUrl = "https://maps.googleapis.com/maps/api/staticmap?" +
		"&size=1900x1080" +
		"&scale=2" +
		"&maptype=roadmap" +
		"%s" +
		"&key=" + apiKey

	markerPart = "&markers=label:%s|%s"
)

func GetUrlWithMarkersMap(markers domain.MarkerList) string {
	markerString := getMarkerArgues(markers)
	return fmt.Sprintf(mapUrl, markerString)
}

func getMarkerArgues(markers domain.MarkerList) string {
	markerString := ""
	for _, marker := range markers.GetMarkers() {
		markerString += fmt.Sprintf(markerPart, marker.Label, url.PathEscape(marker.AddressName))
	}
	return markerString
}

func GetUrlWithPathMap(markers domain.MarkerList, graph []string) string {
	markerString := getMarkerArgues(markers)
	return fmt.Sprintf(mapUrl, markerString) + getPathArgues(markers, graph)
}

func getPathArgues(markers domain.MarkerList, graph []string) string {
	req := ""
	for _, s := range graph {
		split := strings.Split(s, "->")
		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])
		markers := markers.GetMarkers()
		req += "&path=color:0x0000ff|weight:5|" + url.PathEscape(markers[first].AddressName) + "|" + url.PathEscape(markers[second].AddressName)
	}
	return req
}

func GetDistanceUrl(from string, to string) string {
	return fmt.Sprintf(distanceUrl, url.PathEscape(from), url.PathEscape(to))
}
