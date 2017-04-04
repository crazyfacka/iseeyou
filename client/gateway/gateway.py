"""This class posts information to the outside world"""

import Queue

import utils

from . import http

L = utils.Utils().log

class Gateway(object):
    """Gateway class"""
    gateway = None
    queue = Queue.Queue()

    @staticmethod
    def internet():
        """Creates an internet gateway"""
        if Gateway.gateway is None:
            Gateway.gateway = http.HTTP(Gateway.queue)
            Gateway.gateway.start()

    @staticmethod
    def send_update(data):
        """Sends a motion update"""
        if Gateway.gateway is None:
            L('Error: gateway not defined')
            return
        Gateway.queue.put({"type": "update", "data": data})

    @staticmethod
    def send_alive(data):
        """Sends a keepalive message"""
        if Gateway.gateway is None:
            L('Error: gateway not defined')
            return
        Gateway.queue.put({"type": "alive", "data": data})
