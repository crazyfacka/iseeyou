"""This class posts information to the outside world"""

import threading

import utils

from . import http

L = utils.Utils().log

class Gateway(object):
    """Gateway class"""
    gateway = None

    @staticmethod
    def internet():
        """Returns an internet gateway"""
        if Gateway.gateway is None:
            Gateway.gateway = http.HTTP()
        return Gateway.gateway

    @staticmethod
    def send_update(data):
        """Sends a motion update"""
        if Gateway.gateway is None:
            L('Error: gateway not defined')
            return
        thread = threading.Thread(target=Gateway.gateway.send_update, args=(data))
        thread.daemon = True
        thread.start()

    @staticmethod
    def send_alive(data):
        """Sends a keepalive message"""
        if Gateway.gateway is None:
            L('Error: gateway not defined')
            return
        thread = threading.Thread(target=Gateway.gateway.send_alive, args=(data))
        thread.daemon = True
        thread.start()
