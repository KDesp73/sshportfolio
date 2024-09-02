package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sshportfolio/internal/tui"
	"strings"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/joho/godotenv"
)


func getHostAndPort() (string, string) {
	var (
		host string
		port string
		defaultHost = "localhost"
		defaultPort = "23234"
		flagHost string 
		flagPort string
		envHost string 
		envPort string	
	)
	flag.StringVar(&flagHost, "host", flagHost, "Specify ip to serve")
	flag.StringVar(&flagPort, "port", flagPort, "Specify port to serve")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load environment variable: %v", err)
		os.Exit(1)
	}

	envHost = os.Getenv("HOST")
	envPort = os.Getenv("PORT")

	if flagHost != host {
		host = flagHost
	} else if strings.TrimSpace(envHost) != "" {
		host = envHost
	} else {
		host = defaultHost
	}

	if flagPort != port {
		port = flagPort
	} else if strings.TrimSpace(envPort) != "" {
		port = envPort
	} else {
		port = defaultPort
	}

	return host, port
}

func main(){
	host, port := getHostAndPort()

	srv, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(), // Bubble Tea apps usually require a PTY.
			logging.Middleware(),
		),
			
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Starting SSH server", "host", host, "port", port)
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			// We ignore ErrServerClosed because it is expected.
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()

	log.Info("Stopping SSH server")
	if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	// This should never fail, as we are using the activeterm middleware.
	// pty, _, _ := s.Pty()
	// renderer := bubbletea.MakeRenderer(s)
	
	m := tui.NewModel()

	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
