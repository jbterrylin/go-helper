package translatehelper

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Translator struct {
	DefaultLang string // must be one of TranslatorBundle's key
	Localizers  map[string]*i18n.Localizer
}

type TranslatorBundle struct {
	Key           string // e.x: en, zh...
	FilePath      string // file path include file name
	LangTag       language.Tag
	Format        string             // e.x: json
	UnmarshalFunc i18n.UnmarshalFunc // e.x: json.Unmarshal
}

func NewTranslator(defaultLang string, translatorBundles []TranslatorBundle) (translator Translator, err error) {
	if len(translatorBundles) == 0 {
		return
	}
	translator.DefaultLang = defaultLang
	if defaultLang == "" {
		translator.DefaultLang = translatorBundles[0].Key
	}
	translator.Localizers = make(map[string]*i18n.Localizer)

	for _, translatorBundle := range translatorBundles {
		bundle := i18n.NewBundle(translatorBundle.LangTag)
		bundle.RegisterUnmarshalFunc(translatorBundle.Format, translatorBundle.UnmarshalFunc)
		_, err = bundle.LoadMessageFile(translatorBundle.FilePath)
		if err != nil {
			return
		}

		translator.Localizers[translatorBundle.Key] = i18n.NewLocalizer(bundle, translatorBundle.Key)
	}
	return
}

func (t *Translator) TranslateMessage(messageID, language string) (translatedMsg string) {
	if language == "" {
		language = t.DefaultLang
	}

	translatedMsg, err := t.Localizers[language].LocalizeMessage(&i18n.Message{ID: messageID})
	if err != nil {
		return messageID
	}
	return
}

func (t *Translator) TranslateMessageWithVars(messageID string, vars map[string]string, language string) (translatedMsg string) {
	if language == "" {
		language = t.DefaultLang
	}

	translatedMsg, err := t.Localizers[language].Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: vars,
	})
	if err != nil {
		return messageID
	}

	return translatedMsg
}
