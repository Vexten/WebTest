// Возвращает true, если пользователь является автором теста
func CheckAuthorTest(testId in.TestId, id in.UsId) (bool, error) {
	sql := `select * from test_author where test_id = $1 and us_id = $2`
	res, err := db.Exec(context.Background(), sql, testId, id)
	if err != nil {
		return false, err
	}
	if res[len(res)-1] == '0' {
		return false, nil
	}
	return true, nil
}