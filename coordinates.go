package faker

import (
	"strconv"
)

func Latitude() string {
	return string(strconv.AppendFloat([]byte{}, float64(latitude()), 'f', 6, 32))
}

func Longitude() string {
	return string(strconv.AppendFloat([]byte{}, float64(longitude()), 'f', 6, 32))
}

func latitude() float32 {
	return Float32Between(-90, 90)
}

func longitude() float32 {
	return Float32Between(-180, 180)
}

func Coordinates() string {
	return Latitude() + " " + Longitude()
}
