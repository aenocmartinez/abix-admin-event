package domain

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MethodGet struct {
}

func (g *MethodGet) Invoke(c *gin.Context, event Event) (json string) {
	fmt.Println("Entra al GET")
	parameters := g.getParameters(c.Request.URL.Query())
	url := event.ServerSubscriber() + "/" + event.Name() + "?" + parameters

	fmt.Println("url: ", url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

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

func (g *MethodGet) getParameters(data map[string][]string) string {
	strParams := g.convertParamsFormatMapToString(data)
	strParams = strings.ReplaceAll(strParams, "\"", "")
	strParams = strings.ReplaceAll(strParams, "[", "")
	strParams = strings.ReplaceAll(strParams, "]", "")
	strParams = strings.ReplaceAll(strParams, " ", "%20")
	strParams = strParams[:len(strParams)-1]
	return strParams
}

func (g *MethodGet) convertParamsFormatMapToString(m map[string][]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"&", key, value)
	}
	return b.String()
}
