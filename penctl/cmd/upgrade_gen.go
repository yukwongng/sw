// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.
/*
 * Package cmd is a auto generated package.
 * Input file: github.com/pensando/sw/nic/upgrade_manager/upgrade/upgrade.proto
 */

package cmd

import (
	"github.com/spf13/cobra"
)

var upgrademetricsVar string

var upgrademetricsShowCmd = &cobra.Command{
	Use:   "upgrademetrics",
	Short: "Show UpgradeMetrics from Naples",
	Long:  "\n---------------------------------\n Show UpgradeMetrics From Naples \n---------------------------------\n",
	Run:   upgrademetricsShowCmdHandler,
}

func upgrademetricsShowCmdHandler(cmd *cobra.Command, args []string) {
	tabularFormat = false
	jsonFormat = true
	if cmd.Flags().Changed("upgrademetrics") {
		restGet(revProxyPort, "telemetry/v1/metrics/upgrademetrics/default/"+upgrademetricsVar+"/")
	} else {
		restGet(revProxyPort, "telemetry/v1/metrics/upgrademetrics/")
	}
}

func init() {

	metricsShowCmd.AddCommand(upgrademetricsShowCmd)
	upgrademetricsShowCmd.Flags().StringVarP(&upgrademetricsVar, "name", "n", "", "Name/Key for metrics object")

}
