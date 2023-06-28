package main

import (
	"fmt"
	"math"
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

func main() {
	f, err := CreateFinancing(1000, 12, 1.1)
	if err != nil {
		fmt.Println("Error while creating financing:", err)
		return
	}

	fmt.Println(f.FutureValue)
	fmt.Println(f.InstallmentCost)

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
