# Configuration Template.
meta:
    id: TENANT_APP_REDIR

type    : tenant
overlay : vlan

segments:
    - spec  : ref://store/specs/id=SEGMENT_NATIVE_APP_REDIR
      count : 1
    - spec  : ref://store/specs/id=SEGMENT_APP_REDIR
      count : 1

# NFV Endpoints:
# - They will attach to a 'TRUNK' Enic.
# - All the segments will be enabled on these Enics              
nfveps: 0

security_profile: ref://store/objects/id=SEC_PROF_ACTIVE
security_policy: ref://store/specs/id=SECURITY_POLICY_APP_REDIR

sessions:
    unidest:
        ipv4:
            - ref://store/specs/id=SESSION_TCP_APP_REDIR
        ipv6:
            - ref://store/specs/id=SESSION_TCP_APP_REDIR
        mac: None

lif: ref://store/specs/id=LIF_ETH
