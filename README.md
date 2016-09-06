# mackerel-websocket

## how to developments.

1. please download Go-Lang development tools from https://golang.org/dl/
2. installing a package,
  - tar -C /usr/local -xzf goHOGEHOGE.tgz
3. configure environment values.
  - export PATH=$PATH:/usr/local/go/bin (ex: bash)
  - export GOPATH=hogefuga (your settings)
  - for examples...
  ```
  cat /etc/profile.d/golang.sh
  export GOROOT=/usr/local/go
  export PATH=$PATH:/usr/local/go/bin
  export GOPATH=${HOME}/gowork
  ```
4. import a libraries
  - run this cmd:
  ```
  > go get github.com/jessevdk/go-flags
  > go get github.com/mackerelio/checkers
  > go get golang.org/x/net/websocket
  ```
6. enjoy!

## how to build

go build [source file name]

references:<br/>
https://golang.org/doc/install
