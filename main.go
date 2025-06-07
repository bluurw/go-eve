package main

import (
	"io"
	"log"
	"fmt"
	"net/http"
	"strings"
	"go-eve/utils"
	"go-eve/support"
)


// FUNCAO LER O CONTEUDO DO HEADERS


// FUNCAO LER O HTML

//
func headersToMap(h http.Header) map[string]string {
	result := make(map[string]string)
	for k, v := range h {
		result[k] = strings.Join(v, ", ")
	}
	return result
}


func main() {
	url := "https://bb.com.br"
	method := "GET"
	params := ""
	headers := map[string]string{
		"User-Agent": "Go-http-client/1.1",
	}
	timeout := 10

	status, resp, err := utils.Request(url, method, params, headers, timeout)
	if !status {
		log.Printf("Erro na requisição: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	// INFO
	fmt.Printf("\nURL: %s\n", resp.Request.URL)
	fmt.Printf("Protocol: %s\n", resp.Proto)
	fmt.Printf("Method: %s\n", resp.Request.Method)
	fmt.Printf("Status Code: %s\n", resp.Status)
	fmt.Printf("Cookies: %s\n\n", resp.Cookies())
	
	// HEADER
	headerMap := headersToMap(resp.Header)
	
	fmt.Printf("HEADERS COMO MAPA:\n")
	for k, v := range headerMap {
		fmt.Printf("%s : %s\n", k, v)
	}
	
	// fmt.Printf("\nBODY:\n") DEBUG
	fmt.Println()
	body, _ := io.ReadAll(resp.Body)
	html := string(body)
	
	
	techs := support.Crawler(html)
	for _, tech := range techs {
		fmt.Printf("\n %s \n", tech)
		fmt.Println("-----------------------")
		cves := support.SearchExploit(tech, 5)
		for _, cve := range cves {
			fmt.Printf("CVE: %s\n", cve[0])
			fmt.Printf("LINK: %s\n", cve[2])
			fmt.Printf("DESCRIÇÃO:\n%s\n\n", cve[1])
		}
	}
}
