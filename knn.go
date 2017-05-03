package main

/*
 * A program that uses two ‘training’ sets of 2D points and an un-weighted k Nearest Neighbors
 * (kNN) algorithm to classify a third set of ‘testing’ points.
 *
 * Authors: Jordan L. Siaha and Connor Runyan
 * Class: CSc 372 Final Project
 * Written in: go language
 */
import (
    "fmt"
    "io/ioutil"
	"strings"
	"strconv"
	"math"
)
const LENGTH int64 = 50 // The length of the grid
const LENGTH2 int64 = 200 // The amount of a and b points on the grid


/* Struct which represents a Point on this grid.
 * Contains X and Y attributes representing position on the grid.
 * A floating point number representing it's distance from the specific test point we're looking at.
 * A string which is it's textual representation on the grid.
 */

type Point struct{
	X, Y int64
	d float64
	c string
}

func main() {
	var grid [LENGTH][LENGTH]Point // Declare a grid of points, this is the main grid at the beginning.
	var i, j int64
	//-------------------READ INPUT---------------------------------\\
	
	// Read in the text files and store contents into variables.
    set1, err := ioutil.ReadFile("set1.txt")
	set2, err := ioutil.ReadFile("set2.txt")
	test, err := ioutil.ReadFile("test.txt") 
    if err != nil {
        fmt.Print(err)
    }
	
	//------------------CONVERT INPUT TO STRING---------------------\\
	
	// Convert the contents to a string
	s1 := string(set1)
	s2 := string(set2)
	t := string(test)
	
	//---------------SPLIT THE STRINGS BY COMMAS--------------------\\
	
	// Split the strings using the comma as a delimeter
    str := strings.Split(s1, ",") // convert content to a 'string'
	str1 := strings.Split(s2, ",") // convert content to a 'string'
	str2 := strings.Split(t, ",") // convert content to a 'string'
	
	var allPoints [LENGTH2]Point
	
	//-------------FILL GRID WITH DEFAULT VALUES---------------------\\
	// Fill the grid with "empty" points
	for  i = 0; i < int64(len(grid)) ; i++ {
      for j = 0; j < int64(len(grid[0])); j++ {
      	grid[i][j] = Point{i, j, -1, ". "}
      }
   }
   
   //---------------FILL GRID WITH POINTS A--------------------------\\
   // Fill the grid with the first set points.
   for i = 0; i < int64(len(str)); i+=2{
    
	y, err := strconv.ParseInt(str[i], 10, 64)
	x, err := strconv.ParseInt(str[i+1], 10, 64)
	
	if err != nil {
        fmt.Print(err)
    }
	grid[LENGTH-x - 1][y] = Point{y, x, -1, "a "}
   }
   
   //---------------FILL GRID WITH POINTS B--------------------------\\
    for i = 0; i < int64(len(str1)); i+=2{
	y, err := strconv.ParseInt(str1[i], 10, 64)
	x, err := strconv.ParseInt(str1[i+1], 10, 64)
	
	if err != nil {
        fmt.Print(err)
    }
	grid[LENGTH-x - 1][y] = Point{y, x, -1, "b "}
   }
   
   //---------------FILL GRID WITH CANDIDATE POINTS------------------\\
   for i = 0; i < int64(len(str2)); i+=2{
	y, err := strconv.ParseInt(str2[i], 10, 64)
	x, err := strconv.ParseInt(str2[i+1], 10, 64)
	
	if err != nil {
        fmt.Print(err)
    }
	fmt.Printf("LENGTH: %d\n X: %d\n", LENGTH - x - 1, x)
	grid[LENGTH-x - 1][y] = Point{y, x, -1, "x "}
   }
   
	//-------------FILL THE ARRAY OF POINTS WITH A AND B-- ---------\\
	var size int64 = 0
	for  i = 0; i < int64(len(grid)) ; i++ {
      for j = 0; j < int64(len(grid[0])); j++ {
		if strings.ContainsAny(grid[i][j].c, "."){
			// Do nothing
		} else if strings.ContainsAny(grid[i][j].c, "x"){
		    // Do nothing
		} else{
			allPoints[size] = grid[i][j]
			size++
		}
      }
   }

	printGrid(grid)
	
	
	//------------------------ALGORITHM------------------------------\\

	// For each canditate point.....
	for i = 0; i < int64(len(str2)); i+=2{ 
	
	 y, err := strconv.ParseInt(str2[i], 10, 64)
	 x, err := strconv.ParseInt(str2[i+1], 10, 64)
	 if err != nil {
        fmt.Print(err)
    }
	 
	 // Go through the array a points and set each point's distance from the candidate point.
	 for j = 0; j < size; j++{
		allPoints[j].d = distance(allPoints[j].X, allPoints[j].Y, y, x)
	 }
	 var n int64
	 
	 // Now sort the array using bubble sort sorting algorithm on the distances.
	var canSwap bool = true
	for canSwap == true {
		canSwap = false
		for n = 0; n < size - 1; n++ {
			// If elements we are checking are out of order, swap them.
			if (allPoints[n].d > allPoints[n + 1].d) {
				var temp Point = allPoints[n]
				allPoints[n] = allPoints[n + 1]
				allPoints[n + 1] = temp
				// We just swapped, so set canSwap variable to true;
				canSwap = true
			}
		}
	}
	
	var Atype int64 = 0
	var Btype int64 = 0

	// Count the A's and B's in the first three positions of the allPoints array.
	if allPoints[0].c == "a "{
	  Atype++
	} else {
	  Btype++
	}
	  
	if allPoints[1].c == "a " {
	  Atype++
	} else { 
	  Btype++
	}
	  
	if allPoints[2].c == "a " {
	  Atype++
	} else { 
	  Btype++
	}

	var index int64 = 3
	var stillPulling bool = true
	fmt.Printf("\n\n\tCandidate at (%d, %d)\n", y, x)
	
	printPoints(allPoints, size)
	
	// Now deal with ties until done.
	for stillPulling {
	  if floatEquals(allPoints[2].d, allPoints[index].d) {
		if allPoints[index].c == "a " {
		  Atype++
		} else { 
		  Btype++
		}
		index++
	   } else {
		stillPulling = false
	   }
	   // at this point Atype and Btype are correct
	}
	
	// Print out the information and move the candicate point to the appropriate set(Visually).
	fmt.Printf("\t\tATYPE: %d\tBTYPE: %d", Atype, Btype)
	if Atype > Btype {
	  grid[LENGTH-x-1][y].c = "A "
	  fmt.Printf("\tResult: A\n\n")
	} else if Btype > Atype{
	  grid[LENGTH-x-1][y].c = "B "
	  fmt.Printf("\tResult: B\n\n")
	} else {
	  grid[LENGTH-x-1][y].c = "X "
	  fmt.Printf("\tResult: Tie!\n\n")
	}
	printGrid(grid)
	}
	
}

    // float compare
	func floatEquals(a, b float64) bool {
	  if((a - b) < 0.000001 && (b - a) < 0.000001) {
	    return true
	  } else {
	  return false
	  }
	}


	//-------------CALCULATE DISTANCE BETWEEN TWO POINTS--------------\\
	func distance(x1, y1, x2, y2 int64) float64{
		first := math.Pow(float64(x2-x1),2)
		second := math.Pow(float64(y2-y1),2)
		return math.Sqrt(first + second)
	}
	
	//---------------------PRINT THE GRID------------------------------\\
	func printGrid(grid[LENGTH][LENGTH]Point){
		var i, j int64
		for  i = 0; i < LENGTH ; i++ {
		  for j = 0; j < LENGTH; j++ {
			fmt.Printf(grid[i][j].c)
		  }
		  fmt.Printf("\n")
	   }
	}
	//---------------------PRINT THE POINTS-----------------------------\\
	func printPoints(grid[LENGTH2]Point, size int64){
		var i int64
		for  i = 0; i < size ; i++ {
			fmt.Printf("\t\tDistance of %s: %f  ", grid[i].c, grid[i].d)
			fmt.Printf("X and Y (%d, %d)\n", grid[i].X, grid[i].Y)
	   }
	}