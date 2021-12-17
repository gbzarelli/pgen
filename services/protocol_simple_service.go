package services

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// DateFormat The date format used to protocol prefix
const DateFormat = "20060102"

// DefaultProtocolDecimalPlacesAfterDate Number default to decimal places after date
const DefaultProtocolDecimalPlacesAfterDate = 8

// ProtocolServiceImpl struct to manage the protocol service
type ProtocolServiceImpl struct {
	decimalPlacesAfterDate int
	decimalPlacesFormat    string
	maxRandomValue         int
}

// NewProtocolService Create a new instance of NewProtocolService
func NewProtocolService(decimalPlacesAfterDate int) ProtocolService {
	return &ProtocolServiceImpl{
		decimalPlacesAfterDate: decimalPlacesAfterDate,
		decimalPlacesFormat:    "%0" + strconv.Itoa(decimalPlacesAfterDate) + "d",
		maxRandomValue:         maxRandomValueFromDecimalPlaces(decimalPlacesAfterDate),
	}
}

// NewProtocol method to generate a new protocol
func (pService *ProtocolServiceImpl) NewProtocol() (string, error) {
	return time.Now().Format(DateFormat) + pService.getComplementForProtocol(), nil
}

func maxRandomValueFromDecimalPlaces(decimalPlacesAfterDate int) int {
	var stringToConcat string
	for i := 0; i < decimalPlacesAfterDate; i++ {
		stringToConcat += "9"
	}
	value, _ := strconv.Atoi(stringToConcat)
	return value
}

func (pService *ProtocolServiceImpl) getComplementForProtocol() string {
	randomValue := rand.Intn(pService.maxRandomValue)
	return fmt.Sprintf(pService.decimalPlacesFormat, randomValue)
}
