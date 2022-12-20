package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type DatastreamBuffer string

const PartOneNumberDistinctCharacters = 4
const PartTwoNumberDistinctCharacters = 14

func getDatastreamBuffer(reader io.Reader) DatastreamBuffer {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	datastreamBuffer := scanner.Text()

	return DatastreamBuffer(datastreamBuffer)
}

func getNumCharactersProcessedToFindStartOfPacketMarker(datastreamBuffer DatastreamBuffer) int {
	startOfPacketMarkerWrapper := NewMarkerWrapper(PartOneNumberDistinctCharacters)
	for index, character := range datastreamBuffer {
		startOfPacketMarkerWrapper.AddCharacter(string(character))
		if startOfPacketMarkerWrapper.IsMarker() {
			return index + 1
		}
	}
	return -1
}

func getNumCharactersProcessedToFindStartOfMessageMarker(datastreamBuffer DatastreamBuffer) int {
	startOfPacketMarkerWrapper := NewMarkerWrapper(PartTwoNumberDistinctCharacters)
	for index, character := range datastreamBuffer {
		startOfPacketMarkerWrapper.AddCharacter(string(character))
		if startOfPacketMarkerWrapper.IsMarker() {
			return index + 1
		}
	}
	return -1
}

func main() {
	file, err := os.Open("/home/ec2-user/go/src/github.com/iamwillzhu/adventofcode2022day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	datastreamBuffer := getDatastreamBuffer(file)

	numCharactersProcessedToFindStartOfPacketMarker := getNumCharactersProcessedToFindStartOfPacketMarker(datastreamBuffer)
	numCharactersProcessedToFindStartOfMessageMarker := getNumCharactersProcessedToFindStartOfMessageMarker(datastreamBuffer)

	fmt.Printf("The number of characters processed to find the first marker in the datastream buffer is : %d\n", numCharactersProcessedToFindStartOfPacketMarker)
	fmt.Printf("The number of characters processed to find the start of message marker in the datastream buffer is: %d\n", numCharactersProcessedToFindStartOfMessageMarker)

}
