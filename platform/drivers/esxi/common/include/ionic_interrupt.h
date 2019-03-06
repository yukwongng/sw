/*
 * Copyright (c) 2019 Pensando Systems. All rights reserved.
 *
 * This software is provided to you under the terms of the GNU
 * General Public License (GPL) Version 2.0, available from the file
 * [ionic_interrupt.h] in the main directory of this source tree.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
 * BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

/*
 * ionic_interrupt.h --
 *
 * Definitions of interrupt related APIs 
 */

#ifndef _IONIC_INTERRUPT_H_
#define _IONIC_INTERRUPT_H_

#include <vmkapi.h>

VMK_ReturnStatus
ionic_int_register(vmk_ModuleID module_id,
                   vmk_Device vmk_device,
                   vmk_IntrCookie intr_cookie,
                   const char *name,
                   vmk_IntrHandler handler,
                   void *handler_data);

void
ionic_int_unregister(vmk_ModuleID module_id,
                     vmk_IntrCookie intr_cookie,
                     void *handler_data);

VMK_ReturnStatus
ionic_int_alloc(vmk_ModuleID module_id,
                vmk_HeapID heap_id,
                vmk_PCIDevice pci_device,
                vmk_PCIInterruptType int_type,
                vmk_uint32 desired,
                vmk_uint32 required,
                vmk_IntrCookie *intr_cookie_array,
                vmk_uint32 *granted);

void
ionic_int_free(vmk_ModuleID module_id,
               vmk_HeapID heap_id,
               vmk_PCIDevice pci_device,
               vmk_IntrCookie *intr_cookie_array);

#endif /* End of _IONIC_INTERRUPT_H_ */
