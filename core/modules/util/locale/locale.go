package locale

import (
	_r "acs/domain/repository"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type locale struct {
	bundle *i18n.Bundle
}

func New() _r.Locale {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./locale/active.es.toml")
	bundle.MustLoadMessageFile("./locale/active.en.toml")
	bundle.MustLoadMessageFile("./locale/active.pt.toml")

	return &locale{
		bundle: bundle,
	}
}

func (l *locale) MustLocalize(id string, lang string) (res string) {
	localizer := i18n.NewLocalizer(l.bundle, lang)
	res = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
	})
	return
}
