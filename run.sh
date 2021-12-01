go mod tidy
rm -f .gop/gop.cache
cd assets
gop run ..
