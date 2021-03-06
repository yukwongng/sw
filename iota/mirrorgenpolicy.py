#! /usr/bin/python3
import argparse
import json
import os
import collections
import random
import copy

mirrorpolicy_template = {
    "type" : "netagent",
    "rest-endpoint"    : "api/mirror/sessions/",
    "object-key" : "meta.tenant/meta.namespace/meta.name",
    "objects" : [
    {
    "kind"                  : "MirrorSession",
    "meta": {
        "name"              : "mirror-2",
        "tenant"            : "default",
        "namespace"         : "default",
        "creation-time"     : "1970-01-01T00:00:00Z",
        "mod-time"          : "1970-01-01T00:00:00Z"
    },
    "spec": {
        "packet-size": 128,
        "collectors": [
          {
             "type": "erspan_type_3",
             "export-config": {
                 "destination": "192.168.100.101",
            }
          }
        ],
        "match-rules": [
        ]
    },
    "status"                : {}
    }
    ]
}

collector_template = {
    "type": "netagent",
    "rest-endpoint": "api/collectors/",
    "object-key": "meta.tenant/meta.namespace/meta.name",
    "objects": [
        {
            "kind": "Collector",
            "meta": {
                "creation-time": "1970-01-01T00:00:00Z",
                "mod-time": "1970-01-01T00:00:00Z",
                "namespace": "default",
                "name": "lif-collector",
                "tenant": "default"
            },
            "spec": {
                "packet-size": 128,
                "destination": "192.168.100.101",
                "type": "erspan_type_2"
            },
            "status": {},
        }
    ],
}


def get_verif(dst_ip, src_ip, protocol, port, result):
    verif = {}
    verif['dst_ip'] = dst_ip
    verif['src_ip'] = src_ip
    verif['protocol'] = protocol
    verif['port'] = port
    verif['result'] = result
    return verif

def get_appconfig(protocol, port):
    app_config = {}
    app_config['protocol'] = protocol
    app_config['port'] = port
    return app_config

def get_app_proto(protocol, dport):
    proto_ports = []
    app = {}
    app['protocol'] = protocol
    if protocol == 'icmp':
        app['port'] = '0'
        #proto_str = str(protocol).upper() + '/' + '0' + '/' + '0'
    else:
        if dport == 'any':
            app['port'] = '0'
            #proto_str = str(protocol).upper() + '/' + '0'
        else:
            app['port'] = str(dport)
            #proto_str = str(protocol).upper() + '/' + str(dport)
#    app['port'] = str(dport)
#     app['proto-ports'] = []
#     app['proto-ports'].append(proto_str)
    proto_ports.append(app)
    return proto_ports

def get_destination(dst_ip, protocol, port):
    if dst_ip == 'any':
        dst_ip = '0.0.0.0/0'
    dst = {}
    dst['addresses'] = []
    dst['addresses'].append(dst_ip)
    dst['proto-ports'] = get_app_proto(protocol, port)
    #dst['app-protocol-selectors'] = []
    #dst['app-protocol-selectors'].append(get_appconfig(protocol, port))
    return dst

def get_source(src_ip):
    if src_ip == 'any':
        src_ip = '0.0.0.0/0'
    src = {}
    src['addresses'] = []
    src['addresses'].append(src_ip)
    return src

def get_rule(dst_ip, src_ip, protocol, port, action):
    rule = {}
    rule['destination'] = get_destination(dst_ip, protocol, port)
    rule['source'] = get_source(src_ip)
    #rule['app-protocol-selectors'] = get_app_proto(protocol, port)
    #rule['action'] = action
    return rule

parser = argparse.ArgumentParser(description='Naples Mirror Policy Generator')
parser.add_argument('--topology', dest='topology_dir', required = dir,
                    default=None, help='Path to the JSON file having IOTA endpoint information.')

GlobalOptions = parser.parse_args()
GlobalOptions.endpoint_file = GlobalOptions.topology_dir + "/endpoints.json"
GlobalOptions.protocols = ["udp", "tcp", "icmp"]
GlobalOptions.directories = ["udp", "tcp", "icmp", "mixed", "scale", "collector"]
#GlobalOptions.ports = ["10", "22", "24", "30", "50-100", "101-200", "201-250, 251-290", "10000-20000", "65535"]
GlobalOptions.ports = ["120"]
GlobalOptions.actions = ["mirror"]

def StripIpMask(ip_address):
    if '/' in ip_address:
        return ip_address[:-3].encode("ascii")
    return ip_address.encode("ascii")

