package main

import (
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	out := capturer.CaptureStdout(func() {
		err := run()

		assert.NoError(t, err, "it should end with no error")
	})

	assert.Contains(t, out, "All done!", "it should contain 'All done!'.")
}
