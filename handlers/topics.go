import (
	"errors"
	"net/http"
	"strconv"
)

// Обрабатывает создание новой темы
func NewTopic(w http.ResponseWriter, req *http.Request, sessUs *in.SessUs) error {
	if req.Method == "GET" || sessUs.UsId == in.GuestUserId {
		http.Redirect(w, req, `/`, http.StatusFound)
		return nil
	}
	tName := in.TopicName(req.FormValue("TopicName"))
	tDesc := in.TopicDesc(req.FormValue("TopicDesc"))
	strTestId := req.FormValue("TestId")
	intTestId, err := strconv.Atoi(strTestId)
	if err != nil {
		err = errors.New("Ошибка конвертации id теста newTopic " + err.Error())
		return err
	}
	testId := in.TestId(intTestId)
	err = db.CreateTopic(tName, tDesc, testId)
	if err != nil {
		err = errors.New("Ошибка создания темы newTopic " + err.Error())
		return err
	}
	http.Redirect(w, req, `/edit/test/`+strTestId, http.StatusFound)
	return nil
}