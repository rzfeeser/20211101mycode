package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
)

func main() {
	
    client := http.Client{Timeout: 5 * time.Second}

    				// "GET"
    req, err := http.NewRequest(http.MethodGet, "https://example.com:443/secret-protected", http.NoBody)
    if err != nil {
        log.Fatal(err)
    }
			// un     // make believe pw
    req.SetBasicAuth("rzfeeser", "qwerty")

    // make the HTTP GET request (with auth in place)
    res, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer res.Body.Close()

    resBody, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Status: %d\n", res.StatusCode) // contains the HTTP response code
    fmt.Printf("Body: %s\n", string(resBody))  // contents of the HTTP body
}    
