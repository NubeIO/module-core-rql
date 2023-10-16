#!/bin/bash

code_path=$1

if [ "$code_path" = "" ];
then
    code_path="code/go"
    echo NO PATH PROVIDED WILL USE: $code_path
else
    echo PATH PROVIDED WILL USE: $code_path
fi

path=$HOME/$code_path/rubix-os


echo $path
echo "****MODULES****"
ls  $path/data/modules
echo "****MODULES****"

go build -o module-core-rql && mv -f module-core-rql $path/data/rubix-os/data/modules/module-core-rql/v0.0.2 && cd $path && bash build.bash system
