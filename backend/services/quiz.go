package services

import (
	"context"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func CreateQuiz(topic_id string, title string, order int) error {
	q := "INSERT INTO quiz(quiz_id, topic_id, title, order) VALUES (gen_random_uuid(), @TopicID, @Title, @Order)"

	args := pgx.NamedArgs{
		"TopicID": topic_id,
		"Title":   title,
		"Order":   order,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}

func GetTopicQuizzes(topic_id string) ([]structs.Quiz, error) {
	var qs []structs.Quiz
	q := "SELECT * FROM quiz WHERE topic_id = @TopicID"

	args := pgx.NamedArgs{
		"TopicID": topic_id,
	}

	rows, err := config.Dbpool.Query(context.Background(), q, args)
	if err != nil {
		return []structs.Quiz{}, err
	}

	for rows.Next() {
		var q structs.Quiz
		rows.Scan(&q.QuizID, &q.TopicID, &q.Title, &q.Order)
		qs = append(qs, q)
	}

	return qs, nil
}

func UpdateQuiz(quiz structs.Quiz) error {
	q := "UPDATE quiz SET title = @Title, order = @Order WHERE quiz_id = @QuizID AND topic_id = @TopicID"

	args := pgx.NamedArgs{
		"Title":   quiz.Title,
		"Order":   quiz.Order,
		"QuizID":  quiz.QuizID,
		"TopicID": quiz.TopicID,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteQuiz(quiz structs.Quiz) error {
	q := "DELETE FROM quiz WHERE quiz_id = @QuizID AND topic_id = @TopicID"

	args := pgx.NamedArgs{
		"QuizID":  quiz.QuizID,
		"TopicID": quiz.TopicID,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}
