package locale

import (
	"encoding/json"
	"fmt"
	"os"
)

// Locales - map of translations
var Locales = make(map[string]map[string]string)

// LoadTranslations loads translations from files
func LoadTranslations() error {
	languages := []string{"en", "ru"}

	for _, lang := range languages {
		filePath := fmt.Sprintf("locales/%s.json", lang)
		file, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("not found locale %s: %w", filePath, err)
		}

		var translations map[string]string
		if err := json.Unmarshal(file, &translations); err != nil {
			return fmt.Errorf("error while unmarshalling locale %s: %w", filePath, err)
		}

		Locales[lang] = translations
	}

	return nil
}

// GetText returns translation by key
func GetText(lang, key string) string {
	if translation, exists := Locales[lang][key]; exists {
		return translation
	}
	if translation, exists := Locales["en"][key]; exists {
		return translation
	}
	return key
}
