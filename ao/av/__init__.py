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

def main():
    filenames = sys.argv[1:]
    for i in filenames:
        print i

if __name__ == '__main__':
    main()
