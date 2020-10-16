package main

import "fmt"

func distant(lat1, long1, lat2, long2 float64) float64 {
	return 0.0
}

func main() {
	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}
