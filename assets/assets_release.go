//go:build release || wasm

package assets

import "embed"

//go:embed world/*.bin
//go:embed images/Room_Builder_32x32.png
//go:embed images/Theme_Sorter_Black_Shadow_32x32/1_Generic_Black_Shadow_32x32.png
//go:embed images/Premade_Character_32x32_19.png
var embeddedFS embed.FS

func init() {
	FS = embeddedFS
}
