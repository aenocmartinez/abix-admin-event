package domain

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MethodPost struct {
}

func (p *MethodPost) Invoke(c *gin.Context, event Event) (json string) {
	bodyPost, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	strURL := event.ServerSubscriber() + "/" + event.Name()
	request, _ := http.NewRequest("POST", strURL, bytes.NewReader(bodyPost))
	request.Header.Set("Content-Type", "application/json")

	if event.HasToken() {
		request.Header.Set("Authorization", "Bearer "+event.GetTokenRequest(c))
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)
}
