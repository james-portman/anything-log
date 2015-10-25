package main

import (
    // "fmt"
    "code.google.com/p/go-sqlite/go1/sqlite3"
    "time"
    "net/http"
    "log"
)


type Log struct {
	Id int `json:"id"`
	Logtype string `json:"type"`
	Detail string `json:"detail"`
	Timestamp int64 `json:"timestamp"`
}

type Logs []Log
var logs Logs

type Db struct {
    conn *sqlite3.Conn
}
var db Db

func main() {

    // db := &Db{}
    db.open("sqlite.db")
    db.create()
    // db.add("something","somedetail")
    db.get()

    router := NewRouter()
	log.Fatal(http.ListenAndServe(":8081", router))

}



func (db *Db) open(file string) {
    db.conn, _ = sqlite3.Open(file)
    // todo: catch errors
}

func (db *Db) create() {
    db.conn.Exec("CREATE TABLE log(type text, detail text, timestamp int64)")
}

func (db *Db) add(logtype, detail string) {
    args := sqlite3.NamedArgs{"$type": logtype, "$detail": detail, "$timestamp": time.Now().Unix()}
    db.conn.Exec("INSERT INTO log VALUES($type, $detail, $timestamp)", args) // $c will be NULL
}

func (db *Db) get() {
    sql := "SELECT rowid, * FROM log"
    row := make(sqlite3.RowMap)

    // might be easier to do this but might leak memory:
    // logs = logs[:0]
    for range(logs) {
	logs[len(logs)-1] = Log{}
	logs = logs[:len(logs)-1]
    }


    for s, err := db.conn.Query(sql); err == nil; err = s.Next() {
        var rowid int
        s.Scan(&rowid, row)     // Assigns 1st column to rowid, the rest to row
        logtype, _ := row["type"].(string)
        detail, _ := row["detail"].(string)
        timestamp, _ := row["timestamp"].(int64)
        log := Log{Id: rowid, Logtype: logtype, Detail: detail, Timestamp: timestamp}
        logs = append(logs,log)
    }
}
