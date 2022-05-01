package utils

import (
	"log"
	"os"
)

func InitLog() error {
	writer, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	log.SetOutput(writer)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	return nil
}
