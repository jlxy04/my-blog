package extend

import (
	"github.com/satori/go.uuid"
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
