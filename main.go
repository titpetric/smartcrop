package main

import (
	"encoding/json"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/artyom/smartcrop"
)

type CropSize struct {
	x, y int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing parameter: image file")
	}
	filename := os.Args[1]

	fi, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(fi)
	if err != nil {
		log.Fatal(err)
	}

	cropsizes := []CropSize{
		CropSize{1, 1},
		CropSize{4, 3},
		CropSize{3, 4},
		CropSize{16, 9},
	}

	results := map[string]image.Rectangle{}

	for _, cropsize := range cropsizes {
		crop, err := smartcrop.Crop(img, cropsize.x, cropsize.y)
		if err != nil {
			log.Fatal(err)
		}
		results[fmt.Sprintf("%d:%d", cropsize.x, cropsize.y)] = crop
	}

	result, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", result)
}
