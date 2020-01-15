package cmd

import (
	"errors"
	"strings"
	"time"

	"github.com/pensando/sw/iota/test/venice/iotakit"
	"github.com/spf13/cobra"
)

var (
	nodeNames string
)

func init() {
	rootCmd.AddCommand(naplesCmd)
	naplesCmd.AddCommand(naplesAddCmd)
	naplesCmd.AddCommand(naplesDelCmd)
	naplesCmd.AddCommand(naplesUpgradeCmd)
	naplesAddCmd.Flags().StringVarP(&nodeNames, "names", "", "", "Node names")
	naplesDelCmd.Flags().StringVarP(&nodeNames, "names", "", "", "Node names")
}

var naplesCmd = &cobra.Command{
	Use:   "naples",
	Short: "actions on naples",
}

var naplesAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a node with naples personality",
	Run:   naplesAddAction,
}

var naplesDelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a node with naples personality",
	Run:   naplesDeleteAction,
}

var naplesUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade naples",
	Run:   naplesUpgradeAction,
}

func naplesAddAction(cmd *cobra.Command, args []string) {

	if nodeNames == "" {
		errorExit("No node name specified", errors.New("No node name specified"))
	}

	err := setupModel.AddNaplesNodes(strings.Split(nodeNames, ","))

	if err != nil {
		errorExit("Error adding  naples node", err)
	}
}

func naplesDeleteAction(cmd *cobra.Command, args []string) {

	if nodeNames == "" {
		errorExit("No node name specified", errors.New("No node name specified"))
	}

	err := setupModel.DeleteNaplesNodes(strings.Split(nodeNames, ","))

	if err != nil {
		errorExit("Error delete naples node", err)
	}
}

func naplesUpgradeAction(cmd *cobra.Command, args []string) {

	setupModel.ForEachNaples(func(nc *iotakit.NaplesCollection) error {
		setupModel.Action().RunNaplesCommand(nc, "touch /update/upgrade_to_same_firmware_allowed")
		return nil
	})

	//For now upgrade all naples is the only option.
	err := doNaplesUpgrade(0)
	if err != nil {
		errorExit("Naples upgrade failed", err)
	}
}

func doNaplesRemoveAdd(percent int) error {

	naples, err := setupModel.Naples().SelectByPercentage(percent)

	if err != nil {
		return err
	}

	return setupModel.Action().RemoveAddNaples(naples)
}

func doNaplesMgmtLinkFlap(percent int) error {

	naples, err := setupModel.Naples().SelectByPercentage(percent)

	if err != nil {
		return err
	}

	return setupModel.Action().FlapMgmtLinkNaples(naples)
}

func doNaplesUpgrade(percent int) error {

	numNaples := 0
	setupModel.ForEachNaples(func(nc *iotakit.NaplesCollection) error {
		_, err := setupModel.Action().RunNaplesCommand(nc, "touch /data/upgrade_to_same_firmware_allowed")
		numNaples++
		return err
	})

	defer setupModel.ForEachNaples(func(nc *iotakit.NaplesCollection) error {
		_, err := setupModel.Action().RunNaplesCommand(nc, "rm /data/upgrade_to_same_firmware_allowed")
		return err
	})

	setupModel.ForEachFakeNaples(func(nc *iotakit.NaplesCollection) error {
		numNaples++
		return nil
	})

	rollout, err := setupModel.GetRolloutObject(true)
	if err != nil {
		return err
	}

	err = setupModel.Action().PerformRollout(rollout, true)
	if err != nil {
		return err
	}

	TimedOutEvent := time.After(time.Duration(60*numNaples) * time.Second)
	for true {
		select {
		case <-TimedOutEvent:
			return errors.New("Error waiting for upgrade, timed out")
		default:
			err = setupModel.Action().VerifyRolloutStatus(rollout.Name)
			if err == nil {
				//Sleep for while to make sure all naples are connected
				time.Sleep(60 * time.Second)
				return nil
			}
		}
	}

	return nil
}

func generateEvents(rate, count string) error {

	setupModel.Action().StopEventsGenOnNaples(setupModel.Naples())

	return setupModel.Action().StartEventsGenOnNaples(setupModel.Naples(),
		rate, count)
}

func stopEvents(rate, count string) error {

	setupModel.Action().StopEventsGenOnNaples(setupModel.Naples())
	return nil
}

func generateFWLogs(rate, count string) error {

	setupModel.Action().StopFWLogGenOnNaples(setupModel.Naples())

	return setupModel.Action().StartFWLogGenOnNaples(setupModel.Naples(),
		rate, count)
}

func stopFWLogs(rate, count string) error {

	setupModel.Action().StopFWLogGenOnNaples(setupModel.Naples())
	return nil

}
