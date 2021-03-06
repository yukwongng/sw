# Flow generation configuration template.
meta:
    id: SESSION_TCP

proto: tcp

entries:
    - entry:
        label: firewall
        tracking: True
        timestamp: True
        responder: 
            port : const/80
            flow_info:
                action  : allow
                state   : established
            tracking_info:
                tcp_seq_num     : const/1000
                tcp_ack_num     : const/5000
                tcp_win_sz      : const/8192
                tcp_win_scale   : const/1
                tcp_mss         : const/1460
        initiator: 
            port : const/40000
            flow_info:
                action  : allow
                state   : established
            tracking_info:
                tcp_seq_num     : const/5000
                tcp_ack_num     : const/1000
                tcp_win_sz      : const/8192
                tcp_win_scale   : const/1
                tcp_mss         : const/1460

    - entry:
        label: span1
        tracking: False
        timestamp: False
        responder: 
            port : const/81
            span :
                ingress:
                    - ref://store/objects/id=SpanSession0001
                    - ref://store/objects/id=SpanSession0002
                    - ref://store/objects/id=SpanSession0003
                egress:
                    - ref://store/objects/id=SpanSession0004
                    - ref://store/objects/id=SpanSession0005
                    - ref://store/objects/id=SpanSession0006
        initiator: 
            port : const/41000
            span :
                ingress:
                    - ref://store/objects/id=SpanSession0001
                    - ref://store/objects/id=SpanSession0002
                    - ref://store/objects/id=SpanSession0003
                egress:
                    - ref://store/objects/id=SpanSession0004
                    - ref://store/objects/id=SpanSession0005
                    - ref://store/objects/id=SpanSession0006

    - entry:
        label: tcp-proxy
        tracking: False
        timestamp: False
        responder:
            port : const/80
        initiator:
            port : const/47273

    - entry:
        label: ipsec-proxy
        tracking: False
        timestamp: False
        responder:
            port : const/44444
        initiator:
            port : const/44445

    - entry:
        label: raw-redir
        tracking: False
        timestamp: False
        responder:
            port : const/23767
        initiator:
            port : const/23768

    - entry:
        label: raw-redir-flow-miss
        tracking: False
        timestamp: False
        fte: True
        responder:
            port : const/19694
        initiator:
            port : const/19695

    - entry:
        label: raw-redir-known-appid
        tracking: False
        timestamp: False
        fte: True
        responder: 
            port : const/3306 # mysql port
        initiator:
            port : const/46624

    - entry:
        label: raw-redir-known-appid
        tracking: False
        timestamp: False
        fte: True
        responder: 
            port : const/2501 # kismet port
        initiator:
            port : const/46624

    - entry:
        label: proxy-redir
        tracking: False
        timestamp: False
        responder: 
            port : const/23763
        initiator: 
            port : const/23764

    - entry:
        label: networking
        tracking: False
        timestamp: False
        responder: 
            port : const/1
            flow_info:
                eg_qos:    
                    cos_rw  : const/1
                    cos     : const/4
                    dscp_rw : const/1
                    dscp    : const/1
        initiator: 
            port : const/4
            flow_info:
                eg_qos:    
                    cos_rw  : const/1
                    cos     : const/3
                    dscp_rw : const/1
                    dscp    : const/2

    - entry:
        label: recirc 
        tracking: False
        timestamp: False
        responder: 
            port : const/0x704e
            flow_info:
                eg_qos:    
                    cos_rw  : const/1
                    cos     : const/2
                    dscp_rw : const/1
                    dscp    : const/3
        initiator: 
            port : const/2
            flow_info:
                eg_qos:    
                    cos_rw  : const/1
                    cos     : const/1
                    dscp_rw : const/1
                    dscp    : const/4


    - entry:
        label: gft_drop
        tracking: False
        timestamp: False
        responder: 
            port : const/12345
            flow_info:
                action  : drop
        initiator: 
            port : const/54321
            flow_info:
                action  : drop

    - entry:
        label: fte
        fte: True
        tracking: False
        timestamp: False
        responder: 
            port : const/22222
            flow_info:
                eg_qos:    
                    cos_rw  : const/1
                    cos     : const/4
                    dscp_rw : const/1
                    dscp    : const/1
        initiator: 
            port : const/33333
            flow_info:
                eg_qos:    
                    cos_rw  : const/1
                    cos     : const/3
                    dscp_rw : const/1
                    dscp    : const/2

