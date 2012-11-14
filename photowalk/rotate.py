#!/usr/bin/python

import Image
import pyexiv2

import sys

ORIENTATION_KEY = 'Exif.Image.Orientation'

def rotate(image, filename):
    '''Image, filename -> image rotated according to ORIENTATION_KEY.'''
    md = pyexiv2.ImageMetadata(filename)
    md.read()
    o = md[ORIENTATION_KEY].value
    rotation = 0
    # o == 1 is already correctly rotated.
    if o == 8: # left-bottom according to exif command.
        rotation = 90
    elif o == 6: # right-top accordin to exif command.
        rotation = 270
    return image.rotate(rotation)

if __name__ == '__main__':
    fn = sys.argv[1]
    im = Image.open(fn)
    im = rotate(im, fn)
    im.show()
