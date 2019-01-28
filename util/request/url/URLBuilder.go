package url

import (
	"fmt"
	"net/url"
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
)

func GetUrlWithMarkersMap(markers []string) string {
	markerString := ""
	for _, marker := range markers {
		markerString += marker
	}
	return fmt.Sprintf(mapUrl, markerString)
}

func GetUrlWithPathMap(markers []string, paths string) string {
	markerString := ""
	for _, marker := range markers {
		markerString += marker
	}
	return fmt.Sprintf(mapUrl+paths, markerString)
}

func GetDistanceUrl(from string, to string) string {
	return fmt.Sprintf(distanceUrl, url.PathEscape(from), url.PathEscape(to))
}
