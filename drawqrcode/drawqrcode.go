package drawqrcode

import (
	qrcode "github.com/skip2/go-qrcode"
	"image"
	"github.com/disintegration/imaging"
)

func draw(text string)(*image.RGBA){
	var p []byte
	p, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err!=nil{
		panic(err)
	}


	img:=NewImageFromBytes(p)

	border:=img.Bounds().Dy()/6
	nrgba:=imaging.Crop(img,image.Rect(border,border,img.Bounds().Dx()-border,img.Bounds().Dy()-border))

	rgb:=NewFromImage(nrgba)

	return rgb
}

func Get(bgimg []byte,text string,x,y int,w,h int) *image.RGBA {
	cov:=draw(text)
	b:=NewImageFromBytes(bgimg)
	bg:=NewFromImage(b)
	result:=CoverImg(bg,cov,image.Rectangle{image.Point{x,y},image.Point{x+w,y+h}})
	return result
}