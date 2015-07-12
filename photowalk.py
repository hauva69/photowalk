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
import os.path
import shutil

def getimagefiles(directoryname):
    return os.listdir(directoryname)

def handleimages(targetdirname, filename):
    logging.debug(filename)
    fd = open(filename, 'rb')
    tags = EXIF.process_file(fd)
    fd.close()
    datetimestring = str(tags['Image DateTime'])
    logging.debug('datetimestring {0}'.format(datetimestring))
    dt = datetime.datetime.strptime(datetimestring, '%Y:%m:%d %H:%M:%S')
    logging.debug('datetime {0}'.format(dt.isoformat()))
    targetdirname = '{0}/{1}'.format(targetdirname, dt.strftime('%Y/%m%/%d'))
    targetfilename = '{0}/{1}'.format(targetdirname, os.path.basename(filename))
    logging.debug('targetfilename {0}'.format(targetfilename))
    os.makedirs(targetdirname, exist_ok=True)
    shutil.copy2(filename, targetfilename)

if __name__ == '__main__':
    args = docopt(__doc__, version='Photowalk 0.01')
    sourcedir = args['<sourceDir>']
    targetdir = args['<targetDir>']
    logging.basicConfig(level=logging.DEBUG)
    logging.debug('source_dir={0}'.format(sourcedir))
    logging.debug('target_dir={0}'.format(targetdir))
    for i in getimagefiles(sourcedir):
        filename = os.path.join(sourcedir, i)
        handleimages(targetdir, filename)
