package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/neimv/curso-go/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "usuario no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatar/" + perfil.Banner)

	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "error al copiar la imagen", http.StatusBadRequest)
	}
}
