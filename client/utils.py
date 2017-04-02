"""This module handles trivial stuff that should be share across the whole application"""

import time
import datetime

import confs

class Utils(object):
    """Utils class"""
    _cfg = None

    def __init__(self):
        Utils._cfg = confs.Configs()

    @staticmethod
    def log(msg, *arg):
        """Logs output if debug mode is on"""
        if Utils._cfg.debug_mode():
            now = datetime.datetime.fromtimestamp(time.time()).strftime('%Y-%m-%d %H:%M:%S')
            final_msg = msg % arg
            print '[%s] %s' % (now, final_msg)
