package client

import "testing"

func init() {
	relytDatabaseClientConfig := RelytDatabaseClientConfig{
		DmsHost:   "http://127.0.0.1:8180/api/catalog/database",
		AccessKey: "AK8DoEFMRPWBGG0eY1JyNBVj7OnrTO3B6t3uJFyibDcGwz56HrAlg8uKtxf9hQeoHphJzOw",
		SecretKey: "HHJU4NBSLKZVGKTGRM41FCLGZVH4VPWS",
	}
	databaseClient, _ = NewRelytDatabaseClient(relytDatabaseClientConfig)
}

var (
	databaseClient RelytDatabaseClient
)

func TestRelytDatabaseClient_testSign(t *testing.T) {
	databaseClient.testSign()
}
