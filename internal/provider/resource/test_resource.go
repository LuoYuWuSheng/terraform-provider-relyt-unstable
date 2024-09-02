package resource

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"terraform-provider-relyt/internal/provider/client"
	"terraform-provider-relyt/internal/provider/common"
	"time"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &testResource{}
	_ resource.ResourceWithConfigure   = &testResource{}
	_ resource.ResourceWithImportState = &testResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewTestResource() resource.Resource {
	return &testResource{}
}

type TestList struct {
	//basetypes.ObjectType
	Name types.String `tfsdk:"name"`
	//MapValue types.Map    `tfsdk:"map_value"`
}

func (t TestList) name() {

}

type TestResource struct {
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Required types.String `tfsdk:"required"`
	//mmm types.Map    `tfsdk:"mmm"`
	//Mmm      types.Map  `tfsdk:"mmm"`
	//TestList types.List `tfsdk:"self"`
	//TestList []TestList `tfsdk:"self"`
}

// orderResource is the resource implementation.
type testResource struct {
	client *client.RelytClient
}

// Metadata returns the resource type name.
func (r *testResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_test"
}

// Schema defines the schema for the resource.
func (r *testResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"id":       schema.StringAttribute{Computed: true},
			"name":     schema.StringAttribute{Optional: true},
			"required": schema.StringAttribute{Required: true},
			//"mmm": schema.MapAttribute{Computed: true, ElementType: types.StringType},
			//"self": schema.ListNestedAttribute{
			//	Computed: true,
			//	//ElementType: types.StringType,
			//	//NestedObject: schema.StringAttribute{Computed: true},
			//	NestedObject: schema.NestedAttributeObject{
			//		Attributes: map[string]schema.Attribute{"name": schema.StringAttribute{Computed: true}},
			//	},
			//},
		},
	}
}

// Create a new resource.
func (r *testResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan TestResource
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	tflog.Info(ctx, "i am here")
	//if resp.Diagnostics.HasError() {
	//	return
	//}

	//plan.Name = types.StringValue("set value")
	plan.ID = types.StringValue("set values")
	tflog.Info(ctx, "pass Diagnostics")
	resp.State.Set(ctx, plan)

	_, err := common.TimeOutTask(100000, 5, func() (any, error) {
		time.Sleep(1)
		return nil, fmt.Errorf("mock apiFail!")
	})
	if err != nil {
		resp.Diagnostics.AddError("error wait!", err.Error())
		return
	}
	//一旦拿到ID立刻保存
	//objectType := types.ObjectType{
	//	map[string]attr.Type{
	//		"name": types.StringType,
	//	},
	//}
	//objectValues := []TestList{{Name: types.StringValue("abc")}}
	//from, diagnostics := types.ListValueFrom(ctx, objectType, objectValues)
	//resp.Diagnostics.Append(diagnostics...)
	resp.Diagnostics.AddError("err", "err")
	//diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *testResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state TestResource
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Required = types.StringValue("manual set when import")

	//_, err := r.client.GetDwsu(ctx, state.ID.ValueString())
	//if err != nil {
	//	tflog.Error(ctx, "error read dwsu"+err.Error())
	//	return
	//}
	//r.mapRelytModelToTerraform(ctx, &resp.Diagnostics, &state, &model)
	//state.Status = types.StringValue(dwsu.Status)
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "read dwsu succ : "+state.ID.ValueString())
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *testResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("not support", "update dwsu not supported")
	return
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *testResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state TestResource
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	//resp.Diagnostics.AddError("err to delete", "error to delete")
	return
}

// Configure adds the provider configured client to the resource.
func (r *testResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	relytClient, ok := req.ProviderData.(*client.RelytClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *RelytClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = relytClient
}

func (r *testResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// State upgrade implementation from 0 (prior state version) to 2 (Schema.Version)
		0: {
			// Optionally, the PriorSchema field can be defined.
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) { /* ... */
			},
		},
		// State upgrade implementation from 1 (prior state version) to 2 (Schema.Version)
		1: {
			// Optionally, the PriorSchema field can be defined.
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) { /* ... */
			},
		},
	}
}

func (r *testResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
