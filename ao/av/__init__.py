#!/usr/bin/python

PICASA_MAX_SIZE = 2048

import Image
import pyexiv2

import sys

class PhotowalkImage(object):
    '''An image representation class.'''
    image = None
    filename = None
    metadata = None

    def __init__(self, filename):
        '''filename -> PhotowalkImage

        Throws IOError if file type is not identified.'''
        self.image = Image.open(filename)
        self.filename = filename
        self.metadata = pyexiv2.ImageMetadata(filename)
        self.metadata.read()

    def __str__(self):
        s = '''Filename: %s
''' % (self.filename)

        for k, v in self.metadata.iteritems():
            s += '%s: %s\n' % (k, str(v.value))
        return s

def main():
    rc = 0
    prog = sys.argv[0]
    filenames = sys.argv[1:]
    for i in filenames:
        try:
            pwim = PhotowalkImage(i)
            print(pwim)
        except IOError, ex:
            msg = '%s: %s: %s\n' % (prog, i, str(ex))
            sys.stderr.write(msg)
            rc += 1

    sys.exit(rc)

if __name__ == '__main__':
    main()
