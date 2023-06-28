package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
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

	financingFiles, err := listFinancingFiles()
	if err != nil {
		log.Fatal(err)
	}

	if len(financingFiles) == 0 {
		log.Fatal("There is no financing file.")
	}

	for _, fin := range financingFiles {
		fmt.Println("Financing file:", fin)
	}
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

func listFinancingFiles() ([]string, error) {
	files, err := os.ReadDir(DFINANCE_DIR)
	if err != nil {
		return nil, err
	}

	// Filter .dfin files only
	var financingFiles []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if idx := strings.Index(file.Name(), ".dfin"); idx != -1 {
			financingFiles = append(financingFiles, file.Name())
		}
	}

	return financingFiles, nil
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
