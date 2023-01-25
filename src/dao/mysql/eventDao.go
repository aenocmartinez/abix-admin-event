package mysql

import (
	"bytes"
	"fmt"
	"log"
	"pulzo/database"
	"pulzo/src/domain"

	_ "github.com/go-sql-driver/mysql"
)

type EventDao struct {
	db *database.ConnectDB
}

func NewEventDao() *EventDao {
	return &EventDao{
		db: database.Instance(),
	}
}

func (e *EventDao) Create(event domain.Event) error {
	var query bytes.Buffer
	query.WriteString("INSERT INTO events(name, subscriber, method, with_token) VALUES (?, ?, ?, ?)")

	stmt, err := e.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-admin-event / EventDao / Create / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(event.Name(), event.NameSubscriber(), event.Method(), event.HasToken())
	if err != nil {
		log.Println("abix-admin-event / EventDao / Create / stmt.Exec: ", err.Error())
	}
	return err
}

func (e *EventDao) Delete(id int64) error {
	var query bytes.Buffer
	query.WriteString("DELETE FROM events WHERE id = ?")

	stmt, err := e.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-admin-event / EventDao / Delete / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("abix-admin-event / EventDao / Delete / stmt.Exec: ", err.Error())
	}

	return err
}

func (e *EventDao) Update(event domain.Event) error {
	var query bytes.Buffer
	query.WriteString("UPDATE events SET name=?, subscriber=?, method=?, with_token=?, updated_at=NOW() WHERE id=?")
	stmt, err := e.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-admin-event / EventDao / Update / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(event.Name(), event.NameSubscriber(), event.Method(), event.HasToken(), event.Id())
	if err != nil {
		log.Println("abix-admin-event / EventDao / Update / stmt.Exec: ", err.Error())
	}
	return err
}

func (e *EventDao) AllEvents() []domain.Event {
	var events []domain.Event
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT e.id, e.name, e.subscriber, e.method, s.server, e.with_token ")
	strQuery.WriteString("FROM events e ")
	strQuery.WriteString("INNER JOIN subscribers s on s.name = e.subscriber ")
	strQuery.WriteString("order by e.subscriber")
	rows, err := e.db.Source().Conn().Query(strQuery.String())
	if err != nil {
		log.Println("abix-admin-event / SubscriberDao / AllEvents / s.db.Source().Conn().Query: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name, subscriber, method, server string
		var withToken bool
		var id int64
		rows.Scan(&id, &name, &subscriber, &method, &server, &withToken)

		var event domain.Event = *domain.NewEvent(name, method).WithId(id).WithToken(withToken)
		event.WithSubscriber(*domain.NewSubscriber(subscriber).WithServer(server))
		events = append(events, event)
	}

	return events
}

func (e *EventDao) FindById(id int64) domain.Event {
	var event domain.Event
	var name, subscriber, method, server string
	var withToken bool
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT e.id, e.name, e.subscriber, e.method, s.server, e.with_token ")
	strQuery.WriteString("FROM events e ")
	strQuery.WriteString("INNER JOIN subscribers s on s.name = e.subscriber ")
	strQuery.WriteString("WHERE e.id = ?")

	row := e.db.Source().Conn().QueryRow(strQuery.String(), id)
	row.Scan(&id, &name, &subscriber, &method, &server, &withToken)
	event = *domain.NewEvent(name, method).WithId(id).WithToken(withToken)
	event.WithSubscriber(*domain.NewSubscriber(subscriber).WithServer(server))

	return event
}

func (e *EventDao) FindByName(ev domain.Event) domain.Event {
	var event domain.Event
	var name, subscriber, method, server string
	var id int64
	var withToken bool
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT e.id, e.name, e.subscriber, e.method, s.server, e.with_token ")
	strQuery.WriteString("FROM events e ")
	strQuery.WriteString("INNER JOIN subscribers s on s.name = e.subscriber ")
	strQuery.WriteString("WHERE e.name = ? and e.method = ?")

	row := e.db.Source().Conn().QueryRow(strQuery.String(), ev.Name(), ev.Method())
	row.Scan(&id, &name, &subscriber, &method, &server, &withToken)
	event = *domain.NewEvent(name, method).WithId(id).WithToken(withToken)
	event.WithSubscriber(*domain.NewSubscriber(subscriber).WithServer(server))

	return event
}

func (e *EventDao) Exists(event domain.Event) bool {
	var exists string = ""
	var strQuery bytes.Buffer
	strQuery.WriteString("SELECT 'T' FROM events WHERE name = ? and method = ?")

	row := e.db.Source().Conn().QueryRow(strQuery.String(), event.Name(), event.Method())
	err := row.Scan(&exists)
	if err != nil {
		fmt.Println(err)
	}

	return exists != ""
}
