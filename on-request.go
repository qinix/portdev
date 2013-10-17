package portdev

import (
	"net/http"
	"strings"
	"strconv"
	"net/http/httputil"
	"net/url"
	"fmt"
)

func OnRequest(w http.ResponseWriter, r *http.Request) {
	debug("Routing %s%s", r.Host, r.URL)
	port, err := strconv.Atoi(strings.Split(r.Host, ".")[0])
	if err != nil {
		debug("Error: ", err)
		http.NotFound(w, r)
		return
	}

	dest, _ := url.Parse(fmt.Sprintf("http://127.0.0.1:%d%s", port, r.URL.String()))
	server := httputil.NewSingleHostReverseProxy(dest)
	server.ServeHTTP(w, r)
}
