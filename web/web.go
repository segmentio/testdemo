package web

import (
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if v := params.Get("sloths"); v != "" {
		w.Write([]byte(v))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("badwolf"))
	}
}
