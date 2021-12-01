go mod tidy
rm -f .gop/gop.cache
gop build .
rm -rf FlappyBird
gomobile build --tags canvas  -target=android  -androidapi 29