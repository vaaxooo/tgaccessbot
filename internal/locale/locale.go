package locale

import (
	"encoding/json"
	"fmt"
	"os"
)

// The `Storage` struct represents a storage mechanism for
// translations organized by language and key.
// @property translations - The `translations` property in the
// `Storage` struct is a map that stores language translations.
// The keys of the outer map represent the language codes (e.g., "en"
// for English, "fr" for French), and the values are inner maps. The
// keys of the inner maps represent

type Storage struct {
	translations map[string]map[string]string
}

// The NewLocaleStorage function initializes a new Storage instance with an empty translations
// map.
func NewStorage() *Storage {
	return &Storage{translations: make(map[string]map[string]string)}
}

// The `LoadTranslations` method in the `Storage` struct is responsible for loading translations
// from JSON files for different languages.
func (ls *Storage) LoadTranslations() {
	languages := []string{"en", "ru"}

	for _, lang := range languages {
		filePath := fmt.Sprintf("locales/%s.json", lang)

		file, err := os.ReadFile(filePath)
		if err != nil {
			panic(fmt.Errorf("locale file not found: %s, error: %w", filePath, err))
		}

		var translations map[string]string

		if err := json.Unmarshal(file, &translations); err != nil {
			panic(fmt.Errorf("error unmarshalling locale %s: %w", filePath, err))
		}

		ls.translations[lang] = translations
	}
}

// The `GetText` method in the `Storage` struct is a function that retrieves a translation for a
// given language and key.
func (ls *Storage) GetText(lang, key string) string {
	if translation, exists := ls.translations[lang][key]; exists {
		return translation
	}

	if translation, exists := ls.translations["en"][key]; exists {
		return translation
	}

	return key
}
