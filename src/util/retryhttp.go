package util

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func SendPostRequest(body *bytes.Buffer, url string) (*http.Response, error) {
	retryClient := NewRetryClient()
	resp, err := retryClient.Post(url, "application/json", body)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SendPutRequest(body *bytes.Buffer, url string) (*http.Response, error) {
	retryClient := NewRetryClient()
	req, _ := retryablehttp.NewRequest("PUT", url, body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := retryClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CustomErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	resp.Body.Close()
	return resp, err
}

func NewRetryClient() *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.ErrorHandler = CustomErrorHandler

	return retryClient
}

func AfterHTTPRequestHandler(resp *http.Response, method string, tokopediaID int, shopeeID string) {
	orderResponse := ConvertHTTPResponseToOrderResponse(resp.Body)
	if IsFailResponse(resp) {
		fmt.Printf("Failed to send HTTP %s Request with status: %s and error: %s\n", method, resp.Status, orderResponse.Message)
	} else {
		fmt.Printf("Successfully %s order with tokopedia_order_id: %d & shopee_order_id: %s\n", method, tokopediaID, shopeeID)
	}
}

func IsFailResponse(resp *http.Response) bool {
	return (resp.StatusCode < 200 || resp.StatusCode > 299)
}
