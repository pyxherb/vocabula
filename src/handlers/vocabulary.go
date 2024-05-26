package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"vocabula/model"
)

func VocabularyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		queryParams := r.URL.Query()

		if (!queryParams.Has("language")) ||
			(!queryParams.Has("name")) {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		var language string
		var name string

		language = queryParams.Get("language")
		name = queryParams.Get("name")

		var word model.Word
		err := model.QueryWord(language, name, &word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	case "PUT":
		queryParams := r.URL.Query()

		if (!queryParams.Has("language")) ||
			(!queryParams.Has("name")) {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		var language string
		var name string

		language = queryParams.Get("language")
		name = queryParams.Get("name")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var word model.Word
		err = json.Unmarshal(body, &word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = model.UpdateWord(language, name, &word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
