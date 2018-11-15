// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.
/*
 * Package cmd is a auto generated package.
 * Input file: github.com/pensando/sw/nic/proto/nicmgr/metrics.proto
 */

package cmd

import (
	"github.com/spf13/cobra"
)

var lifmetricsShowCmd = &cobra.Command{
	Use:   "lifmetrics",
	Short: "Show LifMetrics from Naples",
	Long:  "\n---------------------------------\n Show LifMetrics From Naples \n---------------------------------\n",
	Run:   lifmetricsShowCmdHandler,
}

func lifmetricsShowCmdHandler(cmd *cobra.Command, args []string) {
	tabularFormat = false
	jsonFormat = true
	restGet(revProxyPort, "telemetry/v1/metrics/lifmetrics/")
}

func init() {

	metricsShowCmd.AddCommand(lifmetricsShowCmd)

}
