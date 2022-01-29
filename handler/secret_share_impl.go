package handler

import "context"

type SecretShareHandler interface {
	Check(ctx context.Context) error
	CreateKeys(ctx context.Context, secret int, minimumKeysRequired int) error
}
