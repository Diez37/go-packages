package repeater

import "context"

type Process interface {
	Process(ctx context.Context) error
}
