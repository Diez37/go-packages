package closer

import (
	"context"
	"github.com/diez37/go-packages/log"
	"os"
	"os/signal"
	"syscall"
)

// OsSignal listener for os signal of application exit, implements Closer
type OsSignal struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	exitSignal chan os.Signal
	informer   log.Informer
}

// NewOsSignal creating and returning a new instance of OsSignal.
// Binds a listener of OS signals and calls cancel function on SIGINT, SIGTERM, SIGABRT
func NewOsSignal(informer log.Informer) Closer {
	closer := &OsSignal{informer: informer}

	closer.ctx, closer.cancelFunc = context.WithCancel(context.Background())
	closer.exitSignal = make(chan os.Signal, 1)

	signal.Notify(closer.exitSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	go func(closer *OsSignal) {
		<-closer.exitSignal

		_ = closer.Close()
	}(closer)

	return closer
}

func (closer *OsSignal) GetContext() context.Context {
	return closer.ctx
}

func (closer *OsSignal) Close() error {
	closer.informer.Info("closer: shutdown")

	closer.cancelFunc()

	return nil
}
