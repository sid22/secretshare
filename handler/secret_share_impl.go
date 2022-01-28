package handler

import "context"

type SecretShareHandler interface {
	Check(ctx context.Context) error
}
