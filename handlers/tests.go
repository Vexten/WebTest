import (
	"errors"
	"net/http"
)

// Обрабатывает страницу редактирования теста
func editTest(w http.ResponseWriter, req *http.Request, sessUs *in.SessUs, testId in.TestId) error {
	ok, err := db.CheckAuthorTest(testId, sessUs.UsId)
	if err != nil {
		err = errors.New("Ошибка получения записи автора из БД editTest " + err.Error())
		return err
	}
	if !ok {
		http.Redirect(w, req, `/`, http.StatusFound)
		return nil
	}
	data := in.DataEditTest{}
	if data.Header, err = renderHeader(sessUs.UsId); err != nil {
		err = errors.New("Ошибка обработки шаблона шапки editTest " + err.Error())
		return err
	}
	data.Topics, err = db.GetTopics(testId)
	if err != nil {
		err = errors.New("Ошибка получения тем editTest " + err.Error())
		return err
	}
	data.TestId = testId
	err = renderTemplate(w, in.EditTestPage, data)
	if err != nil {
		err = errors.New("Ошибка обработки шаблона editTestPage " + err.Error())
	}
	return err
}