package internal

import (
    "encoding/json"
    "net/http"
)

type Response struct {}

func (resp Response) AsJson(w http.ResponseWriter, payload interface{}) {
        response, err := json.Marshal(payload)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(response)
}
