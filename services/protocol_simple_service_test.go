package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateNewProtocol(t *testing.T) {
	decimalPlacesAfterDate := 8
	expectedDecimalPlacesInProtocol := decimalPlacesAfterDate + DefaultProtocolDecimalPlacesAfterDate
	ps := NewProtocolService(DefaultProtocolDecimalPlacesAfterDate)

	protocol, _ := ps.NewProtocol()

	assert.Equal(t, expectedDecimalPlacesInProtocol, len(protocol))
}

func TestGenerateNewProtocolWithCorrectlyDecimalPlaces(t *testing.T) {
	assertions := assert.New(t)
	numOfDecimalPlacesInDate := 8
	var tests = []struct {
		input    int
		expected int
	}{
		{8, 8 + numOfDecimalPlacesInDate},
		{4, 4 + numOfDecimalPlacesInDate},
		{2, 2 + numOfDecimalPlacesInDate},
		{15, 15 + numOfDecimalPlacesInDate},
		{777, 777 + numOfDecimalPlacesInDate},
	}

	for _, test := range tests {
		ps := NewProtocolService(test.input)
		protocol, _ := ps.NewProtocol()
		assertions.Equal(test.expected, len(protocol))
	}
}
