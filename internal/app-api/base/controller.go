package base

import (
	"context"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"project-v/internal/app-api/response"
	"project-v/internal/pkg/ccontext"
	"project-v/pkg/container"
)

// Controller ...
type Controller struct {
}

// Resp ...
func (Controller) Resp() response.IResponse {
	return response.NewResponse()
}

func (Controller) GetMsgI18n(ctx context.Context, key string) string {
	lang := ccontext.GetLang(ctx)
	var bundle *i18n.Bundle
	var err error
	_ = container.ResolverMust[i18n.Bundle]
	msg := key
	if bundle == nil {
		return key
	}

	localizer := i18n.NewLocalizer(bundle, lang)
	msg, err = localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    key,
				Other: key,
			},
		},
	)
	if err != nil {
		msg = key
	}
	return msg
}
