package handlers

import (
	"aiexplorer/data"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAllTags(w http.ResponseWriter, r *http.Request) {

	tags, err := data.GetAllTags()

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(tags)
}

func GetTag(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tag, err := data.GetTag(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if tag != nil {
		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(tag)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(data)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func SaveTag(w http.ResponseWriter, r *http.Request) {
	var tag *data.TagInput
	err := json.NewDecoder(r.Body).Decode(&tag)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := data.SaveTag(tag)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
