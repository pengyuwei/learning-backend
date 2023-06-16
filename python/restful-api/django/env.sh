#!/bin/bash

python3 -m venv env
source env/bin/activate

python3 -m pip install --upgrade pip
pip install django
pip install djangorestframework

django-admin startproject demo_api
cd demo_api

python manage.py runserver

deactivate
