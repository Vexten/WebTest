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