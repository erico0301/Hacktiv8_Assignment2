package structs

type GetOrder struct {
	OrderID      int
	CustomerName string
	OrderAt      string
}

type CrtOrder struct {
	OrderedAt    string `json:"orderedAt"`
	CustomerName string `json:"customerName"`
	Items        []struct {
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	} `json:"items"`
}

type Order struct {
	OrderId      int     `json:"orderId"`
	CustomerName string  `json:"customerName"`
	OrderedAt    string  `json:"orderedAt"`
	Items        []Items `json:"items"`
}

type Items struct {
	ItemId      int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"orderId"`
}
