package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := &model.HealthzResponse{Message: "OK"}// レスポンス用の構造体を作成

	w.Header().Set("Content-Type", "application/json") // レスポンスヘッダーにJSONであることを設定
	if err := json.NewEncoder(w).Encode(res);err != nil {
		log.Println(err)
	}
// ↑の構文がよくわからなかったけど、以下をまとめて書くとこうなるらしい
// 	// 1つ目の処理：JSONエンコードを実行して結果をerrに格納
// err := json.NewEncoder(w).Encode(res)

// // 2つ目の処理：errがnilでないかチェック
// if err != nil {
//     log.Println(err)
// }
}
