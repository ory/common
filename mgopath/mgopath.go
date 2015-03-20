package mgopath

import (
    "gopkg.in/mgo.v2"
    "net/url"
    "errors"
    "log"
)

func Connect(mongoPath string) (*mgo.Database, error) {
    dbConfig, err := url.Parse(mongoPath)
    if err != nil {
        return nil, err
    }

    log.Printf("Connecting to %s", dbConfig.Host)
    sess, err := mgo.Dial(dbConfig.Host)
    if err != nil {
        return nil, err
    }

    dbName := dbConfig.Path
    log.Printf("Using database %s", dbName)
    if len(dbName) < 2 {
        return nil, errors.New("No database name specified.")
    }

    dbName = dbConfig.Path[1:len(dbConfig.Path)]
    return sess.DB(dbName), err
}
