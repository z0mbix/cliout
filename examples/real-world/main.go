// Example: Real-world CLI application
//
// Simulates a realistic deployment tool that uses cliout for user-facing
// output. Demonstrates how levels, themes, prefixes, and format strings
// work together in a practical scenario.
//
// Run:
//
//	go run ./examples/real-world/
//	go run ./examples/real-world/ -verbose
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/z0mbix/cliout"
)

func main() {
	verbose := flag.Bool("verbose", false, "enable debug output")
	flag.Parse()

	cliout.SetTheme(cliout.ThemeCatppuccinoMocha)
	cliout.SetPrefix("»")

	if *verbose {
		cliout.SetLevel(cliout.LevelDebug)
	}

	// --- Initialisation ---
	cliout.Info("starting deployment")
	cliout.Debugf("reading config from %s", "deploy.yaml")
	cliout.Debugf("target environment: %s", "production")
	cliout.Debugf("kubernetes context: %s", "prod-cluster-eu-west-1")

	// --- Validation ---
	cliout.Info("validating configuration")
	sleep()
	cliout.Debug("checking required fields")
	cliout.Debug("validating service definitions")
	cliout.Warnf("service %q has no resource limits defined", "api-gateway")
	cliout.Debug("validation complete")

	// --- Build ---
	cliout.Info("building container images")

	services := []string{"api-gateway", "auth-service", "worker"}
	for i, svc := range services {
		cliout.Infof("building %s (%d/%d)", svc, i+1, len(services))
		sleep()
		cliout.Debugf("docker build -t registry.example.com/%s:v1.2.3 .", svc)
		sleep()
	}
	cliout.Successf("built %d images", len(services))

	// --- Push ---
	cliout.Info("pushing images to registry")
	for _, svc := range services {
		cliout.Debugf("pushing registry.example.com/%s:v1.2.3", svc)
		sleep()
	}
	cliout.Success("all images pushed")

	// --- Deploy ---
	cliout.Info("deploying to cluster")
	for _, svc := range services {
		cliout.Infof("rolling out %s", svc)
		sleep()
		cliout.Debugf("kubectl apply -f k8s/%s/deployment.yaml", svc)
		sleep()
	}

	// --- Health checks ---
	cliout.Info("running health checks")
	for i, svc := range services {
		sleep()
		if i == 1 {
			cliout.Warnf("%s: health check slow (1200ms)", svc)
		} else {
			cliout.Debugf("%s: healthy (45ms)", svc)
		}
	}

	// --- Summary ---
	fmt.Println()
	cliout.Successf("deployment complete: %d services to %s", len(services), "production")
	cliout.Infof("dashboard: https://dashboard.example.com/deployments/v1.2.3")
}

func sleep() {
	time.Sleep(80 * time.Millisecond)
}
