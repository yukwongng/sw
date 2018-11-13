#! /usr/bin/python3
import datetime
import pdb
import sys
import inspect
import threading

import iota.harness.infra.types   as types

prefixes = {
    0: 'NONE',
    1: 'CRIT',
    2: 'ERRR',
    3: 'WARN',
    4: 'INFO',
    5: 'DEBG',
    6: 'VERB',
    7: 'MAX'
}

class LoggerSink:
    def __init__(self, stdout = None, logfile = None):
        self.sink = None
        self.lock = threading.Lock()
        if stdout:
            self.sink = sys.stdout
        elif logfile:
            self.sink = open(logfile, 'w')
        else:
            assert(0)
        return

    def __is_logger_print(self, text):
        for level in range(7):
            pfx = prefixes[level]
            if pfx in text: return True
        return False

    def write(self, text):
        #pdb.set_trace()
        #text = text.replace('\n', ' ')
        #text += '\n'
        self.log(text)
        return

    def log(self, text):
        #if not self.__is_logger_print(text):
        #    return
        self.lock.acquire()
        self.sink.write(text)
        self.lock.release()
        self.flush()
        return

    def flush(self):
        self.sink.flush()
        return

    def isatty(self):
        return True

StdoutLoggerSink = LoggerSink(stdout = True)
#sys.stdout = StdoutLoggerSink
#sys.stderr = StdoutLoggerSink

class _Logger:
    def __init__(self, level, stdout=True, logfile=None):
        self.sinks          = []
        self.indent_enable  = False
        self.level          = level
        self.logfile        = logfile
        self.tsname         = None
        self.tcname         = None
        self.tcid           = None

        if stdout:
            global StdoutLoggerSink
            self.sinks.append(StdoutLoggerSink)

        if logfile:
            self.sinks.append(LoggerSink(logfile = self.logfile))
        return

    def __flush(self, text):
        for s in self.sinks:
            s.log(text)
        return

    def __get_timestamp(self):
        a = datetime.datetime.now()
        return "[%02d:%02d:%02d.%06d] " % (a.hour, a.minute, a.second, a.microsecond)

    def __get_log_prefix(self, level=None):
        prefix = self.__get_timestamp()
        if self.tsname:
            prefix += "[TS:%s]" % self.tsname
        if self.tcname:
            prefix += "[TC:%s_%s]" % (self.tcname, str(self.tcid))
        if level:
            prefix += "[%s]" % prefixes[level]
        else:
            prefix += "[INFO]"
        prefix += " "
        return prefix

    def __args_to_str(self, *args, **kwargs):
        text = ""
        for a in args:
            text = text + str(a) + " "
        return

    def __format(self, *args, **kwargs):
        text = ""
        indent = 0
        if self.indent_enable:
            indent = len(inspect.stack())
            if indent >= defs.LOGGING_DEFAULT_REV_OFFSET:
                indent = indent - defs.LOGGING_DEFAULT_REV_OFFSET

        level = kwargs['level']
        if self.level < level:
            return None

        text = self.__get_log_prefix(level)
        if indent:
            text = text + "  " * indent
        for a in args:
            text = text + str(a) + " "

        text = text.replace('\n', ' ')
        text = text + "\n"
        return text

    def __log(self, *args, **kwargs):
        text = self.__format(*args, **kwargs)
        if text != None:
            self.__flush(text)
        return

    def info(self, *args, **kwargs):
        return self.__log(*args, **kwargs, level=types.loglevel.INFO)

    def debug(self, *args, **kwargs):
        return self.__log(*args, **kwargs, level=types.loglevel.DEBUG)

    def verbose(self, *args, **kwargs):
        return self.__log(*args, **kwargs, level=types.loglevel.VERBOSE)

    def warn(self, *args, **kwargs):
        return self.__log(*args, **kwargs, level=types.loglevel.WARNING)

    def error(self, *args, **kwargs):
        return self.__log(*args, **kwargs, level=types.loglevel.ERROR)

    def critical(self, *args, **kwargs):
        return self.__log(*args, **kwargs, level=types.loglevel.CRITICAL)

    def log(self, level, *args, **kwargs):
        return self.__log(*args, **kwargs, level=level)

    def SetLoggingLevel(self, level):
        self.level = level

    def SetTestsuite(self, tsname):
        self.tsname = tsname
        # Reset the tcname everytime tsname changes.
        self.tcname = None
        self.tcid = None
        return

    def SetTestcase(self, tcname):
        self.tcname = tcname
        return

    def SetTestcaseID(self, tcid):
        self.tcid = tcid
        return

    def GetLogPrefix(self):
        return self.__get_log_prefix()

    def ShowScapyObject(self, scapyobj, build = True):
        if build:
            scapyobj.show2(indent = 0, label_lvl = self.GetLogPrefix())
        else:
            scapyobj.show(indent = 0, label_lvl = self.GetLogPrefix())
        return

    def LogFunctionBegin(self):
        logger.debug("BEG: %s()" % inspect.stack()[1][3])
        return

    def LogFunctionEnd(self, status=0):
        logger.debug("END: %s()  Status:%d" % (inspect.stack()[1][3], status))
        return

    def header(self, string):
        hdr = "-" * 20
        self.info(hdr + string + hdr)
        return

Logger = _Logger(types.loglevel.INFO)
