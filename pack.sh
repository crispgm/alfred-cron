#!/bin/bash

go build
zip alfred-cron.alfredworkflow ./alfred-cron ./info.plist
rm ./alfred-cron
