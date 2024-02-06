#!bash

build_dir="release"
build_name=$(basename $(pwd))

build() {
    echo "Building for $1 $2"
    mkdir -p $build_dir/$build_name-$1-$2
    GOOS=$1 GOARCH=$2 go build -o $build_dir/$build_name-$1-$2/$build_name$(if [ $1 = "windows" ]; then echo ".exe"; fi) ./cmd/main.go
}

rm -rf $build_dir
mkdir -p $build_dir

build linux arm
build linux 386
build linux amd64
build windows 386
build windows amd64
build darwin amd64
build darwin arm64

ls -R $build_dir

cd $build_dir
for file in `ls`
do
    zip -r $file.zip $file
    rm -rf $file
done
