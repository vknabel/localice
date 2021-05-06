package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

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
		platformIndex int
		keyIndex      int
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

		platform := record[platformIndex]
		key := record[keyIndex]
		if key == "" {
			return nil, fmt.Errorf("empty keys not supported: %s", record)
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
