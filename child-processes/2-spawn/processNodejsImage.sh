#!/bin/bash
curl -s https://nodejs.org/static/images/logos/nodejs-new-pantone-black.svg > nodejs-logo.svg
base64 -i nodejs-logo.svg
