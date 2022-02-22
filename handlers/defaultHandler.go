package handler

import "net/http"

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No functionality on root level. Please use paths "+UNI_INFO_PATH+" or "+COUNTRIES_PATH+
		" or "+DIAGNOSTICS_PATH+".", http.StatusOK)
}
