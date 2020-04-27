package feature

import (
	"fmt"
	"image"
	_ "image/jpeg"
)

const left = 0
const up = 1
const right = 2
const down = 3
const und = -1

type Point struct{
  X,Y int
}

type Feature struct
{
	T *featureTree
	M *featureMeta
}

type featureMeta struct
{
	size, avDiff, X1, Y1, X2, Y2 int
}

type featureTree struct
{
	signpost,parent,x,y int
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

func (feat *Feature) PrintBox(){
	fmt.Println("Size:",feat.M.size,"\nCoords:",feat.M.X1,",",feat.M.Y1,"-",feat.M.X2,",",feat.M.Y2)
}

func (feat *Feature) BloomFeatureDiffStrat(x, y, threshold int, m image.Image, visited *map[Point]int) {
	//Initialize a feature tree
	feat.M = &featureMeta{
		size: 1,
		avDiff: 0,
		X1: x,
		Y1: y,
		X2: x,
		Y2: y}
	//SIGNPOST SEARCH NOTE
	//Signpost search is done when signpost points at parent, or for base case
	//if signpost==4 and parent==-1
	feat.T = &featureTree{
		signpost: 0,
		parent: -1,
		x: x,
		y: y,
		children: &[4]*featureTree{nil,nil,nil,nil}}
	(*visited)[Point{x,y}] = 1
	trav := feat.T
	bounds := m.Bounds()

	var r,g,b,r1,g1,b1 uint32
	var curdiff int
	//Loop and build the feature tree
	for trav.signpost < 4 || trav.parent != -1{
		if trav.x < feat.M.X1{
			feat.M.X1 = trav.x
		}
		if trav.x > feat.M.X2{
			feat.M.X2 = trav.x
		}
		if trav.y < feat.M.Y1{
			feat.M.Y1 = trav.y
		}
		if trav.y > feat.M.Y2{
			feat.M.Y2 = trav.y
		}
		if trav.signpost%4 == trav.parent{
			trav.signpost++
			trav = trav.children[trav.parent]
		} else if trav.signpost%4 == left && trav.x-1 >= bounds.Min.X && (*visited)[Point{trav.x-1,trav.y}] == 0{
			//Do things with the pixels to the left
			//Check the pixel to the left for threshold acceptability
			r,g,b,_ = m.At(trav.x, trav.y).RGBA()
			r1,g1,b1,_ = m.At(trav.x-1,trav.y).RGBA()
			curdiff = Coldiff(r,r1) + Coldiff(g,g1) + Coldiff(b,b1)

			if curdiff < threshold{
				(*visited)[Point{trav.x-1,trav.y}] = 1
				trav.children[left] = &featureTree{
					signpost: right+1,
					parent: right,
					x: trav.x-1,
					y: trav.y,
					children: &[4]*featureTree{nil,nil,nil,nil}}
				trav.signpost++
				trav.children[left].children[right] = trav
				trav = trav.children[left]
				feat.M.size++
			} else {
				trav.signpost++
			}
		} else if trav.signpost%4 == up && trav.y-1 >= bounds.Min.Y && (*visited)[Point{trav.x,trav.y-1}] == 0{
			//Do things with the pixels to the top
			//Check the pixel to the left for threshold acceptability
			r,g,b,_ = m.At(trav.x, trav.y).RGBA()
			r1,g1,b1,_ = m.At(trav.x,trav.y-1).RGBA()
			curdiff = Coldiff(r,r1) + Coldiff(g,g1) + Coldiff(b,b1)

			if curdiff < threshold{
				(*visited)[Point{trav.x,trav.y-1}] = 1
				trav.children[up] = &featureTree{
					signpost: down+1,
					parent: down,
					x: trav.x,
					y: trav.y-1,
					children: &[4]*featureTree{nil,nil,nil,nil}}
				trav.signpost++
				trav.children[up].children[down] = trav
				trav = trav.children[up]
				feat.M.size++
			} else {
				trav.signpost++
			}
		} else if trav.signpost%4 == right && trav.x+1 < bounds.Max.X && (*visited)[Point{trav.x+1,trav.y}] == 0{
			//Do things with the pixels to the right
			//Check the pixel to the left for threshold acceptability
			r,g,b,_ = m.At(trav.x, trav.y).RGBA()
			r1,g1,b1,_ = m.At(trav.x+1,trav.y).RGBA()
			curdiff = Coldiff(r,r1) + Coldiff(g,g1) + Coldiff(b,b1)

			if curdiff < threshold{
				(*visited)[Point{trav.x+1,trav.y}] = 1
				trav.children[right] = &featureTree{
					signpost: left+1,
					parent: left,
					x: trav.x+1,
					y: trav.y,
					children: &[4]*featureTree{nil,nil,nil,nil}}
				trav.signpost++
				trav.children[right].children[left] = trav
				trav = trav.children[right]
				feat.M.size++
			} else {
				trav.signpost++
			}
		} else if trav.signpost%4 == down && trav.y+1 < bounds.Max.Y && (*visited)[Point{trav.x,trav.y+1}] == 0{
			//Do things with the pixels to the bottom
			//Check the pixel to the left for threshold acceptability
			r,g,b,_ = m.At(trav.x, trav.y).RGBA()
			r1,g1,b1,_ = m.At(trav.x,trav.y+1).RGBA()
			curdiff = Coldiff(r,r1) + Coldiff(g,g1) + Coldiff(b,b1)

			if curdiff < threshold{
				(*visited)[Point{trav.x,trav.y+1}] = 1
				trav.children[down] = &featureTree{
					signpost: up+1,
					parent: up,
					x: trav.x,
					y: trav.y+1,
					children: &[4]*featureTree{nil,nil,nil,nil}}
				trav.signpost++
				trav.children[down].children[up] = trav
				trav = trav.children[down]
				feat.M.size++
			} else {
				trav.signpost++
			}
		} else {
			trav.signpost++
		}
	}
}
