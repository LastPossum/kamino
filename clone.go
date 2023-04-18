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

// WithExpectedPtrsCount is a functional option for the Clone function that sets the initial capacity for the pointers map used during the cloning process.
// If the amount of pointers in the source object is known and rather big, this setting could reduce allocations and improve performance.
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

// WithForceUnexported is a functional option for the Clone function. When this setting is enabled, unexported fields will be cloned forcefully.
func WithForceUnexported() func(*config) {
	return func(c *config) {
		c.unexportedStrategy = forceDeepCopyUnexportedStrategy
	}
}

// WithZeroUnexported is a functional option for the Clone function.  When this setting is enabled, unexported fields will be forcefully set to zero value of their kind.
func WithZeroUnexported() func(*config) {
	return func(c *config) {
		c.unexportedStrategy = forceZeroUnexported
	}
}

// WithErrOnUnsupported is a functional option for the Clone function. When this setting is enabled, attempting to clone channels and functions, even if they are nested, will result in an error.
func WithErrOnUnsupported() func(*config) {
	return func(c *config) {
		c.errOnUnsuported = true
	}
}

// cloneCtx contains settings and pointers to values mapping for single clone
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

// Clone returns a deep copy of passed argument. The cloned value will be identical to the source value.
// Cloned pointers will point to a new object with the same value as in the source.
// This also works for nested values, such as parts of collections (arrays, slices, and maps) or fields of passed structs.
// It additionally guarantied that if two or more pointers at the source object references to the same value, clones pointers will refers to same vallue.
// But if two or more slices at the source object refer to the same memory area (i.e. the same underlying array), the copy will refer to different arrays.
// By default, unexported struct fields will be shallow copied, but this can be changed with functional options.
// Unsupported types (channels and funcs) will be shallow copied by default, but this can also be changed with functional options.
// Unsafe pointers will be treated like ordinary values.
func Clone[T any](src T, opts ...funcOptions) (T, error) {
	// build a clone config
	cfg := config{}
	for _, o := range opts {
		o(&cfg)
	}

	// clone ctx will be passet to any recursive call cloneNested object
	// it sontains settings and a pointers to values mapping
	ctx := &cloneCtx{
		ptrs:               make(map[unsafe.Pointer]reflect.Value, cfg.expectedPointersCount),
		unexportedStrategy: cfg.unexportedStrategy,
		errOnUnsuported:    cfg.errOnUnsuported,
	}

	// a shallow copy of the source was already made at the moment of passing it as an argument to a Copy function
	// so make an adressable reflect value for it
	valAtPtr := reflect.ValueOf(&src).Elem()
	// and recursively traverse it for deep copying its parts if needed
	err := cloneNested(ctx, valAtPtr)
	return src, err
}

