//go:build cv || all
// +build cv all

package modules

import "e-backend/internal/modules/cv"

func init() {
	m := cv.NewModule()
	EnabledModules = append(EnabledModules, m)
}
