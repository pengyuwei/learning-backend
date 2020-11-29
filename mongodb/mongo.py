#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

import pymongo
import datetime


def test_mongo():
    # connect
    mongo_client = pymongo.MongoClient('127.0.0.1', 27017)
    print(mongo_client.server_info())
    mongo_db = mongo_client['testdb']
    mongo_collection = mongo_db['testcol']

    # insert
    data = {
        'name': '0xFF',
        'descr': 'Some text.',
        'val': {'A': 1, 'B': 'is b'},
        'date': datetime.datetime.now()
    }
    mongo_collection.insert_one(data)

    # update
    where = {'name': '0xFF'}
    data = {
        'name': '0xFF',
        'tags': ['2020', 'blog'],
        'val': {'A': 1, 'B': 'is b'},
        'date': datetime.datetime.now()
    }
    mongo_collection.update_one(where, {'$set': data})

    # find
    where = {
        'name': '0xFF',
        'tags': ['2020', 'blog']
    }
    ret = mongo_collection.find_one(where)
    print(ret)

    where = {
        'name': '0xFF',
    }
    ret = mongo_collection.find_one(
        where,
        projection={'_id': False, 'name': True, 'tags': True})
    print(ret)

    count = mongo_collection.count_documents(where)
    print("count=", count)

    rs = mongo_collection.find(where)
    for itor in rs:
        print(itor)

    # delete
    mongo_collection.delete_one({'name': '0xFF'})

    mongo_client.close()


if __name__ == '__main__':
    test_mongo()
