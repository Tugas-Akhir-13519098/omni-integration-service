package model

type CreateOrderRequest struct {
	TokopediaOrderID int                   `json:"tokopedia_order_id"`
	TokopediaShopID  int                   `json:"tokopedia_shop_id"`
	ShopeeOrderID    string                `json:"shopee_order_id"`
	ShopeeShopID     int                   `json:"shopee_shop_id"`
	TotalPrice       float32               `json:"total_price"`
	Customer         Customer              `json:"customer"`
	OrderStatus      OrderStatus           `json:"order_status"`
	Products         []OrderProductRequest `json:"product"`
}

type Customer struct {
	CustomerName       string `json:"customer_name"`
	CustomerPhone      string `json:"customer_phone"`
	CustomerAddress    string `json:"customer_address"`
	CustomerDistrict   string `json:"customer_district"`
	CustomerCity       string `json:"customer_city"`
	CustomerProvince   string `json:"customer_province"`
	CustomerCountry    string `json:"customer_country"`
	CustomerPostalCode string `json:"customer_postal_code"`
}

type OrderProductRequest struct {
	TokopediaProductID int     `json:"tokopedia_product_id"`
	ShopeeProductID    int     `json:"shopee_product_id"`
	ProductName        string  `json:"product_name"`
	ProductPrice       float32 `json:"product_price"`
	ProductQuantity    int     `json:"product_quantity"`
}

type UpdateOrderStatusRequest struct {
	TokopediaOrderID int         `json:"tokopedia_order_id"`
	ShopeeOrderID    string      `json:"shopee_order_id"`
	OrderStatus      OrderStatus `json:"order_status"`
}

type OrderResponse struct {
	Message string `json:"message"`
	Status  string `json:"failed"`
}
