#!/bin/bash

cd "$(dirname "$0")"

FRONTEND_DIR=../internal/web-app
STATIC_ASSET_DIR=../pkg/services/webRouter/static

cd ${FRONTEND_DIR}
npm run build --silent
cd -

rm -rf ${STATIC_ASSET_DIR}/*
cp -R ${FRONTEND_DIR}/dist/* ${STATIC_ASSET_DIR}/