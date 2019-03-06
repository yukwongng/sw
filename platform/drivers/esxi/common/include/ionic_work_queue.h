/*
 * Copyright (c) 2019 Pensando Systems. All rights reserved.
 *
 * This software is provided to you under the terms of the GNU
 * General Public License (GPL) Version 2.0, available from the file
 * [ionic_work_queue.h] in the main directory of this source tree.
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
 * ionic_work_queue.h --
 *
 * Definitions of functions and data structures that used for work queue
 */

#ifndef _IONIC_WORK_QUEUE_H_
#define _IONIC_WORK_QUEUE_H_

#include <vmkapi.h>

#define IONIC_WORK_QUEUE_MAX_REQUESTS  1024

struct ionic_work_queue {
        vmk_Helper              helper_fun;
        vmk_HeapID              heap_id;
        vmk_Name                name;
        vmk_Bool                is_flush;
        vmk_Lock                flush_lock;
        vmk_atomic64            waiting_last_req;
};


struct ionic_work {
        vmk_HelperRequestFunc   fun;
        vmk_AddrCookie          fun_data;
        struct ionic_work_queue *queue;
};

inline void
ionic_work_init(struct ionic_work *work,
                vmk_HelperRequestFunc fun,
                void *data);

struct ionic_work_queue *
ionic_work_queue_create(vmk_HeapID heap_id,
                        vmk_ModuleID module_id,
                        const char *name);

inline void
ionic_work_queue_destroy(struct ionic_work_queue *work_queue);

inline VMK_ReturnStatus
ionic_work_queue_submit(struct ionic_work_queue *work_queue,
                        struct ionic_work *work,
                        vmk_uint32 delay_ms);

inline vmk_Bool
ionic_work_queue_cancel(struct ionic_work *work);

void
ionic_work_queue_flush(struct ionic_work_queue *work_queue);

#endif /* End of _IONIC_WORK_QUEUE_H_ */
