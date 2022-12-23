package domain

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Event struct {
	repository EventRepository
	id         int64
	name       string
	method     string
	subscriber Subscriber
	withToken  bool
}

func NewEvent(name, method string) *Event {
	return &Event{
		name:      name,
		method:    method,
		withToken: false,
	}
}

func (e *Event) WithId(id int64) *Event {
	e.id = id
	return e
}

func (e *Event) WithName(name string) *Event {
	e.name = name
	return e
}

func (e *Event) WithMethod(method string) *Event {
	e.method = method
	return e
}

func (e *Event) WithToken(withToken bool) *Event {
	e.withToken = withToken
	return e
}

func (e *Event) WithSubscriber(subscriber Subscriber) *Event {
	e.subscriber = subscriber
	return e
}

func (e *Event) WithRepository(repository EventRepository) *Event {
	e.repository = repository
	return e
}

func (e *Event) Id() int64 {
	return e.id
}

func (e *Event) Name() string {
	return e.name
}

func (e *Event) Method() string {
	return e.method
}

func (e *Event) HasToken() bool {
	return e.withToken
}

func (e *Event) NameSubscriber() string {
	return e.subscriber.Name()
}

func (e *Event) ServerSubscriber() string {
	return e.subscriber.Server()
}

func (e *Event) Create() error {
	return e.repository.Create(*e)
}

func (e *Event) Delete() error {
	return e.repository.Delete(e.id)
}

func (e *Event) Update() error {
	return e.repository.Update(*e)
}

func (e *Event) Exists() bool {
	return e.repository.Exists(*e)
}

func (e *Event) HasValidToken(c *gin.Context) bool {
	token := e.GetTokenRequest(c)
	// url := e.ServerSubscriber() + "/validate-token"
	url := "http://localhost:8080/abix360/v1" + "/validate-token"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return string(body) == "{\"isValid\":true}"
}

func (e *Event) GetTokenRequest(c *gin.Context) string {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		return ""
	}
	tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])
	return tokenString
}

func ListEvents(repository EventRepository) []Event {
	return repository.AllEvents()
}

func FindEventById(id int64, repository EventRepository) Event {
	return repository.FindById(id)
}

func FindEventByName(name, method string, repository EventRepository) Event {
	return repository.FindByName(*NewEvent(name, method))
}
