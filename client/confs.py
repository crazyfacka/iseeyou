"""This module loads and handles the application configurations"""

import os
import ConfigParser

CONFIG_FILE = "confs.ini"

class Configs(object):
    """Config class"""
    _confs_loaded = False

    debug = False
    pir_pin = 4
    internet = None
    gateway = None

    def __init__(self):
        if Configs._confs_loaded:
            return

        self._read_confs(self._load_confs())
        Configs._confs_loaded = True
        print 'Loaded configurations'
        print self

    def __str__(self):
        return """Debug: %s
PIR Pin: %s
Internet: %s
Gateway: %s
Is PI: %s
""" % (self.debug, self.pir_pin, self.internet, self.gateway, self.is_pi())

    @staticmethod
    def _load_confs():
        """Loads the configuration from file"""
        cfg = ConfigParser.ConfigParser()
        cfg.read(CONFIG_FILE)
        return cfg

    @staticmethod
    def _read_confs(cfg):
        """Read values from the configuration file"""
        Configs.debug = cfg.get('Main', 'Debug')
        Configs.pir_pin = cfg.get('Main', 'PIRPin')
        Configs.internet = cfg.get('Main', 'Internet')
        Configs.gateway = cfg.get('Main', 'Gateway')

    @staticmethod
    def is_pi():
        """Not 100% validation of Pi hardware, but close"""
        return bool('armv7' in os.uname())

    @staticmethod
    def debug_mode():
        """Informs whether the application is in debug mode"""
        return Configs.debug
