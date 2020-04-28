package main

import (
  "container/list"
	"fmt"
	"image"
	"log"
  "os"
	"image/jpeg"
  "image/color"
  "github.com/rufuskd/featureAnalyzer/feature"
)


func main() {
  files := os.Args[1:]
  if len(files) < 1{
    fmt.Println("Usage: ./featureAnalyzer <filename> <possible other filename> <etc>...")
    return
  }
  for _,filename := range(files){
    reader, err := os.Open(filename)
    if err != nil {
      log.Fatal(err)
    }

    m, _, err := image.Decode(reader)
    bounds := m.Bounds()

    avDiff := feature.AvDiff(m)
    fmt.Println("Average pixel diff:", avDiff)

    avMinDiff := feature.AvMinDiff(m)
    fmt.Println("Average minimum pixel diff:", avMinDiff)

    avTransDiff := feature.AvTransitionDiff(m)
    fmt.Println("Average transition pixel diff:", avTransDiff)

/*
    startingSpotCount := 0
    noStarts := make(map[feature.Point]int)
    var r,g,b,tr,tg,tb uint32
    denom := 0
    diffSum := 0

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
  		for x := bounds.Min.X; x < bounds.Max.X; x++ {
  			r, g, b, _ = m.At(x, y).RGBA()
        diffSum = 0
        denom = 0
        for i1 := -1; i1 <= 1; i1++ {
          for i2 := -1; i2 <= 1; i2++ {
            if x+i2 >= bounds.Min.X && x+i2 < bounds.Max.X && y+i1 >= bounds.Min.Y && y+i1 < bounds.Max.Y{
              tr, tg, tb, _ = m.At(x+i2, y+i1).RGBA()
              diffSum += feature.Coldiff(r,tr) + feature.Coldiff(g,tg) + feature.Coldiff(b,tb)
              denom++
            }
          }
        }
        if diffSum < avDiff{
          startingSpotCount++
        } else {
          noStarts[feature.Point{x,y}] = 1
        }
  		}
  	}
    */

    //Now that we have established the average pixel difference and mapped out
    //the no start zones, we iterate pixel by pixel, bloom, add newly incorporated
    //pixels to the no start zones and let the features rip
    visited := make(map[feature.Point]int)
    featList := list.New()
    featureCount := 0
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
  		for x := bounds.Min.X; x < bounds.Max.X; x++ {
        if visited[feature.Point{x,y}] == 0{
          //Begin a feature bloom here
          newFeature := &(feature.Feature{nil,nil})
          newFeature.BloomFeatureDiffStrat(x,y,avDiff*5,m,&visited)
          featList.PushBack(newFeature)
          featureCount++
        }
      }
    }

    resultImage := image.NewRGBA(bounds)

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
      for x := bounds.Min.X; x < bounds.Max.X; x++ {
        r,g,b,_ := m.At(x,y).RGBA()
        if visited[feature.Point{x,y}] == 1{
          resultImage.SetRGBA(x,y,color.RGBA{uint8(r>>8),uint8(g>>8),uint8(b>>8),255})
        }
      }
    }

    for item := featList.Front(); item != nil; item = item.Next(){
      f := item.Value.(*feature.Feature)
      if f.M.Size > 0{
        //Draw the feature box on the image!
        for bx := f.M.X1; bx <= f.M.X2; bx++{
          resultImage.Set(bx,f.M.Y1,color.RGBA{255,255,255,255})
          resultImage.Set(bx,f.M.Y2,color.RGBA{255,255,255,255})
        }
        for by := f.M.Y1; by <= f.M.Y2; by++{
          resultImage.Set(f.M.X1,by,color.RGBA{255,255,255,255})
          resultImage.Set(f.M.X2,by,color.RGBA{255,255,255,255})
        }
      }
    }
    output, err := os.Create("Output.jpg")
    if err != nil{
      panic(err)
    }

    jpeg.Encode(output,resultImage,nil)

    fmt.Println("Pulled: ",featureCount," features")

    reader.Close()
    output.Close()
  }
}
