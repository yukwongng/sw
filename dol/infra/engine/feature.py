#! /usr/bin/python3

import infra.common.parser      as parser
import infra.common.objects     as objects

from infra.common.logging       import logger
from infra.common.glopts        import GlobalOptions
from infra.engine.modmgr        import ModuleStore

FeatureStore = objects.ObjectDatabase()
BaseTopoExcludeFeatureList = [
    'dolut',
    'fte',
    'fte2',
    'ftevxlan',
    'norm',
    'eth' ,
    'acl',
    'networking',
    'config',
    'vxlan',
    'ipsg',
    'firewall',
    'proxy',
    'ipsec',
    'hostpin',
    'multicast',
    'l4lb',
    'recirc',
    'admin',
    'classic',
    'app_redir',
    'parser',
    'hpvxlan',
    'telemetry',
]

class FeatureObject:
    counts = {}

    def __init__(self, spec):
        self.spec       = spec
        self.id         = spec.feature.id
        self.package    = spec.feature.package
        self.module     = spec.feature.module
        self.enable     = spec.feature.enable
        self.ignore     = spec.feature.ignore
        self.sub        = getattr(spec.feature, 'sub', None)
        self.testspec   = getattr(spec.feature, 'spec', None)
        self.args       = getattr(spec.feature, 'args', None)
        self.selectors  = getattr(spec.feature, 'selectors', None)
        self.runorder   = getattr(spec.feature, 'runorder', 65535)
        self.rtl        = getattr(spec.feature, 'rtl', True)
        self.perf       = getattr(spec.feature, 'perf', False)
        self.pendol     = getattr(spec.feature, 'pendol', False)
        self.tcscale    = getattr(spec.feature, 'tcscale', None)
        self.modscale    = getattr(spec.feature, 'modscale', None)
        return

    def AddModulesFromMlists(self, test_path):
        for m in self.spec.modules:
            mspec = m.module
            if hasattr(mspec, 'mlist'):
                mlist_name = getattr(mspec, 'mlist')
                spec = parser.ParseFile(test_path, mlist_name)
                for mod in spec.modules:
                    self.spec.modules.append(mod)

    def RenameModule(self, mspec):
        if mspec.name in self.counts:
            count = self.counts[mspec.name]
            count += 1
            self.counts[mspec.name] = count
        else:
            self.counts[mspec.name] = 0
        mspec.name = str(self.counts[mspec.name]) + "." + mspec.name

    def LoadOneModule(self, mspec, recurse=False):
        if self.sub is None:
            mspec.feature = self.id
        else:
            mspec.feature = "%s::%s" %(self.id, self.sub)
        mspec.enable    = getattr(mspec, 'enable', self.enable)
        mspec.ignore    = getattr(mspec, 'ignore', self.ignore)
        mspec.module    = getattr(mspec, 'module', self.module)
        mspec.package   = getattr(mspec, 'package', self.package)
        mspec.spec      = getattr(mspec, 'spec', self.testspec)
        mspec.args      = getattr(mspec, 'args', self.args)
        mspec.selectors = getattr(mspec, 'selectors', self.selectors)
        mspec.rtl       = getattr(mspec, 'rtl', self.rtl)
        mspec.perf      = getattr(mspec, 'perf', self.perf)
        mspec.pendol    = getattr(mspec, 'pendol', self.pendol)
        mspec.tcscale   = getattr(mspec, 'tcscale', self.tcscale)
        mspec.modscale  = getattr(mspec, 'modscale', self.modscale)
        mspec.runorder  = self.runorder
        if recurse is True:
            # Rename the module to avoid duplicate cases
            self.RenameModule(mspec)
        ModuleStore.Add(mspec)

    def LoadModulesFromMlist(self, mspec, path):
        mlist_name = getattr(mspec, 'mlist')
        if mlist_name is None:
            return
        spec = parser.ParseFile(path, mlist_name)
        for mod in spec.modules:
            mspec = mod.module
            self.LoadOneModule(mspec, True)

    def LoadModules(self, path):
        for m in self.spec.modules:
            mspec = m.module
            if hasattr(mspec, 'mlist'):
                self.LoadModulesFromMlist(mspec, path)
            else:
                self.LoadOneModule(mspec)
        return

class FeatureObjectHelper(parser.ParserBase):
    def __init__(self):
        self.features = []
        return

    def __is_match(self, feature):
        if GlobalOptions.feature is None:
            if feature.id in BaseTopoExcludeFeatureList:
                return False
            return True
        return feature.id in GlobalOptions.feature

    def __is_subfeature_match(self, feature):
        if GlobalOptions.subfeature is None:
            return True
        return feature.sub in GlobalOptions.subfeature

    def __is_enabled(self, feature):
        if self.__is_subfeature_match(feature) is False:
            logger.info("  - Subfeature Mismatch....Skipping")
            return False

        return self.__is_match(feature)

    def Parse(self, test_path):
        ftlist = super().Parse(test_path, '*.mlist')
        for fspec in ftlist:
            logger.info("Loading Feature Test Module List: %s" % (fspec.feature.id))
            feature = FeatureObject(fspec)
            if self.__is_enabled(feature) is False:
                logger.info("  - Disabled....Skipping")
                continue
            feature.LoadModules(test_path)
            self.features.append(feature)
        return

def Init(test_path):
    if GlobalOptions.feature is not None:
        GlobalOptions.feature = GlobalOptions.feature.split(',')
    if GlobalOptions.subfeature is not None:
        GlobalOptions.subfeature = GlobalOptions.subfeature.split(',')
    FtlHelper = FeatureObjectHelper()
    FtlHelper.Parse(test_path)
    return
