package ubi

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		uuid         string
		expected     string
		wantErr      bool
		wantedErrMsg string
		f            func(string) (string, error)
	}{
		{
			uuid:         "",
			wantErr:      true,
			wantedErrMsg: "invalid UUID length: 0",
			f:            Game,
		},
		{
			uuid:         "incorrectUUID",
			wantErr:      true,
			wantedErrMsg: "invalid UUID length: 13",
			f:            Game,
		},
		{
			uuid:         "77d54c74s7827a4967,9c20-c97c8a8b0e70", // incorrect UUID format
			wantErr:      true,
			wantedErrMsg: "invalid UUID format",
			f:            Game,
		},
		{
			uuid:     "77d54c74-7827-4967-9c20-c97c8a8b0e70", // UUID v4
			expected: "Odd Fatal Frame II: Crimson Butterfly",
			wantErr:  false,
			f:        Game,
		},
		{
			uuid:     "a496d62f-b0bc-4c58-96b9-ed838127d724", // UUID v4
			expected: "Starving Curling 2006",
			wantErr:  false,
			f:        Game,
		},
		{
			uuid:     "7937710f-81a8-4080-a5b9-da8c2ca42c47", // UUID v4
			expected: "Fuzzy Walt Disney's The Jungle Book: Mowgli's Wild Adventure",
			wantErr:  false,
			f:        Game,
		},
		{
			uuid:     "b35c965e-0172-4ba5-9aba-da130e340c32", // UUID v4
			expected: "Sleepy Risk: Urban Assault",
			wantErr:  false,
			f:        Game,
		},
		{
			uuid:     "fa1f95b9-18f1-44c3-99c2-abd4462635a5", // UUID v4
			expected: "Silly Imagine: Makeup Artist",
			wantErr:  false,
			f:        Game,
		},
		{
			uuid:     "eae8aa0c-f1b3-4413-b140-b06dc1da3617", // UUID v4
			expected: "Sturdy Disney's Dinosaur",
			wantErr:  false,
			f:        Game,
		},
		{
			uuid:     "4a7dfcda-0e6d-11ed-861d-0242ac120002", // UUID v1
			expected: "Selfish Roller Champions",
			wantErr:  false,
			f:        Game,
		},
		{
			uuid:     "765aa5a6-0e6d-11ed-861d-0242ac120002", // UUID v1
			expected: "Ugly Roller Champions",
			wantErr:  false,
			f:        Game,
		},
		{
			uuid:     "839ee114-0e6d-11ed-861d-0242ac120002", // UUID v1
			expected: "Fussy Roller Champions",
			wantErr:  false,
			f:        Game,
		},
		{
			uuid:         "",
			wantErr:      true,
			wantedErrMsg: "invalid UUID length: 0",
			f:            Character,
		},
		{
			uuid:         "incorrectUUID",
			wantErr:      true,
			wantedErrMsg: "invalid UUID length: 13",
			f:            Character,
		},
		{
			uuid:         "77d54c74s7827a4967,9c20-c97c8a8b0e70", // incorrect UUID format
			wantErr:      true,
			wantedErrMsg: "invalid UUID format",
			f:            Character,
		},
		{
			uuid:     "77d54c74-7827-4967-9c20-c97c8a8b0e70", // UUID v4
			expected: "Odd Lance Brenner",
			wantErr:  false,
			f:        Character,
		},
		{
			uuid:     "a496d62f-b0bc-4c58-96b9-ed838127d724", // UUID v4
			expected: "Starving Roland",
			wantErr:  false,
			f:        Character,
		},
		{
			uuid:     "7937710f-81a8-4080-a5b9-da8c2ca42c47", // UUID v4
			expected: "Fuzzy Assassin Ray",
			wantErr:  false,
			f:        Character,
		},
		{
			uuid:     "b35c965e-0172-4ba5-9aba-da130e340c32", // UUID v4
			expected: "Sleepy Geon",
			wantErr:  false,
			f:        Character,
		},
		{
			uuid:     "fa1f95b9-18f1-44c3-99c2-abd4462635a5", // UUID v4
			expected: "Silly Synca",
			wantErr:  false,
			f:        Character,
		},
		{
			uuid:     "eae8aa0c-f1b3-4413-b140-b06dc1da3617", // UUID v4
			expected: "Sturdy Castle",
			wantErr:  false,
			f:        Character,
		},
		{
			uuid:     "4a7dfcda-0e6d-11ed-861d-0242ac120002", // UUID v1
			expected: "Selfish Sandbox",
			wantErr:  false,
			f:        Character,
		},
		{
			uuid:     "765aa5a6-0e6d-11ed-861d-0242ac120002", // UUID v1
			expected: "Ugly Sandbox",
			wantErr:  false,
			f:        Character,
		},
		{
			uuid:     "839ee114-0e6d-11ed-861d-0242ac120002", // UUID v1
			expected: "Fussy Sandbox",
			wantErr:  false,
			f:        Character,
		},
	}
	for _, tt := range tests {
		t.Run(tt.uuid, func(t *testing.T) {
			fName := strings.Split(runtime.FuncForPC(reflect.ValueOf(tt.f).Pointer()).Name(), ".")[2]
			gen, err := tt.f(tt.uuid)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("%s() error = %v, wantErr %v", fName, err, tt.wantErr)
					return
				}
				if err.Error() != tt.wantedErrMsg {
					t.Errorf("%s() errorMsg = %v, wantedErrMsg %v", fName, err.Error(), tt.wantedErrMsg)
				}
			}
			if gen != tt.expected {
				t.Errorf("%s() got = %v, expected %v", fName, gen, tt.expected)
			}
		})
	}
}

func TestFuzzUUIDs(t *testing.T) {
	testFunc := func(f func(string) (string, error)) {
		fName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[2]
		gen, err := f(uuid.NewString())
		if err != nil {
			t.Errorf("%s() error = %v", fName, err)
			return
		}
		if len(gen) == 0 {
			t.Errorf("%s() is empty = %v", fName, gen)
			return
		}
	}
	for i := 0; i < 1000; i++ {
		testFunc(Game)
		testFunc(Character)
	}
}
