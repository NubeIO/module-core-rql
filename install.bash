#!/bin/bash

# bash install.bash version=v1.4.0 build=rubix-bios-1.4.0-83b470ea.armv7.zip services=no download=no clean=no ports=yes backup=no restore=no
# if services is "yes" it will stop and remove all the old rubix services
# if download is "yes" it download and unzip the build
# if backup is "yes" make a tmp dir to store some backups of things like FF database file
# if clean is "yes" it delete the /data dir
# if ports is "yes" it open all the ports needed for ROS on UFW
# if restore is "yes" will restore the flow-framework db file to /data/rubix-os/data


for ARGUMENT in "$@"
do
   KEY=$(echo $ARGUMENT | cut -f1 -d=)

   KEY_LENGTH=${#KEY}
   VALUE="${ARGUMENT:$KEY_LENGTH+1}"

   export "$KEY"="$VALUE"
done


if [ "$version" = "" ];
then
    version="v1.4.0"
fi

if [ "$build" = "" ];
then
    build="rubix-bios-1.4.0-83b470ea.armv7.zip"
fi

url="https://github.com/NubeIO/rubix-bios/releases/download/$version/$build"

home=$HOME
zip_build="$build"
echo DOWNLOAD: $url  BUILD: $zip_build

if [ "$download" = "yes" ];
then
  wget $url
  unzip $zip_build
fi


if [ "$ports" = "no" ]; # will open ports
then
  echo "dont open ports"
else
     ufw allow 1659  # BIOS
     ufw allow 1660  # ROS
fi


if [ "$services" = "yes" ];
then
    echo STOP OLD/DISABLE SERVICES
    sudo systemctl stop nubeio-flow-framework
    sudo systemctl disable nubeio-flow-framework

    sudo systemctl stop nubeio-rubix-bios
    sudo systemctl disable nubeio-rubix-bios

    sudo systemctl stop nubeio-rubix-service
    sudo systemctl disable nubeio-rubix-service

    sudo systemctl stop nubeio-wires-plat
    sudo systemctl disable nubeio-wires-plat
fi


if [ "$backup" = "yes" ]; # will make a tmp dir to store some backups
then
  mkdir rubix-tmp
  mkdir rubix-tmp/flow-framework

  # make a copy of FF database file
  cp "/data/flow-framework/data/data.db" "/home/pi/rubix-tmp/flow-framework"

fi

if [ "$clean" = "yes" ]; # will delete the data dir
then

  sudo rm -r /data/flow-framework
  sudo rm -r /data/rubix-service
  sudo rm -r /data/rubix-registry
  sudo rm -r /data/rubix-assist
  sudo rm -r /data/rubix-bios
  sudo rm -r /data/rubix-edge
  sudo rm -r /data/rubix-store
fi


if [ "$install" = "yes" ];
then
  sudo ./rubix-bios install --arch=armv7
fi

if [ "$restore" = "yes" ];
then
  mkdir sudo mkdir -p /data/rubix-os/data
  # make a copy of FF database file
  cp "/home/pi/rubix-tmp/flow-framework/data.db" "/data/rubix-os/data"
fi