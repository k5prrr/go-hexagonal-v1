/*
value, err := Config.Int("name/name")
*/
package Config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var configPath string = "configs/configs.json"
var configJson map[string]interface{}

var cacheString map[string]string
var cacheBool map[string]bool
var cacheInt map[string]int
var cacheFloat map[string]float64

func init() {
	ClearCache()
}
func ClearCache() {
	cacheString = make(map[string]string)
	cacheBool = make(map[string]bool)
	cacheInt = make(map[string]int)
	cacheFloat = make(map[string]float64)
}
func Update() {
	content, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	err = json.Unmarshal(content, &configJson)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return
	}

	ClearCache()
}
func Current(name string) (interface{}, error) {
	if configJson == nil {
		Update()
	}

	keys := strings.Split(name, "/")
	var current interface{} = configJson

	for _, key := range keys {
		if currentMap, ok := current.(map[string]interface{}); ok {
			if val, exists := currentMap[key]; exists {
				current = val
			} else {
				return 0, fmt.Errorf("key '%s' not found in config", key)
			}
		} else {
			return 0, fmt.Errorf("key '%s' is not a map", key)
		}
	}

	return current, nil
}
func String(name string) (string, error) {
	value, exists := cacheString[name]
	if exists {
		return value, nil
	}

	current, err := Current(name)
	if err != nil {
		return "", err
	}

	switch v := current.(type) {
	case string:
		cacheString[name] = v
		return v, nil
	case int:
		cacheString[name] = strconv.Itoa(v)
		return cacheString[name], nil
	case bool:
		if v {
			cacheString[name] = "true"
			return cacheString[name], nil
		} else {
			cacheString[name] = "false"
			return cacheString[name], nil
		}
	case float64:
		cacheString[name] = strconv.FormatFloat(v, 'f', -1, 64)
		return cacheString[name], nil
	default:
		return "", fmt.Errorf("unsupported type '%s' for key '%s'", reflect.TypeOf(v).String(), name)
	}
}
func Bool(name string) (bool, error) {
	value, exists := cacheBool[name]
	if exists {
		return value, nil
	}

	current, err := Current(name)
	if err != nil {
		return false, err
	}

	switch v := current.(type) {
	case bool:
		cacheBool[name] = v
		return v, nil
	case int:
		cacheBool[name] = v != 0
		return cacheBool[name], nil
	case string:
		boolValue, err := strconv.ParseBool(v)
		if err != nil {
			return false, fmt.Errorf("cannot convert '%s' to bool", v)
		}
		cacheBool[name] = boolValue
		return boolValue, nil
	case float64:
		cacheBool[name] = v != 0
		return cacheBool[name], nil
	default:
		return false, fmt.Errorf("unsupported type '%s' for key '%s'", reflect.TypeOf(v).String(), name)
	}
}
func Int(name string) (int, error) {
	value, exists := cacheInt[name]
	if exists {
		return value, nil
	}

	current, err := Current(name)
	if err != nil {
		return 0, err
	}

	switch v := current.(type) {
	case int:
		cacheInt[name] = v
		return v, nil
	case float64:
		cacheInt[name] = int(v)
		return cacheInt[name], nil
	case string:
		intValue, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("cannot convert '%s' to int", v)
		}
		cacheInt[name] = intValue
		return cacheInt[name], nil
	case bool:
		if v {
			cacheInt[name] = 1
		} else {
			cacheInt[name] = 0
		}
		return cacheInt[name], nil
	default:
		return 0, fmt.Errorf("unsupported type '%s' for key '%s'", reflect.TypeOf(v).String(), name)
	}
}
func Float(name string) (float64, error) {
	value, exists := cacheFloat[name]
	if exists {
		return value, nil
	}

	current, err := Current(name)
	if err != nil {
		return 0, err
	}

	switch v := current.(type) {
	case float64:
		cacheFloat[name] = v
		return v, nil
	case int:
		cacheFloat[name] = float64(v)
		return cacheFloat[name], nil
	case string:
		floatValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot convert '%s' to float64", v)
		}
		cacheFloat[name] = floatValue
		return floatValue, nil
	case bool:
		if v {
			return 1.0, nil
		} else {
			return 0.0, nil
		}
	default:
		return 0, fmt.Errorf("unsupported type '%s' for key '%s'", reflect.TypeOf(v).String(), name)
	}
}

