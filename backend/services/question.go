package services

import (
	"context"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func CreateQuestion(quiz_id string, body string) error {
	q := "INSERT INTO question(quiz_id, question_id, body, answer) VALUES(@QuizID, gen_random_uuid(), @Body, 0)"

	args := pgx.NamedArgs{
		"QuizID": quiz_id,
		"Body":   body,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}

func GetQuestions(quiz_id string) ([]structs.Question, error) {
	var questions []structs.Question

	q := "SELECT * FROM question WHERE quiz_id = @QuizID ORDER BY order ASC"

	args := pgx.NamedArgs{
		"QuizID": quiz_id,
	}

	rows, err := config.Dbpool.Query(context.Background(), q, args)
	if err != nil {
		return []structs.Question{}, err
	}

	for rows.Next() {
		var question structs.Question
		rows.Scan(&question.QuizID, &question.QuestionID, &question.Body, &question.Answer)
		questions = append(questions, question)
	}

	return questions, nil
}

func UpdateQuestion(question structs.Question) error {
	q := "UPDATE question SET body = @Body, answer = @Answer WHERE quiz_id = @QuizID AND question_id = @QuestionID"

	args := pgx.NamedArgs{
		"QuizID":     question.QuizID,
		"QuestionID": question.QuestionID,
		"Body":       question.Body,
		"Answer":     question.Answer,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteQuestion(question structs.Question) error {
	q := "DELETE FROM question WHERE quiz_id = @QuizID AND question_id = @QuestionID"

	args := pgx.NamedArgs{
		"QuizID":     question.QuizID,
		"QuestionID": question.QuestionID,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}
