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

#### Go 1.4

Install Go 1.4 to `/usr/local/go1.4`. You'll need it to build Go 1.5.

#### Go 1.5

Install to `/usr/local/go`.

#### Configure environment

    export GOPATH=${HOME}/go
    export GOROOT=/usr/local/go
    export GOROOT_BOOTSTRAP=/usr/local/go1.4
    export PATH="${GOPATH}/bin:${GOROOT}/bin:${PATH}"

#### gomobile

Download and install `gomobile`:

    go get github.com/golang/mobile/cmd/gomobile

#### Build go for arm/android

Might require NDK toolchain

    $ cd $GOROOT/src
    $ CC_FOR_TARGET=$NDK_ROOT/bin/arm-linux-androideabi-gcc GOOS=android GOARCH=arm GOARM=7 CGO_ENABLED=1 ./make.bash
