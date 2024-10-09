package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func doHttpRequest[T any](p *RelytClient, ctx context.Context, host, path, method string,
	respMode *CommonRelytResponse[T],
	request any,
	parameter map[string]string,
	codeHandler func(response *CommonRelytResponse[T], respDumpByte []byte) (*CommonRelytResponse[T], error)) (err error) {
	return doHttpRequestWithHeader(p, ctx, host, path, method, respMode, request, parameter, nil, codeHandler)
}

func doHttpRequestWithHeader[T any](p *RelytClient, ctx context.Context, host, path, method string,
	respMode *CommonRelytResponse[T],
	request any,
	parameter map[string]string,
	header map[string]string,
	codeHandler func(response *CommonRelytResponse[T], respDumpByte []byte) (*CommonRelytResponse[T], error)) (err error) {
	if host == "" {
		host = p.ApiHost
	}
	var jsonData = []byte("")
	if request != nil && "" != request {
		requestJson, err := json.Marshal(request)
		if err != nil {
			tflog.Error(ctx, "fmt request json error:"+err.Error())
		}
		tflog.Info(ctx, "request data :"+string(requestJson))
		jsonData = requestJson // POST请求发送的数据
	}
	hostApi := host + path
	parsedHostApi, err := url.Parse(hostApi)
	if err != nil {
		return err
	}
	queryParams := url.Values{}
	if parameter != nil {
		for k, v := range parameter {
			queryParams.Add(k, v)
		}
	}
	parsedHostApi.RawQuery = queryParams.Encode()

	req, err := http.NewRequest(method, parsedHostApi.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		tflog.Error(ctx, "Error creating request:"+err.Error())
		return err
	}
	req.Header.Set("x-maxone-api-key", p.AuthKey)
	req.Header.Set("x-maxone-role-id", p.Role)
	req.Header.Set("Content-Type", "application/json")
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	requestString, _ := httputil.DumpRequestOut(req, true)
	tflog.Info(ctx, "== request: "+string(requestString))
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, "Error sending request:"+err.Error())
		return err
	}
	defer resp.Body.Close()
	responseString, _ := httputil.DumpResponse(resp, true)
	tflog.Info(ctx, "== response: "+string(responseString))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading responseString body:"+err.Error())
		return err
	}
	if resp.StatusCode != CODE_SUCCESS {
		tflog.Error(ctx, "Error status http code not 200! "+resp.Status)
		//printResp(ctx, resp)
		return fmt.Errorf("Error http code not 200! respCode: %s!\n%s ", resp.Status, string(body))
	}

	err = json.Unmarshal(body, respMode)
	if err != nil {
		tflog.Error(ctx, "read json respFail:"+err.Error())
		return err
	}
	if respMode.Code != CODE_SUCCESS {
		tflog.Warn(ctx, "error call api! resp code not 200: "+string(body))
	}
	if codeHandler != nil {
		tflog.Trace(ctx, "use code handle func!")
		handler, err := codeHandler(respMode, body)
		if handler != nil {
			respMode.Code = handler.Code
			respMode.Data = handler.Data
			respMode.Msg = handler.Msg
		}
		if err != nil {
			return err
		}
	} else {
		if respMode.Code != CODE_SUCCESS {
			tflog.Error(ctx, "error call api! resp code not 200: "+string(body))
			return fmt.Errorf(string(body))
		}
	}
	return nil
}

func signedHttpRequestWithHeader[T any](p *RelytClient, ctx context.Context, host, path, method string,
	respMode *CommonRelytResponse[T],
	request any,
	parameter map[string]string,
	header map[string]string,
	codeHandler func(response *CommonRelytResponse[T], respDumpByte []byte) (*CommonRelytResponse[T], error)) (err error) {
	if host == "" {
		host = p.ApiHost
	}
	var jsonData = []byte("")
	if request != nil && "" != request {
		requestJson, err := json.Marshal(request)
		if err != nil {
			tflog.Error(ctx, "fmt request json error:"+err.Error())
		}
		tflog.Info(ctx, "request data :"+string(requestJson))
		jsonData = requestJson // POST请求发送的数据
	}
	hostApi := host + path
	parsedHostApi, err := url.Parse(hostApi)
	if err != nil {
		return err
	}
	queryParams := url.Values{}
	if parameter != nil {
		for k, v := range parameter {
			queryParams.Add(k, v)
		}
	}
	parsedHostApi.RawQuery = queryParams.Encode()

	req, err := http.NewRequest(method, parsedHostApi.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		tflog.Error(ctx, "Error creating request:"+err.Error())
		return err
	}
	req.Header.Set("x-maxone-api-key", p.AuthKey)
	req.Header.Set("x-maxone-role-id", p.Role)
	req.Header.Set("Content-Type", "application/json")
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	requestString, _ := httputil.DumpRequestOut(req, true)
	tflog.Info(ctx, "== request: "+string(requestString))
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		tflog.Error(ctx, "Error sending request:"+err.Error())
		return err
	}
	defer resp.Body.Close()
	responseString, _ := httputil.DumpResponse(resp, true)
	tflog.Info(ctx, "== response: "+string(responseString))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading responseString body:"+err.Error())
		return err
	}
	if resp.StatusCode != CODE_SUCCESS {
		tflog.Error(ctx, "Error status http code not 200! "+resp.Status)
		//printResp(ctx, resp)
		return fmt.Errorf("Error http code not 200! respCode: %s!\n%s ", resp.Status, string(body))
	}

	err = json.Unmarshal(body, respMode)
	if err != nil {
		tflog.Error(ctx, "read json respFail:"+err.Error())
		return err
	}
	if respMode.Code != CODE_SUCCESS {
		tflog.Warn(ctx, "error call api! resp code not 200: "+string(body))
	}
	if codeHandler != nil {
		tflog.Trace(ctx, "use code handle func!")
		handler, err := codeHandler(respMode, body)
		if handler != nil {
			respMode.Code = handler.Code
			respMode.Data = handler.Data
			respMode.Msg = handler.Msg
		}
		if err != nil {
			return err
		}
	} else {
		if respMode.Code != CODE_SUCCESS {
			tflog.Error(ctx, "error call api! resp code not 200: "+string(body))
			return fmt.Errorf(string(body))
		}
	}
	return nil
}
