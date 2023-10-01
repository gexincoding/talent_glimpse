package smtp

import (
	"talent_glimpse/core/config"
	"testing"
)

func TestSMTP(t *testing.T) {
	err := config.Init()
	if err != nil {
		return
	}

	SendEmail([]string{"gexincoding@gmail.com"}, "Test", "test")

}
