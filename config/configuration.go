package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

type ProjectConfig struct {
	Locales    []string
	Exports    []ExportConfig
	CsvSources []CsvSourceConfig
}

type ExportConfig struct {
	Format        string
	Path          string
	MatchPlatform string
	MatchKey      string
}
type CsvSourceConfig struct {
	Location      string
	Keys          string
	Platforms     string
	LocaleColumns map[string]string
	Replacements  map[string]string
}

type InvalidProjectConfigError struct {
	Fields []string
}

func (err *InvalidProjectConfigError) Error() string {
	return fmt.Sprintf("invalid config fields: %s", err.Fields)
}

func LoadInferredProjectConfig(projectDir string) (*ProjectConfig, error) {
	rawConfigData, err := ioutil.ReadFile(path.Join(projectDir, ".localice.json"))
	if err != nil {
		return nil, err
	}
	return UnmarshalProjectConfig(rawConfigData)
}

func UnmarshalProjectConfig(rawConfigData []byte) (*ProjectConfig, error) {
	var config ProjectConfig
	err := strictUnmarshal(rawConfigData, &config)
	if err != nil {
		return nil, err
	}

	invalidFields := make([]string, 0)
	if len(config.Exports) == 0 {
		invalidFields = append(invalidFields, "exports")
	}
	if len(config.Locales) == 0 {
		invalidFields = append(invalidFields, "locales")
	}
	if len(config.CsvSources) == 0 {
		invalidFields = append(invalidFields, "csvSources")
	}

	if len(invalidFields) != 0 {
		return &config, &InvalidProjectConfigError{invalidFields}
	}
	return &config, nil
}

func strictUnmarshal(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	return decoder.Decode(&v)
}
