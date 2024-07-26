//go:build !release

package assets

import "os"

func init() {
	FS = os.DirFS("./assets")
}
