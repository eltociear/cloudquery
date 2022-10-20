// Code generated by codegen; DO NOT EDIT.

package docdb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterSnapshotAttributes() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_cluster_snapshot_attributes",
		Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterSnapshotAttributesResult.html",
		Resolver:    fetchDocdbClusterSnapshotAttributes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Columns: []schema.Column{
			{
				Name:     "db_cluster_snapshot_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "db_cluster_snapshot_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DBClusterSnapshotAttributes"),
			},
			{
				Name:     "db_cluster_snapshot_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterSnapshotIdentifier"),
			},
		},
	}
}
