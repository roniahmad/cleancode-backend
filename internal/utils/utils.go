package utils

import (
	"bytes"
	"strings"
	"time"

	"math/rand"
)

// Available set "l" for lowercase, "u" for uppercase, "d" digit/numeric, "s" for special character
func RandomString(length int, sets string) string {
	set := map[string][]rune{
		"l": []rune("abcdefghjkmnpqrstuvwxyz"),
		"u": []rune("ABCDEFGHJKMNPQRSTUVWXYZ"),
		"d": []rune("0123456789"),
		"s": []rune("!@#$%&*?"),
	}
	var buf bytes.Buffer
	u := strings.Split(sets, "")
	for _, v := range u {
		buf.WriteString(string(set[v]))
	}

	letters := []rune(buf.String())
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}
