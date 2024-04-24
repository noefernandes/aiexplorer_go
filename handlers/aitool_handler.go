package handlers

import (
	"aiexplorer/data"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	size, err1 := strconv.Atoi(r.URL.Query().Get("size"))
	page, err2 := strconv.Atoi(r.URL.Query().Get("page"))

	if err1 != nil || err2 != nil {
		return
	}

	aitools, totalPages, err := data.GetAll(page, size)

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	type ResponseBody struct {
		Content    []data.AITool `json:"content"`
		TotalPages int64         `json:"totalPages"`
	}

	res := ResponseBody{aitools, int64(totalPages/size) + 1}

	json.NewEncoder(w).Encode(res)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aitool, err := data.Get(id)

	w.Header().Add("Content-Type", "application/json")

	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(aitool)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Save(w http.ResponseWriter, r *http.Request) {
	var aitool *data.AITool

	err := json.NewDecoder(r.Body).Decode(&aitool)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := data.Save(aitool)

	w.Header().Add("Content-Type", "application/json")

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

	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoded, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var aitool data.AITool
	err := json.NewDecoder(r.Body).Decode(&aitool)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := data.Delete(aitool.ID)

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	linesAffected, err := res.RowsAffected()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if linesAffected < 1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
