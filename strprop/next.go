package strprop

import "reflect"

func Next(name string) {
	target, ok := cfg[name]
	if !ok {
		return
	}

	t := reflect.TypeOf(target)
	v := reflect.ValueOf(target)
	if t.Kind() != reflect.Pointer || v.IsNil() {
		return
	}

	typeElem := t.Elem()
	if typeElem.Kind() != reflect.Slice {
		return
	}

	valElem := v.Elem()
	outerPtrType := valElem.Type().Elem()
	if outerPtrType.Kind() != reflect.Pointer {
		return
	}

	outerStructType := outerPtrType.Elem()
	if outerStructType.Kind() != reflect.Struct {
		return
	}

	newOuterPtr := reflect.New(outerStructType)
	newSlice := reflect.Append(valElem, newOuterPtr)
	valElem.Set(newSlice)
}
