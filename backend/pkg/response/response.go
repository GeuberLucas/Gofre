package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse sends a JSON response with the given status code and data
func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}


func ErrorResponse(w http.ResponseWriter, statusCode int, erro error) {
	JSONResponse(w,statusCode,struct{ Erro string `json:"erro"` }{Erro: erro.Error()})
}