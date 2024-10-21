package orderv2

import (
	"context"
)

type service interface {
	UploadOrder(ctx context.Context) error
}
