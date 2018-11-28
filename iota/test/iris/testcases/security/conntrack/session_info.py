#! /usr/bin/python3
import yaml

def get_conntrackinfo(cmd):
    yaml_out = yaml.load_all(cmd.stdout)
    print(type(yaml_out))
    for data in yaml_out:
        if data is not None:
            spec = data['spec']
            init_flow = spec['initiatorflow']
            iseq_num = init_flow['flowdata']['conntrackinfo']['tcpseqnum']
            iack_num = init_flow['flowdata']['conntrackinfo']['tcpacknum']
            iwindosz = init_flow['flowdata']['conntrackinfo']['tcpwinsz']
            iwinscale = init_flow['flowdata']['conntrackinfo']['tcpwinscale']
            print('init flow: seq_num {} ack_num {} iwindosz {} iwinscale {}'.format(iseq_num, iack_num, iwindosz, iwinscale))
            
            resp_flow = spec['responderflow']
            rseq_num = resp_flow['flowdata']['conntrackinfo']['tcpseqnum']
            rack_num = resp_flow['flowdata']['conntrackinfo']['tcpacknum']
            rwindosz = resp_flow['flowdata']['conntrackinfo']['tcpwinsz']
            rwinscale = resp_flow['flowdata']['conntrackinfo']['tcpwinscale']
            print('resp flow: seq_num {} ack_num {} iwindosz {} iwinscale {}'.format(rseq_num, rack_num, rwindosz, rwinscale))
            return iseq_num, iack_num, iwindosz, iwinscale, rseq_num, rack_num, rwindosz, rwinscale
        else:
            return 0,0,0,0,0,0
    return 0,0,0,0,0,0

def get_yaml(cmd):
    yaml_out = yaml.load_all(cmd.stdout)
    for data in yaml_out:
        if data is not None:
            return data

def get_respflow(data):
    return data['spec']['responderflow']

def get_initflow(data):
    return data['spec']['initiatorflow']

def get_conntrack_info(flow):
    return flow['flowdata']['conntrackinfo']

def get_exceptions(conn_info):
    return conn_info['exceptions']
