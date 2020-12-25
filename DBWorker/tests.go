package DBWorker

import (
	"context"

	in "github.com/richkule/prepareTestWeb/init"
)

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

// Создает новый тест
func CreateTest(testName in.TestName, testDesc in.TestDesc, usId in.UsId) error {
	sql := `
with rows as (
INSERT INTO tests (name,"desc") VALUES ($1, $2) RETURNING id
)
INSERT INTO test_author values(
(SELECT id FROM rows),
$3
)`
	_, err := db.Exec(context.Background(), sql, testName, testDesc, usId)
	return err
}
