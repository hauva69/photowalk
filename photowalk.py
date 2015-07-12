#!/usr/bin/env python3

"""Photowalk.

Usage:
    photowalk.py import <sourceDir> <targetDir>
    photowalk.py (-h | --help)
    photowalk.py --version

Options:
  -h --help     Show this screen.
  --version     Show version.
"""

from docopt import docopt
from exif import EXIF

import datetime
import logging
import os

def getimagefiles(directoryname):
    return os.listdir(directoryname)

def handleimages(filename):
    logging.debug(filename)
    fd = open(filename, 'rb')
    tags = EXIF.process_file(fd)
    logging.debug(len(tags))
    for k, v in tags.items():
        print(k)
    datetimestring = str(tags['Image DateTime'])
    #datetimestring = b'2012:11:07 20:19:00'.decode('utf-8')
    logging.debug('datetimestring {0}'.format(datetimestring))
    dt = datetime.datetime.strptime(datetimestring, '%Y:%m:%d %H:%M:%S')
    logging.debug('datetime {0}'.format(dt.isoformat()))
    fd.close()

if __name__ == '__main__':
    args = docopt(__doc__, version='Photowalk 0.01')
    sourcedir = args['<sourceDir>']
    targetdir = args['<targetDir>']
    logging.basicConfig(level=logging.DEBUG)
    logging.debug('source_dir={0}'.format(sourcedir))
    logging.debug('target_dir={0}'.format(targetdir))
    for i in getimagefiles(sourcedir):
        filename = os.path.join(sourcedir, i)
        handleimages(filename)
