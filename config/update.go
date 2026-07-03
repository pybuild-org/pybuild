package config

import "reflect"

func Update(name, field, key, value string) {
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
	valElem := v.Elem()
	switch typeElem.Kind() {

	case reflect.Struct:
		for i := 0; i < typeElem.NumField(); i++ {
			structField := typeElem.Field(i)
			propTag := structField.Tag.Get("prop")

			if propTag == key {
				fieldVal := valElem.Field(i)
				if !fieldVal.CanSet() {
					return
				}

				switch fieldVal.Interface().(type) {

				case string:
					fieldVal.SetString(value)

				case []string:
					newValue := reflect.ValueOf(value)
					fieldVal.Set(reflect.Append(fieldVal, newValue))

				}

				break
			}
		}

	case reflect.Slice:
		sliceLen := valElem.Len()
		if sliceLen == 0 {
			return
		}

		lastItem := valElem.Index(sliceLen - 1)
		if lastItem.Kind() != reflect.Map {
			return
		}

		mapKey := reflect.ValueOf(field)
		mapValue := lastItem.MapIndex(mapKey)

		if !mapValue.IsValid() || mapValue.IsNil() {
			structPtrType := lastItem.Type().Elem()
			if structPtrType.Kind() != reflect.Pointer {
				return
			}

			newStructPtr := reflect.New(structPtrType.Elem())
			lastItem.SetMapIndex(mapKey, newStructPtr)
			mapValue = newStructPtr
		}

		if mapValue.Kind() != reflect.Pointer {
			return
		}

		structVal := mapValue.Elem()
		if structVal.Kind() != reflect.Struct {
			return
		}

		structType := structVal.Type()
		for i := 0; i < structType.NumField(); i++ {
			structField := structType.Field(i)
			propTag := structField.Tag.Get("prop")

			if propTag == key {
				fieldVal := structVal.Field(i)
				if !fieldVal.CanSet() {
					return
				}

				switch fieldVal.Interface().(type) {

				case string:
					fieldVal.SetString(value)

				case []string:
					newValue := reflect.ValueOf(value)
					fieldVal.Set(reflect.Append(fieldVal, newValue))

				}

				break
			}
		}

	}
}

func NewGroup(name string) {
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
