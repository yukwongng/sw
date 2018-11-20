// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.
/*
 * Package cmd is a auto generated package.
 * Input file: accel_metrics.proto
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var accelhwringmetricsShowCmd = &cobra.Command{

	Use:   "accelhwring",
	Short: "Show AccelHwRingMetrics from Naples",
	Long:  "\n---------------------------------\n Show AccelHwRingMetrics From Naples \n---------------------------------\n",
	RunE:  accelhwringmetricsShowCmdHandler,
}

func accelhwringmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	_, err := restGet(revProxyPort, "telemetry/v1/metrics/accelhwringmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

var accelseqqueuemetricsShowCmd = &cobra.Command{

	Use:   "accelseqqueue",
	Short: "Show AccelSeqQueueMetrics from Naples",
	Long:  "\n---------------------------------\n Show AccelSeqQueueMetrics From Naples \n---------------------------------\n",
	RunE:  accelseqqueuemetricsShowCmdHandler,
}

func accelseqqueuemetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	_, err := restGet(revProxyPort, "telemetry/v1/metrics/accelseqqueuemetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func init() {

	metricsShowCmd.AddCommand(accelhwringmetricsShowCmd)

	metricsShowCmd.AddCommand(accelseqqueuemetricsShowCmd)

}
