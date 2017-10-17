package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kandrewc/info344-a17/info344-in-class/zipsvr/models"
)

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/city-name

	cityName := r.URL.Path[len(ch.PathPrefix):]
	cityName = strings.ToLower(cityName)

	if len(cityName) == 0 {
		http.Error(w, "Please provide a cityname", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(accessControlAllowOrigin, "*")
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips)
}
