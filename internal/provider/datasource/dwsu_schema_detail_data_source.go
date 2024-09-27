package datasource

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"terraform-provider-relyt/internal/provider/model"
)

var (
	_ datasource.DataSource              = &DwsuSchemaDetailDataSource{}
	_ datasource.DataSourceWithConfigure = &DwsuSchemaDetailDataSource{}
)

func NewDwsuSchemaDetailDataSource() datasource.DataSource {
	return &DwsuSchemaDetailDataSource{}
}

type DwsuSchemaDetailDataSource struct {
	RelytClientDatasource
	//client *client.RelytClient
}

func (d *DwsuSchemaDetailDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dwsu_schema_detail"
}

// Schema defines the schema for the data source.
func (d *DwsuSchemaDetailDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"auth":          model.DatasourceAuthSchema,
			"dwsu_id":       schema.StringAttribute{Required: true, Description: "The ID of the service unit."},
			"database_name": schema.StringAttribute{Required: true, Description: "The name of the database."},
			"id":            schema.StringAttribute{Required: true, Description: "The ID of the schema."},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *DwsuSchemaDetailDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

}
