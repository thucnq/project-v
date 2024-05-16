package serrors

import (
	"context"

	"project-v/internal/pkg/ccontext"
	"project-v/pkg/container"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GetMessageI18n(ctx context.Context, message string) string {
	lang := ccontext.GetLang(ctx)
	var bundle, _ = container.Resolver[*i18n.Bundle]()
	if bundle == nil {
		return message
	}

	localizer := i18n.NewLocalizer(bundle, lang)

	msg, err := localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    message,
				Other: message,
			},
		},
	)
	if err != nil {
		msg = message
	}
	return msg
}

func MsgUserNotBelongToWorkspace(ctx context.Context) string {
	return GetMessageI18n(ctx, "MsgUserNotBelongToWorkspace")
}
