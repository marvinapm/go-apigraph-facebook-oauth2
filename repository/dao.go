package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	rm "rest-api/model/repository_model"
	er "rest-api/utilities/errors"
	"rest-api/utilities/random"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "XXXXXXXXXXXXXXX",
		ClientSecret: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		RedirectURL:  "https://localhost:7102/ConnectionSocialNetworks/facebook/redirect",
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
	oauthStateString string
	urlUser          string = "https://graph.facebook.com/v8.0/me?access_token="
	urlMedia         string = "https://graph.facebook.com/v8.0/XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/media?fields=id,caption,media_type,media_url,username,owner,permalink,timestamp"
)

func LoginFacebook() string {
	Url, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}

	oauthStateString = random.GetRandom()

	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	Url.RawQuery = parameters.Encode()
	url := Url.String()
	return url
}

func GetFacebookMedia(token string) (rm.ResponseBody, *er.RestErr) {

	resp, err := http.Get(urlMedia + "&access_token=" + token)

	if err != nil {
		log.Fatalln(err)
		return rm.ResponseBody{}, er.NewInternalServerError(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var responseBody rm.ResponseBody
	json.Unmarshal(bodyBytes, &responseBody)
	return responseBody, nil
}

func GetOauthState() string {
	return oauthStateString
}

func GetUrlFacebookUser() string {
	return urlUser
}

func GetOauthConf() *oauth2.Config {
	return oauthConf
}
