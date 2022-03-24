package api

import (
	"fmt"

	"github.com/amillert/hexa-go-ne/internal/ports"
)

type Adapter struct {
	arithmetic ports.ArithmeticPort
}

func NewAdapter(arithmetic ports.ArithmeticPort) *Adapter {
	return &Adapter{arithmetic}
}

func (api *Adapter) GetAddition(a, b int32) (int32, error) {
	res, err := api.arithmetic.Addition(a, b)
	if err != nil {
		fmt.Printf("failed on addition: %v\n", err)
		return 0, err
	}

	return res, nil
}

func (api *Adapter) GetSubtraction(a, b int32) (int32, error) {
	res, err := api.arithmetic.Subtration(a, b)
	if err != nil {
		fmt.Printf("failed on subtraction: %v\n", err)
		return 0, err
	}

	return res, nil
}

func (api *Adapter) GetMultiplication(a, b int32) (int32, error) {
	res, err := api.arithmetic.Multiplication(a, b)
	if err != nil {
		fmt.Printf("failed on multiplication: %v\n", err)
		return 0, err
	}

	return res, nil
}

func (api *Adapter) GetDivision(a, b int32) (int32, error) {
	res, err := api.arithmetic.Division(a, b)
	if err != nil {
		fmt.Printf("failed on division: %v\n", err)
		return 0, err
	}

	return res, nil
}
