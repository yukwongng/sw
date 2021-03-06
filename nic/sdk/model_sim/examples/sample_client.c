//  Hello World client
#include <zmq.h>
#include <string.h>
#include <stdio.h>
#include <unistd.h>
#include <assert.h>
#include "buf_hdr.h"
#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <cstdlib>
#include <assert.h>
#include <iostream>
#include "lib_model_client.h"


int main (void)
{
    lib_model_connect();
    /* Sample pkt 1 */
    uint8_t pkt[] = {
#if 0
            0x00, 0x22, 0x22, 0x22, 0x22, 0x22, 0x00, 0x33, 0x33, 0x33, 0x33, 0x33, 0x81, 0x00, 0x00, 0x02,
            0x08, 0x00, 0x45, 0x00, 0x00, 0x3C, 0x3E, 0xEF, 0x40, 0x00, 0x40, 0x11, 0xFD, 0xCA, 0x40, 0x00,
            0x00, 0x01, 0x40, 0x00, 0x00, 0x02, 0xB8, 0xA9, 0x00, 0x80, 0xBB, 0xB5, 0xE2, 0xA0, 0x00, 0x00,
            0x00, 0x00, 0xA0, 0x02, 0xFF, 0xFF, 0xFE, 0x30, 0x00, 0x00, 0x02, 0x04, 0xFF, 0xD7, 0x04, 0x02,
            0x08, 0x0A, 0x16, 0x38, 0xDC, 0x4F, 0x00, 0x00, 0x00, 0x00, 0x01, 0x03, 0x03, 0x02
#endif
            // pkt2 = Ether(dst="00:22:22:22:22:22",src="00:33:33:33:33:33")/Dot1Q(vlan=2)/IP(src="64.0.0.1",dst="64.0.0.2")/TCP(dport=80,sport=47273,seq=0xbabababa,ack=0xefefefef)
            0x00, 0x22, 0x22, 0x22, 0x22, 0x22, 0x00, 0x33, 0x33, 0x33, 0x33, 0x33, 0x81, 0x00, 0x00, 0x02,
            0x08, 0x00, 0x45, 0x00, 0x00, 0x28, 0x00, 0x01, 0x00, 0x00, 0x40, 0x06, 0xfa, 0xcc, 0x40, 0x00,
            0x00, 0x01, 0x40, 0x00, 0x00, 0x02, 0xb8, 0xa9, 0x00, 0x50, 0xba, 0xba, 0xba, 0xba, 0xef, 0xef,
            0xef, 0xef, 0x50, 0x02, 0x20, 0x00, 0x56, 0xe6, 0x00, 0x00
    };
    std::vector<uint8_t> opkt;
    uint32_t port = 0, cos = 0;
    
    std::vector<uint8_t> ipkt;
    ipkt.resize(sizeof(pkt));
    memcpy(ipkt.data(), pkt, sizeof(pkt));
    std::cout << "Sending packet to model! size: " << ipkt.size() << " on port: " << port << std::endl;
    step_network_pkt(ipkt, port);
    
    get_next_pkt(opkt, port, cos);
    if (!opkt.size()) {
        std::cout << "NO packet back from model! size: " << opkt.size() << std::endl;
        exit(1);
    } else {
        std::cout << "Got packet back from model! size: " << opkt.size() << " on port: " << port << " cos: " << cos << std::endl;
        if (ipkt.size() == opkt.size()) {
            exit(0);
        } else {
            exit(1);
        }
    }

#if 0
    write_reg(0x3400000+0x8010, 0xbadabc);
    uint32_t data;
    read_reg(0x3400000+0x8010, data);
    printf("read_reg data: 0x%x\n", data);

    uint8_t mdata[4096];
    memset(mdata, 0xff, 4096);
        write_mem(0x80000000 , mdata, 176);
        write_mem(0x800000b0 , mdata, 328);
        write_mem(0x800001f8 , mdata, 184);
        write_mem(0x800002b0 , mdata, 160);
        write_mem(0x80000350 , mdata, 456);
        write_mem(0x80000518 , mdata, 112);
        write_mem(0x80000588 , mdata, 104);
        write_mem(0x800005f0 , mdata, 72 );
        write_mem(0x80000638 , mdata, 80 );
        write_mem(0x80000688 , mdata, 160);
        write_mem(0x80000728 , mdata, 144);
        write_mem(0x800007b8 , mdata, 240);
        write_mem(0x800008a8 , mdata, 256);
        write_mem(0x800009a8 , mdata, 832);
        write_mem(0x80000ce8 , mdata, 104);
        write_mem(0x80000d50 , mdata, 104);
        write_mem(0x80000db8 , mdata, 160);
        write_mem(0x80000e58 , mdata, 560);
#endif

    lib_model_conn_close();

    return 0;
}
