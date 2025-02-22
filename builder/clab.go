package builder

import (
	"fmt"
	"github.com/Lachstec/digsinet-ng/types"
	"gopkg.in/yaml.v3"
	"log"
	"os/exec"
)

type ClabBuilder struct {
}

func NewClabBuilder() *ClabBuilder {
	return &ClabBuilder{}
}

func (b *ClabBuilder) DeployTopology(topology types.Topology) error {
	log.Print("Deploying Topology with Containerlab Builder...")

	// Prepare the command
	proc := exec.Command("clab", "deploy", "--topo", "-")

	// Connect stdin for the process
	stdin, err := proc.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	// Start the process
	if err = proc.Start(); err != nil {
		return fmt.Errorf("failed to start process: %w", err)
	}

	// Serialize the topology spec
	topologySpec, err := yaml.Marshal(topology)
	if err != nil {
		return fmt.Errorf("failed to marshal topology: %w", err)
	}

	// Debug: print the topology spec to ensure correctness
	log.Printf("Topology spec to be sent: %s", string(topologySpec))

	// Write the topology spec to stdin
	_, err = stdin.Write(topologySpec)
	if err != nil {
		return fmt.Errorf("failed to write to stdin: %w", err)
	}

	// Close stdin to signal the end of input
	if err := stdin.Close(); err != nil {
		return fmt.Errorf("failed to close stdin: %w", err)
	}

	// Wait for the process to complete
	if err = proc.Wait(); err != nil {
		return fmt.Errorf("process finished with error: %w", err)
	}

	log.Print("Topology deployment completed successfully.")
	return nil
}

func (b *ClabBuilder) DestroyTopology(topology types.Topology) error {
	log.Print("Destroying Topology with Containerlab Builder...")

	// Prepare the command
	proc := exec.Command("clab", "destroy", "--topo", "-")

	// Connect stdin for the process
	stdin, err := proc.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	// Start the process
	if err = proc.Start(); err != nil {
		return fmt.Errorf("failed to start process: %w", err)
	}

	// Serialize the topology spec
	topologySpec, err := yaml.Marshal(topology)
	if err != nil {
		return fmt.Errorf("failed to marshal topology: %w", err)
	}

	// Debug: print the topology spec to ensure correctness
	log.Printf("Topology spec to be sent: %s", string(topologySpec))

	// Write the topology spec to stdin
	_, err = stdin.Write(topologySpec)
	if err != nil {
		return fmt.Errorf("failed to write to stdin: %w", err)
	}

	// Close stdin to signal the end of input
	if err := stdin.Close(); err != nil {
		return fmt.Errorf("failed to close stdin: %w", err)
	}

	// Wait for the process to complete
	if err = proc.Wait(); err != nil {
		return fmt.Errorf("process finished with error: %w", err)
	}

	log.Print("Topology deployment completed successfully.")
	return nil
}
