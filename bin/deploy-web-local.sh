#!/usr/bin/env bash

PRODUCTION_BUNDLE=$(shasum ./build/static/main.js | cut -d ' ' -f 1)
mv ./build/static/main.js ./build/static/${PRODUCTION_BUNDLE}.js
