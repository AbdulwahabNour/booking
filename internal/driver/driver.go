package driver

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)


  
type DB struct{ 
    Sql *sql.DB
}    
 
const (
    maxOpenDbConn =10
    maxIdleConn = 5
    maxDbLifetime = 5 * time.Minute
)

func  ConnectSql(dsn string)(*DB, error){
    db, err := NewDatabase(dsn)
    if err != nil{
        panic(err)
    }
    
    db.SetMaxOpenConns(maxOpenDbConn)
    db.SetMaxIdleConns(maxIdleConn)
    //The connection will expire 5  Minute after it was first created — not 5  Minute after it last became idle.
    db.SetConnMaxLifetime(maxDbLifetime)

 
    
    if err= db.Ping(); err != nil{
        return nil, err
    }
     
    return &DB{Sql: db}, nil
}

func NewDatabase(dsn string)(*sql.DB, error){
    db, err := sql.Open("postgres", dsn)
    if err != nil{
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
} 