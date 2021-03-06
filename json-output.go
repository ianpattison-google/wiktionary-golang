package wiktionary

import (
	"encoding/json"
	"os"
)

type LanguageWord struct {
	Word           string      `json:"word"`
	Meaning        string      `json:"meaning,omitempty"`
	LanguageName   string      `json:"lang"`
	LanguageCode   string      `json:"lang-code"`
	Pronunciations []string    `json:"pron,omitempty"`
	Ipa            string      `json:"ipa,omitempty"`
	Etymologies    []Etymology `json:"etym,omitempty"`
	Anagrams       string      `json:"anag,omitempty"`
}

type Etymology struct {
	Name             string         `json:"name"`
	AlternativeForms string         `json:"alts,omitempty"`
	Text             string         `json:"text,omitempty"`
	Words            []LinkedWord   `json:"words,omitempty"`
	Parts            []PartOfSpeech `json:"parts,omitempty"`
	Pronunciations   []string       `json:"pron,omitempty"`
	Ipa              string         `json:"ipa,omitempty"`
}

// relationships for LinkedWord
const (
	Root       string = "root"
	Inherited  string = "inherited"
	Cognate    string = "cognate"
	Descendant string = "descendant"
)

type LinkedWord struct {
	Relationship    string          `json:"type"`
	Language        string          `json:"lang"`
	Word            string          `json:"word"`
	Meaning         string          `json:"meaning,omitempty"`
	Transliteration string          `json:"translit,omitempty"`
	Attributes      map[string]bool `json:"attrs,omitempty"`
}

type TranslatedWord struct {
	Language        string `json:"lang"`
	Word            string `json:"word"`
	Transliteration string `json:"translit,omitempty"`
}

type PartOfSpeech struct {
	Name         string            `json:"name"`
	Headword     string            `json:"head,omitempty"`
	Attributes   map[string]string `json:"attrs,omitempty"`
	Meanings     []string          `json:"meanings,omitempty"`
	Translations []TranslatedWord  `json:"trans,omitempty"`
	Synonyms     string            `json:"syn,omitempty"`
	Antonyms     string            `json:"ant,omitempty"`
}

func writeJson(word string, langCode string, lw *LanguageWord) error {
	b, err := json.Marshal(lw)
	if err != nil {
		return err
	}
	// write file
	fileName := langCode + "-" + word + ".json"
	errf := os.WriteFile(fileName, b, 0666)
	if errf != nil {
		return errf
	}

	return nil
}
