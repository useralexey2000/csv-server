package handler

import (
	"csv-server/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ServeRecords(rs domain.RecordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			respond(
				w, http.StatusMethodNotAllowed,
				fmt.Errorf("not supported http.metod"))
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			respond(
				w, http.StatusBadRequest,
				fmt.Errorf("no id provided"))
			return
		}
		recs, err := rs.GetRecords(r.Context(), id)
		if err != nil {
			respond(
				w, http.StatusInternalServerError,
				fmt.Errorf("can't get records"))
			return
		}
		respond(
			w, http.StatusOK,
			recs)
	}
}

func respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println(err)
	}
}
