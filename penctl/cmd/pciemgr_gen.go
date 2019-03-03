// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.
/*
 * Package cmd is a auto generated package.
 * Input file: pciemgr.proto
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pciemgrmetricsShowCmd = &cobra.Command{

	Use:   "pciemgr",
	Short: "Show PcieMgrMetrics from Naples",
	Long:  "\n---------------------------------\n Show PcieMgrMetrics From Naples \n---------------------------------\n",
	RunE:  pciemgrmetricsShowCmdHandler,
}

func pciemgrmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pciemgrmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pciemgr object(s) found")
	}
	return nil
}

var pcieportmetricsShowCmd = &cobra.Command{

	Use:   "pcieport",
	Short: "Show PciePortMetrics from Naples",
	Long:  "\n---------------------------------\n Show PciePortMetrics From Naples \n---------------------------------\n",
	RunE:  pcieportmetricsShowCmdHandler,
}

func pcieportmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pcieportmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pcieport object(s) found")
	}
	return nil
}

func init() {

	metricsShowCmd.AddCommand(pciemgrmetricsShowCmd)

	metricsShowCmd.AddCommand(pcieportmetricsShowCmd)

}
