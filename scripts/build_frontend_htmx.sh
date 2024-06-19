#!/bin/bash

cd "$(dirname "$0")"

FRONTEND_BUILD_DIR=../internal/frontend/htmx
STATIC_ASSETS_MIDDLEWARE_DIR=../pkg/services/webRouter/static

cd ${FRONTEND_BUILD_DIR}
npm run build --silent
cd -

rm -rf ${STATIC_ASSETS_MIDDLEWARE_DIR}/*
cp -R ${FRONTEND_BUILD_DIR}/* ${STATIC_ASSETS_MIDDLEWARE_DIR}/