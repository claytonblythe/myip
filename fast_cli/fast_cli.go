package fast_cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Get_urls() {
	js_url := get_js_url()
	token := get_token(js_url)
	client_display, display_strings, _ := get_url_list(token)

	fmt.Printf("Client testing from %s\n\n", client_display)
	fmt.Println("Server locations:")
	for _, display_string := range display_strings {
		fmt.Println(display_string)
	}
}

func get_url_list(token string) (string, []string, []string) {
	s := []string{"https://api.fast.com/netflix/speedtest/v2?https=true&token=", token, "&urlCount=5"}
	url := strings.Join(s, "")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	responseString := string(responseData)
	obj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(responseString), &obj); err != nil {
		log.Fatal(err)
	}
	targets := obj["targets"].([]interface{})
	client_ip := obj["client"].(map[string]interface{})["ip"].(string)
	client_location := obj["client"].(map[string]interface{})["location"]
	client_city := client_location.(map[string]interface{})["city"].(string)
	client_country := client_location.(map[string]interface{})["country"].(string)
	client_display := strings.Join([]string{client_city, client_country, client_ip}, ", ")
	targets_display := []string{}
	target_urls := []string{}
	for _, target := range targets {
		target := target.(map[string]interface{})
		// element is the element from someSlice for where we are
		url := target["url"].(string)
		location := target["location"].(map[string]interface{})
		city := location["city"].(string)
		country := location["country"].(string)
		s := []string{city, country, url}
		target_display := strings.Join(s, ", ")
		targets_display = append(targets_display, target_display)
		target_urls = append(target_urls, url)
	}

	return client_display, targets_display, target_urls

}

func get_token(url string) string {

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	token := GetStringInBetween(responseString, "{https:!0,endpoint:apiEndpoint,token:", ",urlCount:5")
	token = token[1 : len(token)-1]
	return token
}

func get_js_url() string {
	url := "https://fast.com"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	out := GetStringInBetween(responseString, "<script src=", "></script>")
	out = out[1 : len(out)-1]
	s := []string{"https://fast.com", out}
	js_url := strings.Join(s, "")
	return js_url
}
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	if e == -1 {
		return
	}
	return str[s:e]
}
