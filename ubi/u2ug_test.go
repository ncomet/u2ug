package ubi

import (
	"testing"

	"github.com/google/uuid"
)

func TestGame(t *testing.T) {
	tests := []struct {
		uuid         string
		expected     string
		wantErr      bool
		wantedErrMsg string
	}{
		{
			uuid:         "",
			wantErr:      true,
			wantedErrMsg: "invalid UUID length: 0",
		},
		{
			uuid:         "incorrectUUID",
			wantErr:      true,
			wantedErrMsg: "invalid UUID length: 13",
		},
		{
			uuid:         "77d54c74s7827a4967,9c20-c97c8a8b0e70", // incorrect UUID format
			wantErr:      true,
			wantedErrMsg: "invalid UUID format",
		},
		{
			uuid:     "77d54c74-7827-4967-9c20-c97c8a8b0e70", // UUID v4
			expected: "Odd Fatal Frame II: Crimson Butterfly",
			wantErr:  false,
		},
		{
			uuid:     "a496d62f-b0bc-4c58-96b9-ed838127d724", // UUID v4
			expected: "Starving Curling 2006",
			wantErr:  false,
		},
		{
			uuid:     "7937710f-81a8-4080-a5b9-da8c2ca42c47", // UUID v4
			expected: "Fuzzy Walt Disney's The Jungle Book: Mowgli's Wild Adventure",
			wantErr:  false,
		},
		{
			uuid:     "b35c965e-0172-4ba5-9aba-da130e340c32", // UUID v4
			expected: "Sleepy Risk: Urban Assault",
			wantErr:  false,
		},
		{
			uuid:     "fa1f95b9-18f1-44c3-99c2-abd4462635a5", // UUID v4
			expected: "Silly Imagine: Makeup Artist",
			wantErr:  false,
		},
		{
			uuid:     "eae8aa0c-f1b3-4413-b140-b06dc1da3617", // UUID v4
			expected: "Sturdy Disney's Dinosaur",
			wantErr:  false,
		},
		{
			uuid:     "4a7dfcda-0e6d-11ed-861d-0242ac120002", // UUID v1
			expected: "Selfish Roller Champions",
			wantErr:  false,
		},
		{
			uuid:     "765aa5a6-0e6d-11ed-861d-0242ac120002", // UUID v1
			expected: "Ugly Roller Champions",
			wantErr:  false,
		},
		{
			uuid:     "839ee114-0e6d-11ed-861d-0242ac120002", // UUID v1
			expected: "Fussy Roller Champions",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.uuid, func(t *testing.T) {
			game, err := Game(tt.uuid)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Game() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err.Error() != tt.wantedErrMsg {
					t.Errorf("Game() errorMsg = %v, wantedErrMsg %v", err.Error(), tt.wantedErrMsg)
				}
			}
			if game != tt.expected {
				t.Errorf("Game() got = %v, expected %v", game, tt.expected)
			}
		})
	}
}

func TestFuzzUUIDs(t *testing.T) {
	for i := 0; i < 1000; i++ {
		game, err := Game(uuid.NewString())
		if err != nil {
			t.Errorf("Game() error = %v", err)
			return
		}
		if len(game) == 0 {
			t.Errorf("Game() is empty = %v", game)
			return
		}
	}
}
