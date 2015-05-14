package mgopath

import (
    "gopkg.in/mgo.v2"
    "net/url"
    "errors"
)

func Connect(mongoPath string) (session *mgo.Session, dbName string, err error) {
    dbConfig, err := url.Parse(mongoPath)
    dbName = ""
    if err != nil {
        return nil, dbName, err
    }

    if session, err = mgo.Dial(dbConfig.Host); err != nil {
        return nil, dbName, err
    }

    dbName = dbConfig.Path
    if len(dbName) < 2 {
        return nil, dbName, errors.New("No database name specified.")
    }
    dbName = dbConfig.Path[1:len(dbConfig.Path)]
    return session, dbName, err
}
