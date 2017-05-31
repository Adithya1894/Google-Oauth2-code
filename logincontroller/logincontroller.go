package logincontroller

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"./structs"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var conf *oauth2.Config

// RandToken generates a random @l length token.
func RandToken(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func init() {

	conf = &oauth2.Config{
		ClientID:     "Your Client Id",
		ClientSecret: "Your Client Secret Key",
		RedirectURL:  "http://127.0.0.1:9090/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
}

// AuthHandler handles authentication of a user and initiates a session.
func AuthHandler(c *gin.Context) {

	code := c.Request.URL.Query().Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Login failed. Please try again."})
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	u := structs.User{}
	if err = json.Unmarshal(data, &u); err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"message": "Error marshalling response. Please try agian."})
		return
	}
	c.HTML(http.StatusOK, "afterlogin.tmpl", gin.H{"email": u.Email})
}

// LoginHandler handles the login procedure.
func LoginHandler(c *gin.Context) {
	state := RandToken(32)
	session := sessions.Default(c)
	session.Set("state", state)
	log.Printf("Stored session: %v\n", state)
	session.Save()
	link := getLoginURL(state)
	c.HTML(http.StatusOK, "auth.tmpl", gin.H{"link": link})
}
