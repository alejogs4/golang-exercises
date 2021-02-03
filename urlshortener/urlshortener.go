package urlshortener

import (
	"fmt"
	"net/http"
	"os"

	"example.com/alejogs4/learning/urlshortener/handlers"
)

// RunHTTPServer will run http for url shortener
func RunHTTPServer() {
	muxServer := createServer()
	urlsPath := map[string]string{
		"/best-company":   "https://www.s4n.co/",
		"/best-team-ever": "https://www.realmadrid.com/en",
	}

	mapHandler := handlers.MapHandler(urlsPath, muxServer)
	yamlData := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, error := handlers.YamlHandler([]byte(yamlData), mapHandler)
	if error != nil {
		fmt.Println(fmt.Sprintf("Error creating the yaml handler %d", error))
		os.Exit(1)
		return
	}

	fmt.Println("Listening in port 8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func createServer() *http.ServeMux {
	muxServer := http.NewServeMux()
	muxServer.HandleFunc("/", homeHandler)

	return muxServer
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main handler")
}
