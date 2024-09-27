package datasource

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-relyt/internal/provider/model"
)

var (
	_ datasource.DataSource              = &DwsuDatabasesDataSource{}
	_ datasource.DataSourceWithConfigure = &DwsuDatabasesDataSource{}
)

func NewDwsuDatabasesDataSource() datasource.DataSource {
	return &DwsuDatabasesDataSource{}
}

type DwsuDatabasesDataSource struct {
	RelytClientDatasource
}

func (d *DwsuDatabasesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dwsu_databases"
}

// Schema defines the schema for the data source.
func (d *DwsuDatabasesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"auth":    model.DatasourceAuthSchema,
			"dwsu_id": schema.StringAttribute{Required: true, Description: "The ID of the service unit."},
			"ids":     schema.ListAttribute{Computed: true, ElementType: types.StringType, Description: "The ids of database."},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *DwsuDatabasesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state model.DwsuDatabases
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	stringList := []string{"demoDatabase"}
	values, diags := types.ListValueFrom(ctx, types.StringType, stringList)
	resp.Diagnostics.Append(diags...)
	state.IDs = values
	resp.State.Set(ctx, state)
}
