package fastcli

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

func make_request(url string, results chan int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	total_read := 0
	for {
		buf := make([]byte, 16384)
		num_bytes, err := resp.Body.Read(buf)
		if err != nil {
			total_read = total_read + num_bytes
			break
		}
		total_read = total_read + num_bytes
	}
	results <- total_read
}

func My_ip() {
	client_display := get_final_display()
	color.HiGreen("\nFast.com result: \n%s\n\n", client_display)
	res := get_nord_result()
	color.HiGreen("Nord.com result: \n%s\n\n", res)

}

func get_final_display() string {
	js_url := get_js_url()
	token := get_token(js_url)
	client_display := get_client_display(token)
	return client_display
}

func get_nord_result() string {
	url := "https://nordvpn.com/wp-admin/admin-ajax.php?action=get_user_info_data"
	response, err := http.Get(url)
	if err != nil {
		color.HiRed(url)
		log.Fatal(err)
	}
	var data map[string]interface{}
	responseData, err := ioutil.ReadAll(response.Body)
	// responseString := string(responseData)
	err = json.Unmarshal([]byte(responseData), &data)
	s := []string{data["city"].(string), data["country"].(string), data["ip"].(string), data["isp"].(string)}
	final_string := strings.Join(s, ", ")
	return final_string
}

func get_client_display(token string) string {
	s := []string{"https://api.fast.com/netflix/speedtest/v2?https=true&token=", token, "&urlCount=5"}
	fast_endpoint := strings.Join(s, "")
	response, err := http.Get(fast_endpoint)
	if err != nil {
		color.HiRed(fast_endpoint)
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	responseString := string(responseData)
	obj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(responseString), &obj); err != nil {
		log.Fatal(err)
	}
	client_ip := obj["client"].(map[string]interface{})["ip"].(string)
	client_location := obj["client"].(map[string]interface{})["location"]
	client_city := client_location.(map[string]interface{})["city"].(string)
	client_country := client_location.(map[string]interface{})["country"].(string)
	client_display := strings.Join([]string{client_city, client_country, client_ip}, ", ")

	return client_display

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
