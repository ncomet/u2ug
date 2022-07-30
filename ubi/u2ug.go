// Package ubi holds the Game function that returns a random adjective and Ubisoft game name according to a UUID
package ubi

import (
	"fmt"

	ud "github.com/google/uuid"
)

const byteShift = 4

// Game will return a random adjective and Ubisoft game name according to param uuid
// It is an injective function
// error will be returned if UUID is not a valid UUID format, or is a random string
func Game(uuid string) (string, error) {
	return generate(uuid, games)
}

// Character will return a random adjective and Ubisoft character name according to param uuid
// It is an injective function
// error will be returned if UUID is not a valid UUID format, or is a random string
func Character(uuid string) (string, error) {
	return generate(uuid, characters)
}

func generate(uuid string, suffixes []string) (string, error) {
	pUuid, err := ud.Parse(uuid)
	if err != nil {
		return "", err
	}
	aIndex := getIndex(pUuid, 0) % uint64(len(adjectives))
	sIndex := getIndex(pUuid, 8) % uint64(len(suffixes))
	return fmt.Sprintf("%s %s", adjectives[aIndex], suffixes[sIndex]), nil
}

func getIndex(uuid ud.UUID, start uint) uint64 {
	bytes, err := uuid.MarshalBinary()
	if err != nil {
		panic(err)
	}
	s1 := bytes[start : start+byteShift]
	s2 := bytes[start+byteShift : start+2*byteShift]
	var res uint64
	for i, b := range s1 {
		res += uint64(b) * uint64(s2[i])
	}
	return res
}
