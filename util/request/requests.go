package request

import (
	"MapApi/util/request/distance"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type MapRequest struct {
	pointNames []string
	Markers    []string
}

func (m *MapRequest) AddMarker(pointName string) *MapRequest {
	m.pointNames = append(m.pointNames, pointName)
	m.Markers = append(m.Markers, "&Markers=label:"+strconv.Itoa(len(m.Markers))+"|"+url.PathEscape(pointName))
	return m
}

func (m *MapRequest) PrintPoints() {
	for i, n := range m.pointNames {
		fmt.Printf("%d) %s\n", i, n)
	}
}

func (m *MapRequest) GetPaths() string {
	req := ""
	graph := GetGraph(m.GetTimeMatrix())
	for _, s := range graph {
		split := strings.Split(s, "->")
		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])
		req += "&path=color:0x0000ff|weight:5|" + m.pointNames[first] + "|" + m.pointNames[second]
	}
	return req
}

func (m *MapRequest) GetDistances() map[string]int {
	var ans = make(map[string]int)
	for i := range m.pointNames {
		for j := range m.pointNames {
			if i == j {
				continue
			}
			ans[strconv.Itoa(i)+"->"+strconv.Itoa(j)] = distance.GetDurationSecond(m.pointNames[i], m.pointNames[j])
		}
	}
	return ans
}

func (m *MapRequest) GetTimeMatrix() [][]int {
	ans := make([][]int, len(m.pointNames))
	for i := range m.pointNames {
		ans[i] = make([]int, len(m.pointNames))
		for j := range m.pointNames {
			if i == j {
				ans[i][j] = -1
				continue
			}
			duration := distance.GetDurationSecond(m.pointNames[i], m.pointNames[j])
			ans[i][j] = duration
		}
	}
	return ans
}

func (m *MapRequest) GetPathTime() string {
	distances := m.GetDistances()
	graph := GetGraph(m.GetTimeMatrix())

	ans := 0
	for _, w := range graph {
		ans += getTime(distances, w)
	}
	return strconv.Itoa(ans/60) + " mins"
}

func getTime(distances map[string]int, currentWay string) int {
	mins, _ := distances[currentWay]
	return mins
}
