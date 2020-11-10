#!/usr/bin/env bash

go build
zip alfred-cron.alfredworkflow ./alfred-cron ./info.plist ./icon.png
rm ./alfred-cron
