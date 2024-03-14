gop mod tidy
gop build .
rm -rf FlappyBird
gomobile build --tags canvas  -target=android  -androidapi 29
