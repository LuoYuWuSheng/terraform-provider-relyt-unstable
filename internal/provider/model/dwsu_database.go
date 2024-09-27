package model

import (
	resSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	datasourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Auth struct {
	AccessKey types.String `tfsdk:"access_key"`
	SecretKey types.String `tfsdk:"secret_key"`
}

type DwsuDatabase struct {
	Auth   Auth         `tfsdk:"auth"`
	DwsuId types.String `tfsdk:"dwsu_id"`
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Owner  types.String `tfsdk:"owner"`
}

type DwsuExternalSchema struct {
	Auth         Auth                     `tfsdk:"auth"`
	ID           types.String             `tfsdk:"id"`
	DwsuId       types.String             `tfsdk:"dwsu_id"`
	Name         types.String             `tfsdk:"name"`
	DatabaseName types.String             `tfsdk:"database_name"`
	Properties   ExternalSchemaProperties `tfsdk:"properties"`
}

type ExternalSchemaProperties struct {
	MetastoreType         types.String `tfsdk:"metastore_type"`
	TableFormat           types.String `tfsdk:"table_format"`
	GlueAccessControlMode types.String `tfsdk:"glue_access_control_mode"`
	GlueRegion            types.String `tfsdk:"glue_region"`
	S3Region              types.String `tfsdk:"s3_region"`
}

type DwsuDatabases struct {
	Auth   Auth         `tfsdk:"auth"`
	DwsuId types.String `tfsdk:"dwsu_id"`
	IDs    types.List   `tfsdk:"ids"`
}

type DwsuDatabaseDetail struct {
	Auth   Auth         `tfsdk:"auth"`
	DwsuId types.String `tfsdk:"dwsu_id"`
	ID     types.List   `tfsdk:"id"`
}

type DwsuSchemas struct {
	Auth         Auth         `tfsdk:"auth"`
	DwsuId       types.String `tfsdk:"dwsu_id"`
	DatabaseName types.String `tfsdk:"database_name"`
	IDs          types.List   `tfsdk:"ids"`
}
type DwsuSchemaDetail struct {
	Auth         Auth         `tfsdk:"auth"`
	DwsuId       types.String `tfsdk:"dwsu_id"`
	DatabaseName types.String `tfsdk:"database_name"`
	ID           types.List   `tfsdk:"id"`
}

var (
	ResourceAuthSchema = resSchema.SingleNestedAttribute{
		Required: true,
		Attributes: map[string]resSchema.Attribute{
			"access_key": resSchema.StringAttribute{Required: true, Description: "access key"},
			"secret_key": resSchema.StringAttribute{Required: true, Description: "secret key"},
		},
		Description: "The Auth AccessKey And SecretKey.",
	}
	DatasourceAuthSchema = datasourceSchema.SingleNestedAttribute{
		Required: true,
		Attributes: map[string]datasourceSchema.Attribute{
			"access_key": datasourceSchema.StringAttribute{Required: true, Description: "access key"},
			"secret_key": datasourceSchema.StringAttribute{Required: true, Description: "secret key"},
		},
		Description: "The Auth AccessKey And SecretKey.",
	}
)
