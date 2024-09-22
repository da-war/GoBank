package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceFiles(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func ReadBalanceFiles(fileName string) (float64, error) {
	var balanceText, err = os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("Failed to read balance file")
	}

	balanceString := string(balanceText)
	balance, err := strconv.ParseFloat(balanceString, 64)
	if err != nil {
		return 0, errors.New("Failed to convert balance to float")
	}

	return balance, nil
}
