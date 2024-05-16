package intelisysclient

import (
	"context"
	"net/http"
	"net/url"
	"project-v/internal/model/agent"
	"project-v/internal/repository/intelisys"
	"project-v/pkg/httprequest/v2"
	"time"

	serrors "project-v/internal/errors"
	"project-v/pkg/l"
)

const (
	endGetReservations = "reservations?reservationLocator={pnr}"
	endReservations    = "reservations"
)

type Config struct {
	BaseUrl string `json:"base_url" mapstructure:"base_url"`
	APIKey  string `json:"api_key" mapstructure:"api_key"`
}

type ClientImpl struct {
	cfg     Config
	debug   bool
	log     l.Logger
	request *httprequest.HttpClient
}

func NewClientImpl(cfg Config, ll l.Logger) intelisys.IRepository {
	httpClient := httprequest.NewClient()
	return &ClientImpl{
		cfg:     cfg,
		log:     ll,
		request: httpClient,
	}
}

func (r *ClientImpl) GetReservations(
	ctx context.Context, pnr string,
) (*agent.Reservation, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	rc := httprequest.NewRestClient(r.request)
	finalURL, err := httprequest.ResolveURL(r.cfg.BaseUrl, endReservations)
	if err != nil {
		return nil, serrors.ErrInternal(ctx, err)
	}

	// Query params
	q := &url.Values{}
	q.Set("reservationLocator", pnr)

	out := &struct {
		Data *agent.Reservation `json:"data"`
	}{}
	// ll := l.New()
	err = rc.NewRequest().
		WithContext(ctx).
		// Debug(trace.NewTraceLog(ll)).
		WithQuery(q).
		Get(finalURL).
		MustHaveStatus(http.StatusOK).
		Json(out).
		Error()
	if err != nil {
		return nil, err
	}

	return out.Data, nil
}
