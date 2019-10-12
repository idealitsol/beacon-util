package util

import (
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

// MapStructureDecode uses mapstructure library but the reason for this is because of the stringToDate, as mapstructure.Decode(model, modelMap) does not do it well
func MapStructureDecode(model interface{}, out interface{}) {
	stringToDateTimeHook := func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if t == reflect.TypeOf(time.Time{}) {
			return time.Parse(time.RFC3339, data.(string))
		}

		return data, nil
	}

	config := mapstructure.DecoderConfig{
		DecodeHook: stringToDateTimeHook,
		Result:     &model,
	}

	mpdecoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		panic(err)
	}

	mpdecoder.Decode(out)
}
