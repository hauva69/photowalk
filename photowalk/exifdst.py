#!/usr/bin/python

import pyexiv2

import errno
import os
import os.path
import shutil
import sys

def getdate(fn):
    '''directory name, file name -> Exif.Image.Datetime.value'''
    md = pyexiv2.ImageMetadata(fn)
    md.read()
    return md['Exif.Image.DateTime'].value

def exifdst(origdn, treedn):
    '''origdn (the picture directory), treedn (the directory where the 
    files will be created.'''
    for i in os.listdir(origdn):
        infn = os.path.join(origdn, i)
        if not os.path.isfile(infn):
            continue
        try:
            dt = getdate(infn)
        except IOError, ex:
            msg = '%s\n' % str(ex)
            sys.stderr.write(msg)
            continue
        sy = '%d' % dt.year
        sm = '%02d' % dt.month
        sd = '%02d' % dt.day
        dn = os.path.join(treedn, sy, sm, sd)
        try:
            os.makedirs(dn)
        except OSError, ex:
            if ex.errno == errno.EEXIST:
                pass
            else:
                raise ex
        outfn = os.path.join(dn, i.lower())
        try:
            os.link(infn, outfn)
        except OSError, ex:
            if errno.EEXIST == ex.errno:
                continue
            elif errno.EXDEV == ex.errno:
                shutil.copy(infn, outfn)
            else:
                sys.stderr.write('%s\n' % infn)
                raise ex

if __name__ == '__main__':
    origdn = sys.argv[1]
    treedn = sys.argv[2]
    exifdst(origdn, treedn)

