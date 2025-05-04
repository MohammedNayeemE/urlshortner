package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MohammedNayeemE/url-shortner/storage"
	"github.com/MohammedNayeemE/url-shortner/utils"
	"github.com/gorilla/mux"
)

var baseUrl string 

func Init(url string) {
	baseUrl = url;
}

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	ShortURL string `json:shorturl`
}

func ShortenUrl(w http.ResponseWriter , r *http.Request) {

	var req Request;

	json.NewDecoder(r.Body).Decode(&req);

	shortenedurl := utils.GenerateUrl();

	storage.Save(shortenedurl , req.URL);

	ret := Response{ShortURL: baseUrl + "/"+  shortenedurl};

	w.Header().Set("Content-Type" , "application/json");

	json.NewEncoder(w).Encode(ret);
	
}

func RedirectUrl(w http.ResponseWriter , r *http.Request) {

	shortenedurl := mux.Vars(r)["shorturl"];

	originalPath := storage.GetUrl(shortenedurl);

	if originalPath == "" {
		http.Error(w , "url not found" , http.StatusNotFound);
		return;
	}

	http.Redirect(w , r , originalPath , http.StatusFound);
}
