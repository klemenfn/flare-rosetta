package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

const (
	online  = "online"
	offline = "offline"
)

var (
	mode          string
	flareBin      string
	flareConfig   string
	rosettaBin    string
	rosettaConfig string
	bootstrapIP   string
	bootstrapID   string
)

func init() {
	flag.StringVar(&mode, "mode", online, "Operation mode (online/offline)")
	flag.StringVar(&flareBin, "flare-bin", "", "Path to flare binary")
	flag.StringVar(&flareConfig, "flare-config", "", "Path to flare config")
	flag.StringVar(&rosettaBin, "rosetta-bin", "", "Path to rosetta binary")
	flag.StringVar(&rosettaConfig, "rosetta-config", "", "Path to rosetta config")
	flag.StringVar(&bootstrapIP, "bootstrap-ip", "", "Node IP for bootstrapping")
	flag.StringVar(&bootstrapID, "bootstrap-id", "", "Node ID for bootstrapping")
	flag.Parse()

	if !(mode == online || mode == offline) {
		log.Fatal("invalid mode: " + mode)
	}

	if mode == online {
		if flareConfig == "" {
			log.Fatal("flare config path is not provided")
		}
		if flareBin == "" {
			log.Fatal("flare binary path is not provided")
		}
	}

	if rosettaConfig == "" {
		log.Fatal("rosetta config path is not provided")
	}
	if rosettaBin == "" {
		log.Fatal("rosetta binary path is not provided")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	handleSignals([]context.CancelFunc{cancel})

	g, _ := errgroup.WithContext(context.Background())

	if mode == online {
		g.Go(func() error {
			defer cancel()
			return startCommand(ctx, flareBin, "--config-file", flareConfig, "--bootstrap-ips", bootstrapIP, "--bootstrap-ids", bootstrapID)
		})
	}

	g.Go(func() error {
		defer cancel()
		return startCommand(ctx, rosettaBin, "-config", rosettaConfig)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func startCommand(ctx context.Context, path string, opts ...string) (err error) {
	log.Println("starting command:", path, opts)
	defer log.Println("command", path, "finished, error:", err)

	cmd := exec.Command(path, opts...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func() {
		<-ctx.Done()
		if cmd.Process != nil {
			if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
				panic(err)
			}
		}
	}()

	err = cmd.Run()
	return err
}

func handleSignals(listeners []context.CancelFunc) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Println("received signal:", sig)
		for _, listener := range listeners {
			listener()
		}
	}()
}
