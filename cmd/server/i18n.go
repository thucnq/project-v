package main

import (
	"encoding/json"

	"golang.org/x/text/language"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"project-v/pkg/container"
	"project-v/pkg/l"
)

func loadI18n() {
	log, _ := container.Resolver[l.Logger]()

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	_, err := bundle.LoadMessageFile("./i18n/en.json")
	if err != nil {
		log.Fatal("failed to load language", l.Error(err))
		return
	}

	if _, err := bundle.LoadMessageFile("./i18n/vi.json"); err != nil {
		log.Fatal("failed to load language", l.Error(err))
		return
	}

	container.Register(
		func() *i18n.Bundle {
			return bundle
		},
	)
}
