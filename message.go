package elogrus

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

// Msg is key of message to detect message on using {method}J()
//methods of echo Logger.
// e.g context.Logger().InfoJ(log.JSON{elogrus.Msg:"my_message","foo":"bar"})
var Msg = "__msg__"

func extractMsg(j log.JSON) string {
	if msg, ok := j[Msg]; ok {
		delete(j, Msg)

		return fmt.Sprint(msg)
	}

	return ""
}
