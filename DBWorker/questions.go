package DBWorker

import (
	"context"
	in "github.com/richkule/prepareTestWeb/init"
)

// Получает все темы теста
func GetQuestions(id in.TopicId) ([]in.Question, error) {
	sql := `select id, html, q_type
	from questions
	where topic_id = $1`
	qArr := make([]in.Question, 0)
	res, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		q := in.Question{}
		err := res.Scan(&q.QuestionId, q.Body, q.QuestionType)
		if err != nil {
			return nil, err
		}
		qArr = append(qArr, q)
	}
	return qArr, nil
}
