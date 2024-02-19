package meat

import "github.com/today2020-u/logprint"

func MakeDish(name string) {
	msg := "做的荤菜: " + name
	logprint.Info(msg)
}
