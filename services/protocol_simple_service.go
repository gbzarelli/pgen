package services

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const DateFormat = "20060102"
const DefaultProtocolDecimalPlacesAfterDate = 8

type ProtocolServiceImpl struct {
	decimalPlacesAfterDate int
	decimalPlacesFormat    string
	maxRandomValue         int
}

func NewProtocolService(decimalPlacesAfterDate int) ProtocolService {
	return &ProtocolServiceImpl{
		decimalPlacesAfterDate: decimalPlacesAfterDate,
		decimalPlacesFormat:    "%0" + strconv.Itoa(decimalPlacesAfterDate) + "d",
		maxRandomValue:         maxRandomValueFromDecimalPlaces(decimalPlacesAfterDate),
	}
}

func maxRandomValueFromDecimalPlaces(decimalPlacesAfterDate int) int {
	var stringToConcat string
	for i := 0; i < decimalPlacesAfterDate; i++ {
		stringToConcat += "9"
	}
	value, _ := strconv.Atoi(stringToConcat)
	return value
}

func (pService *ProtocolServiceImpl) NewProtocol() (string, error) {
	return time.Now().Format(DateFormat) + pService.getComplementForProtocol(), nil
}

func (pService *ProtocolServiceImpl) getComplementForProtocol() string {
	randomValue := rand.Intn(pService.maxRandomValue)
	return fmt.Sprintf(pService.decimalPlacesFormat, randomValue)
}
