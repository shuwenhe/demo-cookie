package main

import (
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{ // create cookie
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
	}
	cookie2 := http.Cookie{ // create cookie
		Name:     "user",
		Value:    "admin2",
		HttpOnly: true,
	}
	// w.Header().Set("Set-Cookie", cookie.String()) // 将cookie发送给浏览器
	// w.Header().Add("Set-Cookie", cookie2.String())
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

func main() {
	http.HandleFunc("/setCookie", setCookie)

	http.ListenAndServe(":8080", nil)
}
