package export

import (
	"strings"

	"golang.org/x/text/language"
)

func expandExportPath(name string, exportPath string) string {
	path := strings.ReplaceAll(exportPath, "${lowerLocale}", strings.ToLower(name))
	path = strings.ReplaceAll(path, "${upperLocale}", strings.ToUpper(name))
	path = strings.ReplaceAll(path, "${locale}", name)

	if tag, err := language.Parse(name); err == nil {
		base, _ := tag.Base()
		path = strings.ReplaceAll(path, "${base}", base.String())
		path = strings.ReplaceAll(path, "${lowerBase}", strings.ToLower(base.String()))
		path = strings.ReplaceAll(path, "${upperBase}", strings.ToUpper(base.String()))

		script, _ := tag.Script()
		path = strings.ReplaceAll(path, "${script}", script.String())
		path = strings.ReplaceAll(path, "${lowerScript}", strings.ToLower(script.String()))
		path = strings.ReplaceAll(path, "${upperScript}", strings.ToUpper(script.String()))

		region, _ := tag.Region()
		path = strings.ReplaceAll(path, "${region}", region.String())
		path = strings.ReplaceAll(path, "${lowerRegion}", strings.ToLower(region.String()))
		path = strings.ReplaceAll(path, "${upperRegion}", strings.ToUpper(region.String()))

	}
	return path
}
