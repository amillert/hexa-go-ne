package api

import (
	"fmt"

	"github.com/amillert/hexa-go-ne/internal/ports"
)

type Adapter struct {
	arithmetic ports.ArithmeticPort
	db         ports.DbPort
}

func NewAdapter(arithmetic ports.ArithmeticPort, db ports.DbPort) *Adapter {
	return &Adapter{arithmetic: arithmetic, db: db}
}

func (api *Adapter) GetAddition(a, b int32) (int32, error) {
	res, err := api.arithmetic.Addition(a, b)
	if err != nil {
		fmt.Printf("failed on addition: %v\n", err)
		return 0, err
	}

	errQuery := api.db.AddToHistory(res, "addition")
	if errQuery != nil {
		fmt.Printf("failed on adding addition to history: %v\n", err)
		return 0, errQuery
	}

	return res, nil
}

func (api *Adapter) GetSubtraction(a, b int32) (int32, error) {
	res, err := api.arithmetic.Subtraction(a, b)
	if err != nil {
		fmt.Printf("failed on subtraction: %v\n", err)
		return 0, err
	}

	errQuery := api.db.AddToHistory(res, "subtraction")
	if errQuery != nil {
		fmt.Printf("failed on adding subtraction to history: %v\n", err)
		return 0, errQuery
	}

	return res, nil
}

func (api *Adapter) GetMultiplication(a, b int32) (int32, error) {
	res, err := api.arithmetic.Multiplication(a, b)
	if err != nil {
		fmt.Printf("failed on multiplication: %v\n", err)
		return 0, err
	}

	errQuery := api.db.AddToHistory(res, "multiplication")
	if errQuery != nil {
		fmt.Printf("failed on adding multiplication to history: %v\n", err)
		return 0, errQuery
	}

	return res, nil
}

func (api *Adapter) GetDivision(a, b int32) (int32, error) {
	res, err := api.arithmetic.Division(a, b)
	if err != nil {
		fmt.Printf("failed on division: %v\n", err)
		return 0, err
	}

	errQuery := api.db.AddToHistory(res, "division")
	if errQuery != nil {
		fmt.Printf("failed on adding division to history: %v\n", err)
		return 0, errQuery
	}

	return res, nil
}
