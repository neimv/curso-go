package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/neimv/curso-go/bd"
)

func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parémetro página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina, con una valor mayor a cero", http.StatusBadRequest)
	}

	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweets(ID, pag)

	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
