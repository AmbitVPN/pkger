// +build !debug

package pkger

import (
	"github.com/markbates/pkger/pkging"
)

func Apply(pkg pkging.Pkger, err error) error {
	gil.Lock()
	defer gil.Unlock()
	current = pkging.Wrap(current, pkg)
	return nil
}