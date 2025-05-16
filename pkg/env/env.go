// Use
// env := env.New("")
// value, err := env.Int("key")
package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Env struct {
	Path    string
	Strings map[string]string
	Bools   map[string]bool
	Ints    map[string]int
	Floats  map[string]float64
}

func New(envPath string) *Env {
	if envPath == "" {
		envPath = ".env"
	}

	env := Env{Path: envPath}
	env.ClearCache()
	env.Update()
	return &env
}

func (e *Env) ClearCache() {
	e.Strings = make(map[string]string)
	e.Bools = make(map[string]bool)
	e.Ints = make(map[string]int)
	e.Floats = make(map[string]float64)
}

func (e *Env) Update() error {
	text, err := os.ReadFile(e.Path)
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	lines := strings.Split(string(text), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		key := strings.TrimSpace(parts[0])
		if key == "" {
			continue
		}
		var value string
		if len(parts) > 1 {
			value = strings.TrimSpace(parts[1])
		} else {
			value = ""
		}

		e.Strings[key] = value
	}
	return nil
}

func (e *Env) String(name string) (string, error) {
	value, exists := e.Strings[name]
	if exists {
		return value, nil
	}

	return "", fmt.Errorf("variable %s not exists ", name)
}

func (e *Env) Bool(name string) (bool, error) {
	value, exists := e.Bools[name]
	if exists {
		return value, nil
	}

	valueString, err := e.String(name)
	if err != nil {
		return false, err
	}

	valueBool, err := strconv.ParseBool(valueString)
	if err != nil {
		return false, fmt.Errorf("variable %s cannot convert '%s' to bool", name, valueString)
	}
	e.Bools[name] = valueBool
	return valueBool, nil
}

func (e *Env) Int(name string) (int, error) {
	value, exists := e.Ints[name]
	if exists {
		return value, nil
	}

	valueString, err := e.String(name)
	if err != nil {
		return 0, err
	}

	valueInt, err := strconv.Atoi(valueString)
	if err != nil {
		return 0, fmt.Errorf("variable %s cannot convert '%s' to int", name, valueString)
	}
	e.Ints[name] = valueInt
	return valueInt, nil
}

func (e *Env) Float(name string) (float64, error) {
	value, exists := e.Floats[name]
	if exists {
		return value, nil
	}

	valueString, err := e.String(name)
	if err != nil {
		return 0, err
	}

	valueFloat, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		return 0, fmt.Errorf("variable %s cannot convert '%s' to float64", name, valueString)
	}
	e.Floats[name] = valueFloat
	return valueFloat, nil
}
