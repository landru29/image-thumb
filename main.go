package main

import (
	"github.com/landru29/image-thumb/resizer"
)

func main() {

	resizer.ResizePicture("./test/giphy.gif", 150)
	resizer.ResizePicture("./test/jpeg-home.jpg", 150)
	//resizer.ResizePicture("./test/water_PNG3290.png", 150)
}
