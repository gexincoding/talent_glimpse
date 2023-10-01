package smtp

import (
	"fmt"
	"os"
	"talent_glimpse/core/config"
	"testing"
)

func TestSMTP(t *testing.T) {
	fmt.Println(os.Getwd())
	err := config.Init("../../conf/conf.json")
	if err != nil {
		return
	}
	SendEmail([]string{"gexincoding@gmail.com"}, "Test", "test")

}
