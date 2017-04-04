"""This class implements the HTTPS gateway"""

import requests

import confs
import utils

L = utils.Utils().log

class HTTP(object):
    """HTTP class"""
    def __init__(self):
        self.headers = {'user-agent': 'iseeyou/0.1'}
        self.endpoint = confs.Configs().internet
        self.endpoint_update = self.endpoint + '/motion'
        self.endpoint_alive = self.endpoint + '/ping'

    def send_update(self, data):
        """Sends a motion update"""
        rsp = requests.put(self.endpoint_update, data=data, headers=self.headers)
        L('PUT: %s', rsp.text)

    def send_alive(self, data):
        """Sends a keepalive message"""
        rsp = requests.post(self.endpoint_alive, data=data, headers=self.headers)
        L('ALIVE: %s', rsp.text)
