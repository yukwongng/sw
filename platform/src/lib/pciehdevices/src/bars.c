/*
 * Copyright (c) 2018, Pensando Systems Inc.
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <sys/types.h>
#include <sys/param.h>

#include "pci_ids.h"
#include "misc.h"
#include "intrutils.h"
#include "pciemgrutils.h"
#include "pciehdevices.h"
#include "pciehdevices_impl.h"

void
add_common_resource_bar(pciehbars_t *pbars,
                        const pciehdevice_resources_t *pres)
{
    pciehbarreg_t preg;
    pciehbar_t pbar;
    prt_t prt;

    /*****************
     * resource bar - mem64
     * 8*4k regions = 32k bar size
     */
    memset(&pbar, 0, sizeof(pbar));
    pbar.type = PCIEHBARTYPE_MEM64;
    pbar.size = 8 * 0x1000;
    pbar.cfgidx = 0;

    /*****************
     * +0x0000 Device Cmd Regs 4-byte signature read-only */
    memset(&preg, 0, sizeof(preg));
    pmt_bar_enc(&preg.pmt,
                pres->port,
                PMT_TYPE_MEM,
                0x4,    /* pmtsize */
                0x4,    /* prtsize */
                PMT_BARF_WR);
    prt_res_enc(&prt, pres->devcmdpa, 0, PRT_RESF_PMVDIS);
    pciehbarreg_add_prt(&preg, &prt);
    pciehbar_add_reg(&pbar, &preg);

    /*****************
     * +0x0000 Device Cmd Regs
     * +0x1000 Device Cmd Doorbell
     */
    memset(&preg, 0, sizeof(preg));
    pmt_bar_enc(&preg.pmt,
                pres->port,
                PMT_TYPE_MEM,
                0x2000, /* pmtsize */
                0x1000, /* prtsize */
                PMT_BARF_RW);
    pmt_bar_setr_prt(&preg.pmt, 12, 1);

    /* +0x0000 Device Cmd Regs */
    prt_res_enc(&prt, pres->devcmdpa, 0x1000, PRT_RESF_PMVDIS);
    pciehbarreg_add_prt(&preg, &prt);

    /* +0x1000 Device Cmd Doorbell */
    prt_res_enc(&prt, pres->devcmddbpa, 0x4, (PRT_RESF_NOTIFY |
                                              PRT_RESF_PMVDIS));
    pciehbarreg_add_prt(&preg, &prt);
    pciehbar_add_reg(&pbar, &preg);

    /*****************
     * +0x2000 Interrupt Status reg */
    memset(&preg, 0, sizeof(preg));
    preg.baroff = 0x2000;
    pmt_bar_enc(&preg.pmt,
                pres->port,
                PMT_TYPE_MEM,
                0x8,    /* pmtsize */
                0x8,    /* prtsize */
                PMT_BARF_RD);
    prt_res_enc(&prt,
                intr_pba_addr(pres->lif),
                intr_pba_size(pres->intrc),
                PRT_RESF_NONE);
    pciehbarreg_add_prt(&preg, &prt);
    pciehbar_add_reg(&pbar, &preg);

    /*****************
     * +0x3000 Interrupt Control regs */
    {
        const u_int64_t pmtsize= roundup_power2(intr_drvcfg_size(pres->intrc));
        const u_int32_t stride = roundup_power2(intr_drvcfg_size(1));
        const int vfbitb = ffs(intr_drvcfg_size(1)) - 1;
        const int vfbitc = (ffs(pmtsize) - 1) - vfbitb;

        memset(&preg, 0, sizeof(preg));
        preg.baroff = 0x3000;
        pmt_bar_enc(&preg.pmt,
                    pres->port,
                    PMT_TYPE_MEM,
                    pmtsize,
                    stride, /* prtsize */
                    PMT_BARF_RW);
        pmt_bar_set_vfparams(&preg.pmt, vfbitb, vfbitc, 0, pres->intrc);
        prt_res_enc(&prt,
                    intr_drvcfg_addr(pres->intrb),
                    intr_drvcfg_size(pres->intrc),
                    PRT_RESF_NONE);
        prt_res_set_vfstride(&prt, stride);
        pciehbarreg_add_prt(&preg, &prt);
        pciehbar_add_reg(&pbar, &preg);
    }

    /*****************
     * +0x4000 <reserved>
     * +0x5000 <reserved>
     */

    /*****************
     * +0x6000 MSI-X Interrupt Table */
    {
        const u_int64_t pmtsize=roundup_power2(intr_msixcfg_size(pres->intrc));
        const u_int32_t stride =roundup_power2(intr_msixcfg_size(1));
        const int vfbitb = ffs(intr_msixcfg_size(1)) - 1;
        const int vfbitc = (ffs(pmtsize) - 1) - vfbitb;

        memset(&preg, 0, sizeof(preg));
        preg.baroff = 0x6000;
        pmt_bar_enc(&preg.pmt,
                    pres->port,
                    PMT_TYPE_MEM,
                    pmtsize,
                    stride, /* prtsize */
                    PMT_BARF_RW);
        pmt_bar_set_vfparams(&preg.pmt, vfbitb, vfbitc, 0, pres->intrc);
        prt_res_enc(&prt,
                    intr_msixcfg_addr(pres->intrb),
                    intr_msixcfg_size(pres->intrc),
                    PRT_RESF_NONE);
        prt_res_set_vfstride(&prt, stride);
        pciehbarreg_add_prt(&preg, &prt);
        pciehbar_add_reg(&pbar, &preg);
    }

    /*****************
     * +0x7000 MSI-X Interrupt PBA */
    memset(&preg, 0, sizeof(preg));
    preg.baroff = 0x7000;
    pmt_bar_enc(&preg.pmt,
                pres->port,
                PMT_TYPE_MEM,
                0x8,    /* pmtsize */
                0x8,    /* prtsize */
                PMT_BARF_RD);
    prt_res_enc(&prt,
                intr_pba_addr(pres->lif),
                intr_pba_size(pres->intrc),
                PRT_RESF_NONE);
    pciehbarreg_add_prt(&preg, &prt);
    pciehbar_add_reg(&pbar, &preg);

    /*
     * add this bar to our bars
     */
    pciehbars_add_bar(pbars, &pbar);

    /* set msix cap info */
    pciehbars_set_msix_tbl(pbars, 0, 0x6000);
    pciehbars_set_msix_pba(pbars, 0, 0x7000);
}

