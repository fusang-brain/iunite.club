# !/bin/bash

user=alixez
host=7.7.1.226

rsync -avH --progress --exclude=.git ../iunite.club/ ${user}@${host}:~/iunite.club/