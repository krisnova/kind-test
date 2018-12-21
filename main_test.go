package kind_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/kris-nova/logger"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cluster/config"
)

const (
	ClusterName       = "tgik"
	KubernetesInDockerNodeImage = "kindest/node:v1.12.2"
)

// TestMain will set up a Kubernetes cluster to use
func TestMain(m *testing.M) {


	// Configure logger to run in test mode
	logger.TestMode = true

	// Set the logger to most verbose mode
	logger.Level = 4


	// Set up cluster using KIND
	ctx := cluster.NewContext(ClusterName)

	// Defer tear town cluster
	defer func() {
		ctx.Delete()
	}()

	// Set up a new Cluster
	cfg := &config.Config{
		Image: KubernetesInDockerNodeImage,
	}
	retain := false
	wait := time.Duration(0)
	err := ctx.Create(cfg, retain, wait)
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}

	 portStr := strconv.Itoa(ctx.ControlPlaneMeta.APIServerPort)
	 APIPort = portStr

	// Run Go tests
	os.Exit(m.Run())


}
