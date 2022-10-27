package item

//CREATE ITEM

type RequestCreateItem struct {
	ItemCode    string `json:"item_code" example:"hp-01"`
	Description string `json:"description" example:"Samsung s20 FE"`
	Quantity    int    `json:"quantity" example:"69"`
}

type ResponseCreateItem struct {
	message string `json:"message"`
	ID      string `json:"id"`
}
