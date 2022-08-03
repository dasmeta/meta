#!/bin/bash

while true; do
    read -p "First we have to install Meta, is that okay? Y or N? If you choose N this script will exit." yn
    case $yn in
        [Yy]* ) sudo wget https://github.com/dasmeta/meta/releases/download/v0.1.0/meta-linux -O /usr/bin/meta; sudo chmod +x /usr/bin/meta
; break;;
        [Nn]* ) echo Exiting!; exit;;
        * ) echo "Please answer yes or no.";;
    esac
done

<<<<<<< HEAD
sudo mv meta-linux meta
sudo cp meta /usr/local/bin/
sudo chmod +x /usr/local/bin/meta
=======
>>>>>>> origin
echo All done!
exit
