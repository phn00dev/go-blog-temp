package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func Set(ctx *gin.Context, key, value string) {
	session := sessions.Default(ctx)
	session.Set(key, value)
	err := session.Save()
	if err != nil {
		log.Fatal("session writer errors :", err)
		return
	}
}

func Flash(ctx *gin.Context, key string) string {
	session := sessions.Default(ctx)
	response := session.Get(key)
	err := session.Save()
	if err != nil {
		log.Fatal("session writer errors :", err)
		return ""
	}
	session.Delete(key)
	err = session.Save()
	if err != nil {
		log.Fatal("session writer errors :", err)
		return ""
	}
	if response != nil {
		return response.(string)
	}
	return ""
}
