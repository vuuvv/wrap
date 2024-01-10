package wrap

import "github.com/vuuvv/errors"

type Result[T any] struct {
	value T
	err   error
}

func New[T any](value T, err error) Result[T] {
	return Result[T]{value: value, err: errors.WithStack(err)}
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

func Error[T any](err error) Result[T] {
	return Result[T]{err: errors.WithStack(err)}
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(errors.WithStack(r.err))
	}
	return r.value
}

func (r Result[T]) UnwrapOrError() (T, error) {
	return r.value, r.err
}

func (r Result[T]) UnwrapOr(value T) T {
	if r.err != nil {
		return value
	}
	return r.value
}

func (r *Result[T]) Recover() {
	if reason := recover(); reason != nil {
		err, ok := reason.(error)
		if !ok {
			err = errors.Errorf("%v", reason)
		}
		r.err = errors.WithStack(err)
	}
}

func Recover[T any](res *Result[T]) {
	if reason := recover(); reason != nil {
		err, ok := reason.(error)
		if !ok {
			err = errors.Errorf("%v", reason)
		}
		res.err = errors.WithStack(err)
	}
}

func RecoverOr[T any](res *Result[T], val Result[T]) {
	if reason := recover(); reason != nil {
		*res = val
	}
}

func RecoverHandle[T any](res *Result[T], handle func(res *Result[T], err error) Result[T]) {
	if reason := recover(); reason != nil {
		err, ok := reason.(error)
		if !ok {
			err = errors.Errorf("%v", reason)
		}
		*res = handle(res, err)
	}
}
