package utils

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// FUNCAO REQUISICAO
func Request(url string, method string, params string, headers map[string]string, timeout int) (bool, *http.Response, error) {
	
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
