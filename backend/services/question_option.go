package services

import (
	"context"
	"server/config"
	"server/structs"

	"github.com/jackc/pgx/v5"
)

func CreateQuestionOption(question_id string, body string, order int) error {
	q := "INSERT INTO question_option(option_id, question_id, body, order) VALUES(gen_random_uuid(), @QuizID, @Body, @Order)"

	args := pgx.NamedArgs{
		"QuestionID": question_id,
		"Body":       body,
		"Order":      order,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}

func GetQuestionOptions(question_id string) ([]structs.QuestionOption, error) {
	var options []structs.QuestionOption

	q := "SELECT * FROM question_option WHERE question_id = @QuestionID ORDER BY order ASC"

	args := pgx.NamedArgs{
		"QuestionID": question_id,
	}

	rows, err := config.Dbpool.Query(context.Background(), q, args)
	if err != nil {
		return []structs.QuestionOption{}, err
	}

	for rows.Next() {
		var option structs.QuestionOption
		rows.Scan(&option.OptionID, &option.QuestionID, &option.Body, &option.Order)
		options = append(options, option)
	}

	return options, nil
}

func UpdateQuestionOption(option structs.QuestionOption) error {
	q := "UPDATE question_option SET body = @Body, order = @Order WHERE option_id = @OptionID"

	args := pgx.NamedArgs{
		"Body":     option.Body,
		"Order":    option.Order,
		"OptionID": option.OptionID,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteQuestionOption(option_id string) error {
	q := "DELETE FROM question_option WHERE option_id = @OptionID"

	args := pgx.NamedArgs{
		"OptionID": option_id,
	}

	_, err := config.Dbpool.Exec(context.Background(), q, args)
	if err != nil {
		return err
	}

	return nil
}
