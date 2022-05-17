package model

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestSession(t *testing.T) {
	t.Run("测试sessionAdd", TestAddSession)
	t.Run("test Query", TestQuerySession)
	t.Run("test ID GET", TestSessionIdGet)
}
func TestAddSession(t *testing.T) {
	session := &Sessions{
		Sessions_id: uuid.New().String(),
		Username:    "admin2",
		User_id:     1,
	}
	AddSession(session)
}

func TestDelSession(t *testing.T) {
	DelSession("1bfc5f22-71f9-4b02-a8fb-f4499d513b8b")
}

func TestQuerySession(t *testing.T) {
	sess, _ := QuerySession("52f2dbec-4397-4ffc-bd10-0534500f8722")
	fmt.Println(sess)
}

func TestSessionIdGet(t *testing.T) {
	sess, _ := SessionIdGet("8e0469d4-c6ef-4de9-bc7c-0069978a83f6")
	fmt.Println(sess)
}
