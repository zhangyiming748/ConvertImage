package conv

import (
	"path"
	"testing"
)

func TestDir(t *testing.T) {
	fp := "/Users/zen/Downloads/Telegram Desktop/水岛津实/33.mp4"
	ret := path.Dir(fp)
	t.Log(ret)
}
