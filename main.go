package main

import (
	"io"
	"log"
	"fmt"
	"net/http"
	"strings"
	"time"
	"go-eve/support"
)

// FUNCAO REQUISICAO
func request(url string, method string, params string, headers map[string]string, timeout int) (bool, *http.Response, error) {
	
	var body io.Reader
	
	if params != "" {
		body = strings.NewReader(params)
	}
	
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println("Erro ao criar requisição:", err)
		return false, nil, err
	}
	
	// Aplica configuracao de headers
	for k, v := range headers {
		r.Header.Set(k, v)
	}
	
	client := &http.Client {
		Timeout: time.Duration(timeout) * time.Second,
	}
	
	resp, err := client.Do(r)
	if err != nil {
		log.Println("Erro ao enviar requisição:", err)
		return false, nil, err
	}
	
	return true, resp, nil
}

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

	status, resp, err := request(url, method, params, headers, timeout)
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
	
	fmt.Printf("\nBODY:\n")
	fmt.Println()
	body, _ := io.ReadAll(resp.Body)
	html := string(body)
	
	
	techs := support.Crawler(html)
	fmt.Println(techs)
}	
	
	
	// POST
	// data := bytes.NewBufferString(`{"name":"Copilot","age":1}`)
	
	// HEADERS
	// headers := map[string]string{
	// 	"User-Agent": "MeuApp/1.0",
	//	"Authorization": "Bearer TOKEN_AQUI",
	// }

	// UTEIS
	// resp.status
	// resp.URL
	// resp.Method
	// resp.Header
