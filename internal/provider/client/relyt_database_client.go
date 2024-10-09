package client

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"io"
	"net/http"
	"time"
)

func NewRelytDatabaseClient(config RelytDatabaseClientConfig) (RelytDatabaseClient, error) {
	return RelytDatabaseClient{config}, nil
}

type RelytDatabaseClient struct {
	RelytDatabaseClientConfig
}

type RelytDatabaseClientConfig struct {
	DmsHost   string `json:"apiHost"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func (r *RelytDatabaseClientConfig) testSign() {

	body := "{\"name\":\"test_catalog\"}"
	// 创建一个 HTTP 请求
	req, err := http.NewRequest("POST", "http://127.0.0.1:8180/api/catalog/database/create", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	credentials := aws.Credentials{
		AccessKeyID:     "AK8DoEFMRPWBGG0eY1JyNBVj7OnrTO3B6t3uJFyibDcGwz56HrAlg8uKtxf9hQeoHphJzOw",
		SecretAccessKey: "HHJU4NBSLKZVGKTGRM41FCLGZVH4VPWS",
		SessionToken:    "",
		Source:          "",
		CanExpire:       false,
		Expires:         time.Time{},
		AccountID:       "",
	}
	//// 解析时间字符串
	//layout := "20060102T150405Z"
	//parsedTime, err := time.Parse(layout, "20241009T023711Z")
	//if err != nil {
	//	fmt.Println("Error parsing time:", err)
	//	return
	//}
	//// 输出解析后的时间
	//fmt.Println("Parsed Time:", parsedTime)

	// 设置 AWS 静态凭证
	//creds := aws.NewCredentialsCache(aws.NewStaticCredentialsProvider("YOUR_ACCESS_KEY_ID", "YOUR_SECRET_ACCESS_KEY", ""))
	//creds := aws.NewCredentialsCache(aws.AnonymousCredentials("YOUR_ACCESS_KEY_ID", "YOUR_SECRET_ACCESS_KEY", ""))

	// 创建一个签名者
	signer := v4.NewSigner()

	//sha256 := v4.ComputePayloadSHA256{}
	//sha256.ID()
	// 计算 SHA-256 哈希
	hash := sha256.Sum256([]byte(body))

	// 将哈希值转换为十六进制字符串
	payloadHash := hex.EncodeToString(hash[:])

	// 签名请求
	err = signer.SignHTTP(context.TODO(), credentials, req, payloadHash, "relyt", "us-east", time.Now())
	if err != nil {
		fmt.Println("Error signing request:", err)
		return
	}

	// 打印签名后的请求头
	fmt.Println("Signed Request Headers:")
	for key, values := range req.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	// 发送请求（可选）
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("er", err)
		return
	}
	fmt.Println("Response body:", string(all))

}
