package must

import (
	"fmt"
	"github.com/go-stdlib/go-tests"
	"testing"
)

func TestKnownErr(t *testing.T) {
	tests.Table[error, any]{
		"+known":   {Got: ErrPanic, WantPanic: false},
		"-unknown": {Got: fmt.Errorf("random error"), WantPanic: true},
	}.Unit(t, func(t *tests.Test, tc tests.Testcase[error, any]) {
		KnownErr(tc.Got)
	})
}

func TestNotErr(t *testing.T) {
	tests.Table[error, any]{
		"+nil": {Got: nil},
		"-err": {Got: fmt.Errorf("oh no"), WantPanic: true},
	}.Unit(t, func(t *tests.Test, tc tests.Testcase[error, any]) {
		NotErr(tc.Got)
	})
}

func TestNotZero(t *testing.T) {
	tests.Table[any, any]{
		"+int8":       {Got: int8(42)},
		"+int16":      {Got: int16(42)},
		"+int32":      {Got: int32(42)},
		"+int64":      {Got: int64(42)},
		"+int":        {Got: int(42)},
		"+uint8":      {Got: uint8(42)},
		"+uint16":     {Got: uint16(42)},
		"+uint32":     {Got: uint32(42)},
		"+uint64":     {Got: uint64(42)},
		"+uint":       {Got: uint(42)},
		"+string":     {Got: "hello"},
		"+bool":       {Got: true},
		"+float32":    {Got: float32(42.0)},
		"+float64":    {Got: float64(42.0)},
		"+complex64":  {Got: complex64(42.0)},
		"+complex128": {Got: complex128(42.0)},
		"+struct":     {Got: struct{ val int }{val: 42}},
		"+slice":      {Got: []int{}},
		"+map":        {Got: map[string]int{}},
		"+chan":       {Got: make(chan int)},
		"-int8":       {Got: int8(0), WantPanic: true},
		"-int16":      {Got: int16(0), WantPanic: true},
		"-int32":      {Got: int32(0), WantPanic: true},
		"-int64":      {Got: int64(0), WantPanic: true},
		"-int":        {Got: int(0), WantPanic: true},
		"-uint8":      {Got: uint8(0), WantPanic: true},
		"-uint16":     {Got: uint16(0), WantPanic: true},
		"-uint32":     {Got: uint32(0), WantPanic: true},
		"-uint64":     {Got: uint64(0), WantPanic: true},
		"-uint":       {Got: uint(0), WantPanic: true},
		"-string":     {Got: "", WantPanic: true},
		"-bool":       {Got: false, WantPanic: true},
		"-float32":    {Got: float32(0.0), WantPanic: true},
		"-float64":    {Got: float64(0.0), WantPanic: true},
		"-complex64":  {Got: complex64(0.0), WantPanic: true},
		"-complex128": {Got: complex128(0.0), WantPanic: true},
		"-struct":     {Got: struct{}{}, WantPanic: true},
		"-slice":      {Got: []int(nil), WantPanic: true},
		"-map":        {Got: map[string]int(nil), WantPanic: true},
		"-chan":       {Got: (chan int)(nil), WantPanic: true},
	}.Unit(t, func(t *tests.Test, tc tests.Testcase[any, any]) {
		NotZero(tc.Got)
	})
}

func TestFn(t *testing.T) {
	type Got[T any] struct {
		fn func() (T, error)
	}

	tests.Table[Got[any], any]{
		"+no error": {
			Got: Got[any]{
				fn: func() (any, error) { return int8(42), nil },
			},
			WantPanic: false,
		},
		"-error": {
			Got: Got[any]{
				fn: func() (any, error) { return nil, fmt.Errorf("oh no") },
			},
			WantPanic: true,
		},
	}.Unit(t, func(t *tests.Test, tc tests.Testcase[Got[any], any]) {
		Fn(tc.Got.fn)
	})
}

func TestFn1(t *testing.T) {
	type Got[T any] struct {
		fn func() (T, error)
	}

	tests.Table[Got[any], any]{
		"+no error": {
			Got: Got[any]{
				fn: func() (any, error) { return int8(42), nil },
			},
			WantPanic: false,
		},
		"-error": {
			Got: Got[any]{
				fn: func() (any, error) { return nil, fmt.Errorf("oh no") },
			},
			WantPanic: true,
		},
	}.Unit(t, func(t *tests.Test, tc tests.Testcase[Got[any], any]) {
		Fn1(tc.Got.fn)
	})
}

func TestFn2(t *testing.T) {
	type Got[T any] struct {
		fn func() (T, T, error)
	}

	tests.Table[Got[any], any]{
		"+no error": {
			Got: Got[any]{
				fn: func() (any, any, error) { return int8(42), "", nil },
			},
			WantPanic: false,
		},
		"-error": {
			Got: Got[any]{
				fn: func() (any, any, error) { return nil, nil, fmt.Errorf("oh no") },
			},
			WantPanic: true,
		},
	}.Unit(t, func(t *tests.Test, tc tests.Testcase[Got[any], any]) {
		Fn2(tc.Got.fn)
	})
}

func TestV0(t *testing.T) {
	tests.Table[error, any]{
		"+no error": {
			Got:       nil,
			WantPanic: false,
		},
		"-error": {
			Got:       fmt.Errorf("oh no"),
			WantPanic: true,
		},
	}.Unit(t, func(t *tests.Test, tc tests.Testcase[error, any]) {
		V0(tc.Got)
	})
}

func TestV1(t *testing.T) {
	type Got[T any] struct {
		v   T
		err error
	}

	tests.Table[Got[any], any]{
		"+no error": {
			Got: Got[any]{
				v:   "hello world",
				err: nil,
			},
			WantPanic: false,
		},
		"-error": {
			Got: Got[any]{
				v:   nil,
				err: fmt.Errorf("oh no"),
			},
			WantPanic: true,
		},
	}.Unit(t, func(t *tests.Test, tc tests.Testcase[Got[any], any]) {
		V1(tc.Got.v, tc.Got.err)
	})
}

func TestV2(t *testing.T) {
	type Got[T any] struct {
		v1  T
		v2  T
		err error
	}

	tests.Table[Got[any], any]{
		"+no error": {
			Got: Got[any]{
				v1:  "hello world",
				v2:  "another value",
				err: nil,
			},
			WantPanic: false,
		},
		"-error": {
			Got: Got[any]{
				v1:  nil,
				v2:  nil,
				err: fmt.Errorf("oh no"),
			},
			WantPanic: true,
		},
	}.Unit(t, func(t *tests.Test, tc tests.Testcase[Got[any], any]) {
		V2(tc.Got.v1, tc.Got.v1, tc.Got.err)
	})
}
