#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

# pip install sinaweibopy
# Python 3.6.9

import json
from weibo import APIClient

APP_KEY = 'YOUR_APP_KEY'
APP_SECRET = 'YOUR_APP_SECRET'
CALLBACK_URL = 'YOUR_CALLBACK_URL'


def load():
    global APP_KEY, APP_SECRET, CALLBACK_URL
    """
    temp.json:
    {
        "APP_KEY": "",
        "APP_SECRET": "",
        "CALLBACK_URL": "",
    }
    """
    with open('temp.json') as json_file:
        config = json.load(json_file)
    APP_KEY = config['APP_KEY']
    APP_SECRET = config['APP_SECRET']
    CALLBACK_URL = config['CALLBACK_URL']
    print(config)


def main():
    client = APIClient(app_key=APP_KEY, app_secret=APP_SECRET, redirect_uri=CALLBACK_URL)
    url = client.get_authorize_url()
    print(url)


def on_callback(code):
    # TODO: HOW TO callback here?
    r = client.request_access_token(code)
    access_token = r.access_token
    expires = r.expires # UNIX timestamp, e.g., 1384826449.252 (10:01 am, 19 Nov 2013, UTC+8:00)
    print(access_token, expires)
    client = APIClient(app_key=APP_KEY,
            app_secret=APP_SECRET,
            redirect_uri=CALLBACK_URL,
            access_token=YOUR_ACCESS_TOKEN,
            expires=EXPIRES_TIME)
    invoke_api(client)


def invoke_api(client):
    print(client.statuses.user_timeline.get())
    print(client.statuses.update.post(status=u'test weibo'))
    print(client.statuses.upload.post(status=u'test weibo with picture',
                                    pic=open('./test.png')))
    # publish a picture weibo
    f = open('./test.png', 'rb')
    r = client.statuses.upload.post(status=u'test weibo with picture', pic=f)
    f.close()


if __name__ == '__main__':
    load()
    main()
