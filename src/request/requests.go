package request

import (
	"MapApi/src/domain"
	"MapApi/src/request/distance"
	"MapApi/src/request/url"
	"strconv"
)

type MapRequest struct {
	Markers domain.MarkerList
}

func (m *MapRequest) AddMarker(pointName string) *MapRequest {
	m.Markers.AddMarker(pointName)
	return m
}

func (m *MapRequest) GetDistances() map[string]int {
	var ans = make(map[string]int)
	addressList := m.Markers.GetAddressList()
	for i := range addressList {
		for j := range addressList {
			if i == j {
				continue
			}
			ans[strconv.Itoa(i)+"->"+strconv.Itoa(j)] = distance.GetDurationSecond(addressList[i], addressList[j])
		}
	}
	return ans
}

func (m *MapRequest) GetTimeMatrix() [][]int {
	ans := make([][]int, m.Markers.Count())
	addressList := m.Markers.GetAddressList()
	for i := range addressList {
		ans[i] = make([]int, m.Markers.Count())
		for j := range addressList {
			if i == j {
				ans[i][j] = -1
				continue
			}
			duration := distance.GetDurationSecond(addressList[i], addressList[j])
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

func (m *MapRequest) GetLinkMapUrl() string {
	return url.GetLinkMapUrl(m.Markers, GetGraph(m.GetTimeMatrix()))
}

func getTime(distances map[string]int, currentWay string) int {
	mins, _ := distances[currentWay]
	return mins
}
