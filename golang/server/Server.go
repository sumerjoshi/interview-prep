package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{
}

func getGreeting() string {
	content := map[int]rune {
		13: 'r',
		15: 'a',
		17: 'a',
		16: 'd',
		19: '!',
		4: 'o',
		5: 'm',
		8: 't',
		18: ' ',
		6: 'e',
		1: 'e',
		2: 'l',
		3: 'c',
		14: 'k',
		10: ' ',
		7: ' ',
		12: 'e',
		9: 'o',
		0: 'W',
		11: 'V',
	}
	result := []rune{}
        for i := 0; i < len(content); i++ {
        	result = append(result, content[i])
        }
        return string(result)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Printf("%s request with path: %s", r.Method, r.URL.Path)
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		// main endpoint for the web clients to get the logs
		case "/greeting":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, getGreeting())))
		default:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error": "not found"}`))
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "not found"}`))
	}
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
