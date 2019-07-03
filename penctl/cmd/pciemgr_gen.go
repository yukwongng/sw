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

//cmd - pciemgr
//rootCmd = pcie
//longHelpStr = Pcie Mgr Metrics information:\nValue Description:\nnot_intr : notify total intrs.\nnot_spurious : notify spurious intrs.\nnot_cnt : notify total txns.\nnot_max : notify max txns per intr.\nnot_cfgrd : notify config reads.\nnot_cfgwr : notify config writes.\nnot_memrd : notify memory reads.\nnot_memwr : notify memory writes.\nnot_iord : notify io reads.\nnot_iowr : notify io writes.\nnot_unknown : notify unknown type.\nnot_rsrv0 : notify rsrv0.\nnot_rsrv1 : notify rsrv1.\nnot_msg : notify pcie message.\nnot_unsupported : notify unsupported.\nnot_pmv : notify pgm model violation.\nnot_dbpmv : notify doorbell pmv.\nnot_atomic : notify atomic trans.\nnot_pmtmiss : notify PMT miss.\nnot_pmrmiss : notify PMR miss.\nnot_prtmiss : notify PRT miss.\nnot_bdf2vfidmiss : notify bdf2vfid table miss.\nnot_prtoor : notify PRT out-of-range.\nnot_vfidoor : notify vfid out-of-range.\nnot_bdfoor : notify bdf out-of-range.\nnot_pmrind : notify PMR force indirect.\nnot_prtind : notify PRT force indirect.\nnot_pmrecc : notify PMR ECC error.\nnot_prtecc : notify PRT ECC error.\nind_intr : indirect total intrs.\nind_spurious : indirect spurious intrs.\nind_cfgrd : indirect config reads.\nind_cfgwr : indirect config writes.\nind_memrd : indirect memory reads.\nind_memwr : indirect memory writes.\nind_iord : indirect io reads.\nind_iowr : indirect io writes.\nind_unknown : indirect unknown type.\n\n
//shortHelpStr = Pcie Mgr Metrics information
var pciemgrpcieShowCmd = &cobra.Command{
	Use:   "pciemgr",
	Short: "Pcie Mgr Metrics information",
	Long:  "\n---------------------------------\n Pcie Mgr Metrics information:\nValue Description:\nnot_intr : notify total intrs.\nnot_spurious : notify spurious intrs.\nnot_cnt : notify total txns.\nnot_max : notify max txns per intr.\nnot_cfgrd : notify config reads.\nnot_cfgwr : notify config writes.\nnot_memrd : notify memory reads.\nnot_memwr : notify memory writes.\nnot_iord : notify io reads.\nnot_iowr : notify io writes.\nnot_unknown : notify unknown type.\nnot_rsrv0 : notify rsrv0.\nnot_rsrv1 : notify rsrv1.\nnot_msg : notify pcie message.\nnot_unsupported : notify unsupported.\nnot_pmv : notify pgm model violation.\nnot_dbpmv : notify doorbell pmv.\nnot_atomic : notify atomic trans.\nnot_pmtmiss : notify PMT miss.\nnot_pmrmiss : notify PMR miss.\nnot_prtmiss : notify PRT miss.\nnot_bdf2vfidmiss : notify bdf2vfid table miss.\nnot_prtoor : notify PRT out-of-range.\nnot_vfidoor : notify vfid out-of-range.\nnot_bdfoor : notify bdf out-of-range.\nnot_pmrind : notify PMR force indirect.\nnot_prtind : notify PRT force indirect.\nnot_pmrecc : notify PMR ECC error.\nnot_prtecc : notify PRT ECC error.\nind_intr : indirect total intrs.\nind_spurious : indirect spurious intrs.\nind_cfgrd : indirect config reads.\nind_cfgwr : indirect config writes.\nind_memrd : indirect memory reads.\nind_memwr : indirect memory writes.\nind_iord : indirect io reads.\nind_iowr : indirect io writes.\nind_unknown : indirect unknown type.\n\n\n---------------------------------\n",
	RunE:  pciemgrpcieShowCmdHandler,
}

func pciemgrpcieShowCmdHandler(cmd *cobra.Command, args []string) error {
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pciemgrmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pciemgrmetrics object(s) found")
	}
	return nil
}

//cmd - port
//rootCmd = pcie
//longHelpStr = Pcie Port Metrics information:\nValue Description:\nintr_total : total port intrs.\nintr_ltssmst_early : link train before linkup.\nintr_ltssmst : link train after  linkup.\nintr_linkup2dn : link down.\nintr_linkdn2up : link up.\nintr_rstup2dn : mac up.\nintr_rstdn2up : mac down.\nintr_secbus : secondary bus set.\nlinkup : link is up.\nhostup : host is up (secbus).\nphypolllast : phy poll count (last).\nphypollmax : phy poll count (max).\nphypollperstn : phy poll lost perstn.\nphypollfail : phy poll failed.\ngatepolllast : gate poll count (last).\ngatepollmax : gate poll count (max).\nmarkerpolllast : marker poll count (last).\nmarkerpollmax : marker poll count (max).\naxipendpolllast : axipend poll count (last).\naxipendpollmax : axipend poll count (max).\nfaults : link faults.\n\n
//shortHelpStr = Pcie Port Metrics information
var portpcieShowCmd = &cobra.Command{
	Use:   "port",
	Short: "Pcie Port Metrics information",
	Long:  "\n---------------------------------\n Pcie Port Metrics information:\nValue Description:\nintr_total : total port intrs.\nintr_ltssmst_early : link train before linkup.\nintr_ltssmst : link train after  linkup.\nintr_linkup2dn : link down.\nintr_linkdn2up : link up.\nintr_rstup2dn : mac up.\nintr_rstdn2up : mac down.\nintr_secbus : secondary bus set.\nlinkup : link is up.\nhostup : host is up (secbus).\nphypolllast : phy poll count (last).\nphypollmax : phy poll count (max).\nphypollperstn : phy poll lost perstn.\nphypollfail : phy poll failed.\ngatepolllast : gate poll count (last).\ngatepollmax : gate poll count (max).\nmarkerpolllast : marker poll count (last).\nmarkerpollmax : marker poll count (max).\naxipendpolllast : axipend poll count (last).\naxipendpollmax : axipend poll count (max).\nfaults : link faults.\n\n\n---------------------------------\n",
	RunE:  portpcieShowCmdHandler,
}

func portpcieShowCmdHandler(cmd *cobra.Command, args []string) error {
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pcieportmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pcieportmetrics object(s) found")
	}
	return nil
}

//cmd - pcie
//rootCmd =
//longHelpStr = Metrics for pciemgr
//shortHelpStr =
var pcieShowCmd = &cobra.Command{
	Use:   "pcie",
	Short: "",
	Long:  "\n---------------------------------\n Metrics for pciemgr\n---------------------------------\n",
}

func init() {

	pcieShowCmd.AddCommand(pciemgrpcieShowCmd)

	pcieShowCmd.AddCommand(portpcieShowCmd)

	//cmd - pcie
	//rootCmd =
	//longHelpStr = Metrics for pciemgr
	//shortHelpStr =

	metricsShowCmd.AddCommand(pcieShowCmd)

}
