package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum uint) float64 {
	const SidesTriangle uint = 3
	const SidesSquare uint = 4
	const SidesCircle uint = 0

	if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else if sidesNum == SidesTriangle {
		return sideLen * sideLen * math.Sqrt(3) / 4
	} else if sidesNum == SidesCircle {
		return sideLen * sideLen * math.Pi
	} else {
		return 0
	}
}
