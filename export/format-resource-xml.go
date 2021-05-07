package export

import (
	"encoding/xml"
	"io"
	"log"

	"github.com/vknabel/localice/internal"
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

type ResourceXmlLocalizationExporter struct {
	w io.Writer
}

func NewResourceXmlLocalizationExporter(w io.Writer) ResourceXmlLocalizationExporter {
	return ResourceXmlLocalizationExporter{w}
}

func (resourceWriter ResourceXmlLocalizationExporter) Export(localization internal.Localization) error {
	_, err := io.WriteString(resourceWriter.w, `<?xml version="1.0" encoding="utf-8"?>
`)
	if err != nil {
		return err
	}

	xmlEncoder := xml.NewEncoder(resourceWriter.w)
	xmlEncoder.Indent("", "    ")
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

	_, err = io.WriteString(resourceWriter.w, "\n")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
