package controller

import (
	InitTemp "code/temps"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "index", act)
}

func InitIndex(w http.ResponseWriter, r *http.Request) {
	act.Code = r.FormValue("code")
	act.Code = dechiffrage(act.Code)
	http.Redirect(w, r, "/index", http.StatusMovedPermanently)
}


