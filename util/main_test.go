package util

import (
	"MapApi/util/request"
	"MapApi/util/request/url"
	"fmt"
	"testing"
)

func Test_Main(t *testing.T) {
	mapRequest := request.BuildMapRequest4()
	mapRequest.AddMarker("plzeňská 558/24")
	fmt.Println("Точки:")
	mapRequest.PrintPoints()
	fmt.Println("Граф")
	matrix := request.GetGraph(mapRequest.GetTimeMatrix())
	for _, a := range matrix {
		fmt.Println(a)
	}
	fmt.Println("Время в пути:")
	fmt.Println(mapRequest.GetPathTime())
	fmt.Println("Карта:")
	fmt.Println(url.GetUrlWithPathMap(mapRequest.Markers, mapRequest.GetPaths()))
}

func TestMapRequest_GetPathTime(t *testing.T) {
	mapRequest := request.BuildMapRequest4()
	mapRequest.GetPathTime()
}
