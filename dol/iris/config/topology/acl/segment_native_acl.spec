# Segment Configuration Spec
meta:
    id  : SEGMENT_NATIVE_ACL

type        : tenant
fabencap    : vlan
native      : True
broadcast   : flood
multicast   : replicate
l4lb        : False
endpoints   :
    useg    : 0
    pvlan   : 2
    direct  : 0
    remote  : 4 # 1 EP per uplink
