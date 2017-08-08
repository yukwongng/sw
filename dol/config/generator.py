#! /usr/bin/python3
import pdb

import infra.common.defs        as defs
import infra.common.parser      as parser
import config.hal.api           as halapi

from config.objects.uplink              import UplinkHelper
from config.objects.uplinkpc            import UplinkPcHelper
from config.objects.tenant              import TenantHelper
from config.objects.session             import SessionHelper
from config.objects.security_profile    import SecurityProfileHelper
from config.objects.tcp_proxy_cb     import TcpCbHelper

from infra.common.logging import cfglogger as cfglogger

def process(topospec):
    # Security Profiles
    SecurityProfileHelper.main()
    # Uplinks
    UplinkHelper.main(topospec)
    # UplinkPC
    UplinkPcHelper.main(topospec)
    # Generate and Configure Tenants
    TenantHelper.main(topospec)

    # Enable all segments on Uplinks and UplinkPcs
    UplinkHelper.ConfigureAllSegments()
    UplinkPcHelper.ConfigureAllSegments()

    # Generate all sessions
    SessionHelper.main()
    
    TcpCbHelper.Configure(TcpCbHelper.objlist) 
    # Generate and Configure TCP and TLS Proxy CBs
    #tcp_proxy_cb_helper = tcp_proxy_cb.TenantObjectHelper()
    #tcp_proxy_cb_objlist = tcp_proxy_cb_helper.main(topospec)
    #tls_proxy_cb_helper = tls_proxy_cb.TenantObjectHelper()
    #tls_proxy_cb_objlist = tls_proxy_cb_helper.main(topospec)

    return

def main(topofile):
    halapi.init()
    cfglogger.info("Generating Config Objects for Topology = %s" % topofile)
    topospec = parser.ParseFile('config/topology/', topofile)
    process(topospec)
    return

