//go:build admins || all || cp
// +build admins all cp

package modules

import "e-backend/internal/modules/admins"

func init() {
	m := admins.NewModule()
	EnabledModules = append(EnabledModules, m)
}
