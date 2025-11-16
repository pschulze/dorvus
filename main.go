package main

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"math/big"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/favicon.ico", fs)

	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	seed := fiveMinInterval(time.Now()).Format(time.Stamp)
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, unlockCode(seed))
}

func fiveMinInterval(t time.Time) time.Time {
	minute := (t.Minute() / 5) * 5
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, 0, 0, t.Location())
}

func unlockCode(foo string) string {
	bar := "secret" + foo
	hash := sha256.Sum256([]byte(bar))
	hashInt := new(big.Int).SetBytes(hash[:])
	unlockCodeInt := new(big.Int).Mod(hashInt, big.NewInt(100_000))
	unlockCode := fmt.Sprintf("%05s", unlockCodeInt.Text(10))
	return unlockCode
}
