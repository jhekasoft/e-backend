package cmd

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	ColorSuccess = color.New(color.FgHiGreen)
	ColorError   = color.New(color.FgHiRed)
)

// Banner returns a formatted banner string with the provided version.
func Banner(subtitle, version, buildTime string) string {
	cLogo := color.New(color.FgBlue)
	cLogo2 := color.New(color.FgYellow)

	banner := cLogo.Sprintf(`
░█▀▀░█▀▄░█▀█░█▀▀░█░█░█▀▀░█▀█░█▀▄
░█▀▀░█▀▄░█▀█░█░░░█▀▄░█▀▀░█░█░█░█
░▀▀▀░▀▀░░▀░▀░▀▀▀░▀░▀░▀▀▀░▀░▀░▀▀`) + cLogo2.Sprintf("%s\n", subtitle)

	if version != "" {
		banner += cLogo2.Sprintf("Version: %s", version)
	}

	if buildTime != "" {
		// Add a separator if version is also provided
		if version != "" {
			banner += cLogo2.Sprintf(" | ")
		}
		banner += cLogo2.Sprintf("Build Time: %s", buildTime)
	}

	// Remove the first newline for better formatting
	return strings.Replace(banner, "\n", "", 1)
}

// CheckErr prints the msg with the prefix 'Error:' and exits with error code 1.
// If the msg is nil, it does nothing.
func CheckErr(msg interface{}) {
	if msg != nil {
		ColorError.Fprintln(os.Stderr, "Error:", msg)
		os.Exit(1)
	}
}
