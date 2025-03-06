package model

// A HealthzResponse expresses health check message.
type HealthzResponse struct{
	//   Message という Field が JSON にシリアライズされた時に message として出力されるように struct tag を追加
	Message string `json:"message"`
}
