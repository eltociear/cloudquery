// Code generated by codegen; DO NOT EDIT.

package organization

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "snyk_organizations",
		Description: `https://pkg.go.dev/github.com/pavel-snyk/snyk-sdk-go/snyk#Organization`,
		Resolver:    fetchOrganizations,
		Multiplex:   client.SingleOrganization,
		Columns: []schema.Column{
			{
				Name:     "group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Group"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "slug",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Slug"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
		},
	}
}