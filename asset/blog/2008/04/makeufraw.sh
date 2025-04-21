#!/bin/bash

NEF=$1 #"./test.pef"
TIF="./test.tif"

# Common options for running the profiling runs
UFRAW="./ufraw-batch --wavelet-denoising-threshold=30 --out-type=tiff --overwrite"

# Add/remove options to configure, if necessary
CONFIGURE="./configure --enable-extras --with-exiv2 --enable-mime"

echo "Running autogen..."
sh autogen.sh

# I use CC and CXX for older versions of UFRaw which did not allow C(XX)FLAGS to be overridden
# Generate code for the machine's native architecture, generate profiling info
export CC="gcc -march=native -fprofile-generate"
export CXX="g++ -march=native -fprofile-generate"
echo "configuring..."
$CONFIGURE
echo "making with profile-generate..."
make

for restore in clip lch hsv
do
  echo $UFRAW --restore=$restore --exposure=-1 $NEF
  $UFRAW --restore=$restore --exposure=-1 $NEF
  rm $TIF
done

for clip in digital film
do
  echo $UFRAW --clip=$clip --exposure=1 $NEF
  $UFRAW --clip=$clip --exposure=1 $NEF
  rm $TIF
done


for interpolation in eahd ahd vng four-color ppg bilinear
do
  echo $UFRAW --interpolation=$interpolation $NEF
  $UFRAW --interpolation=$interpolation $NEF
  rm $TIF
done

make clean
# Generate code for the machine's native architecture, use profiling info
export CC="gcc -march=native -fprofile-use"
export CXX="g++ -march=native -fprofile-use"
echo "configuring..."
$CONFIGURE
echo "making with profile-use..."
make
rm *.gcno *.gcda
