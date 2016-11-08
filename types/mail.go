// mail
package types

import (
	"bytes"
	"regexp"
)

type Mail string

func (mail Mail) getName() (name []byte) {
	for k, v := range bytes.Split([]byte(mail), []byte("@")) {
		if k == 0 {
			name = v
		}
	}
	return name
}

func (mail Mail) getDomain() (domain []byte) {
	for k, v := range bytes.Split([]byte(mail), []byte("@")) {
		if k == 1 {
			domain = v
		}
	}
	return domain
}

func (mail Mail) String() string {
	return string(mail)
}

//Matched that  whether s is a email address or not
func parseMail(s string) Mail {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, s); !m {
		return Mail(0)
	}
	return Mail(s)
}
