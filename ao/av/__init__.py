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
        '''filename -> PhotowalkImage'''
        self.image = Image.open(filename)
        self.filename = filename
        self.metadata = pyexiv2.ImageMetadata(filename)

    def __str__(self):
        s = '''Filename: %s
''' % (self.filename)
        return s

def main():
    filenames = sys.argv[1:]
    for i in filenames:
        pwim = PhotowalkImage(i)
        print(pwim)

if __name__ == '__main__':
    main()
