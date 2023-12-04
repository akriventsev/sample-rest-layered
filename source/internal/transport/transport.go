package transport

import "context"

type Transport interface {
	Start(ctx context.Context) error
}
