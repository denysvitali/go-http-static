package main

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/denysvitali/go-http-static/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	port          string
	listenAddress string
	enableTLS     bool
	certFile      string
	certKey       string
	logger        *logrus.Logger
)

var serveCmd = &cobra.Command{
	Use:   "serve [path]",
	Short: "Start the HTTP static file server",
	Long: `Start the HTTP static file server to serve files from the specified directory.

The server will serve all files in the given directory and its subdirectories.
Use --tls flag along with --cert and --key to enable HTTPS.`,
	Args: cobra.ExactArgs(1),
	Run:  runServe,
}

func runServe(cmd *cobra.Command, args []string) {
	path := args[0]

	// Convert to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		logger.WithError(err).Fatal("Failed to resolve absolute path")
	}

	// Validate TLS configuration
	if enableTLS {
		if certFile == "" {
			logger.Fatal("Certificate file is required when TLS is enabled (use --cert)")
		}
		if certKey == "" {
			logger.Fatal("Certificate key is required when TLS is enabled (use --key)")
		}
	}

	// Create server
	srv := server.New(absPath, port, listenAddress, enableTLS, certFile, certKey)

	// Display styled startup messages
	displayStartupMessage(srv)

	// Log server start
	logger.WithFields(logrus.Fields{
		"address": srv.Address(),
		"path":    absPath,
		"tls":     enableTLS,
	}).Info("Starting HTTP server")

	// Start server
	if err := srv.Start(); err != nil {
		logger.WithError(err).Fatal("Server failed to start")
	}
}

func displayStartupMessage(srv *server.Server) {
	// Styled startup message using lipgloss
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ADD8")).
		Bold(true).
		Padding(1, 0, 0, 0)

	infoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Italic(true)

	pathStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F7DF1E")).
		Bold(true)

	addrStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#61DAFB")).
		Bold(true)

	fmt.Println(headerStyle.Render("🚀 Go HTTP Static Server"))

	fmt.Printf("%s Serving directory: %s\n",
		infoStyle.Render("📂"),
		pathStyle.Render(srv.Path))

	protocol := "http"
	icon := "🌐"
	if srv.TLS {
		protocol = "https"
		icon = "🔒"
	}

	fmt.Printf("%s %s server listening on: %s\n",
		infoStyle.Render(icon),
		protocol,
		addrStyle.Render(fmt.Sprintf("%s://%s", protocol, srv.Address())))
}

func init() {
	// Initialize logger
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	rootCmd.AddCommand(serveCmd)

	// Flags
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to listen on")
	serveCmd.Flags().StringVarP(&listenAddress, "listen", "l", "", "Listen address (default: all interfaces)")
	serveCmd.Flags().BoolVarP(&enableTLS, "tls", "t", false, "Enable HTTPS/TLS")
	serveCmd.Flags().StringVarP(&certFile, "cert", "c", "", "TLS certificate file")
	serveCmd.Flags().StringVarP(&certKey, "key", "k", "", "TLS private key file")

	// Mark TLS flags as required when TLS is enabled
	serveCmd.MarkFlagsRequiredTogether("tls", "cert", "key")
}