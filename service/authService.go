package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var AUTH_SERVER_HOST = os.Getenv("AUTH_SERVER_HOST")
var AUTH_SERVER_IP = os.Getenv("AUTH_SERVER_IP")

func Authorize(r *http.Request) (bool, string, error) {
	cookie, err := r.Cookie("token")

	if err != nil {
		log.Println("Token not found")
		return false, "", nil
	}

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	postData := strings.NewReader("")

	fmt.Println("Making request to auth server for token validation: ", "http://"+AUTH_SERVER_HOST+":"+AUTH_SERVER_IP+"/login")
	req, err := http.NewRequest("POST", "http://"+AUTH_SERVER_HOST+":"+AUTH_SERVER_IP+"/login", postData)

	if err != nil {
		log.Println("Error while creating request to auth server", err)
		return false, "", err
	}

	req.AddCookie(cookie)

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error in response from auth server", err)
		return false, "", err
	}

	fmt.Println(resp)

	if resp.StatusCode != http.StatusOK {
		log.Println("Unexpected status code recieved from auth server", resp.StatusCode)
		return false, "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	role := string(bodyBytes)
	log.Println("user role", role)

	return true, role, nil
}
