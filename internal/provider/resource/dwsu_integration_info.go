package resource

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"terraform-provider-relyt/internal/provider/client"
	"terraform-provider-relyt/internal/provider/common"
	tfModel "terraform-provider-relyt/internal/provider/model"
)

var (
	_ resource.Resource              = &dwsuIntegrationInfoResource{}
	_ resource.ResourceWithConfigure = &dwsuIntegrationInfoResource{}
)

func NewDwsuIntegrationInfoResource() resource.Resource {
	return &dwsuIntegrationInfoResource{}
}

type dwsuIntegrationInfoResource struct {
	//client *client.RelytClient
	RelytClientResource
}

func (r *dwsuIntegrationInfoResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dwsu_integration_info"
}

// Schema defines the schema for the data source.
func (r *dwsuIntegrationInfoResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dwsu_id": schema.StringAttribute{Required: true},
			"integration_info": schema.SingleNestedAttribute{
				Required:    true,
				Description: "used to set Integration Info. Empty block will use the system default Integration Info ",
				Attributes: map[string]schema.Attribute{
					"external_id":     schema.StringAttribute{Optional: true, Computed: true, Description: "The externalId of dwsu", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
					"relyt_principle": schema.StringAttribute{Computed: true, Description: "The relyt Principal", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
					"relyt_vpc":       schema.StringAttribute{Computed: true, Description: "The relyt VPC", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
				},
			},
		},
	}
}

// Create a new resource.
func (r *dwsuIntegrationInfoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan tfModel.IntegrationModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//if !plan.IntegrationInfo.IsNull() && !plan.IntegrationInfo.IsUnknown() {
	//
	//}
	info := tfModel.IntegrationInfo{}
	types.Object.As(plan.IntegrationInfo, ctx, &info, basetypes.ObjectAsOptions{})
	if !info.ExternalId.IsUnknown() {
		meta := common.RouteRegionUri(ctx, plan.DwsuId.ValueString(), r.client, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		regionUri := meta.URI
		r.updateIntegrationInfo(ctx, regionUri, plan.DwsuId.ValueString(), &info, &resp.Diagnostics)
	}
	meta := common.RouteRegionUri(ctx, plan.DwsuId.ValueString(), r.client, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	regionUri := meta.URI
	r.readIntegrationInfo(ctx, regionUri, &plan, &resp.Diagnostics)
	resp.State.Set(ctx, &plan)
}

// Read resource information.
func (r *dwsuIntegrationInfoResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state tfModel.IntegrationModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if state.DwsuId.IsNull() || state.DwsuId.IsUnknown() || state.DwsuId.ValueString() == "" {
		resp.Diagnostics.AddError("require dwsu id", "dwsu_id must be supplied")
	}
	meta := common.RouteRegionUri(ctx, state.DwsuId.ValueString(), r.client, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	regionUri := meta.URI
	r.readIntegrationInfo(ctx, regionUri, &state, &resp.Diagnostics)
	resp.State.Set(ctx, &state)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *dwsuIntegrationInfoResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "user try delete integration info! doing nothing")
}

func (r *dwsuIntegrationInfoResource) readIntegrationInfo(ctx context.Context, regionUri string,
	state *tfModel.IntegrationModel, diagnostics *diag.Diagnostics) {
	info, err := common.CommonRetry(ctx, func() (*client.IntegrationInfo, error) {
		return r.client.GetIntegration(ctx, regionUri, state.DwsuId.ValueString())
	})
	if err != nil || info == nil {
		msg := " read info is nil"
		if err != nil {
			msg = err.Error()
		}
		diagnostics.AddError("failed to get integration info", "get dwsu integration info get err:"+msg)
		return
	}
	integrationInfo := tfModel.IntegrationInfo{
		ExternalId:     types.StringValue(info.ExternalId),
		RelytPrinciple: types.StringValue(info.RelytPrincipal),
		RelytVpc:       types.StringValue(info.RelytVpc),
	}
	tfObject, d := types.ObjectValueFrom(ctx, map[string]attr.Type{
		"external_id":     types.StringType,
		"relyt_principle": types.StringType,
		"relyt_vpc":       types.StringType,
	}, integrationInfo)
	diagnostics.Append(d...)
	state.IntegrationInfo = tfObject
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *dwsuIntegrationInfoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//resp.Diagnostics.AddError("not support", "update account not supported")
	var plan tfModel.IntegrationModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// read old status
	var stat tfModel.IntegrationModel
	req.State.Get(ctx, &stat)
	if resp.Diagnostics.HasError() {
		return
	}
	stateInfo := tfModel.IntegrationInfo{}
	planInfo := tfModel.IntegrationInfo{}
	types.Object.As(plan.IntegrationInfo, ctx, &planInfo, basetypes.ObjectAsOptions{})
	types.Object.As(stat.IntegrationInfo, ctx, &stateInfo, basetypes.ObjectAsOptions{})
	if !stateInfo.ExternalId.Equal(planInfo.ExternalId) {
		meta := common.RouteRegionUri(ctx, stat.DwsuId.ValueString(), r.client, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		regionUri := meta.URI
		r.updateIntegrationInfo(ctx, regionUri, stat.DwsuId.ValueString(), &planInfo, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		r.readIntegrationInfo(ctx, regionUri, &stat, &resp.Diagnostics)
		resp.State.Set(ctx, stat)
	}
}

func (r *dwsuIntegrationInfoResource) updateIntegrationInfo(ctx context.Context, regionUri, dwsuId string, info *tfModel.IntegrationInfo, diagnostic *diag.Diagnostics) {
	integrationInfo := client.IntegrationInfo{
		ExternalId: info.ExternalId.ValueString(),
	}
	_, err := common.CommonRetry(ctx, func() (*client.CommonRelytResponse[string], error) {
		return r.client.PatchIntegration(ctx, regionUri, dwsuId, integrationInfo)
	})
	if err != nil {
		diagnostic.AddError("failed update integration info", "update integration info get err"+err.Error())
		return
	}
}
