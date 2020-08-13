package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{ // create cookie
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
		MaxAge:   60,
	}
	cookie2 := http.Cookie{ // create cookie
		Name:     "user2",
		Value:    "admin2",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

func getCookie(w http.ResponseWriter, r *http.Request) { // 获取cookie
	// cookies := r.Header["Cookie"] // 获取请求头中所有cookie
	cookie, _ := r.Cookie("user") //如果想得到某一个cookie，可以直接调用
	fmt.Println("得到的Cookie有：", cookie)
}

func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)

	http.ListenAndServe(":8080", nil)
}
