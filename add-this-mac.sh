#!/bin/bash

brew install python@3.9 
brew install golang 
brew install golang-migrate 
brew install curl 
brew install wget
curl https://api.github.com/repos/dasmeta/pre-commit-terraform/releases/tags/v2.1.1 | grep "tarball_url" | grep -Eo 'https://[^\"]*' | sed -n '1p' | xargs wget -O - | tar -xz
cd dasmeta-pre*/
chmod +x install.sh
./meta.sh
alias meta=`go list -f "{{.Target}}"`/meta
