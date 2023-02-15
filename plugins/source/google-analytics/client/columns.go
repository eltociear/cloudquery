package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	AccountIDColumn = schema.Column{
		Name:        "account_id",
		Type:        schema.TypeString,
		Description: "Account ID",
		Resolver:    ResolveAccountID,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: true,
			NotNull:    true,
		},
	}
	WebPropertyIDColumn = schema.Column{
		Name:        "web_property_id",
		Type:        schema.TypeString,
		Description: "Web property ID",
		Resolver:    ResolveWebPropertyID,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: true,
			NotNull:    true,
		},
	}
)
