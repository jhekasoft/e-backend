//go:build doc || all
// +build doc all

package modules

import "e-backend/modules/doc"

func init() {
	m := doc.NewModule()
	EnabledModules = append(EnabledModules, m)
}
