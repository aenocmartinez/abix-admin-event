package mysql

import (
	"abix360/database"
	"abix360/src/domain"
	"bytes"
	"log"

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
	query.WriteString("INSERT INTO events(name, subscriber, method) VALUES (?, ?, ?)")

	stmt, err := e.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-admin-event / EventDao / Create / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(event.Name(), event.NameSubscriber(), event.Method())
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
	query.WriteString("UPDATE events SET name=?, subscriber=?, method=?, updated_at=NOW() WHERE id=?")
	stmt, err := e.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-admin-event / EventDao / Update / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(event.Name(), event.NameSubscriber(), event.Method(), event.Id())
	if err != nil {
		log.Println("abix-admin-event / EventDao / Update / stmt.Exec: ", err.Error())
	}
	return err
}

func (e *EventDao) AllEvents() []domain.Event {
	var events []domain.Event
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT e.id, e.name, e.subscriber, e.method, s.server ")
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
		var id int64
		rows.Scan(&id, &name, &subscriber, &method, &server)

		var event domain.Event = *domain.NewEvent(name, method, *domain.NewSubscriber(subscriber).WithServer(server)).WithId(id)
		events = append(events, event)
	}

	return events
}

func (e *EventDao) FindById(id int64) domain.Event {
	var event domain.Event
	var name, subscriber, method, server string
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT e.id, e.name, e.subscriber, e.method, s.server ")
	strQuery.WriteString("FROM events e ")
	strQuery.WriteString("INNER JOIN subscribers s on s.name = e.subscriber ")
	strQuery.WriteString("WHERE e.id = ?")

	row := e.db.Source().Conn().QueryRow(strQuery.String(), id)
	row.Scan(&id, &name, &subscriber, &method, &server)
	event = *domain.NewEvent(name, method, *domain.NewSubscriber(subscriber).WithServer(server)).WithId(id)

	return event
}

func (e *EventDao) FindByName(name string) domain.Event {
	var event domain.Event
	var subscriber, method, server string
	var id int64
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT e.id, e.name, e.subscriber, e.method, s.server ")
	strQuery.WriteString("FROM events e ")
	strQuery.WriteString("INNER JOIN subscribers s on s.name = e.subscriber ")
	strQuery.WriteString("WHERE e.name = ?")

	row := e.db.Source().Conn().QueryRow(strQuery.String(), name)
	row.Scan(&id, &name, &subscriber, &method, &server)
	event = *domain.NewEvent(name, method, *domain.NewSubscriber(subscriber).WithServer(server)).WithId(id)

	return event
}
