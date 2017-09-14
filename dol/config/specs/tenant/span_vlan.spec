# Configuration Template.
meta:
    id: TENANT_SPAN_VLAN

type    : span
overlay : vlan

segments:
    - spec  : ref://store/specs/id=SEGMENT_SPAN
      count : 1

security_profile: ref://store/objects/id=SEC_PROF_DEFAULT

span_sessions:
    - spec  : ref://store/specs/id=SPAN_TYPE_LOCAL
    - spec  : ref://store/specs/id=SPAN_TYPE_RSPAN
    - spec  : ref://store/specs/id=SPAN_TYPE_ERSPAN
    - spec  : ref://store/specs/id=SPAN_TYPE_LOCAL
    - spec  : ref://store/specs/id=SPAN_TYPE_RSPAN
    - spec  : ref://store/specs/id=SPAN_TYPE_ERSPAN

sessions: None

lif: ref://store/specs/id=LIF_ETH_RDMA
