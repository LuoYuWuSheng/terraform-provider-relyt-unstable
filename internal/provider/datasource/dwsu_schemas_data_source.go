package datasource

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-relyt/internal/provider/model"
)

var (
	_ datasource.DataSource              = &DwsuSchemasDataSource{}
	_ datasource.DataSourceWithConfigure = &DwsuSchemasDataSource{}
)

func NewDwsuSchemasDataSource() datasource.DataSource {
	return &DwsuSchemasDataSource{}
}

type DwsuSchemasDataSource struct {
	RelytClientDatasource
	//client *client.RelytClient
}

func (d *DwsuSchemasDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dwsu_schemas"
}

// Schema defines the schema for the data source.
func (d *DwsuSchemasDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"auth":          model.DatasourceAuthSchema,
			"dwsu_id":       schema.StringAttribute{Required: true, Description: "The ID of the service unit."},
			"database_name": schema.StringAttribute{Required: true, Description: "The database name of the schema."},
			"ids":           schema.ListAttribute{Computed: true, ElementType: types.StringType, Description: "The ids of schema."},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *DwsuSchemasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state model.DwsuSchemas
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	stringList := []string{"demoSchema"}
	values, diags := types.ListValueFrom(ctx, types.StringType, stringList)
	resp.Diagnostics.Append(diags...)
	state.IDs = values
	resp.State.Set(ctx, state)
}
