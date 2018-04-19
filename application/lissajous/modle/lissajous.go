// Lissajous generates GIF animations of random Lissajous figures.
package modle

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

const (
	whiteIndex = iota
	blackIndex
)

var palette = []color.Color{color.White, color.RGBA{88, 33, 44, 4}} // []color.Color{color.White, color.Black}

func Lissajous(out io.Writer) {
	const (
		cycles  = 5     // 振荡器转速 rad/s
		res     = 0.001 // 角分辨率
		size    = 100   // 画板size
		nframes = 64    // 动画帧数
		delay   = 8     // 帧数之间的延时，以 10ms 为基数
	)

	freq := rand.Float64() * 3.0 // y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference（相位？）
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误
}
