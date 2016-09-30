package cmd

import (
	"errors"

	"github.com/getcarina/carina/console"
	"github.com/spf13/cobra"
)

func newGrowCommand() *cobra.Command {
	var options struct {
		name     string
		template string
		nodes    int
		wait     bool
	}

	var cmd = &cobra.Command{
		Use:   "grow <cluster-name>",
		Short: "Add nodes to a cluster",
		Long:  "Add nodes to a cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if options.nodes < 1 {
				return errors.New("--nodes must be >= 1")
			}

			return bindName(args, &options.name)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cluster, err := cxt.Client.GrowCluster(cxt.Account, options.name, options.nodes, options.wait)
			if err != nil {
				return err
			}

			console.WriteCluster(cluster)

			return nil
		},
	}

	cmd.ValidArgs = []string{"cluster-name"}
	cmd.Flags().IntVar(&options.nodes, "nodes", 1, "number of nodes to increase the cluster by")
	cmd.Flags().BoolVar(&options.wait, "wait", false, "wait for the previous cluster operation to complete")

	return cmd
}

func init() {
	rootCmd.AddCommand(newGrowCommand())
}