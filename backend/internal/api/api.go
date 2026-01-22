package api

import (
	"backend/internal/schema"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
)

func App(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "error reading body"})
		return
	}

	var data schema.MsgPack

	if err := msgpack.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
