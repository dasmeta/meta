#!/bin/bash

while true; do
    read -p "First we have to install Meta, is that okay? Y or N? If you choose N this script will exit." yn
    case $yn in
        [Yy]* ) /bin/bash -c "$(wget https://github.com/dasmeta/meta/releases/download/v0.1.0/meta.go)"; break;;
        [Nn]* ) echo Exiting!; exit;;
        * ) echo "Please answer yes or no.";;
    esac
done

brew install python@3.9 golang golang-migrate curl wget
go mod init meta
go build ./meta.go
cp meta /usr/bin
echo All done!
exit