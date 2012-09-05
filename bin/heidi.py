#!/usr/bin/python

# Heidin pyytama koko:
# 1778x1181

import errno
import os
import os.path

DN = '/home/hauva/Pictures/2012/New_York/08'
HDN = '/home/hauva/Pictures/heidille'

def dofiles(dname, fnames):
    todname = dname.replace(DN, HDN)
    try:
        os.makedirs(todname)
    except OSError, ex:
        if errno.EEXIST == ex.errno:
            pass
        else:
            raise ex
    for i in fnames:
        fn = '%s/%s' % (dname, i)
        tofn = '%s/%s' % (todname, i.lower())
        print fn, ' => ', tofn
        im = Image.open(fn)

def visit(arg, dname, fnames):
    if dname.endswith('/cand'):
        dofiles(dname, fnames)

def main():
    try:
        os.mkdir(HDN)
    except OSError, ex:
        if ex.errno == errno.EEXIST:
            pass
        else:
            raise ex
    os.path.walk(DN, visit, None)

if __name__ == '__main__':
    main()

