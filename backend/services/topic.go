package services

import (
	"context"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func CreateTopic(name string) error {
	q := "INSERT INTO topic(topic_id, name) VALUES(gen_random_uuid(), @Name)"

	args := pgx.NamedArgs{
		"Name": name,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}

func GetTopics() ([]structs.Topic, error) {
	var topics []structs.Topic

	q := "SELECT * FROM question WHERE quiz_id = @QuizID ORDER BY order ASC"

	rows, err := config.Dbpool.Query(context.Background(), q)
	if err != nil {
		return []structs.Topic{}, err
	}

	for rows.Next() {
		var topic structs.Topic
		rows.Scan(&topic.TopicID, &topic.Name)
		topics = append(topics, topic)
	}

	return topics, nil
}

func UpdateTopic(topic structs.Topic) error {
	q := "UPDATE topic SET name = @Name WHERE topic_id = @TopicID"

	args := pgx.NamedArgs{
		"TopicID": topic.TopicID,
		"Name":    topic.Name,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTopic(topic_id string) error {
	q := "DELETE FROM topic WHERE topic_id = @TopicID"

	args := pgx.NamedArgs{
		"TopicID": topic_id,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}
