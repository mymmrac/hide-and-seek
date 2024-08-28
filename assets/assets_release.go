//go:build release || wasm

package assets

import "embed"

//go:embed *
var embeddedFS embed.FS

func init() {
	FS = embeddedFS
}
