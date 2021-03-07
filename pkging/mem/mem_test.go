package mem_test

import (
	"testing"

	"github.com/ambitvpn/pkger/pkging"
	"github.com/ambitvpn/pkger/pkging/mem"
	"github.com/ambitvpn/pkger/pkging/pkgtest"
)

func Test_Pkger(t *testing.T) {
	pkgtest.All(t, func(ref *pkgtest.Ref) (pkging.Pkger, error) {
		return mem.New(ref.Info)
	})
}
