package dummy

import (
	"fmt"
	"net/http"
)

func DummyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain; charset=utf-8")
	for i, v := range r.Header {
		fmt.Fprintf(w, "%s %s\n", i, v)
	}

}
