#!/usr/bin/python

from __future__ import division

import Image
import pyexiv2

import errno
import getopt
import os
import os.path
import sys

def scale(image, size):
    '''Returns a scaled image where the length of the longest edge is size.'''
    w, h = image.size
    if w == h:
        w = size
        h = size
    elif w > h:
        h = int(size * h / w)
        w = size
    else:
        w = int(size * w/h)
        h = size
    return image.resize((w, h), Image.ANTIALIAS)

if __name__ == '__main__':
    frmd = sys.argv[1]
    tod = sys.argv[2]
    try:
        os.mkdir(tod)
    except OSError, ex:
        if errno.EEXIST == ex.errno:
            pass
        else:
            raise ex
    files = os.listdir(frmd)
    for i in files:
        fn = '%s/%s' % (frmd, i)
        try:
            im = Image.open(fn)
        except IOError, ex:
            # FIXME
	    continue
        out = scale(im, 800)
        outfn = '%s/%s' % (tod, i)
        out.save(outfn)

