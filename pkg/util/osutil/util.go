package osutil

import (
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
)

var envMap map[string]string

func initEnvValues() {
	// TODO
	envMap = make(map[string]string)
	envs := os.Environ()
	for _, e := range envs {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			continue
		} else {
			envMap[parts[0]] = parts[1]
		}
	}
}

// Get environment values
func GetEnvValues() map[string]string {
	if envMap == nil {
		initEnvValues()
	}

	return envMap
}

func ReplaceEnvValues(text string) string {
	reg := regexp.MustCompile(`\${[\w.]*\}`)
	keys := reg.FindAllString(text, -1)

	if len(keys) == 0 {
		return text
	}

	if envMap == nil {
		initEnvValues()
	}

	var mapKey string

	for _, key := range keys {
		mapKey = key[2 : len(key)-1]
		if envMap[mapKey] != "" {
			text = strings.Replace(text, key, envMap[mapKey], -1)
		}
	}

	return text
}
func ReplaceEnvValuesHookFunc() mapstructure.DecodeHookFunc {
	return func(rf reflect.Type, rt reflect.Type, data interface{}) (interface{}, error) {
		if rf.Kind() == reflect.String && (rt.Kind() == reflect.String || rt.Kind() == reflect.Int || rt.Kind() == reflect.Interface) {
			if rf.Kind() == reflect.String && (rt.Kind() == reflect.Int) {
				valInt, err := strconv.Atoi(ReplaceEnvValues(data.(string)))
				if err == nil {
					return valInt, nil
				} else {
					return 0, nil
				}
			} else if reflect.TypeOf(data).Kind() == reflect.String {
				return ReplaceEnvValues(data.(string)), nil
			} else {
				return data, nil
			}
		} else {
			return data, nil
		}
	}
}
