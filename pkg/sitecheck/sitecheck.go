package sitecheck

import (
	"fmt"
	"net/http"
	"time"
)

func Check(url string) string {
	client := &http.Client{
		Timeout: (800 * time.Millisecond),
	}

	_, err := client.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		return "unhealthy"
	}

	return "healthy"
}
