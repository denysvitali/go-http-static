package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:     "go-http-static",
	Short:   "A simple HTTP static file server",
	Long: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ADD8")).
		Bold(true).
		Render(`🚀 Go HTTP Static Server

A fast, simple HTTP server for serving static files with optional TLS support.
Built with Go's standard library for maximum performance and minimal dependencies.`),
	Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Bold(true)

		fmt.Fprintf(os.Stderr, "%s %v\n", errorStyle.Render("❌ Error:"), err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}