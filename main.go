package main

import (
	"C"
	"fmt"

	"github.com/gookit/color"
)
import (
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type colorRGB struct {
	r uint8
	g uint8
	b uint8
}

var (
	firePixelsArray   []int
	fireWidth         int
	fireHeight        int
	debug             bool
	fireColorsPalette [37]colorRGB
)

const (
	minFire    int  = 0
	maxFire    int  = 36
	printIndex bool = true
	printValue bool = false
)

func initiateFireColorsPalette() {
	fireColorsPalette[0] = colorRGB{7, 7, 7}
	fireColorsPalette[1] = colorRGB{31, 7, 7}
	fireColorsPalette[2] = colorRGB{47, 15, 7}
	fireColorsPalette[3] = colorRGB{71, 15, 7}
	fireColorsPalette[5] = colorRGB{103, 31, 7}
	fireColorsPalette[4] = colorRGB{87, 23, 7}
	fireColorsPalette[6] = colorRGB{119, 31, 7}
	fireColorsPalette[7] = colorRGB{143, 39, 7}
	fireColorsPalette[8] = colorRGB{159, 47, 7}
	fireColorsPalette[9] = colorRGB{175, 63, 7}
	fireColorsPalette[10] = colorRGB{191, 71, 7}
	fireColorsPalette[11] = colorRGB{199, 71, 7}
	fireColorsPalette[12] = colorRGB{223, 79, 7}
	fireColorsPalette[13] = colorRGB{223, 87, 7}
	fireColorsPalette[14] = colorRGB{223, 87, 7}
	fireColorsPalette[15] = colorRGB{215, 95, 7}
	fireColorsPalette[16] = colorRGB{215, 95, 7}
	fireColorsPalette[17] = colorRGB{215, 103, 15}
	fireColorsPalette[18] = colorRGB{207, 111, 15}
	fireColorsPalette[19] = colorRGB{207, 119, 15}
	fireColorsPalette[20] = colorRGB{207, 127, 15}
	fireColorsPalette[21] = colorRGB{207, 135, 23}
	fireColorsPalette[22] = colorRGB{199, 135, 23}
	fireColorsPalette[23] = colorRGB{199, 143, 23}
	fireColorsPalette[24] = colorRGB{199, 151, 31}
	fireColorsPalette[25] = colorRGB{191, 159, 31}
	fireColorsPalette[26] = colorRGB{191, 159, 31}
	fireColorsPalette[27] = colorRGB{191, 167, 39}
	fireColorsPalette[28] = colorRGB{191, 167, 39}
	fireColorsPalette[29] = colorRGB{191, 175, 47}
	fireColorsPalette[30] = colorRGB{183, 175, 47}
	fireColorsPalette[31] = colorRGB{183, 183, 47}
	fireColorsPalette[32] = colorRGB{183, 183, 55}
	fireColorsPalette[33] = colorRGB{207, 207, 111}
	fireColorsPalette[34] = colorRGB{223, 223, 159}
	fireColorsPalette[35] = colorRGB{239, 239, 199}
	fireColorsPalette[36] = colorRGB{255, 255, 255}
}

func instatiateDataArray() {
	var numberOfPixels int = fireWidth * fireHeight
	for index := 0; index < numberOfPixels; index++ {
		firePixelsArray = append(firePixelsArray, minFire)
	}
	createFireSource()
}

func printDataAsArray(debug bool) {
	for index := 0; index < fireWidth*fireHeight; index++ {
		if debug {
			fmt.Print(index, " ")
		} else {
			fmt.Print(firePixelsArray[index], " ")
		}
	}
	fmt.Print("\n")
}

func printDataAsMatrix(debug bool) {
	for row := 0; row < fireHeight; row++ {
		for column := 0; column < fireWidth; column++ {
			pixelIndex := column + (fireWidth * row)
			fireIntensity := firePixelsArray[pixelIndex]
			if debug {
				fmt.Print("(", row, " , ", column, " )")
			} else {
				fmt.Print(fireIntensity, " ")
			}

		}
		fmt.Print("\n")
	}
}

func setFire(width int, height int) {
	fireWidth = width
	fireHeight = height
}

func createFireSource() {
	for column := 0; column < fireWidth; column++ {
		pixelIndex := (fireWidth*fireHeight - fireWidth) + column
		firePixelsArray[pixelIndex] = maxFire
	}
}

func calculateFirePropagation() {
	for column := 0; column < fireWidth; column++ {
		for row := 0; row < fireHeight; row++ {
			pixelIndex := column + (fireWidth * row)
			updateFireIntensityPerPixel(pixelIndex)
		}
	}
	clearTerminal()
	renderFire()
}

func updateFireIntensityPerPixel(currentIndex int) {
	belowPixelIndex := currentIndex + fireWidth

	if belowPixelIndex >= fireWidth*fireHeight {
		return
	}

	decay := rand.Intn(7)
	if decay < 0 {
		decay *= -1
	}

	//decay := 1
	belowPixelFireIntensity := firePixelsArray[belowPixelIndex]
	newFireIntensity := belowPixelFireIntensity - decay
	if newFireIntensity < 0 {
		newFireIntensity = 0
	}

	//firePixelsArray[currentIndex-decay] = newFireIntensity
	firePixelsArray[currentIndex] = newFireIntensity
}

func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func renderFire() {
	for row := 0; row < fireHeight; row++ {
		for column := 0; column < fireWidth; column++ {

			pixelIndex := column + (fireWidth * row)
			fireIntensity := firePixelsArray[pixelIndex]
			colorDefinition := fireColorsPalette[fireIntensity]

			//fmt.Print("| ", colorDefinition.r, " ", colorDefinition.g, " ", colorDefinition.b, " |")

			colorRender := color.RGB(colorDefinition.r, colorDefinition.g, colorDefinition.b, true)
			colorRender.Print("  ")
		}
		fmt.Print("\n")
	}
}

func main() {
	setFire(100, 30)
	initiateFireColorsPalette()
	instatiateDataArray()
	//printDataAsArray(printIndex)
	//printDataAsMatrix(printValue)

	for {
		calculateFirePropagation()
		time.Sleep(time.Millisecond * 10)
	}

}
