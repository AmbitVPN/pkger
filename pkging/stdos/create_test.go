package stdos

import (
	"os"
	"testing"

	"github.com/ambitvpn/pkger/pkging/pkgtest"
	"github.com/stretchr/testify/require"
)

func Test_Pkger_Create(t *testing.T) {
	r := require.New(t)

	ref, err := pkgtest.NewRef()
	r.NoError(err)
	defer os.RemoveAll(ref.Dir)

	pkg, err := New(ref.Info)
	r.NoError(err)

	pkgtest.CreateTest(t, ref, pkg)
}
