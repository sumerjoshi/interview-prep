#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd $DIR
cd api
LOG_LEVEL=DEBUG FLASK_APP=api flask run -p 8080 --host=0.0.0.0
