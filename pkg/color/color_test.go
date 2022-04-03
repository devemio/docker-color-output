package color

import (
	"testing"

	"docker-color-output/pkg/utils/assert"
)

//nolint:funlen
func TestColors(t *testing.T) {
	msg := "Message"

	t.Run("Black", func(t *testing.T) {
		assert.Equal(t, "\u001B[0;30m"+msg+"\u001B[0m", Black(msg))
	})

	t.Run("DarkGray", func(t *testing.T) {
		assert.Equal(t, "\u001B[1;30m"+msg+"\u001B[0m", DarkGray(msg))
	})

	t.Run("Red", func(t *testing.T) {
		assert.Equal(t, "\u001B[0;31m"+msg+"\u001B[0m", Red(msg))
	})

	t.Run("LightRed", func(t *testing.T) {
		assert.Equal(t, "\u001B[1;31m"+msg+"\u001B[0m", LightRed(msg))
	})

	t.Run("Green", func(t *testing.T) {
		assert.Equal(t, "\u001B[0;32m"+msg+"\u001B[0m", Green(msg))
	})

	t.Run("LightGreen", func(t *testing.T) {
		assert.Equal(t, "\u001B[1;32m"+msg+"\u001B[0m", LightGreen(msg))
	})

	t.Run("Brown", func(t *testing.T) {
		assert.Equal(t, "\u001B[0;33m"+msg+"\u001B[0m", Brown(msg))
	})

	t.Run("Yellow", func(t *testing.T) {
		assert.Equal(t, "\u001B[1;33m"+msg+"\u001B[0m", Yellow(msg))
	})

	t.Run("Blue", func(t *testing.T) {
		assert.Equal(t, "\u001B[0;34m"+msg+"\u001B[0m", Blue(msg))
	})

	t.Run("LightBlue", func(t *testing.T) {
		assert.Equal(t, "\u001B[1;34m"+msg+"\u001B[0m", LightBlue(msg))
	})

	t.Run("Purple", func(t *testing.T) {
		assert.Equal(t, "\u001B[0;35m"+msg+"\u001B[0m", Purple(msg))
	})

	t.Run("LightPurple", func(t *testing.T) {
		assert.Equal(t, "\u001B[1;35m"+msg+"\u001B[0m", LightPurple(msg))
	})

	t.Run("Cyan", func(t *testing.T) {
		assert.Equal(t, "\u001B[0;36m"+msg+"\u001B[0m", Cyan(msg))
	})

	t.Run("LightCyan", func(t *testing.T) {
		assert.Equal(t, "\u001B[1;36m"+msg+"\u001B[0m", LightCyan(msg))
	})

	t.Run("LightGray", func(t *testing.T) {
		assert.Equal(t, "\u001B[0;37m"+msg+"\u001B[0m", LightGray(msg))
	})

	t.Run("White", func(t *testing.T) {
		assert.Equal(t, "\u001B[1;37m"+msg+"\u001B[0m", White(msg))
	})
}
