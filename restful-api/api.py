#!/usr/bin/env python
#-*- coding: UTF-8 -*-
#coding=utf8

from flask import Flask, jsonify, abort, make_response, request, url_for
from flask_cors import CORS
import logging
#from flask.ext.httpauth import HTTPBasicAuth

app = Flask(__name__)
# support cross-origin
CORS(app, resources=r'/*', supports_credentials=True)

objects = [
    {
        'id': 1,
        'title': u'标题1',
        'description': u'描述1'
    },
    {
        'id': 2,
        'title': u'标题2',
        'description': u'描述2'
    }
]

token = ''


# 从数据库查询对象信息
def get_object_db(req_new):
    ret = objects

    return ret


def get_new_object():
    return {}


def make_public_object(object):
    new_object = {}
    for field in object:
        if field == 'id':
            new_object['uri'] = url_for('get_object', object_id=object['id'], _external=True)
        else:
            new_object[field] = object[field]
    return new_object


def valid_token(req_token):
    if req_token == 'ABCDEF0123456789':
        return True
    else:
        return False


def json_contents(ret):
    response = make_response(jsonify(ret))
    response.headers['Access-Control-Allow-Origin'] = '*'
    response.headers['Access-Control-Allow-Methods'] = 'OPTIONS,HEAD,GET,POST'
    response.headers['Access-Control-Allow-Headers'] = 'x-requested-with,content-type'
    return response


@app.route('/api/v1.0/objects', methods=['GET'])
def get_objects():
    req_new = request.args.get('new', u'')
    logging.info("req_new=" + req_new)
    req_token = request.args.get('token', u'')
    if req_token != 'ABCDEF0123456789':
        return jsonify({'error': 'invalid token!! Please re-login!!'})
    else:
        return jsonify({'objects': get_object_db(req_new)})


@app.route('/api/v1.0/objects/<int:object_id>', methods=['GET'])
def get_object(object_id):
    req_token = request.args.get('token', u'')
    if valid_token(req_token):
        if object_id == '0':
            ret = get_new_object()
            return jsonify({'object': ret})
        object = filter(lambda t: t['id'] == object_id, get_object_db(None))
        if len(object) == 0:
            abort(404)
        return jsonify({'object': object[0]})
    else:
        return jsonify({'error': 'invalid token!! Please re-login!!'})


@app.errorhandler(404)
def not_found(error):
    return make_response(jsonify({'error': 'Not found'}), 404)


@app.route('/api/v1.0/get_token', methods=['POST'])
def get_token():
    if not request.form.get('user'):
       abort(400)
    
    user = request.form.get('user')
    passwd = request.form.get('passwd')
    
    if user == 'test' and passwd == 'test':
        token = 'ABCDEF0123456789'
        return json_contents({'token': token})
    else:
        return json_contents({'error': 'Invalid user or password!! Please re-login!!'})


@app.route('/api/v1.0/login', methods=['POST'])
def login_verify():
    if not request.json or not 'user' in request.json:
        abort(400)
    if request.json['user'] == 'test' and request.json['passwd'] == 'test':
        token = 'ABCDEF0123456789'
        return jsonify({'token': token})
    else:
        return None


@app.route('/api/v1.0/objects', methods=['POST'])
def create_object():
    global objects
    # if not request.json or not 'title' in request.json or not 'token' in request.json:
    #     abort(400)
    if not valid_token(request.args['token']):
        return jsonify({'error': 'invalid token!! Please re-login!!'})
    
    title = request.json.get('title')
    description = request.json.get('description', "")
    newobject = {
        'id': objects[-1]['id'] + 1,
        'title': title,
        'description': description,
    }
    objects.append(newobject)
    return jsonify({'object': newobject}), 201


@app.route('/api/v1.0/objects/<int:object_id>', methods=['PUT'])
def update_object(object_id):
    object = filter(lambda t: t['id'] == object_id, objects)
    if len(object) == 0:
        abort(404)
    if not request.json:
        abort(400)
    if 'title' in request.json and type(request.json['title']) != unicode:
        abort(400)
    if 'description' in request.json and type(request.json['description']) is not unicode:
        abort(400)
    object[0]['title'] = request.json.get('title', object[0]['title'])
    object[0]['description'] = request.json.get('description', object[0]['description'])
    return jsonify({'object': object[0]})


@app.route('/api/v1.0/objects/<int:object_id>', methods=['DELETE'])
def delete_object(object_id):
    global objects
    object = filter(lambda t: t['id'] == object_id, objects)
    if len(object) == 0:
        abort(404)
    objects.remove(object[0])
    return jsonify({'result': True})


if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0", port=8081)