package main

import "fmt"

func main() {
	//Rendering Doughnut(Torus) in Go-Lang
	minor_radius := 3 //minor_radius is the radius of the circle being revolved.
	major_radius := 4 //major_radius is the radius of the revolution.
	k1 := 1.0         //the screen distance z’.(plane away from the viewer)
	k2 := 5.0         //k2 is the distance of the doughnut from the viewer

	fmt.Println(minor_radius, major_radius, k1, k2)

}
