package main

import (
	"fmt"
	"os"
)

const image_width int = 256
const image_height int = 256

func main() {
	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)
	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", j)
		for i := 0; i < image_width; i++ {
			r := float64(i) / float64(image_width-1)
			g := float64(j) / float64(image_height-1)
			b := 0.25

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
	fmt.Fprint(os.Stderr, "\nDone\n")

}
