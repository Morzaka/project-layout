package messaging

import (
	"log"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Body struct {
	Text string
}

type Message struct {
	Id       uuid.UUID
	SendTime time.Time
}

var validationRegexp *regexp.Regexp

func validateRegexp(v *regexp.Regexp) {
	valid, err := regexp.Compile(`([0-9]+)|(^(Fizz){1}|(Buzz){1}$)`)
	v = valid

	if err != nil {
		log.Panicln("validateRegexp error:", err)
	}
}
