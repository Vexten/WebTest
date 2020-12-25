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