void
add_common_doorbell_bar(pciehbars_t *pbars,
                        const pciehdevice_resources_t *pres,
                        const u_int8_t upd[8])
{
    pciehbarreg_t preg;
    pciehbar_t pbar;
    prt_t prt;
    u_int32_t npids;

    /*****************
     * doorbell bar - mem64
     */
    memset(&pbar, 0, sizeof(pbar));

    /* at least 1 pid even if none specified */
    npids = MAX(pres->npids, 1);

    pbar.type = PCIEHBARTYPE_MEM64;
    pbar.size = 0x1000 * roundup_power2(npids);
    pbar.cfgidx = 2;

    memset(&preg, 0, sizeof(preg));
    pmt_bar_enc(&preg.pmt,
                pres->port,
                PMT_TYPE_MEM,
                pbar.size,  /* pmtsize */
                pbar.size,  /* prtsize */
                PMT_BARF_WR);
    pmt_bar_set_qtype(&preg.pmt, 3, 0x7);

    prt_db64_enc(&prt, pres->lif, upd);
    pciehbarreg_add_prt(&preg, &prt);

    pciehbar_add_reg(&pbar, &preg);
    pciehbars_add_bar(pbars, &pbar);
}

void
add_common_cmb_bar(pciehbars_t *pbars,
                   const pciehdevice_resources_t *pres)
{
    pciehbarreg_t preg;
    pciehbar_t pbar;
    prt_t prt;

    /*****************
     * optional cmb bar
     */
    if (pres->cmbsz) {
        memset(&pbar, 0, sizeof(pbar));
        pbar.type = PCIEHBARTYPE_MEM64;
        pbar.size = roundup_power2(pres->cmbsz);
        pbar.cfgidx = 4;

        memset(&preg, 0, sizeof(preg));
        pmt_bar_enc(&preg.pmt,
                    pres->port,
                    PMT_TYPE_MEM,
                    pbar.size,  /* pmtsize */
                    pbar.size,  /* prtsize */
                    PMT_BARF_RW);

        prt_res_enc(&prt, pres->cmbpa, pres->cmbsz, PRT_RESF_PMVDIS);
        pciehbarreg_add_prt(&preg, &prt);
        pciehbar_add_reg(&pbar, &preg);
        pciehbars_add_bar(pbars, &pbar);
    }
}
