// Code generated by codegen; DO NOT EDIT.

package artifactregistry

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/artifactregistry/apiv1"
)

func Files() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_files",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.files#File`,
		Resolver:    fetchFiles,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "size_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SizeBytes"),
			},
			{
				Name:     "hashes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Hashes"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:     "owner",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner"),
			},
		},
	}
}

func fetchFiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListFilesRequest{
		Parent: parent.Item.(*pb.Repository).Name,
	}
	gcpClient, err := artifactregistry.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListFiles(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}