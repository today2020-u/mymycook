package vegetable

import "github.com/today2020-u/logprint"

func MakeDish(name string) {
	msg := "做的素菜: " + name
	logprint.Info(msg)
}
