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
	mapType := valElem.Type().Elem()
	if mapType.Kind() != reflect.Map {
		return
	}

	newMap := reflect.MakeMap(mapType)
	newSlice := reflect.Append(valElem, newMap)
	valElem.Set(newSlice)
}
