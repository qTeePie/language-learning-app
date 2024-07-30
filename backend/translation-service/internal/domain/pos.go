package domain

import (
	"fmt"
)

// PartOfSpeech represents a part of speech
type PartOfSpeech int

const (
	InvalidPOS PartOfSpeech = iota - 1 // Define a specific invalid value
	Noun
	Verb
	Adjective
	// Add more as needed
)

var posMap = map[string]PartOfSpeech{
	"noun":      Noun,
	"verb":      Verb,
	"adjective": Adjective,
	// Add more as needed
}

// ParsePOS converts a string to a PartOfSpeech value
func ParsePOS(pos string) (PartOfSpeech, error) {
	p, ok := posMap[pos]
	if !ok {
		return InvalidPOS, fmt.Errorf("invalid part of speech: %s", pos)
	}
	return p, nil
}

func IsValidPOS(pos string) bool {
	_, ok := posMap[pos]
	return ok
}
