//go:build auth || all
// +build auth all

package modules

import "e-backend/modules/auth"

func init() {
	m := auth.NewModule()
	EnabledModules = append(EnabledModules, m)
}
