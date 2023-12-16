package main

import (
	"encoding/json"
	"net/http"
)

// 아래 코드 사용 금지
// http.NewServeMux()
// ORM, 최근에 추가된 api 사용금지
// sqlite 데이터 베이스

func main() {

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := make(map[string]string)
		body["text"] = "hello"
		res, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		w.Write(res)
	})

	http.ListenAndServe(":5000", nil)
}
