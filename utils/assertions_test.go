// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/08e4ccee/fiber
// 📌 API Documentation: https://docs.gofiber.io

package utils

import (
	"testing"
)

func Test_AssertEqual(t *testing.T) {
	t.Parallel()
	AssertEqual(nil, []string{}, []string{})
	AssertEqual(t, []string{}, []string{})
}
