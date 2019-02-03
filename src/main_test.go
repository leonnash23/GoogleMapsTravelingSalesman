package src

import (
	"MapApi/src/request"
	"fmt"
	"testing"
)

func Test_Main(t *testing.T) {
	mapRequest := request.BuildMapRequest4()
	//mapRequest.AddMarker("plzeňská 558/24")
	fmt.Println("Точки:")
	list := mapRequest.Markers.GetAddressList()
	for i, m := range list {
		fmt.Printf("%d) %s\n", i, m)
	}
	fmt.Println("Граф")
	matrix := request.GetGraph(mapRequest.GetTimeMatrix())
	for _, a := range matrix {
		fmt.Println(a)
	}
	fmt.Println("Время в пути:")
	fmt.Println(mapRequest.GetPathTime())
	fmt.Println("Карта:")
	fmt.Println(mapRequest.GetLinkMapUrl())
}

func TestMapRequest_GetPathTime(t *testing.T) {
	mapRequest := request.BuildMapRequest4()
	mapRequest.GetPathTime()
}
