package handler

import (
	"context"
	"fmt"
)

type BasicPolynomialHandler struct {
	version string
}

func NewBasicPolynomialHandler(
	version string,
) (*BasicPolynomialHandler, error) {
	return &BasicPolynomialHandler{
		version: version,
	}, nil
}

func (handler *BasicPolynomialHandler) Check(
	ctx context.Context,
) error {
	fmt.Printf("inside basic poly check for version %s\n", handler.version)
	return nil
}
