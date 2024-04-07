# ARM Computer
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.2.linux-arm64.tar.gz

# x86_64 Computer
# rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

go version
