#!/usr/bin/python

'''
Google API client for photowalk.

https://developers.google.com/picasa-web/docs/1.0/developers_guide_python
'''

import gdata.photos.service


def main():
    '''The main function.'''
    username = 'hauva69@gmail.com'
    gd_client = gdata.photos.service.PhotosService()
    albums = gd_client.GetUserFeed(user=username)
    for album in albums.entry:
        print 'title: %s, number of photos: %s, id: %s' % (album.title.text,
                                                           album.numphotos.text,
                                                           album.gphoto_id.text)

if __name__ == '__main__':
    main()
