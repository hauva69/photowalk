#!/usr/bin/python

# Heidin pyytama koko:
# 1778x1181

import errno
import os
import os.path

DN = '/home/hauva/Pictures/2012/New_York/08'
HDN = '/home/hauva/Pictures/heidille'

def visit(arg, dname, fnames):
    if dname.endswith('/cand'):
        print dname

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

