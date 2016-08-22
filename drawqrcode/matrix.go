package drawqrcode

import (
	"image"
	"github.com/disintegration/imaging"
	"image/jpeg"
	"image/png"
	"image/gif"
	"bytes"
	"bufio"
)

func NewImageFromBytes(bs []byte) image.Image {
	r := bytes.NewReader(bs)

	img, err := jpeg.Decode(r)
	if err == nil {
		return img
	}

	r = bytes.NewReader(bs)
	img, err = png.Decode(r)
	if err == nil {
		return img
	}

	r = bytes.NewReader(bs)
	img, err = gif.Decode(r)
	if err == nil {
		return img
	}
	return nil
}

func NewFromImage(img image.Image) *image.RGBA {

	rect := img.Bounds()

	this := image.NewRGBA(rect)

	for y := 0; y < rect.Dy(); y++ {
		for x := 0; x < rect.Dx(); x++ {
			this.Set(x, y, img.At(x, y))
		}
	}
	return this
}
func ImageSave(img *image.RGBA)(*bytes.Buffer){

	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriterSize(b, img.Stride)
	err := png.Encode(w,img)
	if err != nil {
		panic(err)
	}
	return b
}
func  CoverImg(this *image.RGBA,coverimg *image.RGBA,rect image.Rectangle)(*image.RGBA){
	dstImage := imaging.Resize(coverimg,rect.Dx(),rect.Dy(), imaging.Lanczos)

	for y:=0;y<dstImage.Rect.Dy();y++{
		for x:=0;x<dstImage.Rect.Dx();x++{
			this.Set(rect.Min.X+x,rect.Min.Y+y,dstImage.At(x,y))
		}
	}

	return this
}