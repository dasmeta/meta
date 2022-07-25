#!/bin/bash

while true; do
    read -p "First we have to install Meta, is that okay? Y or N? If you choose N this script will exit." yn
    case $yn in
        [Yy]* ) /bin/bash -c "$(wget https://github.com/dasmeta/meta/releases/download/v0.1.0/meta-mac)"; break;;
        [Nn]* ) echo Exiting!; exit;;
        * ) echo "Please answer yes or no.";;
    esac
done

mv meta-mac meta
cp meta /usr/bin
chmod +x /usr/bin/meta
echo All done!
exit
