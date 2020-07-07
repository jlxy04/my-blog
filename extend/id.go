package extend

import (
	uuid "github.com/iris-contrib/go.uuid"
	"log"
	"strings"
)

func GenerateId() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
	}
	uuidStr := strings.Replace(uuid.String(), "-", "", len(uuid.String()))
	return uuidStr
}
