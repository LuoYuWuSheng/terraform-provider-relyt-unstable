package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"terraform-provider-relyt/internal/provider/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &dwUserResource{}
	_ resource.ResourceWithConfigure   = &dwUserResource{}
	_ resource.ResourceWithImportState = &dwUserResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewdwUserResource() resource.Resource {
	return &dwUserResource{}
}

// orderResource is the resource implementation.
type dwUserResource struct {
	client *client.RelytClient
}

// Metadata returns the resource type name.
func (r *dwUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dwuser"
}

// Schema defines the schema for the resource.
func (r *dwUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"dwsu_id":                             schema.StringAttribute{Required: true, Description: "The ID of the service unit."},
			"id":                                  schema.StringAttribute{Computed: true, Description: "The ID of the DW user."},
			"account_name":                        schema.StringAttribute{Required: true, Description: "The name of the DW user, which is unique in the instance. The name is the email address."},
			"account_password":                    schema.StringAttribute{Required: true, Description: "initPassword"},
			"datalake_aws_lakeformation_role_arn": schema.StringAttribute{Optional: true, Computed: true, Description: ""},
			"async_query_result_location_prefix":  schema.StringAttribute{Optional: true, Computed: true},
			"async_query_result_location_aws_role_arn": schema.StringAttribute{Optional: true, Computed: true},
		},
	}
}

