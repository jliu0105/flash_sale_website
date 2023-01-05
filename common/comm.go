package common

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

//Map data to the structure according to the sql tag in the structure and convert the type
func DataToStructByTagSql(data map[string]string, obj interface{}) {
	objValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < objValue.NumField(); i++ {
		// get the value of the sql
		value := data[objValue.Type().Field(i).Tag.Get("sql")]
		// get the name of the corespond string
		name := objValue.Type().Field(i).Name
		// get the type of the corresponding string
		structFieldType := objValue.Field(i).Type()
		// get the type of the variable
		val := reflect.ValueOf(value)
		var err error
		if structFieldType != val.Type() {
			// typecasting
			val, err = TypeConversion(value, structFieldType.Name())
			if err != nil {

			}
		}
		// set the value of type
		objValue.FieldByName(name).Set(val)
	}
}

// typecasting
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	return reflect.ValueOf(value), errors.New("unknown type" + ntype)
}
