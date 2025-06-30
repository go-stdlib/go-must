package must

import (
	"github.com/go-stdlib/go-errors"
	"reflect"
)

// ErrPanic is returned when a 'must' assertion fails and we must panic.
var ErrPanic = errors.Defined{
	Code:      "panic",
	Message:   "assertion failed",
	Namespace: "go-stdlib/go-must",
}

// Alias panics if given value cannot be aliased to the type 'T'.
func Alias[T any](v any) T {
	t, ok := v.(T)
	if !ok {
		panic(ErrPanic.Wrapf("must.Alias[%s](%T)", name[T](), v))
	}
	return t
}

// KnownErr panics if given error contains any unknown errors.
func KnownErr(err error) {
	var ke errors.Error
	if !errors.As(err, &ke) {
		panic(ErrPanic.Wrapf("must.KnownErr(%w)", err))
	}
}

// NotErr panics if given error is set.
//
// This is equivalent to 'V0'.
func NotErr(err error) {
	if err != nil {
		panic(ErrPanic.Wrapf("must.NotErr(%w)", err))
	}
}

// NotZero panics if given value is equal to the zero value of the type.
func NotZero[T any](t T) T {
	if reflect.ValueOf(t).IsZero() {
		panic(ErrPanic.Wrapf("must.NotZero[%s]", name[T]()))
	}
	return t
}

// Fn is an alias for 'Fn1' which is the most common case.
func Fn[T any](fn func() (T, error)) T {
	v, err := fn()
	if err != nil {
		panic(ErrPanic.Wrapf("must.Fn[%s] error=%w", name[T](), err))
	}
	return v
}

// Fn1 panics if the given function returns an error otherwise the value is returned.
func Fn1[T any](fn func() (T, error)) T {
	v, err := fn()
	if err != nil {
		panic(ErrPanic.Wrapf("must.Fn1[%s] error=%w", name[T](), err))
	}
	return v
}

// Fn2 panics if the given function returns an error otherwise the values are returned.
func Fn2[T1 any, T2 any](fn func() (T1, T2, error)) (T1, T2) {
	v1, v2, err := fn()
	if err != nil {
		panic(ErrPanic.Wrapf("must.Fn2[%s, %s] error=%w", name[T1](), name[T2](), err))
	}
	return v1, v2
}

// V0 panics if given error is set.
//
// This is equivalent to 'NotErr'.
func V0(err error) {
	if err != nil {
		panic(ErrPanic.Wrapf("must.V0 error=%w", err))
	}
}

// V1 panics if given error is set otherwise the value is returned.
func V1[T any](t T, err error) T {
	if err != nil {
		panic(ErrPanic.Wrapf("must.V1[%s] error=%w", name[T](), err))
	}
	return t
}

// V2 panics if given error is set otherwise the value is returned.
func V2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		panic(ErrPanic.Wrapf("must.V2[%s, %s] error=%w", name[T1](), name[T2](), err))
	}
	return t1, t2
}

// name returns the name of the generic type.
//
// This is helpful since '%T' in printf doesn't support generic types.
func name[T any]() string {
	return reflect.TypeFor[T]().String()
}
