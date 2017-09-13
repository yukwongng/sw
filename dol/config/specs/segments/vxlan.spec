# Segment Configuration Spec
meta:
    id: SEGMENT_VXLAN

type        : tenant
fabencap    : vxlan
native      : False
broadcast   : flood
multicast   : replicate
l4lb        : True
endpoints   :
    useg    : 2
    pvlan   : 2
    direct  : 0
    remote  : 4 # 1 EP per uplink
