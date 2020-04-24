package feature

import (
	"image"
	_ "image/png"
)

const left = 0
const up = 1
const right = 2
const down = 3
const und = -1

type Feature struct
{
	t *featureTree
	m *featureMeta
}

type featureMeta struct
{
	threshold, size, x1, y1, x2, y2 int
}

type featureTree struct
{
	signpost,parent int
	children *[4]*featureTree
}

func Coldiff(x, y uint32) int {
  diff := int(x) - int(y)
  if diff < 0{
    return diff*(-1)
  } else{
    return diff
  }
}

func (feat *Feature) BloomFeatureDiffStrat(x, y int, m *image.Image) {
	//Initialize a feature tree
	feat.m = &featureMeta{
		threshold: 0,
		size: 1,
		x1: 0,
		y1: 0,
		x2: 0,
		y2: 0}
	//SIGNPOST SEARCH NOTE
	//Signpost search is done when signpost points at parent, or for base case
	//if signpost==4 and parent==-1
	feat.t = &featureTree{
		signpost: 0,
		parent: -1,
		children: nil}
	trav := feat
	bounds := m.Bounds()
	r,g,b,_ := m.At(x, y).RGBA()
	var gotOne bool
	//Loop and build the feature tree
	for trav.signpost != 4 || parent != -1{
		gotOne = false
		
		//left
		if x-1 >= bounds.min.X{
			//Do things with the pixels to the left
			//Check the pixel to the left for threshold acceptability
		}
		trav.signpost++

		//up
		if y-1 >= bounds.min.Y{
			//Do things with the pixels to the top

		}
		trav.signpost++

		//right
		if x+1 < bounds.max.X{
			//Do things with the pixels to the right

		}
		trav.signpost++

		//down
		if y+1 < bounds.max.Y{
			//Do things with the pixels to the bottom

		}
		trav.signpost++

		if gotOne == false{
			//Check size, if we aren't there yet fix the threshold
		}
	}
}

func (feat *Feature) BloomFeatureCrawlingStrat(x, y int, m *image.Image) {
	feat.m = &featureMeta{ 0, 0, 0, 0, 0, 0 }
	feat.t = &featureTree{ 0, -1, nil }
}
