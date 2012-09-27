#!/usr/bin/python

import pyexiv2

import errno
import os
import os.path
import sys

def getfilenameanddate(dn, fn):
    fn = os.path.join(dn, fn)
    md = pyexiv2.ImageMetadata(fn)
    md.read()
    return fn, md['Exif.Image.DateTime'].value

def main():
    origdn = sys.argv[1]
    treedn = sys.argv[2]
    for i in os.listdir(origdn):
        infn, dt = getfilenameanddate(origdn, i)
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
        os.link(infn, outfn)

if __name__ == '__main__':
    main()

