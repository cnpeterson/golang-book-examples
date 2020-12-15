package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
    "strconv"
)

// var palette = []color.Color{color.Black,
//     color.RGBA{0x00, 0x80, 0x00, 0xff}, // Green
//     color.RGBA{0xff, 0xd7, 0x00, 0xff}, // Gold
//     color.RGBA{0x00, 0x00, 0x80, 0xff}, // Navy
//     color.RGBA{0x4b, 0x00, 0x82, 0xff}, // Indigo
//     color.RGBA{0x40, 0xe0, 0xd0, 0xff}, // Turquoise
//     // color.White,
// }

var palette = []color.Color{color.Black}
var cycles int = 5   // number of complete x oscillator revolutions
var res float64 = 0.001  // angular resolution
var size int = 100   // image canvas covers [-size..+size]
var nframes int = 64 // number of animation frames
var delay int = 8    // delay between frames in 10ms units

const (
    blackIndex = 0
    greenIndex = 1
    whiteIndex = 2
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        lissajous(w, r)
    })
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {
    cparam := r.URL.Query().Get("cycles")
    if cparam != "" {
        sv, err := strconv.Atoi(cparam)
        if err == nil {
            cycles = sv
        }
    }

    for i := 0; i < nframes; i++ {
        palette = append(palette, color.RGBA{0x00, uint8(i), 0x00, 0xff}) // Fades out green
    }

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < float64(cycles)*2.0*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(i))
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
