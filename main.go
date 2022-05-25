package main

import (
	"fmt"
	"math"
)

func main() {
	//Rendering Doughnut(Torus) in Go-Lang
	screen_width := 500.0                                                                        // The variable storing the width of the screen
	screen_height := 500.0                                                                       // The variable storing the height of the screen
	minor_radius := 3                                                                            //minor_radius is the radius of the circle being revolved.
	major_radius := 4                                                                            //major_radius is the radius of the revolution.
	k2 := 5                                                                                      //k2 is the distance of the doughnut from the viewer
	k1 := screen_width * float64(k2) * 3 / (float64(8) * float64((minor_radius + major_radius))) //the screen distance z’.(plane away from the viewer)
	rotate_x := 0                                                                                //rotation on x axis
	rotate_z := 0                                                                                //roatation on z axis
	sin_rx, cos_rx := math.Sincos(float64(rotate_x))                                             //get the sin and cos value of rotation on x axis
	sin_rz, cos_rz := math.Sincos(float64(rotate_z))                                             //get the sin and cos value of rotation on z axis
	//Dummy Prints to let main.go compile
	fmt.Println("screen_width:", screen_width)
	fmt.Println("screen_height:", screen_height)
	fmt.Println("minor_radius:", minor_radius)
	fmt.Println("major_radius:", major_radius)
	fmt.Println("rotate_x:", rotate_x)
	fmt.Println("rotate_z:", rotate_z)
	fmt.Println("sin_rx:", sin_rx)
	fmt.Println("cos_rx:", cos_rx)
	fmt.Println("sin_rz:", sin_rz)
	fmt.Println("cos_rz:", cos_rz)
	fmt.Println("k1:", k1)
	fmt.Println("k2:", k2)

}
