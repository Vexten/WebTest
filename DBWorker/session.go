import (
	"context"
	"time"
)

// Создает новую запись сессии
func CreateSessId(id *in.SessUs) error {
	sql := `insert into session(sess_id,user_id,last_activity,active) values($1,(select id from login where id = $2),$3,'true')`
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec(context.Background(), sql, id.SessId, id.UsId, timeNow)
	return err
}