package mgopath

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/ory-am/dockertest"
    "github.com/ory-am/common/rand/sequence"
)

func TestConnect(t *testing.T) {
    containerID, ip, port := dockertest.SetupMongoContainer(t)
    dbName, err := sequence.RuneSequence(22, []rune("abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ1234567890"))
    assert.Nil(t, err)
    defer containerID.KillRemove(t)
    path := fmt.Sprintf("mongodb://%s:%d/%s", ip, port, string(dbName))
    db, err := Connect(path)
    assert.NotNil(t, db)
    assert.Nil(t, err)
}
