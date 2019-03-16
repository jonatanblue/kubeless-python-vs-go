package kubeless

import (
	"fmt"
	"github.com/kubeless/kubeless/pkg/functions"
	"github.com/sirupsen/logrus"
	"time"
)

// Hello sample function with dependencies
func Hello(event functions.Event, context functions.Context) (string, error) {
	logrus.Info(event)
	fmt.Println(time.Now())
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum++
	}
	fmt.Println(sum)
	fmt.Println(time.Now())
	return "Hello world!", nil
}
