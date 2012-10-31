#!/usr/bin/python

#http://ubuntuforums.org/showthread.php?t=363666
# growisofs -dvd-compat -Z /dev/sr0=/path/to/image.iso

import ConfigParser
import os

CONF = '%s/.aburnrc' % (os.environ['HOME'])

def main():
    print CONF

if __name__ == '__main__':
    main()

