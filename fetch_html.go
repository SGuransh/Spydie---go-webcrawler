// Starting comments
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://medium.com/@datajournal/parse-html-in-golang-83c882576a0a"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
