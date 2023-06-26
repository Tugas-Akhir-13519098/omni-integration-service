package util

import (
	"bytes"
	"encoding/json"
	"io"
	"omni-integration-service/src/model"
)

func ConvertKafkaOrderMessageToCreateOrderRequest(orderMessage model.KafkaOrderMessage) *bytes.Buffer {
	customer := model.Customer{
		CustomerName:       orderMessage.CustomerName,
		CustomerPhone:      orderMessage.CustomerPhone,
		CustomerAddress:    orderMessage.CustomerAddress,
		CustomerDistrict:   orderMessage.CustomerDistrict,
		CustomerCity:       orderMessage.CustomerCity,
		CustomerProvince:   orderMessage.CustomerProvince,
		CustomerCountry:    orderMessage.CustomerCountry,
		CustomerPostalCode: orderMessage.CustomerPostalCode,
	}

	var products []model.OrderProductRequest
	for _, product := range orderMessage.Products {
		products = append(products, model.OrderProductRequest(product))
	}

	createOrderRequest := model.CreateOrderRequest{
		TokopediaOrderID: orderMessage.TokopediaOrderID,
		TokopediaShopID:  orderMessage.TokopediaShopID,
		ShopeeOrderID:    orderMessage.ShopeeOrderID,
		ShopeeShopID:     orderMessage.ShopeeShopID,
		TotalPrice:       orderMessage.TotalPrice,
		Customer:         customer,
		OrderStatus:      orderMessage.OrderStatus,
		Products:         products,
	}

	body, _ := json.Marshal(createOrderRequest)
	responseBody := bytes.NewBuffer(body)

	return responseBody
}

func ConvertHTTPResponseToOrderResponse(body io.ReadCloser) model.OrderResponse {
	respBody, _ := io.ReadAll(body)
	var orderResponse model.OrderResponse
	_ = json.Unmarshal(respBody, &orderResponse)

	return orderResponse
}

func ConvertKafkaOrderMessageToUpdateOrderStatusRequest(orderMessage model.KafkaOrderMessage) *bytes.Buffer {
	updateOrderStatusRequest := model.UpdateOrderStatusRequest{
		TokopediaOrderID: orderMessage.TokopediaOrderID,
		ShopeeOrderID:    orderMessage.ShopeeOrderID,
		OrderStatus:      orderMessage.OrderStatus,
	}

	body, _ := json.Marshal(updateOrderStatusRequest)
	responseBody := bytes.NewBuffer(body)

	return responseBody
}

func ConvertToErrorMessage(method string, url string, req string, err string, status string, reqTime string) []byte {
	message := model.KafkaErrorMessage{
		Method:      method,
		Url:         url,
		RequestBody: req,
		Error:       err,
		Status:      status,
		RequestTime: reqTime,
	}
	messageByte, _ := json.Marshal(message)

	return messageByte
}
