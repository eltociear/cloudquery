// Code generated by codegen; DO NOT EDIT.

package cloudiot

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func DeviceConfigs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudiot_device_configs",
		Description: `https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries.devices.configVersions#DeviceConfig`,
		Resolver:    fetchDeviceConfigs,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudiot.googleapis.com"),
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
				Name:     "device_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "cloud_update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CloudUpdateTime"),
			},
			{
				Name:     "device_ack_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("DeviceAckTime"),
			},
			{
				Name:     "binary_data",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("BinaryData"),
			},
		},
	}
}