package model

import "github.com/hashicorp/terraform-plugin-framework/types"

type IntegrationModel struct {
	DwsuId          types.String     `tfsdk:"dwsu_id"`
	IntegrationInfo *IntegrationInfo `tfsdk:"integration_info"`
}

type IntegrationInfo struct {
	ExternalId     types.String `tfsdk:"external_id"`
	RelytPrinciple types.String `tfsdk:"relyt_principle"`
	RelytVpc       types.String `tfsdk:"relyt_vpc"`
}
