package common

import log "github.com/sirupsen/logrus"

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
