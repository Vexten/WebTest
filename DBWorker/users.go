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