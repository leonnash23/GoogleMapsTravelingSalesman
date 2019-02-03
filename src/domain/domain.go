package domain

import "strconv"

type MarkerList struct {
	markers []Marker
}

type Marker struct {
	Label       string
	AddressName string
}

func (m *MarkerList) AddMarker(addressName string) {
	marker := Marker{}
	marker.Label = strconv.Itoa(len(m.markers))
	marker.AddressName = addressName
	m.markers = append(m.markers, marker)
}

func (m *MarkerList) GetMarkers() []Marker {
	cp := make([]Marker, m.Count())
	copy(cp, m.markers)
	return cp
}

func (m *MarkerList) GetAddressList() []string {
	addresses := make([]string, m.Count())
	for i := range m.markers {
		addresses[i] = m.markers[i].AddressName
	}
	return addresses
}

func (m *MarkerList) Count() int {
	return len(m.markers)
}
