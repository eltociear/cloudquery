package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"google.golang.org/api/analyticsreporting/v4"
	"google.golang.org/api/option"
)

type Client struct {
	*analyticsreporting.Service
	backend.Backend

	accountID string
	profile   string

	logger zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return "google-analytics:account:{" + c.accountID + "}:profile:{" + c.profile + "}"
}

func Configure(ctx context.Context, logger zerolog.Logger, srcSpec specs.Source, options source.Options) (schema.ClientMeta, error) {
	var spec Spec
	if err := srcSpec.UnmarshalSpec(&spec); err != nil {
		return nil, err
	}

	opts := []option.ClientOption{
		option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the same side with telemetry
		option.WithTelemetryDisabled(),
	}
	if len(spec.APIKey) > 0 {
		opts = append(opts, option.WithAPIKey(spec.APIKey))
	}

	svc, err := analyticsreporting.NewService(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	svc.UserAgent = "cloudquery:source-google-analytics/" + srcSpec.Version

	c := &Client{
		Service:   svc,
		Backend:   options.Backend,
		accountID: "", //TODO
		profile:   "", //TODO
		logger:    logger.With().Str("plugin", "google-analytics").Logger(),
	}

	return c, nil
}