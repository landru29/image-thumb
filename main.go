package main

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func createThumbFolder(path string) (folder string, err error) {
	folder = filepath.Join(path, "thumb")
	info, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	if _, errFound := os.Stat(folder); os.IsNotExist(errFound) {
		err = os.Mkdir(folder, info.Mode())
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

func getImageFormat(filename string) string {
	format := strings.ToLower(filepath.Ext(filename))
	return format[1:len(format)]
}

func resizeMe(filename string, size uint) (err error) {
	var source image.Image
	var target image.Image
	var conf image.Config

	fmt.Printf("Open %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	folder, err := createThumbFolder(filepath.Dir(filename))
	if err != nil {
		log.Fatal(err)
	}

	format := getImageFormat(filename)

	switch format {
	case ".jpg", ".jpeg":
		source, err = jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		conf, err = jpeg.DecodeConfig(file)
		if err != nil {
			log.Fatal(err)
		}
	case ".png":
		source, err = png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		conf, err = png.DecodeConfig(file)
		if err != nil {
			log.Fatal(err)
		}
	case ".gif":
		source, err = gif.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		conf, err = gif.DecodeConfig(file)
		if err != nil {
			log.Fatal(err)
		}
	default:
		err = errors.New("Not a compatible picture")
	}

	fmt.Printf("Resizing ...\n")
	if conf.Height > conf.Width {
		target = resize.Resize(0, size, source, resize.Lanczos3)
	} else {
		target = resize.Resize(size, 0, source, resize.Lanczos3)
	}

	outFilename := filepath.Join(folder, filepath.Base(filename))
	fmt.Printf("Writing %s\n", outFilename)

	out, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	switch format {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(out, target, nil)
		source, err = jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	case ".png":
		err = jpeg.Encode(out, target, nil)
		if err != nil {
			log.Fatal(err)
		}
	case ".gif":
		err = jpeg.Encode(out, target, nil)
		if err != nil {
			log.Fatal(err)
		}
	default:
		err = errors.New("Not a compatible picture")
	}

	return
}

func main() {

	//resizeMe("test/giphy.gif", 150)
	resizeMe("test/jpeg-home.jpg", 150)
	resizeMe("test/water_PNG3290.png", 150)

}
