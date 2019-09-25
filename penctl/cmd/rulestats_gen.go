// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.
/*
 * Package cmd is a auto generated package.
 * Input file: rulestats.proto
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rulemetricsShowCmd = &cobra.Command{

	Use:   "rule",
	Short: "Show RuleMetrics from Distributed Service Card",
	Long:  "\n---------------------------------\n Show RuleMetrics From Distributed Service Card \n---------------------------------\n",
	RunE:  rulemetricsShowCmdHandler,
}

func rulemetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/rulemetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No rule object(s) found")
	}
	return nil
}

func init() {

	metricsShowCmd.AddCommand(rulemetricsShowCmd)

}
