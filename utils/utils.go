package utils

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
)

var utilObj Utils

type Utils interface {
	GenerateUUID() (string, error)
	TimeNow() time.Time
	GenerateRandomNumber() string
}

type utils struct {
	cwd string
}

func NewUtils() Utils {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return &utils{
		cwd: pwd,
	}
}

func GetUtils() Utils {

	if utilObj == nil {
		utilObj = NewUtils()
	}
	return utilObj
}

func (u *utils) GenerateUUID() (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func (t *utils) TimeNow() time.Time {
	return time.Now().UTC()
}

func (u *utils) GenerateRandomNumber() string {
	min := 100000
	max := 999999
	number := rand.Intn(max-min) + min
	return fmt.Sprintf("%v", number)
}
