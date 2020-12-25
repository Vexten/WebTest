import "context"

// Получает все темы теста
func GetTopics(testId in.TestId) ([]in.Topic, error) {
	sql := `select id, name, "desc"
	from topic
	where test_id = $1`
	tsArr := make([]in.Topic, 0)
	res, err := db.Query(context.Background(), sql, testId)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		ts := in.Topic{}
		err := res.Scan(&ts.TopicId, &ts.TopicName, &ts.TopicDesc)
		if err != nil {
			return nil, err
		}
		tsArr = append(tsArr, ts)
	}
	return tsArr, nil
}

// Возвращает true, если пользователь является автором темы
func CheckAuthorTopic(topicId in.TopicId, usId in.UsId) (bool, error) {
	sql := `
select *
from topic join test_author on topic.test_id = test_author.test_id
where id = $1 and us_id = $2`
	res, err := db.Exec(context.Background(), sql, topicId, usId)
	if err != nil {
		return false, err
	}
	if res[len(res)-1] == '0' {
		return false, nil
	}
	return true, nil
}
