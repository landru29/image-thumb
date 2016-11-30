package resizer

import (
	"errors"
	"image"
	"image/jpeg"
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

// ResizePicture resize a picture
func ResizePicture(filename string, size uint) (err error) {
	var source image.Image
	var target image.Image

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	source, _, err = image.Decode(file)
	if err != nil {
		log.Fatal(errors.New("Decode: " + err.Error()))
	}

	file.Seek(0, 0)
	conf, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(errors.New("Config: " + err.Error()))
	}

	folder, err := createThumbFolder(filepath.Dir(filename))
	if err != nil {
		log.Fatal(errors.New("Folder: " + err.Error()))
	}

	if conf.Height > conf.Width {
		target = resize.Resize(0, size, source, resize.Lanczos3)
	} else {
		target = resize.Resize(size, 0, source, resize.Lanczos3)
	}

	outFilename := filepath.Join(folder, filepath.Base(filename))

	out, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(errors.New("Create: " + err.Error()))
	}
	defer out.Close()

	err = jpeg.Encode(out, target, nil)
	if err != nil {
		log.Fatal(errors.New("Encode: " + err.Error()))
	}

	return
}
