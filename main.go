package main

import (
	"MapApi/util/request"
	"MapApi/util/request/url"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	testMap()
}

func testMap() {
	markersCount := getDotCount()
	mapRequest := request.MapRequest{}
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < markersCount; i++ {
		println(fmt.Sprintf("Точка #%d:", i))
		line, _, _ := reader.ReadLine()
		mapRequest.AddMarker(string(line))
	}
	println("Граф")
	matrix := request.GetGraph(mapRequest.GetTimeMatrix())
	for _, a := range matrix {
		println(a)
	}
	time := mapRequest.GetPathTime()
	println("Время в пути:")
	println(time)
	println("Карта:")
	println(url.GetUrlWithPathMap(mapRequest.Markers, mapRequest.GetPaths()))
}

func getDotCount() int {
	reader := bufio.NewReader(os.Stdin)
	println("Введи количество точек:")
	line, _, _ := reader.ReadLine()
	markersCount, _ := strconv.Atoi(string(line))
	return markersCount
}
