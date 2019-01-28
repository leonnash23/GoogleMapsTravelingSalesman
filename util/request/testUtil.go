package request

func BuildMapRequest4() *MapRequest {
	mapRequest := MapRequest{}
	mapRequest.AddMarker("5 просека 99Б")
	mapRequest.AddMarker("Ново-садовая 333")
	mapRequest.AddMarker("Гастелло 43а")
	mapRequest.AddMarker("Дыбенко 30")
	return &mapRequest
}

func BuildMapRequest3() *MapRequest {
	mapRequest := MapRequest{}
	mapRequest.AddMarker("5 просека 99Б")
	mapRequest.AddMarker("Ново-садовая 333")
	mapRequest.AddMarker("Гастелло 43а")
	return &mapRequest
}
