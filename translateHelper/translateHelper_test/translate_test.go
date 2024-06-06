package translatehelper_test

import (
	"encoding/json"
	"testing"

	translatehelper "github.com/jbterrylin/go-helper/translateHelper"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TestNewTranslator(t *testing.T) {
	translatorBundles := []translatehelper.TranslatorBundle{
		{
			Key:           "en",
			FilePath:      "test.en.json",
			LangTag:       language.English,
			Format:        "json",
			UnmarshalFunc: i18n.UnmarshalFunc(json.Unmarshal),
		},
		{
			Key:           "zh",
			FilePath:      "test.zh.json",
			LangTag:       language.Chinese,
			Format:        "json",
			UnmarshalFunc: i18n.UnmarshalFunc(json.Unmarshal),
		},
	}

	translator, err := translatehelper.NewTranslator("en", translatorBundles)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if translator.DefaultLang != "en" {
		t.Errorf("Expected default language 'en', got %v", translator.DefaultLang)
	}

	if len(translator.Localizers) != 2 {
		t.Errorf("Expected 2 localizers, got %d", len(translator.Localizers))
	}
}

func TestTranslateMessage(t *testing.T) {
	translatorBundles := []translatehelper.TranslatorBundle{
		{
			Key:           "en",
			FilePath:      "test.en.json",
			LangTag:       language.English,
			Format:        "json",
			UnmarshalFunc: i18n.UnmarshalFunc(json.Unmarshal),
		},
		{
			Key:           "zh",
			FilePath:      "test.zh.json",
			LangTag:       language.Chinese,
			Format:        "json",
			UnmarshalFunc: i18n.UnmarshalFunc(json.Unmarshal),
		},
	}

	translator, err := translatehelper.NewTranslator("en", translatorBundles)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	translatedMsg := translator.TranslateMessage("hello", "en")
	expectedMsg := "Hello" // This should match the content of active.en.json
	if translatedMsg != expectedMsg {
		t.Errorf("Expected %v, got %v", expectedMsg, translatedMsg)
	}

	translatedMsg = translator.TranslateMessage("hello", "zh")
	expectedMsg = "你好" // This should match the content of active.zh.json
	if translatedMsg != expectedMsg {
		t.Errorf("Expected %v, got %v", expectedMsg, translatedMsg)
	}
}

func TestTranslateMessageWithVars(t *testing.T) {
	translatorBundles := []translatehelper.TranslatorBundle{
		{
			Key:           "en",
			FilePath:      "test.en.json",
			LangTag:       language.English,
			Format:        "json",
			UnmarshalFunc: i18n.UnmarshalFunc(json.Unmarshal),
		},
		{
			Key:           "zh",
			FilePath:      "test.zh.json",
			LangTag:       language.Chinese,
			Format:        "json",
			UnmarshalFunc: i18n.UnmarshalFunc(json.Unmarshal),
		},
	}

	translator, err := translatehelper.NewTranslator("en", translatorBundles)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	vars := map[string]string{
		"Name": "World",
	}

	translatedMsg := translator.TranslateMessageWithVars("greeting", vars, "en")
	expectedMsg := "Hello, World" // This should match the content of active.en.json with variables replaced
	if translatedMsg != expectedMsg {
		t.Errorf("Expected %v, got %v", expectedMsg, translatedMsg)
	}

	translatedMsg = translator.TranslateMessageWithVars("greeting", vars, "zh")
	expectedMsg = "你好, World" // This should match the content of active.zh.json with variables replaced
	if translatedMsg != expectedMsg {
		t.Errorf("Expected %v, got %v", expectedMsg, translatedMsg)
	}
}
