package helper

import "log"

func HandleError(err error) error {
	if err != nil {
		log.Println(err)
	}
	return nil
}
