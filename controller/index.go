package controller

import (
	InitTemp "code/temps"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "index", nil)
}

func InitIndex(w http.ResponseWriter, r *http.Request){
	act.Code = r.FormValue("code")
	dechiffrage(act.Code)
	http.Redirect(w, r, "/result", http.StatusMovedPermanently)
}

func Result(w http.ResponseWriter, r *http.Request){
	InitTemp.Temp.ExecuteTemplate(w, "result", act)
}