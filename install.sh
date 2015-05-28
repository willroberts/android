#!/bin/bash
rm android.apk
gomobile build -v
/usr/local/android-sdk/platform-tools/adb push 'android.apk' '/mnt/sdcard/golang/'
