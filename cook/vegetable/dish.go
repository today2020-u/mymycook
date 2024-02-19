package vegetable

import "github.com/dotbalo/logprint"

func MakeDish(name string) {
	msg := "做的素菜: " + name
	logprint.Info(msg)
}
