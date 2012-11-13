#!/usr/bin/python

PICASA_MAX_SIZE = 2048

import Image
import sys

class PhotowalkImage(object):
    image = None
    filename = None

    def __init__(self, filename):
        image = Image.open(filename)
        self.filename = filename

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
