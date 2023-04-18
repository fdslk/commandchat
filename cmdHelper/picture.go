package cmdHelper

import (
	"fmt"
	"image"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/nsf/termbox-go"
)

func draw(img image.Image) {
	width, height, whratio := canvasSize()

	bounds := img.Bounds()
	imgW, imgH := bounds.Dx(), bounds.Dy()

	imgScale := scale(imgW, imgH, width, height, whratio)

	width, height = int(float64(imgW)/imgScale), int(float64(imgH)/(imgScale*whratio))

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			startX, startY, endX, endY := imgArea(x, y, imgScale, whratio)

			r, g, b := avgRGB(img, startX, startY, endX, (startY+endY)/2)
			colorUp := termbox.Attribute(termColor(r, g, b))

			r, g, b = avgRGB(img, startX, (startY+endY)/2, endX, endY)
			colorDown := termbox.Attribute(termColor(r, g, b))

			termbox.SetCell(x, y, '▄', colorDown, colorUp)
		}
	}
	termbox.Flush()
}

func Display(imagePath string) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	img, err := LoadImageFile(imagePath)
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)

	fmt.Println("Close the image with <ESC> or by pressing 'q'.")

	draw(img)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Ch == 'q' {
				return
			}
		case termbox.EventResize:
			draw(img)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
