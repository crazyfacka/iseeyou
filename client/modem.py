"""This script is to monitor the health and state of the 3G modem
installed on the RaspberryPI
"""

import urllib2
import subprocess

CMDS = {
    'test': 'echo test',
    'lsusb': 'lsusb',
    'usb_switch': '/usr/sbin/usb_modeswitch -v 12d1 -p 1446 -c /usr/share/usb_modeswitch/12d1:1446'
}

def get_cmd_as_args(cmd):
    """Checks for and returns a command as an argument list"""
    if cmd in CMDS:
        return CMDS[cmd].split()
    else:
        return ['echo', 'chicken']

def get_output_of_cmd(cmd):
    """Runs a specific command and returns its output in an array format"""
    try:
        return subprocess.check_output(get_cmd_as_args(cmd), stderr=subprocess.STDOUT).split('\n')
    except subprocess.CalledProcessError:
        print 'Error executing', CMDS[cmd]
        raise

def check_if_modem_enabled(out):
    """Given an output, checks if the modem is in the correct mode"""
    for line in out:
        if '12d1:1446' in line:
            return 0
        if '12d1:1001' in line:
            return 1
    return -1

def switch_modem():
    """Switches the modem to its supposed state"""
    try:
        get_output_of_cmd('usb_switch')
        return True
    except subprocess.CalledProcessError:
        return False

def check_internet_connection():
    """Checks if connected to internet"""
    try:
        urllib2.urlopen('http://163.172.170.201', timeout=1)
        return True
    except urllib2.URLError:
        return False

MODEM_ENABLED = check_if_modem_enabled(get_output_of_cmd('lsusb'))
if MODEM_ENABLED == 0:
    print 'Enabling modem'
    if not switch_modem():
        print 'Failed to switch modem'
        # Exit
elif MODEM_ENABLED == -1:
    print 'Modem not detected'
    # Exit

print 'Modem enabled'
print 'Checking internet connection...'

if not check_internet_connection():
    print 'Connecting modem'

print 'Connected'
