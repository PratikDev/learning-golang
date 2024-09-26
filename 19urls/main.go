package main

import (
	"fmt"
	"net/url"
)

const URL = "https://honest-feedback.vercel.app:3000/dashboard?theme=dark&nav=open"

func main() {
	parsedURL, err := url.Parse(URL)
	checkErrNil(err)

	qparams := parsedURL.Query()

	fmt.Println("URL Port:", parsedURL.Port())
	fmt.Println("URL path:", parsedURL.Path)

	for key, param := range qparams {
		fmt.Println(key, ":", param[0])
	}

	urlOne := &url.URL{
		Scheme:   "https",
		Host:     "somethinng.com",
		Path:     "dashboard",
		RawQuery: "theme=dark&user=null",
	}

	fmt.Println(urlOne.String())
}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
