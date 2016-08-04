package ltkgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLtkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ltkgo Suite")
}
