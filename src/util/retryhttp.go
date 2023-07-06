package util

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"omni-integration-service/config"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/segmentio/kafka-go"
)

func SendPostRequest(body *bytes.Buffer, url string) (*http.Response, error) {
	retryClient := NewRetryClient()
	req, _ := retryablehttp.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Get().AdminToken)
	resp, err := retryClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SendPutRequest(body *bytes.Buffer, url string) (*http.Response, error) {
	retryClient := NewRetryClient()
	req, _ := retryablehttp.NewRequest("PUT", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Get().AdminToken)
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

func CustomRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// don't propagate other errors
	shouldRetry, _ := BaseRetryPolicy(resp, err)
	return shouldRetry, nil
}

func BaseRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		// The error is likely recoverable so retry.
		return true, nil
	}

	// 429 Too Many Requests is recoverable.
	if resp.StatusCode == http.StatusTooManyRequests {
		return true, nil
	}

	// 408 Request Timeout is recoverable.
	if resp.StatusCode == http.StatusRequestTimeout {
		return true, nil
	}

	// Retry on 5xx status code
	if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != http.StatusNotImplemented) {
		return true, fmt.Errorf("unexpected HTTP status %s", resp.Status)
	}

	return false, nil
}

func NewRetryClient() *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.ErrorHandler = CustomErrorHandler
	retryClient.CheckRetry = CustomRetryPolicy

	return retryClient
}

func AfterHTTPRequestHandler(req string, resp *http.Response, method string, httpMethod string, tokopediaID int, shopeeID string, url string) {
	orderResponse := ConvertHTTPResponseToOrderResponse(resp.Body)
	cfg := config.Get()

	if IsFailResponse(resp) {
		fmt.Printf("Failed to send HTTP %s Request with status: %s and error: %s\n", httpMethod, resp.Status, orderResponse.Message)

		kafkaMessage := ConvertToErrorMessage(httpMethod, url, req, orderResponse.Message, resp.Status, time.Now().Format("2006-01-02 15:04:05"))

		// Publish to Kafka Error Topic
		config := kafka.WriterConfig{
			Brokers: []string{fmt.Sprintf("%s:%s", cfg.KafkaHost, cfg.KafkaPort)},
			Topic:   cfg.KafkaErrorTopic,
		}
		writer := kafka.NewWriter(config)

		var key string
		if tokopediaID == 0 {
			key = shopeeID
		} else {
			key = strconv.Itoa(tokopediaID)
		}
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(key),
			Value: []byte(kafkaMessage),
		})
		if err != nil {
			fmt.Println("Failed to send request to Kafka Error Topic")
		}
	} else {
		fmt.Printf("Successfully %s order with tokopedia_order_id: %d & shopee_order_id: %s\n", method, tokopediaID, shopeeID)
	}
}

func IsFailResponse(resp *http.Response) bool {
	return (resp.StatusCode < 200 || resp.StatusCode > 299)
}
