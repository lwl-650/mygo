# 1 生成go.mod文件
`` go mod init "文件名"

# 2 下载vscode依赖 看情况
``>>> 
go get -v github.com/mdempsky/gocode;
go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs;
go get -v github.com/ramya-rao-a/go-outline;
go get -v github.com/acroca/go-symbols;
go get -v golang.org/x/tools/cmd/guru;
go get -v golang.org/x/tools/cmd/gorename;
go get -v github.com/cweill/gotests/gotests;
go get -v github.com/fatih/gomodifytags;
go get -v github.com/josharian/impl;
go get -v github.com/davidrjenni/reftools/cmd/fillstruct;
go get -v github.com/haya14busa/goplay/cmd/goplay;
go get -v github.com/godoctor/godoctor;
go get -v github.com/go-delve/delve/cmd/dlv;
go get -v github.com/stamblerre/gocode;
go get -v github.com/rogpeppe/godef;
go get -v github.com/sqs/goreturns;
go get -v golang.org/x/lint/golint


# 3
`` go get github.com/pilu/fresh