func cloneNested(ctx *cloneCtx, v reflect.Value) error {
	kind := v.Kind()
	// check if the value should be copied, i.e. it's neither of basic type, nor zero
	if isBasicKind(kind) || v.IsZero() {
		return nil
	}
	//  handle according to original's Kind
	switch kind {
	case reflect.Struct:
		// for structs iterate over it's fields
		for i := 0; i < v.NumField(); i++ {
			wField := v.Field(i)

			// if it can be set (exported)
			if wField.CanSet() {
				// recursively clone it
				if err := cloneNested(ctx, wField); err != nil {
					return err
				}
				continue
			}

			// if it is unexported it can be treaten according to on of following strategies
			switch ctx.unexportedStrategy {
			// do nothing (i.e shallow copy)
			case shallowCopyUnexportedStrategy:
				continue
			// forcely deep copy it
			case forceDeepCopyUnexportedStrategy:
				newAt := reflect.NewAt(wField.Type(), unsafe.Pointer(wField.UnsafeAddr()))
				if err := cloneNested(ctx, newAt.Elem()); err != nil {
					return err
				}
			// forcely turn it to zero value
			case forceZeroUnexported:
				typ := wField.Type()
				newAt := reflect.NewAt(typ, unsafe.Pointer(wField.UnsafeAddr()))
				newAt.Elem().Set(reflect.Zero(typ))
			}
		}
	case reflect.Array:
		// for arrays allocate the new one of the elem type
		elem := v.Type().Elem()
		res := reflect.New(reflect.ArrayOf(v.Len(), elem)).Elem()
		// and copy values from source at once
		reflect.Copy(res, v)

		// if an elem kind is basic just return
		if isBasicKind(elem.Kind()) {
			v.Set(res)
			return nil
		}

		// otherwise recursively clone the elems
		for i := 0; i < res.Len(); i++ {
			if err := cloneNested(ctx, res.Index(i)); err != nil {
				return err
			}
		}

		// replace the source with the copy
		v.Set(res)
	case reflect.Slice:
		// for slices allocate the new one of the elem type
		typ := v.Type()
		res := reflect.MakeSlice(typ, v.Len(), v.Cap())
		// and copy values from source at once
		reflect.Copy(res, v)

		// if an elem kind is basic just return
		if isBasicKind(typ.Elem().Kind()) {
			v.Set(res)
			return nil
		}

		// otherwise recursively clone the elems
		for i := 0; i < res.Len(); i++ {
			if err := cloneNested(ctx, res.Index(i)); err != nil {
				return err
			}
		}

		// replace the source with the copy
		v.Set(res)
	case reflect.Map:
		// for maps allocate the new one of the source type
		typ := v.Type()
		res := reflect.MakeMapWithSize(typ, v.Len())

		// create new values for iterating over the map
		iter := v.MapRange()
		newK := reflect.New(typ.Key()).Elem()
		newV := reflect.New(typ.Elem()).Elem()

		// check if key and value maps are of basic type just once
		keyIsBasic := isBasicKind(typ.Key().Kind())
		valueIsBasic := isBasicKind(typ.Elem().Kind())

		for iter.Next() {
			k := iter.Key()
			// if key needs to be copied, copy it recursively
			if !keyIsBasic && !k.IsZero() {
				newK.Set(k)
				if err := cloneNested(ctx, newK); err != nil {
					return err
				}
				k = newK
			}
			v := iter.Value()
			// if value needs to be copied, copy it recursively
			if !valueIsBasic && !k.IsZero() {
				newV.Set(v)
				if err := cloneNested(ctx, newV); err != nil {
					return err
				}
				v = newV
			}
			// put key and value to a map copy
			res.SetMapIndex(k, v)
		}

		// replace the source with the copy
		v.Set(res)
	case reflect.Pointer:
		// check if pointer has already been met
		ptr := v.UnsafePointer()
		if newV, ok := ctx.resolvePointer(ptr); ok {
			// if it has been use previous copy
			v.Set(newV)
			return nil
		}

		// if this pointer has not been met create a new value
		newV := reflect.New(v.Elem().Type())
		// and put it to context
		ctx.setPointer(ptr, newV)

		// put the source value to new pointer and recursively copy it
		newV.Elem().Set(v.Elem())
		if err := cloneNested(ctx, newV.Elem()); err != nil {
			return err
		}
		// replace the source with the copy
		v.Set(newV)
	case reflect.Interface:
		// if interface underlying value needs to be copied
		el := v.Elem()
		if isBasicKind(el.Kind()) || el.IsZero() {
			return nil
		}
		// copy it recursively
		newV := reflect.New(v.Elem().Type()).Elem()
		newV.Set(v.Elem())
		if err := cloneNested(ctx, newV); err != nil {
			return err
		}
		v.Set(newV)
	default:
		// if unsupported type strategy has been setted to err - return an error
		if ctx.errOnUnsuported {
			return fmt.Errorf("unsupported type: %s", v.Type().Name())
		}
	}

	return nil
}

func isBasicKind(k reflect.Kind) bool {
	switch k {
	case
		reflect.Bool,
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
		reflect.String,
		reflect.UnsafePointer:
		return true
	default:
		return false
	}
}
