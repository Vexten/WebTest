import (
	"errors"
	"net/http"
)

// Обрабатывает создание нового теста
func NewTest(w http.ResponseWriter, req *http.Request, sessUs *in.SessUs) error {
	if req.Method == "GET" {
		http.Redirect(w, req, `/`, http.StatusFound)
		return nil
	}
	if sessUs.UsId == in.GuestUserId {
		http.Redirect(w, req, `/`, http.StatusFound)
		return nil
	}
	tName := in.TestName(req.FormValue("TestName"))
	tDesc := in.TestDesc(req.FormValue("TestDesc"))
	err := db.CreateTest(tName, tDesc, sessUs.UsId)
	if err != nil {
		err = errors.New("Ошибка создания теста newTest " + err.Error())
		return err
	}
	http.Redirect(w, req, `/create`, http.StatusFound)
	return nil
}

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