package nameservice

import (
	"net/http"
)

func StartServer() {
	smux := http.NewServeMux()
	smux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		email := q.Get("email")

		if email == "test@test.com" {
			rw.Write([]byte(`{"name":"ricardo","surname":"linck"}`))
			return
		}

		rw.WriteHeader(http.StatusNotFound)
	})
	http.ListenAndServe("localhost:8787", smux)
}
