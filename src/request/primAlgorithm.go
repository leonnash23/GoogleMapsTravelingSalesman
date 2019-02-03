package request

import (
	set "github.com/deckarep/golang-set"
	"math"
	"strconv"
)

var ENF = math.MaxInt32

func GetGraph(timeMatrix [][]int) []string {
	var ans []string
	graph := set.NewSet()
	graph.Add(0)

	minWayIndexY := -1
	minWayValue := ENF

	dotsCount := len(timeMatrix)
	for graphLen := 1; graphLen < dotsCount; {
		element := -1
		for j := 0; j < graphLen; j++ {
			element = j
			for i := 0; i < dotsCount; i++ {
				if i == element {
					continue
				}
				durationFromXtoY := getDurationFromXtoY(timeMatrix, element, i)
				if minWayValue > durationFromXtoY && !graph.Contains(i) {
					minWayIndexY = i
					minWayValue = durationFromXtoY
				}
			}
		}
		if !graph.Contains(minWayIndexY) {
			graphLen += 1
			minWayValue = ENF
		}
		graph.Add(minWayIndexY)

		ans = append(ans, formatPath(element, minWayIndexY))
	}

	return ans
}

func getDurationFromXtoY(matrix [][]int, element int, i int) int {
	return matrix[element][i]
}

func formatPath(first int, second int) string {
	return strconv.Itoa(first) + "->" + strconv.Itoa(second)
}
