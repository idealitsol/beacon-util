/*
 * @author    Emmanuel Kofi Bessah
 * @email     bekinsoft@gmail.com
 */

package util

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/structs"
)

// GetGormModelPrimaryKeyField returns the name of gorm primary key from the tag
func GetGormModelPrimaryKeyField(st interface{}) string {
	s := structs.New(st)
	for _, f := range s.Fields() {
		if strings.Contains(f.Tag("gorm"), "primary_key") {
			return f.Name()
		}
	}
	// val := reflect.ValueOf(s)
	// for i := 0; i < val.Type().NumField(); i++ {
	// 	if tag := val.Type().Field(i).Tag.Get("gorm"); tag != "" && tag == "primary_key" {
	// 		return val.Type().Field(i).Name
	// 	}
	// }

	return ""
}

// TypeOfField gets the type of interface
func TypeOfField(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// InvokeModelConstraintFunction calls the "ConstraintError" function on a model
func InvokeModelConstraintFunction(val interface{}, err error) []reflect.Value {
	v := reflect.ValueOf(val)

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(err)

	// return v.Call(in)
	return v.MethodByName("ConstraintError").Call(in)
	// return v.MethodByName("ConstraintError").Call([]reflect.Value{})

	// return v.MethodByName("ConstraintError").Call([]reflect.Value{})

	// for i := 0; i < t.NumMethod(); i++ {
	// 	m := t.Method(i)

	// 	if m.Name == "GetSomeAdditionalData" {

	// 	}
	// 	// fmt.Printf("Method %s\n", m)
	// 	// fmt.Printf("\tName: %s\n", m.Name)
	// 	// fmt.Printf("\tType: %s\n", m.Type)
	// 	// fmt.Printf("\tFunc: %s\n", m.Func)
	// 	// fmt.Printf("\tPackage path: %s\n", m.PkgPath)
	// }

	// return nil
}

// ValidateGORMFields validates gorm model fields which has either 'primary' or 'not null' properties
func ValidateGORMFields(st interface{}, onlyPrimaryKeys ...bool) error {
	s := structs.New(st)

	for _, f := range s.Fields() {
		condition := getCondition(f, onlyPrimaryKeys)
		switch f.Kind() {
		case reflect.String:
			if condition && f.Value().(string) == "" {
				return fmt.Errorf(`Missing required field '` + f.Tag("json") + `'`)
			}
		case reflect.Int:
			if condition && f.Value().(int) == 0 {
				return fmt.Errorf(`Missing required field '` + f.Tag("json") + `'`)
			}
		case reflect.Int32:
			if condition && f.Value().(int32) == 0 {
				return fmt.Errorf(`Missing required field '` + f.Tag("json") + `'`)
			}
		case reflect.Int64:
			if condition && f.Value().(int64) == 0 {
				return fmt.Errorf(`Missing required field '` + f.Tag("json") + `'`)
			}
		case reflect.Struct: // 25
			if condition && f.Value().(time.Time).IsZero() {
				return fmt.Errorf(`Missing required field '` + f.Tag("json") + `'`)
			}
		}
	}

	return nil
}

func getCondition(f *structs.Field, onlyPrimaryKeys []bool) bool {
	if len(onlyPrimaryKeys) > 0 && onlyPrimaryKeys[0] {
		return strings.Contains(f.Tag("gorm"), "primary_key") && !strings.Contains(f.Tag("gorm"), "default")
	}

	return (strings.Contains(f.Tag("gorm"), "primary_key") && !strings.Contains(f.Tag("gorm"), "default")) || strings.Contains(f.Tag("gorm"), "not null")
}
