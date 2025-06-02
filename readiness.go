package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)//405
		w.Write([]byte("Method Not Allowed"))
		return
	}
	
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}