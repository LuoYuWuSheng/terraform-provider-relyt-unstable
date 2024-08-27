package model

import "github.com/hashicorp/terraform-plugin-framework/types"

type IntegrationModel struct {
	DwsuId          types.String `tfsdk:"dwsu_id"`
	IntegrationInfo types.Object `tfsdk:"integration_info"`
	//IntegrationInfo *IntegrationInfo `tfsdk:"integration_info"`
}

type IntegrationInfo struct {
	ExternalId     types.String `tfsdk:"external_id"`
	RelytPrincipal types.String `tfsdk:"relyt_principal"`
	RelytVpc       types.String `tfsdk:"relyt_vpc"`
}
