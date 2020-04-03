package contextimpl

import (
	"errors"
	"time"
)

type emptyCtx int

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct{}

	Err() error

	Value(key interface{}) interface{}
}

func (ctx *cancelCtx) Deadline() (deadline time.Time, ok bool) {
	return ctx.parent.Deadline()
}

func (ctx *cancelCtx) Done() <-chan struct{} {
	return ctx.done
}

func (ctx *cancelCtx) Err() error {
	return ctx.err
}

func (ctx *cancelCtx) Value(key interface{}) interface{} {
	return ctx.parent.Value(key)
}

func Background() *emptyCtx {
	return background
}

func TODO() *emptyCtx {
	return todo
}

type cancelCtx struct {
	parent Context
	done chan struct{}
	err error
}

var Canceled = errors.New("context canceled")

type CancelFunc func()

func WithCancel(parent Context) (Context , CancelFunc)  {
	ctx := &cancelCtx{
		parent: parent,
		done:   make(chan struct{}),
	}

	cancel := func() {
		ctx.err = Canceled
		close(ctx.done)
	}

	return ctx , cancel

}