package linode

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/alexsacr/linode/_third_party/mapstructure"
)

func marshallArgs(s interface{}) (map[string]interface{}, error) {
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		return nil, errors.New("cannot marshall non-struct")
	}

	ret := make(map[string]interface{})

	sVal := reflect.ValueOf(s)

	for i := 0; i < sVal.NumField(); i++ {
		tField := sVal.Type().Field(i)
		tag := tField.Tag.Get("args")
		if tag == "" {
			return nil, fmt.Errorf("No args tag set on %s", tField.Name)
		}

		vField := sVal.Field(i)

		// Special handle int-encoded bools
		if strings.Contains(tag, ",int") {
			if vField.IsNil() {
				continue
			}

			b, ok := vField.Elem().Interface().(bool)
			if !ok {
				return nil, fmt.Errorf("%s has an int tag, but is not of type bool", tField.Name)
			}

			tag = strings.Split(tag, ",")[0]

			if b == true {
				ret[tag] = 1
			} else {
				ret[tag] = 0
			}

			continue
		}

		ret[tag] = vField.Interface()
	}

	return ret, nil
}

func unmarshalSingle(data json.RawMessage, name string, out interface{}) error {
	t := reflect.TypeOf(out)
	if t.Kind() != reflect.Ptr {
		return errors.New("output must be a pointer")
	}

	var dataMap map[string]interface{}
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		return err
	}

	val, found := dataMap[name]
	if !found {
		return fmt.Errorf("%s not found in map, contents: %v", name, dataMap)
	}

	o := reflect.ValueOf(out).Elem()

	switch out.(type) {
	case *int:
		v, ok := val.(float64)
		if !ok {
			return fmt.Errorf("cannot assert %s to an integer via float, value: %v", name, val)
		}
		o.Set(reflect.ValueOf(int(v)))
	case *string:
		v, ok := val.(string)
		if !ok {
			return fmt.Errorf("cannot assert %s to a string, value: %v", name, val)
		}
		o.Set(reflect.ValueOf(v))
	case *bool:
		v, ok := val.(bool)
		if !ok {
			return fmt.Errorf("cannot assert %s to a bool, value: %v", name, val)
		}
		o.Set(reflect.ValueOf(v))
	default:
		return fmt.Errorf("%s is not an int, string, or bool, value: %v", name, val)
	}

	return nil
}

func unmarshalMultiMap(data json.RawMessage, out interface{}) error {
	t := reflect.TypeOf(out)
	if t.Kind() != reflect.Ptr {
		return errors.New("output must be a pointer")
	}

	var mapSlice []map[string]interface{}
	err := json.Unmarshal(data, &mapSlice)
	if err != nil {
		return err
	}

	decodeHook := func(from reflect.Kind, to reflect.Kind, v interface{}) (interface{}, error) {
		if from == reflect.String && to == reflect.Int {
			val := v.(string)
			if val == "" {
				return "0", nil
			}
		}
		return v, nil
	}

	v := reflect.ValueOf(out).Elem()
	for _, m := range mapSlice {
		tmp := reflect.New(v.Type().Elem()).Interface()

		config := &mapstructure.DecoderConfig{
			DecodeHook:       decodeHook,
			WeaklyTypedInput: true,
			Result:           tmp,
		}

		// Error can be safely ignored
		decoder, _ := mapstructure.NewDecoder(config)

		err = decoder.Decode(m)
		if err != nil {
			return err
		}

		vTmp := reflect.ValueOf(tmp).Elem()

		v.Set(reflect.Append(v, vTmp))
	}

	return nil
}
