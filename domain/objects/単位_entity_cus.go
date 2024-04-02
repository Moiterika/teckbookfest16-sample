package objects

import (
	"fmt"
	"unicode/utf8"
)

func (e *Ent単位) Validate() error {
	// please, write validation logic here
	// これはエラーにさせる例
	if utf8.RuneCountInString(e.Getコード.String()) > 8 {
		return fmt.Errorf("単位コードは8文字までにしてください。: %w", ErrOverMaxLength)
	}

	return nil
}
