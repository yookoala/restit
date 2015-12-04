package restit_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yookoala/restit/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandString generate fix length random strings
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// dummyJSONStr returns dummy JSON string for test
func dummyJSONStr() string {
	return `{
		"foo": "bar",
		"hello": "world",
		"answer": 42
	}`
}

// dummyJSONTest tests the *restit.JSON created from json string
// by dummyJSONStr(). Returns error if the test fail
func dummyJSONTest(j *restit.JSON) (err error) {

	var n *restit.JSON

	if n, err = j.Get("foo"); err != nil {
		err = fmt.Errorf("unable to find foo (%s)\nraw: %#v",
			err.Error(), string(j.Raw()))
		return
	} else if want, have := "bar", n.String(); want != have {
		err = fmt.Errorf(".foo expected %#v, got %#v", want, have)
		return
	}

	if n, err = j.Get("hello"); err != nil {
		err = fmt.Errorf("unable to find hello (%s)\nraw: %#v",
			err.Error(), string(j.Raw()))
		return
	} else if want, have := "world", n.String(); want != have {
		err = fmt.Errorf(".hello expected %#v, got %#v", want, have)
		return
	}

	if n, err = j.Get("answer"); err != nil {
		err = fmt.Errorf("unable to find answer (%s)\nraw: %#v",
			err.Error(), string(j.Raw()))
		return
	} else if want, have := float64(42), n.Number(); want != have {
		err = fmt.Errorf(".answer expected %#v, got %#v", want, have)
		return
	}
	return
}

// dummyJSONStr2 returns dummy JSON string for test
func dummyJSONStr2() string {
	return `{
    "number": 1234.56,
    "string": "foo bar",
    "arrayOfString": [
      "one",
      "two",
      "three",
      "four"
    ],
    "object": {
      "foo": "bar",
      "hello": "world",
      "answer": 42
    },
    "true": true,
    "false": false,
    "null": null
  }`
}
