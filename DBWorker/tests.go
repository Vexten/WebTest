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



// Возвращает структуру всех тестов доступных автору по его id
func GetTestsById(authorId in.UsId) ([]in.Test, error) {
	sql := `
select tests.id, tests.name, tests.desc, login.login, AVG(test_rate.rate) 
from tests join test_author on tests.id = test_author.test_id join login on test_author.us_id = login.id left join test_rate on tests.id = test_rate.tes_id 
where login.id = $1
group by tests.id, tests.name, tests.desc, login.login 
order by id
`
	tsArr := make([]in.Test, 0)
	res, err := db.Query(context.Background(), sql, authorId)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var floatPoint *float32
		ts := in.Test{}
		err := res.Scan(&ts.TestId, &ts.TestName, &ts.TestDesc, &ts.AuthorName, &floatPoint)
		if err != nil {
			return nil, err
		}
		if floatPoint == nil {
			ts.TestRate = 0
		} else {
			ts.TestRate = in.TestRate(*floatPoint)
		}
		tsArr = append(tsArr, ts)
	}
	return tsArr, nil
}
