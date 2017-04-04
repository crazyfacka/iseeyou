"""This class implements the HTTPS gateway"""

import threading

import requests

import confs
import utils

L = utils.Utils().log

class HTTP(threading.Thread):
    """HTTP class"""
    def __init__(self, queue):
        threading.Thread.__init__(self)
        self._queue = queue
        self.daemon = True

        self.headers = {'user-agent': 'iseeyou/0.1'}
        self.endpoint = confs.Configs().internet
        self.endpoint_update = self.endpoint + '/motion'
        self.endpoint_alive = self.endpoint + '/ping'

    def run(self):
        while True:
            self._dispatch_message(self._queue.get())

    def _dispatch_message(self, msg):
        msg_type = msg['type']
        data = msg['data']
        if msg_type == "update":
            self._send_update(data)
        elif msg_type == "alive":
            self._send_alive(data)

    def _send_update(self, data):
        """Sends a motion update"""
        rsp = requests.put(self.endpoint_update, data=data, headers=self.headers)
        L('PUT: %s', rsp.text)

    def _send_alive(self, data):
        """Sends a keepalive message"""
        rsp = requests.post(self.endpoint_alive, data=data, headers=self.headers)
        L('ALIVE: %s', rsp.text)
