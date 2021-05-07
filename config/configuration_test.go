package config

import (
	"strconv"
	"strings"
	"testing"
)

func TestInvalidProjectConfigErrorsContainFields(t *testing.T) {
	err := InvalidProjectConfigError{[]string{"MyField"}}
	message := err.Error()
	if !strings.Contains(message, "MyField") {
		t.Errorf("message must contain all fields: %s", strconv.Quote(message))
	}
}

func TestNonJsonConfigsAreInvalid(t *testing.T) {
	_, err := UnmarshalProjectConfig([]byte(``))
	if err == nil {
		t.Error("non valid json configs are not valid")
	}
}

func TestEmptyConfigsAreInvalid(t *testing.T) {
	_, err := UnmarshalProjectConfig([]byte(`
	{}
	`))
	if err == nil {
		t.Error("empty configs are not valid")
	}
}

func TestCanLoadExampleProjectConfig(t *testing.T) {
	_, err := LoadInferredProjectConfig("../ExampleProject")
	if err != nil {
		t.Error(err)
	}
}

func TestCannotLoadWrongConfigPaths(t *testing.T) {
	_, err := LoadInferredProjectConfig("this/folder/does/not/exist")
	if err == nil {
		t.Error("found invalid config")
	}
}
