package core

import "math"

func RadiansToDegrees(radians float64) float64 {
	return math.Mod(radians, 2*math.Pi) * 180 / math.Pi
}

func DegreesToRadians(degrees float64) float64 {
	return math.Mod(degrees, 360) * math.Pi / 180
}
