package main

import (
	"fmt"
	"math"
)

func render_doughnut(rotate_x, rotate_z float64) {
	for {
		rotate_x += 0.07
		rotate_z += 0.03
		var sin_rx, cos_rx float64 = math.Sincos(rotate_x) //get the sin and cos value of rotation on x axis
		var sin_rz, cos_rz float64 = math.Sincos(rotate_z) //get the sin and cos value of rotation on z axis
		fmt.Println("sin_rx:", sin_rx)
		fmt.Println("cos_rx:", cos_rx)
		fmt.Println("sin_rz:", sin_rz)
		fmt.Println("cos_rz:", cos_rz)

	}
}

func main() {
	//Rendering Doughnut(Torus) in Go-Lang
	var screen_width, screen_height float64 = 500.0, 500.0 // The variable storing the width and height of the screen
	// theta_spacing := 0.07
	// phi_spacing := 0.02
	var minor_radius, major_radius int = 3, 4                                                    //minor_radius is the radius of the circle being revolved, and major_radius is the radius of revolution                                                                         //major_radius is the radius of the revolution.
	k2 := 5                                                                                      //k2 is the distance of the doughnut from the viewer
	k1 := screen_width * float64(k2) * 3 / (float64(8) * float64((minor_radius + major_radius))) //the screen distance z’.(plane away from the viewer)
	var rotate_x, rotate_z float64 = 0.0, 0.0                                                    //rotation on x and z axis

	render_doughnut(rotate_x, rotate_z)
	//Dummy Prints to let main.go compile
	fmt.Println("screen_width:", screen_width)
	fmt.Println("screen_height:", screen_height)
	fmt.Println("minor_radius:", minor_radius)
	fmt.Println("major_radius:", major_radius)
	fmt.Println("k1:", k1)
	fmt.Println("k2:", k2)

}
