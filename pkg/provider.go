package provider

import (
	"context"
)

type Provider interface {
	IsOnline(ctx context.Context) (bool, error)
}
