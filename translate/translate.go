package translate

import (
	"context"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type Translator interface {
	Translate(targetLanguage string, text string) (string, error)
}

type GoogleTranslator struct {
	client *translate.Client
}

func NewGoogleTranslator(apiKey string) *GoogleTranslator {
	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	return &GoogleTranslator{
		client: client,
	}
}

func (t GoogleTranslator) Translate(targetLanguage string, text string) (string, error) {
	ctx := context.Background()
	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", err
	}
	res, err := t.client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", err
	}
	return res[0].Text, nil
}
