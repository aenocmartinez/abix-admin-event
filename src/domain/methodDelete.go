package domain

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MethodDelete struct {
}

func (g *MethodDelete) Invoke(c *gin.Context, event Event) (json string) {
	strURL := event.ServerSubscriber() + "/" + event.Name() + "/" + c.Param("id")
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", strURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	if event.HasToken() {
		req.Header.Set("Authorization", "Bearer "+event.GetTokenRequest(c))
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	return string(body)
}
