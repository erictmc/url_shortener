#!/usr/bin/env bash

cd ./testing || exit
yarn install
yarn test
cd ..
