"""This application monitors a PIR for motion, sending information
through the Internets to inform about its readings
"""

import time

import confs
import utils

import lib

CONF = confs.Configs()
if CONF.is_pi is True:
    import RPi.GPIO as GPIO #pylint: disable=I0011,F0401
    SLEEP = 1
else:
    import random
    SLEEP = 5

L = utils.Utils().log

def setup_pins():
    """Initializes PIR pins"""
    if CONF.is_pi is True:
        GPIO.setmode(GPIO.BCM)
        GPIO.setup(4, GPIO.IN, pull_up_down=GPIO.PUD_DOWN)

def read_pin():
    """Reads pin input, or generates random input if not on PI"""
    if CONF.is_pi is True:
        return GPIO.input(CONF.pir_pin)
    else:
        return random.randint(0, 1)

def capture_motion(trigger):
    """Infinite loop that monitors the PIR activity"""
    while True:
        motion_detected = read_pin()
        trigger.eval_state(motion_detected)
        if motion_detected:
            L('%s: Moving', motion_detected)
        else:
            L('%s: Not moving', motion_detected)

        time.sleep(SLEEP)

setup_pins()

lib.Alive('iseeyou').begin_keep_alive()
capture_motion(lib.Trigger())
