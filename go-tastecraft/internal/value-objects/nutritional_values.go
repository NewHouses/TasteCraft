package valueobjects

type NutritionalValues struct {
	Kilocalories  int     `json:"kilocalories"`
	Proteins      float32 `json:"proteins"`
	Carbohydrates float32 `json:"carbohydrates"`
	Fat           float32 `json:"fat"`
	Fiber         float32 `json:"fiber"`
	Calcium       float32 `json:"calcium"`
	Iron          float32 `json:"iron"`
	Zinc          float32 `json:"zinc"`
	VitaminA      float32 `json:"vitamin-a"`
	VitaminB      float32 `json:"vitamin-b"`
	VitaminB1     float32 `json:"vitamin-b1"`
	VitaminB2     float32 `json:"vitamin-b2"`
	VitaminB3     float32 `json:"vitamin-b3"`
	VitaminB6     float32 `json:"vitamin-b6"`
	VitaminB12    float32 `json:"vitamin-b12"`
	VitaminC      float32 `json:"vitamin-c"`
	VitaminD      float32 `json:"vitamin-d"`
	VitaminE      float32 `json:"vitamin-e"`
	VitaminK      float32 `json:"vitamin-k"`
}
