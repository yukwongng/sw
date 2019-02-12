
#ifndef __ENDPOINT_HPP__
#define __ENDPOINT_HPP__

#include "lib/indexer/indexer.hpp"

#include "hal.hpp"
#include "enic.hpp"

typedef std::tuple<HalL2Segment *, mac_t> ep_key_t;

class HalEndpoint : public HalObject
{
public:
  static HalEndpoint *Factory(HalL2Segment *l2seg, mac_t mac, Enic *enic);
  static hal_irisc_ret_t  Destroy(HalEndpoint *ep);
  static HalEndpoint *Lookup(HalL2Segment *l2seg, mac_t mac);

  hal_irisc_ret_t HalEndpointCreate();
  hal_irisc_ret_t HalEndpointDelete();
  HalL2Segment *GetL2Seg();
  mac_t GetMac();

private:
  HalEndpoint(HalL2Segment *l2seg, mac_t mac, Enic *enic);
  ~HalEndpoint() {};

  mac_t mac;
  HalL2Segment *l2seg;
  Enic *enic;

  static std::map<ep_key_t, HalEndpoint*> ep_db;
};

#endif /* __ENDPOINT_HPP__ */
