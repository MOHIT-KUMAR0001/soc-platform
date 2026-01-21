package api

import (
	"backend/internal/ingest"
	"backend/internal/schema"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
)

var fakeDatabase = make(map[string]schema.ScanIP)

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

	for log := range len(data.Events) {
		dec, _ := ingest.ParseRawLog(data.Events[log].Raw,
			data.Events[log].AgentID,
			data.Events[log].Hostname,
		)

		// Check if entry exists
		if entry, ok := fakeDatabase[dec.SourceIP]; ok {
			entry.Source.Count++
			if entry.Source.Count > 2 {
				fmt.Println("High alert Generated someone is try to bruiteforcing the application .....")
				os.Exit(1)
			}
			fmt.Println("this is entry", entry.Source.Count)
			// Key exists - update the count
			entry.Source.LastSeen = time.Now().Format(time.RFC3339) // optional
			fakeDatabase[dec.SourceIP] = entry
		} else {
			// Key doesn't exist - create a new entry
			fakeDatabase[dec.SourceIP] = schema.ScanIP{
				Source: schema.Element{
					Count:     1,
					FirstSeen: time.Now().Format(time.RFC3339),
					LastSeen:  time.Now().Format(time.RFC3339),
				},
			}
		}
	}

	// Iterate over the map
	for key, value := range fakeDatabase {
		fmt.Printf("IP: %s, Count: %d\n", key, value.Source.Count)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
