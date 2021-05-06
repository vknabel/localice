package internal

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

type ProjectConfig struct {
	Locales    []string
	Platforms  []PlatformConfig
	CsvSources []CsvSourceConfig
}

type PlatformConfig struct {
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
}

func LoadProjectConfig(projectDir string) (ProjectConfig, error) {
	rawConfigData, err := ioutil.ReadFile(path.Join(projectDir, ".localice.json"))
	if err != nil {
		return ProjectConfig{}, err
	}
	var config ProjectConfig
	err = json.Unmarshal(rawConfigData, &config)
	if err != nil {
		return ProjectConfig{}, err
	}

	return config, nil
}
