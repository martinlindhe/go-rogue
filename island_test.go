package rogue

import (
	"fmt"
	"image/png"
	"os"
	"testing"
)

func BenchmarkGenerateIsland(b *testing.B) {
	for n := 0; n < b.N; n++ {

		GenerateIsland(666, 220, 140)
	}
}

func TestGenerateIsland(t *testing.T) {

	seed := int64(123)
	island := GenerateIsland(seed, 200, 100)
	island.FillWithCritters()

	islandColImgFile, _ := os.Create("test-island.png")
	png.Encode(islandColImgFile, island.ColoredHeightMapAsImage())

	fmt.Println("oi ")

	island.PrintSpawns()
}
