package common

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"strconv"
	"terraform-provider-relyt/internal/provider/client"
	"testing"
	"time"
)

func TestScroll(t *testing.T) {
	relytDatabaseClientConfig := client.RelytDatabaseClientConfig{
		DmsHost:   "http://127.0.0.1:8180",
		AccessKey: "AK8DoEFMRPWBGG0eY1JyNBVj7OnrTO3B6t3uJFyibDcGwz56HrAlg8uKtxf9hQeoHphJzOw",
		SecretKey: "HHJU4NBSLKZVGKTGRM41FCLGZVH4VPWS",
	}
	databaseClient, _ := client.NewRelytDatabaseClient(relytDatabaseClientConfig)
	start := time.Now()
	records, _ := ScrollPageRecords(&diag.Diagnostics{}, func(pageSize, pageNum int) ([]*client.SchemaMeta, error) {
		listRecords, err := CommonRetry(context.TODO(), func() (*client.CommonPage[client.SchemaMeta], error) {
			start = time.Now()
			schemas, err := databaseClient.ListSchemas(context.TODO(), client.SchemaPageQuery{
				PageQuery: client.PageQuery{
					PageSize:   pageSize,
					PageNumber: pageNum,
				},
				Database: "exp",
			})
			msg := ""
			if err != nil {
				msg += err.Error()
			}
			println("http call cost:" + (time.Now().Sub(start)).String() + " err" + msg)
			return schemas, err
		})
		if err != nil {
			return nil, err
		}
		if listRecords == nil {
			return nil, fmt.Errorf(" shouldn't get nil CommonPage resp")
		}
		return listRecords.Records, nil
	})
	fmt.Println("get all cost: " + (time.Now().Sub(start).String()))
	marshal, _ := json.Marshal(records)
	fmt.Println("size:" + strconv.Itoa(len(records)) + "body" + string(marshal))
}
