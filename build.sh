#!/bin/sh

go build .

if [ ! -d "deploy" ]; then
    mkdir deploy
fi

mkdir -p deploy/common/config
cp -f common/config/gocloud.toml deploy/common/config/

mkdir -p deploy/static/css
cp -f static/css/* deploy/static/css

mkdir -p deploy/static/fonts
cp -f static/fonts/* deploy/static/fonts

mkdir -p deploy/static/i
cp -f static/i/* deploy/static/i

mkdir -p deploy/static/js
cp -f static/js/* deploy/static/js

mkdir -p deploy/views
cp -f  views/* deploy/views

case "`uname`" in
  CYGWIN*)
    cp -f gocloud.exe deploy
    ;;
  *)
    cp -f gocloud deploy
esac

cd deploy
tar zcvf gocloud.tar.gz *
