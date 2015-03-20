package env

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "os"
    "code.google.com/p/go-uuid/uuid"
)

func TestIfFallbackWorks(t *testing.T) {
    f := "bar"
    k := uuid.NewRandom().String()
    v := Getenv(k, f)
    assert.Equal(t, v, f)
}

func TestIfEnvWorks(t *testing.T) {
    f := "bar"
    ev := "foo"
    k := uuid.NewRandom().String()

    os.Setenv(k, ev)
    defer os.Unsetenv(k)

    v := Getenv(k, f)
    assert.NotEqual(t, v, f)
    assert.Equal(t, v, ev)
}

func BenchmarkGetEnv(b *testing.B) {
    k := uuid.NewRandom().String()
    for i := 0; i < b.N; i++ {
        _ =  Getenv(k, "")
    }
}