def GetIpRange(ip_address):
    stripped_ip = ip_address[:ip_address.rfind('.')]
    return "{}.1-{}.100".format(stripped_ip, stripped_ip)

def GetIpCidr(ip_address):
    stripped_ip = ip_address[:ip_address.rfind('.')]
    return "{}.0/24".format(stripped_ip)

def Main():
    #import pdb; pdb.set_trace()
    with open(GlobalOptions.endpoint_file, 'r') as fp:
        obj = json.load(fp)
    EP = []
    EP_nodes = []

    for i in range(0, len(obj["objects"])):
        print("EP[%d] : %s" % (i, obj["objects"][i]["spec"]["ipv4-addresses"][0]))
        EP.append(StripIpMask(obj["objects"][i]["spec"]["ipv4-addresses"][0]))
        EP_nodes.append(obj["objects"][i]["spec"]["node-uuid"])

    GlobalOptions.topology_dir = GlobalOptions.topology_dir + '/gen/telemetry/mirror'
    for dir in GlobalOptions.directories:
        if not os.path.exists(GlobalOptions.topology_dir + "/{}".format(dir)):
            os.makedirs(GlobalOptions.topology_dir + "/{}".format(dir))

    # Specific match policy
    for protocol in GlobalOptions.protocols:
        for action in GlobalOptions.actions:
            mirrorpolicy = mirrorpolicy_template
            match_rules = mirrorpolicy_template['objects'][0]['spec']['match-rules']
            del match_rules[:]
            verif =[]
            loop = 1
            for i in range(0, len(EP)):
                for j in range(i+1, len(EP)):
                    for k in GlobalOptions.ports:
                        if EP_nodes[i] == EP_nodes[j]:
                            continue
                        if loop == 0:
                            continue
                        rule = get_rule(EP[i], EP[j], protocol, k, action)
                        match_rules.append(rule)
                        verif.append(get_verif(EP[i], EP[j], protocol, k, action))
                        rule = get_rule(EP[j], EP[i], protocol, k, action)
                        match_rules.append(rule)
                        verif.append(get_verif(EP[j], EP[i], protocol, k, action))
                        loop = 0
            if len(verif) > 0:
                json.dump(verif, open(GlobalOptions.topology_dir +"/{}/{}_{}_specific_verif.json".format(protocol, protocol, action), "w"), indent=4)
                json.dump(mirrorpolicy, open(GlobalOptions.topology_dir + "/{}/{}_{}_specific_policy.json".format(protocol, protocol, action), "w"), indent=4)

    EP_SUBNET = []
    EP_SUBNET.append(GetIpCidr(EP[0]))
    # Subnet policy
    for protocol in GlobalOptions.protocols:
        for action in GlobalOptions.actions:
            mirrorpolicy = mirrorpolicy_template
            match_rules = mirrorpolicy_template['objects'][0]['spec']['match-rules']
            del match_rules[:]
            verif =[]
            loop = 1
            for i in range(0, len(EP_SUBNET)):
                for j in range(0, len(EP)):
                    for k in GlobalOptions.ports:
                        if EP_nodes[i] == EP_nodes[j]:
                            continue
                        if loop == 0:
                            continue
                        rule = get_rule(EP_SUBNET[i], EP[j], protocol, k, action)
                        match_rules.append(rule)
                        verif.append(get_verif(EP_SUBNET[i], EP[j], protocol, k, action))
                        rule = get_rule(EP[j], EP_SUBNET[i], protocol, k, action)
                        match_rules.append(rule)
                        verif.append(get_verif(EP[j], EP_SUBNET[i], protocol, k, action))
                        loop = 0
            if len(verif) > 0:
                json.dump(mirrorpolicy, open(GlobalOptions.topology_dir +"/{}/{}_{}_subnet_policy.json".format(protocol, protocol, action), "w"), indent=4)
                json.dump(verif, open(GlobalOptions.topology_dir + "/{}/{}_{}_subnet_verif.json".format(protocol, protocol, action), "w"), indent=4)

    EP_ANY = []
    EP_ANY.append("any")
    GlobalOptions.ports.append("any")
    # Generic (any) Policy
    for protocol in GlobalOptions.protocols:
        for action in GlobalOptions.actions:
            mirrorpolicy = mirrorpolicy_template
            match_rules = mirrorpolicy_template['objects'][0]['spec']['match-rules']
            del match_rules[:]
            verif =[]
            loop = 1
            for i in range(0, len(EP_ANY)):
                for j in range(i, len(EP)):
                    for k in GlobalOptions.ports:
                        if EP_nodes[i] == EP_nodes[j]:
                            continue
                        if loop == 0:
                            continue
                        rule = get_rule(EP_ANY[i], EP[j], protocol, k, action)
                        match_rules.append(rule)
                        verif.append(get_verif(EP_ANY[i], EP[j], protocol, k, action))
                        rule = get_rule(EP[j], EP_ANY[i], protocol, k, action)
                        match_rules.append(rule)
                        verif.append(get_verif(EP[j], EP_ANY[i], protocol, k, action))
                        loop = 0
            if len(verif) > 0:
                json.dump(mirrorpolicy, open(GlobalOptions.topology_dir + "/{}/{}_{}_any_policy.json".format(protocol, protocol, action), "w"), indent=4)
                json.dump(verif, open(GlobalOptions.topology_dir + "/{}/{}_{}_any_verif.json".format(protocol, protocol, action), "w"), indent=4)

    # Mixed Config
    del match_rules[:]
    verif =[]
    for protocol in GlobalOptions.protocols:
        for action in GlobalOptions.actions:
            mirrorpolicy = mirrorpolicy_template
            match_rules = mirrorpolicy_template['objects'][0]['spec']['match-rules']
            loop = 1
            for i in range(0, len(EP)):
                for j in range(i+1, len(EP)):
                    for k in GlobalOptions.ports:
                        if EP_nodes[i] == EP_nodes[j]:
                            continue
                        if loop == 0:
                            continue
                        rule = get_rule(EP[i], EP[j], protocol, k, action)
                        match_rules.append(rule)
                        verif.append(get_verif(EP[i], EP[j], protocol, k, action))
                        rule = get_rule(EP[j], EP[i], protocol, k, action)
                        match_rules.append(rule)
                        verif.append(get_verif(EP[j], EP[i], protocol, k, action))
                        loop = 0
    if len(verif) > 0:
        json.dump(verif, open(GlobalOptions.topology_dir +"/{}/mirror_mixed_verif.json".format('mixed'), "w"), indent=4)
        json.dump(mirrorpolicy, open(GlobalOptions.topology_dir + "/{}/mirror_mixed_policy.json".format('mixed'), "w"), indent=4)

    # Scale Config
    verif = []
    for action in GlobalOptions.actions:
        mirrorpolicies = copy.deepcopy(mirrorpolicy_template)
        objects = mirrorpolicies['objects']
        mirorpolicy = copy.deepcopy(objects[0])
        collector = copy.deepcopy(mirorpolicy['spec']['collectors'][0])
        collectors = []
        del objects[:]
        proto_index = 0
        ep_index = 0
        for i in range(0, min(len(EP), 8)):   # Collector loop
            collector['export-config']['destination'] = EP[i-7]
            collectors.append(copy.deepcopy(collector))
            objects.append(copy.deepcopy(mirorpolicy))
            objects[-1]['spec']['collectors'] = copy.deepcopy(collectors)
            objects[-1]['meta']['name'] = "mirror-scale-%s"%i
            match_rules = objects[-1]['spec']['match-rules']
            del match_rules[:]
            protocol =  GlobalOptions.protocols[proto_index]
            for j in range(ep_index+1, len(EP)):
                for k in GlobalOptions.ports:
                    rule = get_rule(EP[ep_index], EP[j], protocol, k, action)
                    match_rules.append(rule)
                    #verif.append(get_verif(EP[i], EP[j], protocol, k, collector, action))
                    rule = get_rule(EP[j], EP[ep_index], protocol, k, action)
                    match_rules.append(rule)
                    #verif.append(get_verif(EP[j], EP[i], protocol, k, collector, action))

            proto_index = (proto_index + 1) % len(GlobalOptions.protocols)
            if proto_index == 0:
                ep_index += 1
    json.dump(mirrorpolicies, open(GlobalOptions.topology_dir + "/{}/mirror_scale_policy.json".format('scale'), "w"), indent=4)

    # collector Config
    collectors = collector_template
    objects = collectors['objects']
    collector = objects[0]
    del objects[:]
    for i in range(0, min(len(EP), 8)):   # Collector loop
        objects.append(copy.deepcopy(collector))
        objects[-1]['meta']['name'] = "lif-collector-%s"%(i)
        objects[-1]['spec']['destination'] = EP[i]
    json.dump(collectors, open(GlobalOptions.topology_dir + "/{}/collector_scale.json".format('collector'), "w"), indent=4)
if __name__ == '__main__':
    Main()
