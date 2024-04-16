package main

import (
	"fmt"
	"net/http"

	ctoai "github.com/cto-ai/sdk-go/v2"
)

func main() {
	client := ctoai.NewClient()
	client.Ux.Print("starting the server")

	event := map[string]interface{}{
		"event_name":   "deployment",
		"event_action": "success",
		"branch":       "review_test",
		"repo":         "nitishkumar71/cto_staging_test",
	}

	err := client.Sdk.Track([]string{}, "", event)
	if err != nil {
		panic(err)
	}

	fmt.Println("Dummy Line")
	fmt.Println("Fibbo of 9: ", fibonacciDP(9))
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func fibonacciDP(n int) int {
	f := make([]int, n+2)
	f[0], f[1] = 0, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
