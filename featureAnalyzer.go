package main

import (
  "math/rand"
	"fmt"
	"image"
	"log"
  "time"
  "os"
	_ "image/png"
  "github.com/rufuskd/featureAnalyzer/feature"
)



type point struct{
  x,y int
}


func main() {
  var fuck feature.Feature
  fmt.Println("I'm losing my mind",fuck)
  rand.Seed(time.Now().UnixNano())
  reader, err := os.Open("/home/rufus/featureAnalyzer/0ltqoml2xrt41.png")
  if err != nil {
    log.Fatal(err)
  }
  defer reader.Close()

  m, _, err := image.Decode(reader)
  bounds := m.Bounds()

  pixelCount := 0
  avDiff := 0
  diffSum := 0
  denom := 0
  start := time.Now()
  var r,g,b,tr,tg,tb uint32
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ = m.At(x, y).RGBA()
      diffSum = 0
      denom = 0
      for i1 := -1; i1 <= 1; i1++ {
        for i2 := -1; i2 <= 1; i2++ {
          if x+i2 >= bounds.Min.X && x+i2 < bounds.Max.X && y+i1 >= bounds.Min.Y && y+i1 < bounds.Max.Y{
            tr, tg, tb, _ = m.At(x+i2, y+i1).RGBA()
            diffSum += coldiff(r,tr) + coldiff(g,tg) + coldiff(b,tb)
            denom++
          }
        }
      }
      avDiff += diffSum/denom
      pixelCount++
		}
	}
  avDiff = avDiff/pixelCount
  t := time.Now()

  fmt.Println("Pixel count:", pixelCount)
  fmt.Println("Average pixel diff:", avDiff)
  fmt.Println("Time: ",t.Sub(start))

  startingSpotCount := 0
  noStarts := make(map[point]int)
  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ = m.At(x, y).RGBA()
      diffSum = 0
      denom = 0
      for i1 := -1; i1 <= 1; i1++ {
        for i2 := -1; i2 <= 1; i2++ {
          if x+i2 >= bounds.Min.X && x+i2 < bounds.Max.X && y+i1 >= bounds.Min.Y && y+i1 < bounds.Max.Y{
            tr, tg, tb, _ = m.At(x+i2, y+i1).RGBA()
            diffSum += coldiff(r,tr) + coldiff(g,tg) + coldiff(b,tb)
            denom++
          }
        }
      }
      if diffSum < avDiff{
        startingSpotCount++
      } else {
        noStarts[point{x,y}] = 1
      }
		}
	}
  fmt.Println("There are:",startingSpotCount," possible starting pixels")
  rx := rand.Int()%bounds.Max.X
  ry := rand.Int()%bounds.Max.Y
  fmt.Println("Can I start at:",rx,ry," ",noStarts[point{rx,ry}]==0)

  //Now that we have established the average pixel difference and mapped out
  //the no start zones, we iterate pixel by pixel, bloom, add newly incorporated
  //pixels to the no start zones and let the features rip
  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
      if noStarts[point{x,y}] == 0{
        //Begin a feature bloom here

      }
    }
  }
}
