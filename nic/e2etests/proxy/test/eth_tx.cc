//  Hello World client
#include <zmq.h>
#include <string.h>
#include <stdio.h>
#include <unistd.h>
#include <assert.h>
#include <cmath>
#include <cstdlib>
#include <iostream>
#include <iomanip>

#include "nic/sdk/model_sim/include/buf_hdr.h"
#include "nic/e2etests/driver/lib_driver.hpp"
#include "nic/sdk/model_sim/include/lib_model_client.h"
#include "nic/e2etests/proxy/flows.hpp"


int main(int argc, char *argv[]) {
  if (argc < 2) {
    std::cout << "Usage:" << argv[0] << " <lif>" << std::endl;
    exit(1);
  }

  int ret;
  //uint32_t i;
  uint16_t lif_id = (uint16_t) (atoi(argv[1]) & 0xffff);
  uint32_t port = 0, cos = 0;

  //
  uint8_t *pkt = enic_to_uplink;
  std::vector<uint8_t> ipkt, opkt;

  ipkt.resize(sizeof(enic_to_uplink));
  opkt.resize(sizeof(enic_to_uplink));
  memcpy(ipkt.data(), pkt, sizeof(enic_to_uplink));

  lib_model_connect();

  // --------------------------------------------------------------------------------------------------------//

  // Create Queues
  alloc_queue(lif_id, TX, 0, 1024);
  alloc_queue(lif_id, RX, 0, 1024);

  // --------------------------------------------------------------------------------------------------------//

  // Post tx buffer
  uint8_t *buf = alloc_buffer(ipkt.size());
  assert(buf != NULL);
  memcpy(buf, ipkt.data(), ipkt.size());

  // Transmit Packet
  std::cout << "Writing packet to model! size: " << ipkt.size() << " on port: " << port << std::endl;
  post_buffer(lif_id, TX, 0, buf, ipkt.size());

  // Wait for packet to come out of port
  get_next_pkt(opkt, port, cos);
  if (!opkt.size()) {
    std::cout << "NO packet back from model! size: " << opkt.size() << std::endl;
    exit(1);
  } else {
    std::cout << "Got packet back from model! size: " << opkt.size() << " on port: " << port << " cos " << cos
              << std::endl;
    if (ipkt.size() == opkt.size()) {
      if (memcmp(ipkt.data(), opkt.data(), sizeof(enic_to_uplink))) {
        printf("Received packet does not match Sent packet!\n");
        ret = -1;
      } else {
        printf("Received packet matches Sent packet!\n");
        ret = 0;
      }
    } else {
      printf("Received packet is smaller than Sent packet.\n");
      ret = -1;
    }
  }

  lib_model_conn_close();

  return ret;
}
