package locale

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Provider interface {
	GetText(lang, key string) string
}

// The `Storage` struct is a storage for translations.
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
		filePath := filepath.Join("assets", "locales", lang+".json")

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
	if translations, exists := ls.translations[lang]; exists {
		if translation, ok := translations[key]; ok {
			return translation
		}
	}

	if translations, exists := ls.translations["en"]; exists {
		if translation, ok := translations[key]; ok {
			return translation
		}
	}

	return key
}
