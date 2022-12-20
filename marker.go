package main

import (
	"container/list"
	"fmt"
	"strings"
)

type MarkerWrapper struct {
	Marker                   *list.List
	CharacterSet             map[string]int
	NumberDistinctCharacters int
}

func (m *MarkerWrapper) IsMarker() bool {
	return m.Marker.Len() == m.NumberDistinctCharacters && len(m.CharacterSet) == m.NumberDistinctCharacters
}

func (m *MarkerWrapper) AddCharacter(character string) {
	if m.Marker.Len() >= m.NumberDistinctCharacters {
		front := m.Marker.Front()
		frontValue := front.Value.(string)
		m.Marker.Remove(front)

		m.CharacterSet[frontValue] -= 1

		if m.CharacterSet[frontValue] == 0 {
			delete(m.CharacterSet, frontValue)
		}

	}
	m.Marker.PushBack(character)
	m.CharacterSet[character] += 1
}

func (m *MarkerWrapper) ToString() string {
	characters := make([]string, 0)

	currentElement := m.Marker.Front()

	for currentElement != nil {
		characters = append(characters, currentElement.Value.(string))
		currentElement = currentElement.Next()
	}

	return fmt.Sprintf("Marker: %v, CharacterSet: %v", strings.Join(characters, ""), m.CharacterSet)
}

func NewMarkerWrapper(numberOfDistinctCharacters int) *MarkerWrapper {
	characterSet := make(map[string]int)
	return &MarkerWrapper{
		Marker:                   list.New(),
		CharacterSet:             characterSet,
		NumberDistinctCharacters: numberOfDistinctCharacters,
	}
}
