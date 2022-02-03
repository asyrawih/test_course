package helper

import "log"

// HandleError function    
func HandleError(err error) error {
	if err != nil {
		log.Println(err)
	}
	return nil
}
