package resource

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"terraform-provider-relyt/internal/provider/client"
	"terraform-provider-relyt/internal/provider/common"
	"terraform-provider-relyt/internal/provider/model"
)

func updateDps(ctx context.Context, relytClient *client.RelytClient, state, plan *model.Dps, diag *diag.Diagnostics, dwsuId, dpsId string) {
	diagnostics := diag
	meta := common.RouteRegionUri(ctx, dwsuId, relytClient, diagnostics)
	if diagnostics.HasError() {
		return
	}
	regionUri := meta.URI
	patchDps := client.DpsMode{
		Description: plan.Description.ValueString(),
		Engine:      plan.Engine.ValueString(),
		Name:        plan.Name.ValueString(),
		Spec:        &client.Spec{Name: plan.Size.ValueString()},
	}
	_, err := relytClient.PatchDps(ctx, regionUri, dwsuId, dpsId, patchDps)
	if err != nil {
		tflog.Error(ctx, "error update dps"+err.Error())
		diagnostics.AddError("update dps failed!", "error update dps!"+err.Error())
		return
	}
	//读一下最新状态，写入Status，告诉用户正在变配。注意这时候不能写size，否则会导致TF认为已经是目标Size不再重入到Update逻辑
	readDps(ctx, dwsuId, dpsId, relytClient, diagnostics, state)
	//dig := resp.State.Set(ctx, &state)
	//resp.Diagnostics.Append(dig...)
	if diagnostics.HasError() {
		return
	}
	_, err = WaitDpsReady(ctx, relytClient, regionUri, dwsuId, dpsId, diagnostics)
	if err != nil {
		tflog.Error(ctx, "error wait dps ready after update"+err.Error())
		diagnostics.AddError("update dps failed!", "error wait dps ready after update!"+err.Error())
		return
	}
	//mapRelytDpsToTFModel(dps, state)
	//更改成功，则将Size设置为目标Size
	state.Size = plan.Size
}

func readDps(ctx context.Context, dwsuId, dpsId string, r *client.RelytClient, diagnostics *diag.Diagnostics, dpsModel *model.Dps) {
	meta := common.RouteRegionUri(ctx, dwsuId, r, diagnostics)
	if diagnostics.HasError() {
		return
	}
	regionUri := meta.URI
	dps, err := common.CommonRetry(ctx, func() (*client.DpsMode, error) {
		return r.GetDps(ctx, regionUri, dwsuId, dpsId)
	})
	//_, err := r.client.GetDps(ctx, regionUri, state.DwsuId.ValueString(), state.ID.ValueString())
	if err != nil || dps == nil {
		msg := "read dps get nil"
		if err != nil {
			msg = err.Error()
		}
		tflog.Error(ctx, "error read dps"+msg)
		diagnostics.AddError("error read", "error read dps!"+msg)
		return
	}
	mapRelytDpsToTFModel(dps, dpsModel)
}

func mapRelytDpsToTFModel(dps *client.DpsMode, dpsModel *model.Dps) {
	if dps != nil && dpsModel != nil {
		//注意，为了实现变配阻塞，自动重新发起变配，这里不能更新Size。由update判断是否更新成功后更新size
		dpsModel.Status = types.StringValue(dps.Status)
	}
}

func WaitDpsReady(ctx context.Context, relytClient *client.RelytClient, regionUri string, dwsuId, dpsId string, diagnostics *diag.Diagnostics) (*client.DpsMode, error) {
	queryDpsMode, err := common.TimeOutTask(relytClient.CheckTimeOut, relytClient.CheckInterval, func() (any, error) {
		dps, err2 := relytClient.GetDps(ctx, regionUri, dwsuId, dpsId)
		if err2 != nil {
			//这里判断是否要充实
			return dps, err2
		}
		if dps != nil && dps.Status == client.DPS_STATUS_READY {
			return dps, nil
		}
		return dps, fmt.Errorf("dps is not Ready")
	})
	if err != nil {
		tflog.Error(ctx, "error wait dps ready"+err.Error())
		diagnostics.AddError(
			"create dps failed!", "error wait dps ready! "+err.Error(),
		)
		return nil, err
		//fmt.Println(fmt.Sprintf("drop dwsu%s", err.Error()))
	}
	//上面已经判断了，这里其实不可能为nil
	var convertType *client.DpsMode
	if queryDpsMode != nil {
		convertType = queryDpsMode.(*client.DpsMode)
	}
	return convertType, err
}
