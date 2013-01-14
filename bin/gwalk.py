#!/usr/bin/python

import pygtk
pygtk.require('2.0')
import gtk

class Photowalk(object):
    def __init__(self):
        self.window = gtk.Window(gtk.WINDOW_TOPLEVEL)
        self.window.show()

    def main(self):
        gtk.main()

if __name__ == '__main__':
    mw = Photowalk()
    mw.main()
