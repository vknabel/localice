package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"regexp"
)

func ReadSource(projectDir string, source CsvSourceConfig) ([]Localization, error) {
	localizations, err := ReadCsvSource(projectDir, source)
	if err != nil {
		return nil, err
	}

	for l10nIndex := range localizations {
		for translationIndex := range localizations[l10nIndex].Translations {
			reference := &localizations[l10nIndex].Translations[translationIndex]
			for pattern, replacement := range source.Replacements {
				expr, err := regexp.Compile(pattern)
				if err != nil {
					return nil, err
				}
				reference.Text = expr.ReplaceAllString(reference.Text, replacement)
			}
		}
	}
	return localizations, err
}

func ReadCsvSource(projectDir string, source CsvSourceConfig) ([]Localization, error) {
	sourcePath := path.Join(projectDir, source.Location)
	rawFile, err := os.Open(sourcePath)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(rawFile)

	header, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	var (
		platformIndex int                     = -1
		keyIndex      int                     = -1
		localeIndices map[string]int          = make(map[string]int)
		localizations map[string]Localization = make(map[string]Localization)
	)

	for index, column := range header {
		switch column {
		case source.Platforms:
			platformIndex = index
		case source.Keys:
			keyIndex = index
		default:
			if localeName, ok := source.LocaleColumns[column]; ok {
				localeIndices[column] = index
				localizations[column] = Localization{
					Name:         localeName,
					Translations: make([]Translation, 0),
				}
			}
		}
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		key := record[keyIndex]
		if key == "" {
			return nil, fmt.Errorf("empty keys not supported: %s", record)
		}

		var platform string
		if platformIndex > -1 {
			platform = record[platformIndex]
		}
		for sourceCol, recordIndex := range localeIndices {
			localization := localizations[sourceCol]
			localization.Translations = append(localizations[sourceCol].Translations, Translation{
				Key:      key,
				Platform: platform,
				Text:     record[recordIndex],
			})
			localizations[sourceCol] = localization
		}

	}

	listOfLocalizations := make([]Localization, 0, len(localizations))
	for _, value := range localizations {
		listOfLocalizations = append(listOfLocalizations, value)
	}
	return listOfLocalizations, nil
}
