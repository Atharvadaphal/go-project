package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go-project/checker" // Importing our custom package
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://redhat.com",
		"https://fake-url.invalid",
		"https://httpbin.org/delay/5",
		"https://redhat.com", // This site intentionally delays for 5 seconds to test our timeout
	}

	ch := make(chan checker.Result)
	history := make(map[string]checker.Result)

	// Create a pointer to our configuration
	config := &checker.Config{
		Timeout: 3 * time.Second,
	}

	// Fire off goroutines
	for _, url := range urls {
		go checker.CheckURL(url, config, ch)
	}

	// Wait for results with a global timeout
	// If checking takes longer than 4 seconds total, we abort the remaining checks.
	globalTimeout := time.After(4 * time.Second)

CollectLoop:
	for i := 0; i < len(urls); i++ {
		select {
		case res := <-ch:
			history[res.URL] = res
			fmt.Println(res)
		case <-globalTimeout:
			fmt.Println("\n[!] Global timeout reached! Aborting remaining checks.")
			break CollectLoop // Breaks out of the for-loop entirely
		}
	}

	// Save to JSON
	saveJSON("status.json", history)
}

func saveJSON(filename string, data map[string]checker.Result) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("File error:", err)
		return
	}
	defer file.Close()

	// json.MarshalIndent makes the JSON pretty and readable
	// The "" is the prefix, and "    " (4 spaces) is the indent.
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}

	file.Write(jsonData)
	fmt.Println("\nResults saved neatly to", filename)
}
