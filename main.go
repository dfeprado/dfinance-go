package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

type InvalidInterestError float64

func (err InvalidInterestError) Error() string {
	return fmt.Sprintf("Invalid interest rate. It must be between 0 and 1. %v was given.", float64(err))
}

type Financing struct {
	Principal       float64
	Installments    uint16
	Interest        float64
	FutureValue     float64
	InstallmentCost float64
}

var DFINANCE_DIR string = fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".dfinance")

func main() {
	createDFinanceHome()
}

func createDFinanceHome() {
	err := os.Mkdir(DFINANCE_DIR, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	} else if err != nil && os.IsExist(err) {
		log.Default().Printf("DFinance home already exists.")
	} else {
		log.Default().Printf("DFinance home created at %v", DFINANCE_DIR)
	}
}

func CreateFinancing(cost float64, installmentCount uint16, interest float64) (*Financing, error) {
	if interest < 0 || interest > 1 {
		return nil, InvalidInterestError(interest)
	}

	futureValue := cost * math.Pow(1+interest, float64(installmentCount))
	installmentCost := futureValue / float64(installmentCount)

	return &Financing{
		cost,
		installmentCount,
		interest,
		futureValue,
		installmentCost,
	}, nil
}
