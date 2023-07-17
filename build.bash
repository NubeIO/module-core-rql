#!/bin/bash
echo PATH $1/rubix-os
echo ****MODULES****
ls  $1/rubix-os/data/modules
echo ****MODULES****
go build -o module-core-rql && mv -f module-core-rql $1/rubix-os/data/modules && cd $1/rubix-os && bash build.bash
