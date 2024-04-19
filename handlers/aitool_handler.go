package handlers

import (
	"aiexplorer/data"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	aitools, err := data.GetAll()

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(aitools)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aitool, err := data.Get(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if aitool != nil {
		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(aitool)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(data)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func Save(w http.ResponseWriter, r *http.Request) {
	var aitool *data.AITool
	err := json.NewDecoder(r.Body).Decode(&aitool)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := data.Save(aitool)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var aitool *data.AITool
	err := json.NewDecoder(r.Body).Decode(&aitool)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := data.Update(aitool)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoded, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result != nil {
		w.WriteHeader(http.StatusCreated)
		w.Write(encoded)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
