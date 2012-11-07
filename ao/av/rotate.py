#!/usr/bin/python

import pyexiv2

import sys

ORIENTATION_KEY = 'Exif.Image.Orientation'

def rotate(filename):
    md = pyexiv2.ImageMetadata(filename)
    md.read()
    print md[ORIENTATION_KEY].value
    

if __name__ == '__main__':
    rotate(sys.argv[1])
