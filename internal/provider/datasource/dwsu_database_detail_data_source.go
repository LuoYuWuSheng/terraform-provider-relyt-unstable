package datasource

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"terraform-provider-relyt/internal/provider/model"
)

var (
	_ datasource.DataSource              = &DwsuDatabaseDetailDataSource{}
	_ datasource.DataSourceWithConfigure = &DwsuDatabaseDetailDataSource{}
)

func NewDwsuDatabaseDetailDataSource() datasource.DataSource {
	return &DwsuDatabaseDetailDataSource{}
}

type DwsuDatabaseDetailDataSource struct {
	RelytClientDatasource
	//client *client.RelytClient
}

func (d *DwsuDatabaseDetailDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dwsu_database_detail"
}

// Schema defines the schema for the data source.
func (d *DwsuDatabaseDetailDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"auth":    model.DatasourceAuthSchema,
			"dwsu_id": schema.StringAttribute{Required: true, Description: "The ID of the service unit."},
			"id":      schema.StringAttribute{Required: true, Description: "The ID of the database."},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *DwsuDatabaseDetailDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}
