"""This class will handle the keep alive process"""

import json
import time
import threading

import utils

L = utils.Utils().log

class Alive(object):
    """Alive class"""
    def __init__(self, name):
        self._name = name
        self._thread = None

    def begin_keep_alive(self):
        """This method begins the keep alive thread"""
        self._thread = threading.Thread(target=self._post_keep_alive, args=())
        self._thread.daemon = True
        self._thread.start()

    def _post_keep_alive(self):
        """This method builds and posts a keep alive message"""
        while True:
            data = {
                'ping': True,
                'name': self._name,
                'start': time.time()
            }

            L('%s', json.dumps(data))
            time.sleep(60)