// Create a new resource.
func (r *dwUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from dwUserModel
	var dwUserModel DWUserModel
	diags := req.Plan.Get(ctx, &dwUserModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	meta := RouteRegionUri(ctx, dwUserModel.DwsuId.ValueString(), r.client, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	regionUri := meta.URI
	relytAccount := client.Account{
		InitPassword: dwUserModel.AccountPassword.ValueString(),
		Name:         dwUserModel.AccountName.ValueString(),
	}
	// Create new order
	createResult, err := r.client.CreateAccount(ctx, regionUri, dwUserModel.DwsuId.ValueString(), relytAccount)
	if err != nil || createResult.Code != 200 {
		resp.Diagnostics.AddError(
			"Error creating dwuser",
			"Could not create dwuser, unexpected error: "+err.Error(),
		)
		return
	}
	dwUserModel.ID = types.StringValue(relytAccount.Name)
	r.handleAccountConfig(ctx, &dwUserModel, regionUri, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		err := r.client.DropAccount(ctx, regionUri, dwUserModel.DwsuId.ValueString(), dwUserModel.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error rollback create dwuser",
				"Could not rollback dwuser! please clear it with destroy or manual! userId: "+dwUserModel.ID.ValueString()+""+err.Error(),
			)
		}
	}
	if resp.Diagnostics.HasError() {
		//如果有异常，dwuser不要写状态
		return
	}
	diags = resp.State.Set(ctx, dwUserModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *dwUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state DWUserModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if state.ID.IsNull() {
		resp.Diagnostics.AddError("can't read dwuser", "dwuser id is nil")
		return
	}
	//state.ID = state.AccountName
	meta := RouteRegionUri(ctx, state.DwsuId.ValueString(), r.client, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	config, err := r.client.GetAsyncAccountConfig(ctx, meta.URI, state.DwsuId.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error get dwuser asyncAccountConfig",
			"Could not get config asyncAccountConfig, unexpected error: "+err.Error(),
		)
		return
	}
	if config != nil {
		state.AsyncQueryResultLocationPrefix = types.StringValue(config.S3LocationPrefix)
		state.AsyncQueryResultLocationAwsRoleArn = types.StringValue(config.AwsIamArn)

	}

	lakeInfo, err := r.client.GetLakeFormationConfig(ctx, meta.URI, state.DwsuId.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error get dwuser LakeFormationConfig",
			"Could not get config LakeFormationConfig, unexpected error: "+err.Error(),
		)
		return
	}
	if lakeInfo != nil {
		state.DatalakeAwsLakeformationRoleArn = types.StringValue(lakeInfo.IAMRole)
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	return
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *dwUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//resp.Diagnostics.AddError("not support", "update account not supported")
	var plan DWUserModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// read old status
	var stat DWUserModel
	req.State.Get(ctx, &stat)
	if resp.Diagnostics.HasError() {
		return
	}
	if stat.AccountName.ValueString() != plan.AccountName.ValueString() {
		resp.Diagnostics.AddError("not support", "can't update account name!")
	}
	if stat.AccountPassword.ValueString() != plan.AccountPassword.ValueString() {
		resp.Diagnostics.AddError("not support", "can't update init password!")
	}
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = plan.AccountName
	meta := RouteRegionUri(ctx, plan.DwsuId.ValueString(), r.client, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	regionUri := meta.URI
	tflog.Info(ctx, "accountId:"+plan.ID.ValueString())
	r.handleAccountConfig(ctx, &plan, regionUri, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	return
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *dwUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state DWUserModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	meta := RouteRegionUri(ctx, state.DwsuId.ValueString(), r.client, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	regionUri := meta.URI

	// Delete existing order
	err := r.client.DropAccount(ctx, regionUri, state.DwsuId.ValueString(), state.ID.ValueString())
	if err != nil {
		//要不要加error
		resp.Diagnostics.AddError(
			"Error Deleting dwuser",
			"Could not delete dwuser, unexpected error: "+err.Error(),
		)
	}
}

// Configure adds the provider configured client to the resource.
func (r *dwUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	relytClient, ok := req.ProviderData.(*client.RelytClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *RelytClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = relytClient
}

func (r *dwUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *dwUserResource) handleAccountConfig(ctx context.Context, dwUserModel *DWUserModel, regionUri string, diagnostics *diag.Diagnostics) {
	//dwUserModel.ID = dwUserModel.AccountName
	asyncResult := client.AsyncResult{
		AwsIamArn:        dwUserModel.AsyncQueryResultLocationAwsRoleArn.ValueString(),
		S3LocationPrefix: dwUserModel.AsyncQueryResultLocationPrefix.ValueString(),
	}
	lakeFormation := client.LakeFormation{
		IAMRole: dwUserModel.DatalakeAwsLakeformationRoleArn.ValueString(),
	}
	tflog.Info(ctx, fmt.Sprintf("=======uknown %t nil %t", dwUserModel.AsyncQueryResultLocationAwsRoleArn.IsUnknown(), dwUserModel.AsyncQueryResultLocationAwsRoleArn.IsNull()))
	if dwUserModel.AsyncQueryResultLocationPrefix.IsUnknown() {
		dwUserModel.AsyncQueryResultLocationPrefix = types.StringNull()
	}
	if dwUserModel.AsyncQueryResultLocationAwsRoleArn.IsUnknown() {
		dwUserModel.AsyncQueryResultLocationAwsRoleArn = types.StringNull()
	}
	if !dwUserModel.AsyncQueryResultLocationAwsRoleArn.IsNull() && !dwUserModel.AsyncQueryResultLocationPrefix.IsNull() {
		_, err := r.client.AsyncAccountConfig(ctx, regionUri, dwUserModel.DwsuId.ValueString(), dwUserModel.ID.ValueString(), asyncResult)
		if err != nil {
			diagnostics.AddError(
				"Error config dwuser",
				"Could not config dwuser async, unexpected error: "+err.Error(),
			)
			//return
		}
	} else if dwUserModel.AsyncQueryResultLocationPrefix.IsNull() && dwUserModel.AsyncQueryResultLocationAwsRoleArn.IsNull() {
		//config, err := r.client.GetAsyncAccountConfig(ctx, regionUri, dwUserModel.DwsuId.ValueString(), dwUserModel.ID.ValueString())
		//if err != nil {
		//	diagnostics.AddError(
		//		"Error read dwuser",
		//		"Could not read asyncResult before drop dwuser async config, unexpected error: "+err.Error(),
		//	)
		//} else {
		//	if config != nil && config.AwsIamArn != "" {
		//
		//	}
		//}
		_, err := r.client.DeleteAsyncAccountConfig(ctx, regionUri, dwUserModel.DwsuId.ValueString(), dwUserModel.ID.ValueString())
		if err != nil {
			diagnostics.AddError(
				"Error config dwuser",
				"Could not drop dwuser async config, unexpected error: "+err.Error(),
			)
			//return
		}
	} else if dwUserModel.AsyncQueryResultLocationPrefix.IsNull() || dwUserModel.AsyncQueryResultLocationAwsRoleArn.IsNull() {
		//只有一个属性的时候报给用户异常
		diagnostics.AddError(
			"Error config dwuser",
			"Could not config dwuser async, arn and prefix should be set together",
		)
	}
	if dwUserModel.DatalakeAwsLakeformationRoleArn.IsUnknown() {
		dwUserModel.DatalakeAwsLakeformationRoleArn = types.StringNull()
	}
	if !dwUserModel.DatalakeAwsLakeformationRoleArn.IsNull() {
		_, err := r.client.LakeFormationConfig(ctx, regionUri, dwUserModel.DwsuId.ValueString(), dwUserModel.ID.ValueString(), lakeFormation)
		if err != nil {
			diagnostics.AddError(
				"Error config dwuser",
				"Could not config dwuser lakeformation, unexpected error: "+err.Error(),
			)
			//return
		}
	} else if dwUserModel.DatalakeAwsLakeformationRoleArn.IsNull() {
		_, err := r.client.DeleteLakeFormationConfig(ctx, regionUri, dwUserModel.DwsuId.ValueString(), dwUserModel.ID.ValueString())
		if err != nil {
			diagnostics.AddError(
				"Error config dwuser",
				"Could not delete dwuser lakeformation, unexpected error: "+err.Error(),
			)
			//return
		}
	}
}
