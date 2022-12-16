package mysql

import (
	"abix360/database"
	"abix360/src/domain"
	"bytes"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type SubscriberDao struct {
	db *database.ConnectDB
}

func NewSubscriberDao() *SubscriberDao {
	return &SubscriberDao{
		db: database.Instance(),
	}
}

func (s *SubscriberDao) Create(subscriber domain.Subscriber) error {
	var query bytes.Buffer
	query.WriteString("INSERT INTO subscribers(name, server) VALUES (?,?)")

	stmt, err := s.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-admin-event / SubscriberDao / Create / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(subscriber.Name(), subscriber.Server())
	if err != nil {
		log.Println("abix-admin-event / SubscriberDao / Create / stmt.Exec: ", err.Error())
	}

	return err
}

func (s *SubscriberDao) Delete(name string) error {
	var query bytes.Buffer
	query.WriteString("DELETE FROM subscribers WHERE name = ?")

	stmt, err := s.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-admin-event / SubscriberDao / Delete / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(name)
	if err != nil {
		log.Println("abix-admin-event / SubscriberDao / Delete / stmt.Exec: ", err.Error())
	}

	return err
}

func (s *SubscriberDao) FindByName(name string) domain.Subscriber {
	var subscriber domain.Subscriber
	var server string
	var cad bytes.Buffer

	cad.WriteString("SELECT name, server FROM subscribers WHERE name = ?")
	row := s.db.Source().Conn().QueryRow(cad.String(), name)
	row.Scan(&name, &server)
	subscriber = *domain.NewSubscriber(name, server)

	return subscriber
}

func (s *SubscriberDao) AllSubscribers() []domain.Subscriber {
	var subscribers []domain.Subscriber
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT name, server FROM subscribers order by name")
	rows, err := s.db.Source().Conn().Query(strQuery.String())
	if err != nil {
		log.Println("abix-admin-event / SubscriberDao / AllSubscribers / s.db.Source().Conn().Query: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name, server string
		rows.Scan(&name, &server)

		subscribers = append(subscribers, *domain.NewSubscriber(name, server))
	}

	return subscribers
}
