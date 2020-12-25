package helpFun

import (
	"encoding/hex"
	"errors"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	db "github.com/richkule/prepareTestWeb/DBWorker"
	in "github.com/richkule/prepareTestWeb/init"
	"net/http"
)

// Генерирует сессию без коллизий в БД
func genSessId(id in.UsId) (*in.SessUs, error) {
	for true {
		sid := in.SessId(hex.EncodeToString(securecookie.GenerateRandomKey(32)))
		uid, err := db.GetUserId(sid)
		if err != nil {
			return nil, err
		}
		if uid == nil {
			uid = &in.SessUs{SessId: sid, UsId: id}
			p := len(sid)
			_ = p
			return uid, nil
		}
	}
	err := errors.New("Невозможный выход из цикла ")
	return nil, err
}
