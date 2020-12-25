import "context"

// Возвращает идентефикатор пользователя, если его сессия активна и существует
func GetUserId(sid in.SessId) (*in.SessUs, error) {
	var uid in.SessUs
	var isActive bool
	res := db.QueryRow(context.Background(), `select * from session where sess_id = $1 and active = 'true'`, sid)
	err := res.Scan(nil, &uid.SessId, &uid.UsId, nil, &isActive)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	if isActive {
		return &uid, nil
	} else {
		return nil, nil
	}
}

// Получает данные пользователя
func GetUserData(name in.UserName) (id in.UsId, pass string, err error) {
	res := db.QueryRow(context.Background(), `select * from login where login = $1`, name)
	passArr := make([]byte, 32)
	err = res.Scan(nil, &passArr, &id)
	if err != nil {
		return 0, "", err
	}
	return id, string(passArr[:]), nil
}
// Получает логин пользователя по его id
func GetUserName(id in.UsId) (in.UserName, error) {
	res := db.QueryRow(context.Background(), `select * from login where id = $1`, id)

	var uName in.UserName
	err := res.Scan(&uName, nil, nil)
	if err != nil {
		return "", err
	}
	return uName, nil
}