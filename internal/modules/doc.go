//go:build doc || all
// +build doc all

package modules

import "e-backend/internal/modules/doc"

func init() {
	m := doc.NewModule()
	EnabledModules = append(EnabledModules, m)
}
