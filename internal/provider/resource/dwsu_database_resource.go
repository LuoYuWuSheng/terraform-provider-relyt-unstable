package resource

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-relyt/internal/provider/model"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &DwsuDatabaseResource{}
	_ resource.ResourceWithConfigure   = &DwsuDatabaseResource{}
	_ resource.ResourceWithImportState = &DwsuDatabaseResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewDwsuDatabaseResource() resource.Resource {
	return &DwsuDatabaseResource{}
}

// orderResource is the resource implementation.
type DwsuDatabaseResource struct {
	RelytClientResource
}

// Metadata returns the resource type name.
func (r *DwsuDatabaseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dwsu_database"
}

// Schema defines the schema for the resource.
func (r *DwsuDatabaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"auth":    model.ResourceAuthSchema,
			"dwsu_id": schema.StringAttribute{Required: true, Description: "The ID of the service unit."},
			"name":    schema.StringAttribute{Required: true, Description: "The Name of the database."},
			"id":      schema.StringAttribute{Computed: true, Description: "The ID of the database."},
			"owner":   schema.StringAttribute{Computed: true, Description: "The owner of the database."},
		},
	}
}

// Create a new resource.
func (r *DwsuDatabaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	database := model.DwsuDatabase{}
	diags := req.Plan.Get(ctx, &database)
	database.ID = database.Name
	database.Owner = types.StringValue("mockOwner")
	resp.Diagnostics.Append(diags...)
	resp.State.Set(ctx, database)
}

// Read resource information.
func (r *DwsuDatabaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *DwsuDatabaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *DwsuDatabaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *DwsuDatabaseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
