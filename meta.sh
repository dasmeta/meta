#!/bin/bash

while true; do
    read -p "First we have to install Meta, is that okay? Y or N? If you choose N this script will exit." yn
    case $yn in
        [Yy]* ) /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/dasmeta/pre-commit-terraform/v2.1.1/install.sh)"; break;;
        [Nn]* ) echo Exiting!; exit;;
        * ) echo "Please answer yes or no.";;
    esac
done

echo All done!
exit
