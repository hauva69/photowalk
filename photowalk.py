#!/usr/bin/env python2

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
import pyexiv2

import datetime
import logging
import os
import os.path
import shutil

def getimagefiles(directoryname):
    return os.listdir(directoryname)

def handleimages(targetdirname, filename):
    logging.debug(filename)
    md = pyexiv2.ImageMetadata(filename)
    md.read()
    dt = md['Exif.Image.DateTime'].value
    logging.debug('datetime {0}'.format(dt.isoformat()))
    targetdirname = '{0}/{1}'.format(targetdirname, dt.strftime('%Y/%m%/%d'))
    targetfilename = '{0}/{1}'.format(targetdirname, os.path.basename(filename).lower())
    logging.debug('targetfilename {0}'.format(targetfilename))
    if not os.path.exists(targetdirname):
        os.makedirs(targetdirname)
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
