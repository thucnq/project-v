package iamclient

import (
	"context"
	"net/http"
	"time"

	"project-v/internal/constant"
	serrors "project-v/internal/errors"
)

// LoadServices ...
func (o *repositoryImpl) loadServices() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rc := httprequest.NewRestClient(o.req)

	finalURL, err := httprequest.ResolveURL(
		o.cfg.BaseUrl, "/service-keys",
	)
	if err != nil {
		return serrors.ErrInternal(ctx, err)
	}
	out := &struct {
		Data []*ServiceInfo `json:"data"`
	}{}

	err = rc.NewRequest().
		WithContext(ctx).
		AddHeaders(constant.XGapoRole, xgaporole.Service).
		AddHeaders(constant.XGapoApiKey, o.cfg.SecretKey).
		AddHeaders(constant.HeaderContentType, constant.MIMEApplicationJson).
		Get(finalURL).
		MustHaveStatus(http.StatusOK).
		Json(out).
		Error()
	if err != nil {
		return err
	}

	o.listServices = out.Data

	// fmt.Printf("%+v", o.listServices)
	o.mapServiceKey = make(map[string]bool)
	for _, item := range out.Data {
		o.mapServiceKey[item.APIKey] = true
	}

	return nil
}

// IsAllow ...
func (o *repositoryImpl) IsAllow(apiKey string) bool {
	if o.mapServiceKey == nil {
		return false
	}
	if len(o.cfg.DebugKey) > 0 && o.cfg.DebugKey == apiKey {
		return true
	}
	return o.mapServiceKey[apiKey]
}
