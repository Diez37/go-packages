package closer

import (
	"context"
	"io"
)

const (
	// ExitCodeError exit code status on error
	ExitCodeError int = 1
)

// Closer general interface for closing the context to stop the application
type Closer interface {
	io.Closer

	// GetContext return context.Context
	GetContext() context.Context
}
