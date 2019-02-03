package request

import (
	"testing"
)

func TestGetGraph(t *testing.T) {
	test4(t)
	test3(t)
}

func test4(t *testing.T) {
	mapRequest := BuildMapRequest4()
	graph := GetGraph(mapRequest.GetTimeMatrix())
	if mapRequest.Markers.Count() != len(graph)+1 {
		t.Error("Количество вершин - 1")
	}
	if graph[0] != "0->1" {
		t.Error("ER:0->1, AR:" + graph[0])
	}
	if graph[1] != "1->2" {
		t.Error("ER:1->2, AR:" + graph[1])
	}
	if graph[2] != "2->3" {
		t.Error("ER:2->3, AR:" + graph[2])
	}
}

func test3(t *testing.T) {
	mapRequest3 := BuildMapRequest3()
	graph3 := GetGraph(mapRequest3.GetTimeMatrix())
	if mapRequest3.Markers.Count() != len(graph3)+1 {
		t.Error("Количество вершин - 1")
	}
	if graph3[0] != "0->1" {
		t.Error("ER:0->1, AR:" + graph3[0])
	}
	if graph3[1] != "1->2" {
		t.Error("ER:1->2, AR:" + graph3[1])
	}
}
