package timeHelper

import "math"

func RoundTime(input float64) int {
	var result float64
	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}
	i, _ := math.Modf(result)
	return int(i)
}
