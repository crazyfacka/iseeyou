"""This class will handle motion changes and decide when to trigger stuff"""

import json
import time

import utils

L = utils.Utils().log

class Trigger(object):
    """Trigger class"""
    def __init__(self):
        self._moving = False
        self._last_change = time.time()
        self._first_pass_done = False

    def eval_state(self, state):
        """Evaluates current state to see if it computes as a change in motion"""
        if self._first_pass_done is False:
            self._first_pass_done = True
            self._moving = state
            return
        if state != self._moving:
            self._post_change()
            self._moving = state
            self._last_change = time.time()
            L('Motion changed to %s', 'moving' if bool(state) else 'not moving')

    def _post_change(self):
        """Creates the object with the information to send outside"""
        data = {
            'motion': self._moving,
            'duration': (time.time() - self._last_change),
            'start': self._last_change
        }

        L('%s', json.dumps(data))
