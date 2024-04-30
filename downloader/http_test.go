package downloader

import (
	"testing"

	"github.com/itzngga/fake-useragent/setting"
)

func TestDownload_Get(t *testing.T) {
	downloader := Download{
		Delay:   setting.HTTP_DELAY,
		Timeout: setting.HTTP_TIMEOUT,
	}

	_, err := downloader.Get("https://developers.whatismybrowser.com")
	if err != nil {
		t.Errorf("downloader.Get err: %v", err)
	}
}
