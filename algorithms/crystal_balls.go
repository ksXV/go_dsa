package algorithms

import "math"

func FindTheFirstFloorThatBreaksTheCrystalBalls(arr []bool) int {
	distanceToJump := math.Sqrt(float64(len(arr)))
	i := 0
	found := false
	for !found && i < len(arr) {
		if !arr[i] {
			i += int(distanceToJump)
		} else {
			found = true
		}
	}
	if !found {
		return -1
	}
	for arr[i] == true {
		i--
	}
	return i + 1
}
