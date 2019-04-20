package main

import (
	"fmt"
	"os"
	"image/png"
	"github.com/skreimeyer/PlanReview/pkg/planreview"
)

func main() {
	addr := "500 W Markham"
	loc := planreview.Geocode(addr)
	fmt.Println("Location:",loc)
	ring := planreview.FetchParcel(loc)
	fmt.Println("Ring found:",ring)
	x1,y1,x2,y2 := planreview.MakeEnvelope(ring, 0.1)
	fmt.Println("Envelope:",x1,y1,x2,y2)
	img := planreview.FetchMap(x1,y1,x2,y2)
	fmt.Println("Got image")
	f,err := os.Create("map.png")
	if err != nil {
		panic(err)
	}
	if err :=  png.Encode(f,img); err != nil {
		f.Close()
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
}

