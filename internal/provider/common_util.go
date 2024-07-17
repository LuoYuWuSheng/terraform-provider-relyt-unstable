package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"math"
	"strconv"
	"terraform-provider-relyt/internal/provider/client"
	"time"
)

func RouteRegionUri(ctx context.Context, dwsuId string, relytClient *client.RelytClient,
	diag *diag.Diagnostics) *client.OpenApiMetaInfo {
	meta, err := RetryFunction[client.OpenApiMetaInfo](ctx, 5, 1, 1.0,
		func() (*client.OpenApiMetaInfo, error) {
			return relytClient.GetDwsuOpenApiMeta(ctx, dwsuId)
		})
	if err != nil || meta == nil {
		errMsg := "get RegionApi is nil"
		if err != nil {
			errMsg = err.Error()
		}
		diag.AddError("error get region api", "fail to get Region uri address dwsuID:"+
			""+dwsuId+" error: "+errMsg)
		return meta
	}
	return meta
}

func RetryFunction[T any](ctx context.Context, retryNum, intervalSecond int,
	backoffCoefficient float64,
	retryableFunc func() (*T, error)) (*T, error) {
	var err error
	var result *T
	for i := 0; i < retryNum; i++ {
		result, err = retryableFunc()
		if err == nil {
			return result, nil
		}
		time.Sleep(time.Duration(intervalSecond))
		intervalSecond = int(math.Ceil(backoffCoefficient * float64(intervalSecond)))
		tflog.Warn(ctx, "retry failed func! backoff second:"+strconv.Itoa(intervalSecond)+"error msg!"+err.Error())
	}

	return result, err
}
