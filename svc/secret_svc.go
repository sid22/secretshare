package svc

import (
	"fmt"

	handler "github.com/sid22/secretshare/handler"
)

type SecretAlgoTypes string

const (
	BasicPolynomial = SecretAlgoTypes("basicpolynomial")
)

func GetSecretHandler(
	algoType SecretAlgoTypes,
	version string,
) (handler.SecretShareHandler, error) {
	switch algoType {
	case BasicPolynomial:
		return handler.NewBasicPolynomialHandler(
			version,
		)
	default:
		return nil, fmt.Errorf("unknown algo type %v", algoType)
	}
}
