package sequence

import (
    "regexp"
    "github.com/stretchr/testify/assert"
    "testing"
    "fmt"
)

func TestRuneSequenceMatchesPattern(t *testing.T) {
    length := 25
    patterns := [][]rune{
        []rune("abcdefghijklmnopqrstuvwxyz"),
        []rune("abcdefghijklmnopqrstuvwxyz1234567890"),
    }
    regx := []string{
        fmt.Sprintf("[a-z]{%d}", length),
        fmt.Sprintf("[a-z0-9]{%d}", length),
    }

    for k, v := range patterns {
        seq, err := RuneSequence(length, v)
        assert.Nil(t, err)
        assert.Equal(t, length, len(seq))

        valid, err := regexp.Match(regx[k], []byte(string(seq)))
        assert.Nil(t, err)
        assert.True(t, valid)
    }
}

func TestRuneSequenceIsPseudoUnique(t *testing.T) {
    // 1 in 100 probability of collision
    times := 9000
    runes := []rune("ab")
    length := 32
    s := make(map[string]bool)

    for i := 0; i < times; i++ {
        k, err := RuneSequence(length, runes)
        assert.Nil(t, err)
        ks := string(k)
        _, ok := s[ks]
        assert.False(t, ok)
        if ok {
            return
        }
        s[ks] = true
    }
}

func BenchmarkTestInt64(b *testing.B) {
    length := 25
    pattern := []rune("abcdefghijklmnopqrstuvwxyz")
    for i := 0; i < b.N; i++ {
        RuneSequence(length, pattern)
    }
}
