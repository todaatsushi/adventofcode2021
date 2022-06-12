package aoc5

import (
	"log"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func convertXAndY(xAndY string) (int, int) {
	xAndY = strings.TrimSpace(xAndY)
	rawCoords := strings.Split(xAndY, ",")

	var err error
	x, err := strconv.Atoi(rawCoords[0])
	if err != nil {
		log.Fatalf("Couldn't convert %v to int", rawCoords[0])
	}
	y, err := strconv.Atoi(rawCoords[1])
	if err != nil {
		log.Fatalf("Couldn't convert %v to int", rawCoords[1])
	}
	return x, y
}

func readCoordinates(coordinates string) (*Coordinate, *Coordinate) {
	cleanCoordinates := make([]*Coordinate, 2)
	splitCoordinate := strings.Split(coordinates, "->")

	var x int
	var y int
	for i, input := range splitCoordinate {
		x, y = convertXAndY(input)
		cleanCoordinates[i] = &Coordinate{x: x, y: y}
	}
	return cleanCoordinates[0], cleanCoordinates[1]
}
