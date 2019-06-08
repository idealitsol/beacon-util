package util

import (
	"fmt"
	"reflect"
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GoTimeToGrpcTime converts golang time.Time to google protobuf timestamp.Timestamp
func GoTimeToGrpcTime(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}

	if t.IsZero() {
		return nil
	}
	return &timestamp.Timestamp{Seconds: t.Unix()}
}

// GrpcTimeToGoTime converts google protobuf timestamp.Timestamp to golang time.Time
func GrpcTimeToGoTime(t *timestamp.Timestamp) *time.Time {
	if t.GetSeconds() == 0 {
		return nil
	}

	tt := time.Unix(t.GetSeconds(), 0)
	return &tt
}

// TransformGrpcToGo transforms or copy over values from gRPC to Golang
/*
	adminUsers := iam.AdminUsers{}
	for _, data := range res.GetAdminUsers() {
		adminUsers = append(adminUsers, transformerGrpcToGo(data, &iam.AdminUser{}).(iam.AdminUser))
	}
*/
func TransformGrpcToGo(grpcStruct interface{}, golangStruct interface{}) interface{} {
	// fmt.Println(grpcStruct)
	// fmt.Println(golangStruct)
	grpcs := reflect.ValueOf(grpcStruct).Elem()
	// fmt.Println("Done 1")
	golangs := reflect.ValueOf(golangStruct).Elem()
	// fmt.Println("Done 2")

	// fmt.Println("PASSED......")

	for i := 0; i < grpcs.NumField(); i++ {
		field := golangs.FieldByName(grpcs.Type().Field(i).Name)
		// fmt.Println(field.Kind())
		switch field.Kind() {
		case reflect.Int:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
			golangs.FieldByName(grpcs.Type().Field(i).Name).SetInt(grpcs.Field(i).Interface().(int64))
		case reflect.Float32:
		case reflect.Float64:
			golangs.FieldByName(grpcs.Type().Field(i).Name).SetFloat(grpcs.Field(i).Interface().(float64))
		case reflect.String:
			golangs.FieldByName(grpcs.Type().Field(i).Name).SetString(grpcs.Field(i).Interface().(string))
		case reflect.Bool:
			golangs.FieldByName(grpcs.Type().Field(i).Name).SetBool(grpcs.Field(i).Interface().(bool))
		case reflect.Ptr:
			if grpcs.Field(i).Interface() != nil {
				fmt.Println(field.Type().String())
				if field.Type().String() == "*time.Time" {
					ptime := GrpcTimeToGoTime(grpcs.Field(i).Interface().(*timestamp.Timestamp))
					golangs.FieldByName(grpcs.Type().Field(i).Name).Set(reflect.ValueOf(ptime))
				} else if field.Type().String() == "*string" {
					pstr := grpcs.Field(i).Interface().(string)
					golangs.FieldByName(grpcs.Type().Field(i).Name).Set(reflect.ValueOf(&pstr))
				} else {

					// typeee := reflect.TypeOf(golangs.Field(i).Type())
					// fmt.Println(grpcs.Field(i).Type())
					// fmt.Println(grpcs.Field(i).Interface())
					// fmt.Println(golangs.Field(i).Type())
					// pstr := grpcs.Field(i).Interface()
					// return TransformGrpcToGo(grpcs.Field(i).Interface(), reflect.New(golangs.Field(i).Type().Elem()).Interface())
					// return TransformGrpcToGo(grpcs.Field(i).Interface(), reflect.New(golangs.Field(i).Type().Elem()).Interface())

					// golangs.FieldByName(grpcs.Type().Field(i).Name).Set(reflect.ValueOf(pstr))
					// val := TransformGrpcToGo(grpcs.Field(i).Interface(), golangs.Field(i).Interface())
					// golangs.FieldByName(grpcs.Type().Field(i).Name).Set(reflect.ValueOf(pstr))
				}
			}
		default:
			// Invalid fields which we're not interested in
		}
	}
	return golangs.Interface()
}

// TransformGoToGrpc transforms or copy over values from Golang to gRPC
// Example: transformerGoToGrpc(&data, &pbx.AdminUser{}).(pbx.AdminUser)
/*
	pbxAdminUsers := []*pbx.AdminUser{}
	for _, data := range adminUsers {
		// pbxAdminUsers = append(pbxAdminUsers, transformAdminUserRPC(data))
		data := transformerGoToGrpc(&data, &pbx.AdminUser{}).(pbx.AdminUser)
		pbxAdminUsers = append(pbxAdminUsers, &data)
	}
*/
func TransformGoToGrpc(golangStruct interface{}, grpcStruct interface{}) interface{} {
	grpcs := reflect.ValueOf(grpcStruct).Elem()
	golangs := reflect.ValueOf(golangStruct).Elem()

	for i := 0; i < grpcs.NumField(); i++ {
		field := golangs.FieldByName(grpcs.Type().Field(i).Name)
		switch field.Kind() {
		case reflect.Int:
		case reflect.Int32:
		case reflect.Int64:
			grpcs.FieldByName(grpcs.Type().Field(i).Name).SetInt(golangs.Field(i).Interface().(int64))
		case reflect.String:
			grpcs.FieldByName(grpcs.Type().Field(i).Name).SetString(golangs.Field(i).Interface().(string))
		case reflect.Bool:
			grpcs.FieldByName(grpcs.Type().Field(i).Name).SetBool(golangs.Field(i).Interface().(bool))
		case reflect.Ptr:
			if !golangs.Field(i).IsNil() {
				if field.Type().String() == "*time.Time" {
					ptime := GoTimeToGrpcTime(golangs.Field(i).Interface().(*time.Time))
					grpcs.FieldByName(grpcs.Type().Field(i).Name).Set(reflect.ValueOf(ptime))
				} else if field.Type().String() == "*string" {
					pstr := golangs.Field(i).Interface().(*string)
					grpcs.FieldByName(grpcs.Type().Field(i).Name).SetString(*pstr)
				}
			}
		default:
			// Invalid fields which we're not interested in
		}
	}
	return grpcs.Interface()
}

// GrpcError return gRPC error
func GrpcError(code codes.Code, err error) error {
	return status.Errorf(
		codes.Internal,
		err.Error(),
	)
}
