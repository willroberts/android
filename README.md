# android

Android Apps in Go

## Requirements

### Android SDK and NDK

Download and install each from Google. I've stored them here:

    /usr/local/android-sdk
    /usr/local/android-ndk

#### Configure environment

    export SDK_ROOT=/usr/local/android-sdk
    export NDK_ROOT=/usr/local/android-ndk
    export PATH="${PATH}:${SDK_ROOT}/tools"
    export PATH="${PATH}:${SDK_ROOT}/platform-tools"

#### Build the NDK toolchain

    $ $NDK_ROOT/build/tools/make-standalone-toolchain.sh --install-dir=$NDK_ROOT --arch=arm --toolchain=arm-linux-androideabi-4.9 --platform=android-21

### Go

Most of this comes from [this doc](https://godoc.org/golang.org/x/mobile/cmd/gomobile).

#### Go 1.4

Install in `/usr/local/go1.4`. You'll need it to build Go 1.5.

#### Go 1.5

Install in `/usr/local/go`. If you already have Go 1.5 installed, go
ahead and rebuild it anyway (from HEAD of master). Things are being
fixed rapidly and it's worth the update.

#### Configure environment

    export GOPATH=${HOME}/go
    export GOROOT=/usr/local/go
    export GOROOT_BOOTSTRAP=/usr/local/go1.4
    export PATH="${GOPATH}/bin:${GOROOT}/bin:${PATH}"

#### gomobile

Download and install `gomobile`:

    $ go get github.com/golang/mobile/cmd/gomobile

### Install dependencies

    $ sudo apt-get -y install libgles2-mesa-dev libegl1-mesa-dev

#### Build go for android/arm

Might require NDK toolchain

    $ cd $GOROOT/src
    $ CC_FOR_TARGET=$NDK_ROOT/bin/arm-linux-androideabi-gcc GOOS=android GOARCH=arm GOARM=7 CGO_ENABLED=1 ./make.bash --no-clean

## Testing Your Setup

This will download, build, and install the [`basic`](https://godoc.org/golang.org/x/mobile/example/basic) example app.

Haven't made this work yet.

### Test the app locally

#### Set your CC back to normal (?)

    $ cd $GOROOT/src
    $ CC_FOR_TARGET= GOOS=linux GOARCH=amd64 ./make.bash

#### Install the app

    $ go install golang.org/x/mobile/example/basic
    $ basic

### Build the app for Android

    $ go get -d golang.org/x/mobile/example/basic
    $ gomobile build golang.org/x/mobile/example/basic

### Push to phone

    $ gomobile install golang.org/x/mobile/example/basic
    - or -
    $ ./install basic.apk
