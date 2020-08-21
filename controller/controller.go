package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"rest-api/model/client_model"
	"rest-api/repository"
	"rest-api/utilities/errors"

	"golang.org/x/oauth2"
)

var errorRedirect = "/ConnectionSocialNetworks/facebook/error?error="
var tokenRedirect = "/ConnectionSocialNetworks/facebook/media?code="

func Error(w http.ResponseWriter, r *http.Request) {
	errorMsg := r.FormValue("error")
	var jsonError = errors.NewInternalServerError(errorMsg)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonError)
}

func Login(w http.ResponseWriter, r *http.Request) {
	url := repository.LoginFacebook()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	oauthStateString := repository.GetOauthState()

	if state != oauthStateString {
		var er = ("invalid oauth state, expected '" + oauthStateString + "', got '" + state)
		log.Printf(er)
		http.Redirect(w, r, (errorRedirect + er), http.StatusTemporaryRedirect)
	}

	code := r.FormValue("code")

	oauthConf := repository.GetOauthConf()

	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("\noauthConf.Exchange() failed with '" + err.Error() + "\n")
		http.Redirect(w, r, errorRedirect+err.Error(), http.StatusTemporaryRedirect)
	}

	resp, err := http.Get(repository.GetUrlFacebookUser() + url.QueryEscape(token.AccessToken))
	if err != nil {
		log.Printf("\n Get: " + err.Error() + "\n")
		http.Redirect(w, r, (errorRedirect + err.Error()), http.StatusTemporaryRedirect)
	}
	defer resp.Body.Close()

	http.Redirect(w, r, (tokenRedirect + token.AccessToken), http.StatusTemporaryRedirect)

}

func Media(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("code")

	var out client_model.InformacionResponse
	var informacion client_model.Informacion

	resp, err := repository.GetFacebookMedia(token)
	if err != nil {
		fmt.Printf("*****ERROR****** :\n%+v\n", err)
		out.Parametros = client_model.Parametros{
			"Error",
			err.Error,
		}
		json.NewEncoder(w).Encode(out)
	}

	for i := 0; i < len(resp.Data); i++ {
		informacion = client_model.Informacion{
			resp.Data[i].ID,
			resp.Data[i].Caption,
			resp.Data[i].MediaURL,
			resp.Data[i].Username,
			resp.Data[i].Permalink,
			resp.Data[i].Timestamp,
		}

		out.Informacion = append(out.Informacion, informacion)
	}

	out.Parametros = client_model.Parametros{
		"OK",
		"Respuesta exitosa",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)

}
