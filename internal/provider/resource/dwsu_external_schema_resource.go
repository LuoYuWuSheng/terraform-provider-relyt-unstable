package resource

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"terraform-provider-relyt/internal/provider/model"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &DwsuExternalSchemaResource{}
	_ resource.ResourceWithConfigure   = &DwsuExternalSchemaResource{}
	_ resource.ResourceWithImportState = &DwsuExternalSchemaResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewDwsuExternalSchemaResource() resource.Resource {
	return &DwsuExternalSchemaResource{}
}

// orderResource is the resource implementation.
type DwsuExternalSchemaResource struct {
	RelytClientResource
}

// Metadata returns the resource type name.
func (r *DwsuExternalSchemaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dwsu_external_schema"
}

// Schema defines the schema for the resource.
func (r *DwsuExternalSchemaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"auth":          model.ResourceAuthSchema,
			"dwsu_id":       schema.StringAttribute{Required: true, Description: "The ID of the service unit."},
			"name":          schema.StringAttribute{Required: true, Description: "The Name of the schema."},
			"database_name": schema.StringAttribute{Required: true, Description: "The Name of the database."},
			"id":            schema.StringAttribute{Computed: true, Description: "The ID of the database."},

			"properties": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"metastore_type":           schema.StringAttribute{Computed: true, Optional: true, Description: "metastore_type", Default: stringdefault.StaticString("Glue")},
					"table_format":             schema.StringAttribute{Computed: true, Optional: true, Description: "table_format", Default: stringdefault.StaticString("DELTA")},
					"glue_access_control_mode": schema.StringAttribute{Computed: true, Optional: true, Description: "glue_access_control_mode", Default: stringdefault.StaticString("Lake Formation")},
					"glue_region":              schema.StringAttribute{Computed: true, Optional: true, Description: "glue_region", Default: stringdefault.StaticString("ap-east-1")},
					"s3_region":                schema.StringAttribute{Computed: true, Optional: true, Description: "s3_region", Default: stringdefault.StaticString("ap-east-1")},
				},
				Description: "The properties of the schema."},
		},
	}
}

// Create a new resource.
func (r *DwsuExternalSchemaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	externalSchema := model.DwsuExternalSchema{}
	diags := req.Plan.Get(ctx, &externalSchema)
	resp.Diagnostics.Append(diags...)
	externalSchema.ID = externalSchema.Name
	resp.State.Set(ctx, externalSchema)
}

// Read resource information.
func (r *DwsuExternalSchemaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *DwsuExternalSchemaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *DwsuExternalSchemaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *DwsuExternalSchemaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
