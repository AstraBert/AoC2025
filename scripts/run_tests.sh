for dir in ./day-*
do
    cd $dir
    go test ./...
    cd ..
done