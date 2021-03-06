#! /usr/bin/python3
import sys
import os
import pdb

asic = os.environ.get('ASIC', 'capri')
paths = [
    '/dol/',
    '/dol/third_party/',
    '/nic/sdk',
    '/nic/build/aarch64/apulu/' + asic + '/gen/proto',
]

ws_top = os.path.dirname(sys.argv[0]) + '/../'
ws_top = os.path.abspath(ws_top)
os.environ['WS_TOP'] = ws_top


for path in paths:
    fullpath = ws_top + path
    print("Adding Path: %s" % fullpath)
    sys.path.insert(0, fullpath)


APULU_PROTO_PATH = os.environ['WS_TOP'] + '/nic/build/aarch64/apulu/' + asic + '/gen/proto/'

sys.path.insert(0, APULU_PROTO_PATH)
