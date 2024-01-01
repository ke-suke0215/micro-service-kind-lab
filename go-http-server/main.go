package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/test", sampleHandler)

    log.Println("サーバー起動 :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

// CORS設定
func setupCORS(w *http.ResponseWriter, req *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "*")
    (*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
    // CORS設定を適用
    setupCORS(&w, r)
    // OPTIONSメソッドのリクエストに対しては、処理をここで終了
    if r.Method == "OPTIONS" {
        return
    }

    // 以下は元のハンドラーのロジック
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    
    response := map[string]string{"key": "value"}
    json.NewEncoder(w).Encode(response)
}
