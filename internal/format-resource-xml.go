package internal

import (
	"encoding/xml"
	"io"
	"log"
)

type androidXMLResource struct {
	XMLName xml.Name           `xml:"resources"`
	Strings []androidXMLString `xml:",any"`
}
type androidXMLString struct {
	XMLName xml.Name `xml:"string"`
	Name    string   `xml:"name,attr"`
	Text    string   `xml:",chardata"`
}

type ResourceXmlLocalizationWriter struct {
	w io.Writer
}

func NewResourceXmlLocalizationWriter(w io.Writer) ResourceXmlLocalizationWriter {
	return ResourceXmlLocalizationWriter{w}
}
func (resourceWriter ResourceXmlLocalizationWriter) Write(localization Localization) error {
	_, err := io.WriteString(resourceWriter.w, `<?xml version="1.0" encoding="utf-8"?>`)
	if err != nil {
		return err
	}

	xmlEncoder := xml.NewEncoder(resourceWriter.w)
	resources := androidXMLResource{
		Strings: make([]androidXMLString, len(localization.Translations)),
	}
	for index, translation := range localization.Translations {
		resources.Strings[index] = androidXMLString{
			Name: translation.Key,
			Text: translation.Text,
		}
	}
	err = xmlEncoder.Encode(resources)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
