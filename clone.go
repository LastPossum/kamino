package kamino

import (
	"fmt"
	"reflect"
	"unsafe"
)

type config struct {
	expectedPointersCount int
	unexportedStrategy    unexportedStrategy
	errOnUnsuported       bool
}

type funcOptions func(*config)

func WithExpectedPtrsCount(cnt int) func(*config) {
	return func(c *config) {
		c.expectedPointersCount = cnt
	}
}

type unexportedStrategy int

const (
	shallowCopyUnexportedStrategy unexportedStrategy = iota
	forceDeepCopyUnexportedStrategy
	forceZeroUnexported
)

func WithForceUnexported() func(*config) {
	return func(c *config) {
		c.unexportedStrategy = forceDeepCopyUnexportedStrategy
	}
}

func WithZeroUnexported() func(*config) {
	return func(c *config) {
		c.unexportedStrategy = forceZeroUnexported
	}
}

func WithErrOnUnsupported() func(*config) {
	return func(c *config) {
		c.errOnUnsuported = true
	}
}

type cloneCtx struct {
	ptrs               map[unsafe.Pointer]reflect.Value
	unexportedStrategy unexportedStrategy
	errOnUnsuported    bool
}

func (ctx *cloneCtx) resolvePointer(ptr unsafe.Pointer) (reflect.Value, bool) {
	v, ok := ctx.ptrs[ptr]
	return v, ok
}

func (ctx *cloneCtx) setPointer(ptr unsafe.Pointer, value reflect.Value) {
	ctx.ptrs[ptr] = value
}

func Clone[T any](src T, opts ...funcOptions) (T, error) {
	if isBoxedNil(src) {
		return src, nil
	}

	cfg := config{}
	for _, o := range opts {
		o(&cfg)
	}

	ctx := &cloneCtx{
		ptrs:               make(map[unsafe.Pointer]reflect.Value, cfg.expectedPointersCount),
		unexportedStrategy: cfg.unexportedStrategy,
		errOnUnsuported:    cfg.errOnUnsuported,
	}
	valAtPtr := reflect.ValueOf(&src).Elem()
	err := cloneNested(ctx, valAtPtr)
	return src, err
}

func cloneNested(ctx *cloneCtx, v reflect.Value) error {
	if !needCp(v) {
		return nil
	}
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			wField := v.Field(i)

			if wField.CanSet() {
				if err := cloneNested(ctx, wField); err != nil {
					return err
				}
				continue
			}

			switch ctx.unexportedStrategy {
			case shallowCopyUnexportedStrategy:
				continue
			case forceDeepCopyUnexportedStrategy:
				newAt := reflect.NewAt(wField.Type(), unsafe.Pointer(wField.UnsafeAddr()))
				if err := cloneNested(ctx, newAt.Elem()); err != nil {
					return err
				}
			case forceZeroUnexported:
				typ := wField.Type()
				newAt := reflect.NewAt(typ, unsafe.Pointer(wField.UnsafeAddr()))
				newAt.Elem().Set(reflect.Zero(typ))
			}
		}
	case reflect.Array:
		elem := v.Type().Elem()
		res := reflect.New(reflect.ArrayOf(v.Len(), elem)).Elem()
		reflect.Copy(res, v)

		if isBasicKind(elem.Kind()) {
			v.Set(res)
			return nil
		}

		for i := 0; i < res.Len(); i++ {
			if err := cloneNested(ctx, res.Index(i)); err != nil {
				return err
			}
		}

		v.Set(res)
	case reflect.Slice:
		typ := v.Type()
		res := reflect.MakeSlice(typ, v.Len(), v.Cap())
		reflect.Copy(res, v)

		if isBasicKind(typ.Elem().Kind()) {
			v.Set(res)
			return nil
		}

		for i := 0; i < res.Len(); i++ {
			if err := cloneNested(ctx, res.Index(i)); err != nil {
				return err
			}
		}

		v.Set(res)
	case reflect.Map:
		typ := v.Type()
		res := reflect.MakeMapWithSize(typ, v.Len())

		iter := v.MapRange()
		newK := reflect.New(typ.Key()).Elem()
		newV := reflect.New(typ.Elem()).Elem()

		for iter.Next() {
			k := iter.Key()
			if needCp(k) {
				newK.Set(k)
				if err := cloneNested(ctx, newK); err != nil {
					return err
				}
				k = newK
			}
			v := iter.Value()
			if needCp(v) {
				newV.Set(v)
				if err := cloneNested(ctx, newV); err != nil {
					return err
				}
				v = newV
			}
			res.SetMapIndex(k, v)
		}

		v.Set(res)
	case reflect.Pointer:

		ptr := v.UnsafePointer()
		if newV, ok := ctx.resolvePointer(ptr); ok {
			v.Set(newV)
			return nil
		}

		newV := reflect.New(v.Elem().Type())
		ctx.setPointer(ptr, newV)

		newV.Elem().Set(v.Elem())
		if err := cloneNested(ctx, newV.Elem()); err != nil {
			return err
		}

		v.Set(newV)
	case reflect.Interface:
		if !needCp(v.Elem()) {
			return nil
		}
		newV := reflect.New(v.Elem().Type()).Elem()
		newV.Set(v.Elem())
		if err := cloneNested(ctx, newV); err != nil {
			return err
		}
		v.Set(newV)
	default:
		if ctx.errOnUnsuported {
			return fmt.Errorf("unsupported type: %s", v.Type().Name())
		}
	}

	return nil
}

func needCp(v reflect.Value) bool {
	k := v.Kind()
	if isBasicKind(k) {
		return false
	}
	if v.IsZero() {
		return false
	}
	return true
}

func isBoxedNil(src any) bool {
	return src == nil || src == any(nil)
}

func isBasicKind(k reflect.Kind) bool {
	switch k {
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.String:
		return true
	default:
		return false
	}
}
