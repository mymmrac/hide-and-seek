//go:build !release && !wasm

package assets

import "os"

func init() {
	FS = os.DirFS("./assets")
}
