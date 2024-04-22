package page

import (
	"context"
	"time"

	"github.com/go-rod/rod"
)

type PageWithOnError struct {
	rod.Page
	OnError          func(p *PageWithOnError)
	cancel           context.CancelFunc
	quitTimeoutCheck chan struct{} // Channel to signal quitting the timeout goroutine
}

// NewPageWithOnError
// has all of the functions of a rod page, but if you call WithTimeout, it will
// call OnError and close the page if the timeout is reached.
func NewPageWithOnError(page *rod.Page, onErrorFunc func(p *PageWithOnError)) *PageWithOnError {
	return &PageWithOnError{
		Page:             *page,
		OnError:          onErrorFunc,
		quitTimeoutCheck: make(chan struct{}),
	}
}

func (p *PageWithOnError) WithTimeout(d time.Duration) *PageWithOnError {
	// optionally quit exiting timeout goroutine
	p.CancelTimeout()

	page, cancel := p.WithCancel()
	p.cancel = cancel
	p.Page = *page

	pctx := page.GetContext()
	go func(p *PageWithOnError) {
		select {
		case <-time.After(d):
			p.OnError(p)
		case <-pctx.Done():
		case <-p.quitTimeoutCheck:
			return
		}
		cancel()
	}(p)
	return p
}

func (p *PageWithOnError) CancelTimeout() *PageWithOnError {
	select {
	case p.quitTimeoutCheck <- struct{}{}:
	default:
	}
	return p
}
