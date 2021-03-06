# localice

Generate `Localizable.strings`, `InfoPlist.strings` and `strings.xml` from multiple CSV files.

> **Attention:** This project is in early development and not yet meant to be used in production.
> It lacks validation, tests and might break configs in future versions. Also: I am still learning Go. 

## Getting Started

Create a `.localice.json` file in your project root.

```json
{
    "locales": ["de", "en"],
    "csvSources": [
        {
            "location": "input.csv",
            "keys": "Key",
            "platforms": "Platform",
            "localeColumns": {
                "DE": "de",
                "EN": "en"
            }
        },
         {
            "location": "secrets.csv",
            "keys": "Key",
            "platforms": "platform",
            "localeColumns": {
                "DE": "de",
                "EN": "en"
            }
        }
    ],
    "exports": [
        {
            "format": "resource-xml",
            "path": "android/${resourceLocale}.xml",
            "matchPlatform": "^(Android)?$"
        },
        {
            "format": "strings",
            "path": "${lowerLocale}.lproj/Localizable.strings",
            "matchPlatform": "^(iOS)?$",
            "matchKey": "^[^A-Z]"
        },
         {
            "format": "strings",
            "path": "${lowerLocale}.lproj/InfoPlist.strings",
            "matchPlatform": "^(iOS|InfoPlist)?$",
            "matchKey": "^[A-Z]{2,}"
        }
    ]
}
```

## Installation

```bash
brew install vknabel/install/localice
```

## Future Development

* improved logging feedback
* parsing json sources
* Better and multiple examples
* Documentation
* Test coverage
* code refactorings
* `localice init` command
* Parallelization

## License

localice is available under the [MIT](./LICENSE) license.
