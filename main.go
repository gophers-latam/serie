package main

import (
	"flag"
	. "github.com/gophers-latam/serie/app"
	. "github.com/gophers-latam/serie/math"
)

func main() {
	x := flag.Int("x", 1, "dado X retorna Y siguiendo la serie especificada en la documentacion")
	flag.Parse()
	env := setupApp()
	y := Compute(*x)
	env.InfoLog.Printf("y = %#v", y)
}

func setupApp() *Application {
	infoLogger := NewInfoLogger()
	errorLogger := NewErrorLogger()
	env := &Application{
		ErrorLog: errorLogger,
		InfoLog:  infoLogger,
	}
	return env
}
