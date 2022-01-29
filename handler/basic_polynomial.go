package handler

import (
	"context"
	cryp "crypto/rand"
	"fmt"
	"io"
	"math"
	"math/big"
)

var maxCoord = math.Pow(10, 6)

type polynomial struct {
	constant int
	degree   int
}

type BasicPolynomialHandler struct {
	version  string
	equation []polynomial
}

func NewBasicPolynomialHandler(
	version string,
) (*BasicPolynomialHandler, error) {
	return &BasicPolynomialHandler{
		version:  version,
		equation: []polynomial{},
	}, nil
}

func (handler *BasicPolynomialHandler) Check(
	ctx context.Context,
) error {
	fmt.Printf("inside basic poly check for version %s\n", handler.version)
	return nil
}

func getRandomPoint(
	reader io.Reader,
) (int64, error) {
	point, err := cryp.Int(reader, big.NewInt(int64(maxCoord)))
	if err != nil {
		return 0, err
	}
	return point.Int64(), nil
}

func (handler *BasicPolynomialHandler) CreateKeys(
	ctx context.Context,
	secret int,
	minimumKeysRequired int,
) error {
	var xAxisPoints []int
	var yAxisPoints []int
	n := 0
	xAxisPoints = append(xAxisPoints, 0)
	yAxisPoints = append(yAxisPoints, secret)
	for n < minimumKeysRequired {
		crypReader := cryp.Reader
		pnt, err := getRandomPoint(crypReader)
		if err != nil {
			return err
		}
		fmt.Printf("x axis point %v\n", pnt)
		xAxisPoints = append(xAxisPoints, int(pnt))
		pnt, err = getRandomPoint(crypReader)
		if err != nil {
			return err
		}
		fmt.Printf("y axis point %v\n", pnt)
		yAxisPoints = append(yAxisPoints, int(pnt))
		n++
	}
	return nil
}
