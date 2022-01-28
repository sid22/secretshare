package main

import (
	"context"
	"fmt"

	"github.com/sid22/secretshare/svc"
)

func main() {
	handlr, err := svc.GetSecretHandler(
		svc.SecretAlgoTypes(svc.BasicPolynomial),
		"v1",
	)
	if err != nil {
		fmt.Printf("error is %v\n", err)
		return
	}
	ctx := context.Background()
	err = handlr.Check(ctx)
	if err != nil {
		fmt.Printf("error in check %v\n", err)
	}

}
