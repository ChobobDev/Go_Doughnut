package main

// https://github.com/Vlad-Shevliakov/ASCII-render/blob/master/donut.go 참조

import (
	"fmt"
	"math"
	"time"
)

const (
	delay      = 40 * time.Millisecond // 60 fps hahaha
	coreString = ".,-~:;=!*#$@"
	pi         = math.Pi
)

type sliceType interface {
	len() int
}

func floatMemset(arr []float64, v float64) {
	for i := range arr {
		arr[i] = v
	}
}

func byteMemset(arr []string, v string) {
	for i := range arr {
		arr[i] = v
	}
}

func main() {
	var rotate_x, rotate_z float64 = 0.0, 0.0 //rotation on x and z axis
	z := make([]float64, 1760)
	b := make([]string, 1760)
	for {
		byteMemset(b, " ")
		floatMemset(z, 0)
		for theta := 0.0; theta < 2*pi; theta += 0.07 {
			for phi := 0.0; phi < 2*pi; phi += 0.02 {
				var sin_theta, cos_theta float64 = math.Sincos(theta)
				var sin_phi, cos_phi float64 = math.Sincos(phi)
				var sin_rx, cos_rx float64 = math.Sincos(rotate_x) //get the sin and cos value of rotation on x axis
				var sin_rz, cos_rz float64 = math.Sincos(rotate_z) //get the sin and cos value of rotation on z axis
				var denominator float64 = 1 / (sin_phi*sin_rx*(cos_theta+2) + sin_theta*cos_rx + 5)
				t := sin_phi*(cos_theta+2.0)*cos_rx - sin_theta*sin_rx //this is a clever factoring of some of the terms in x' and y'
				var x int = int(40 + 30*denominator*(cos_phi*(cos_theta+2)*cos_rz-t*sin_rz))
				var y int = int(12 + 12*denominator*(cos_phi*(cos_theta+2)*sin_rz-t*cos_rz))
				var o int = x + 80*y
				var N int = int(8 * ((sin_theta*sin_rx-sin_phi*cos_theta*cos_rx)*cos_rz - sin_phi*cos_theta*sin_rx - sin_theta*cos_rx - cos_phi*cos_theta*sin_rz))
				if y < 22 && y > 0 && x > 0 && x < 80 && denominator > z[o] {
					z[o] = denominator
					point := 0
					if N > 0 {
						point = N
					}

					b[o] = string(coreString[point])
				}
			}
		}
		print("\x1b[H")

		for k := 0; k < 1761; k++ {

			v := "\n"

			if k%80 > 0 {
				v = string(b[k])
			}

			fmt.Printf(v)

			rotate_x += 0.00004
			rotate_z += 0.00002
		}

		time.Sleep(delay)

	}
}
