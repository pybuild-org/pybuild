package strprop

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

		lastItemPtr := valElem.Index(sliceLen - 1)
		if lastItemPtr.Kind() != reflect.Pointer || lastItemPtr.IsNil() {
			return
		}

		outerStructVal := lastItemPtr.Elem()
		outerStructType := outerStructVal.Type()

		var innerStructVal reflect.Value
		foundField := false

		for i := 0; i < outerStructType.NumField(); i++ {
			structField := outerStructType.Field(i)

			if structField.Tag.Get("prop") == field {
				innerStructVal = outerStructVal.Field(i)
				if innerStructVal.Kind() != reflect.Struct {
					return
				}

				foundField = true
				break
			}
		}

		if !foundField {
			return
		}

		innerStructType := innerStructVal.Type()
		for i := 0; i < innerStructType.NumField(); i++ {
			structField := innerStructType.Field(i)

			if structField.Tag.Get("prop") == key {
				fieldVal := innerStructVal.Field(i)
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
