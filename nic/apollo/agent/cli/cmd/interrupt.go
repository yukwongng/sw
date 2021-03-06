package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pensando/sw/nic/apollo/agent/cli/utils"
	"github.com/pensando/sw/nic/apollo/agent/gen/pds"
)

var interruptShowCmd = &cobra.Command{
	Use:   "interrupts",
	Short: "show interrupt details",
	Long:  "show interrupt details information",
	Run:   interruptShowCmdHandler,
}

var interruptClearCmd = &cobra.Command{
	Use:   "interrupts",
	Short: "clear interrupt details",
	Long:  "clear interrupt details information",
	Run:   interruptClearCmdHandler,
}

func init() {
	showCmd.AddCommand(interruptShowCmd)
	clearCmd.AddCommand(interruptClearCmd)
}

func interruptClearCmdHandler(cmd *cobra.Command, args []string) {
	c, err := utils.CreateNewGRPCClient()
	if err != nil {
		fmt.Printf("Could not connect to the sysmon, is sysmon running?\n")
		return
	}
	defer c.Close()

	// handle command
	cmdResp, err := HandleSvcReqCommandMsg(pds.Command_CMD_INTR_CLEAR, nil)
	if err != nil {
		fmt.Printf("Command failed with %v error\n", err)
		return
	}
	if cmdResp.ApiStatus != pds.ApiStatus_API_STATUS_OK {
		fmt.Printf("Command failed with %v error\n", cmdResp.ApiStatus)
		return
	}
}

func interruptShowCmdHandler(cmd *cobra.Command, args []string) {
	c, err := utils.CreateNewGRPCClient()
	if err != nil {
		fmt.Printf("Could not connect to the sysmon, is sysmon running?\n")
		return
	}
	defer c.Close()

	// handle command
	cmdResp, err := HandleSvcReqCommandMsg(pds.Command_CMD_INTR_DUMP, nil)
	if err != nil {
		fmt.Printf("Command failed with %v error\n", err)
		return
	}
	if cmdResp.ApiStatus != pds.ApiStatus_API_STATUS_OK {
		fmt.Printf("Command failed with %v error\n", cmdResp.ApiStatus)
		return
	}
}
