// Go Api server
// @jeffotoni

package handler

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/jeffotoni/go.standard.project.layout/project.web/standard.libray/crud.user/pkg/pkg/util"

	"github.com/jeffotoni/go.standard.project.layout/project.web/standard.libray/crud.user/pkg/pkg/zerolog"
)

var rp = rand.New(rand.NewSource(time.Now().UnixNano()))

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	defer r.Body.Close()
	delay := int32(rp.Float64() * 500)
	time.Sleep(time.Millisecond * time.Duration(delay))
	if strings.ToUpper(r.Method) != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		jsonstr := `{"msg":"O método permitido é Post!"}`
		w.Write([]byte(jsonstr))
		zerolog.Error(
			"1.0.0",
			"ping.go",
			27,
			"api.crud.user.com.br",
			"/ping",
			"Ping Test Msg Error, tem que ser method POST")
		return
	}

	jsonstr := util.BuildJSON(`{"msg":"pong","method":"`, strings.ToUpper(r.Method), `"}`)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonstr))
	return
}