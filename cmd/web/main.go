package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"zametki/pkg/models/mysql"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	zametkis      *mysql.ZametModel
	templateCache map[string]*template.Template
	ID            map[int]int
	ChanAuth      chan int
}

func main() {

	addr := flag.String("addr", ":8000", "Сетевой адрес веб-сервера")
	dsn := flag.String("mysql", "root:@/zametki?parseTime=true", "Название MySQL источника данных")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Инициализируем новый кэш шаблона...
	templateCache, err := newTemplateCache("ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	ID := make(map[int]int)
	ChanAuth := make(chan int)

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		zametkis:      &mysql.ZametModel{DB: db},
		templateCache: templateCache,
		ID:            ID,
		ChanAuth:      ChanAuth,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск сервера на http://127.0.0.1%s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
