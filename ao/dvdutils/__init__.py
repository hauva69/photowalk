#!/usr/bin/python

#http://ubuntuforums.org/showthread.php?t=363666
# growisofs -dvd-compat -Z /dev/sr0=/path/to/image.iso

import ConfigParser
import os
import os.path
import sys

CONF = '%s/.aburnrc' % (os.environ['HOME'])

def perror(msg):
    sys.stderr.write(msg)

def main():
    config = ConfigParser.SafeConfigParser()
    if not os.path.exists(CONF):
        msg = '%s: missing configuration file.\n' % CONF
        perror(msg)
        sys.exit(1)
    config.read(CONF)

if __name__ == '__main__':
    main()

