import "context"

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