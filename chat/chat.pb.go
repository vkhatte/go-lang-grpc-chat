Last login: Fri Jul 22 09:09:08 on ttys004

The default interactive shell is now zsh.
To update your account to use zsh, please run `chsh -s /bin/zsh`.
For more details, please visit https://support.apple.com/kb/HT208050.
Veerabhushans-MacBook-Pro:~ vhatte$ pwd
/Users/vhatte
Veerabhushans-MacBook-Pro:~ vhatte$ cd go
go/      go-lang/
Veerabhushans-MacBook-Pro:~ vhatte$ cd go
Veerabhushans-MacBook-Pro:go vhatte$ ls
bin	pkg	src
Veerabhushans-MacBook-Pro:go vhatte$ cd pkg/
Veerabhushans-MacBook-Pro:pkg vhatte$ ls
darwin_amd64	mod		sumdb
Veerabhushans-MacBook-Pro:pkg vhatte$ cd mod/
Veerabhushans-MacBook-Pro:mod vhatte$ ls
cache							google.golang.org
github.com						gopkg.in
go.starlark.net@v0.0.0-20200821142938-949cc6f4b097	honnef.co
golang.org						mvdan.cc
Veerabhushans-MacBook-Pro:mod vhatte$ ls -ltr
total 0
drwxr-xr-x   3 vhatte  staff   96 Jul 14 15:42 cache
drwxr-xr-x   3 vhatte  staff   96 Jul 14 15:42 golang.org
drwxr-xr-x   3 vhatte  staff   96 Jul 14 15:42 gopkg.in
dr-xr-xr-x  19 vhatte  staff  608 Jul 14 15:42 go.starlark.net@v0.0.0-20200821142938-949cc6f4b097
drwxr-xr-x   3 vhatte  staff   96 Jul 14 15:42 honnef.co
drwxr-xr-x   4 vhatte  staff  128 Jul 14 15:42 mvdan.cc
drwxr-xr-x  25 vhatte  staff  800 Jul 22 13:44 github.com
drwxr-xr-x   5 vhatte  staff  160 Jul 22 13:44 google.golang.org
Veerabhushans-MacBook-Pro:mod vhatte$ cd google.golang.org
Veerabhushans-MacBook-Pro:google.golang.org vhatte$ ls
genproto@v0.0.0-20200526211855-cb27e3aa2013	protobuf@v1.27.1
grpc@v1.48.0
Veerabhushans-MacBook-Pro:google.golang.org vhatte$ cd grpc\@v1.48.0/
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ ls
AUTHORS						go.sum
CODE-OF-CONDUCT.md				grpc_test.go
CONTRIBUTING.md					grpclog
Documentation					health
GOVERNANCE.md					interceptor.go
LICENSE						internal
MAINTAINERS.md					interop
Makefile					keepalive
NOTICE.txt					metadata
README.md					peer
SECURITY.md					picker_wrapper.go
admin						picker_wrapper_test.go
attributes					pickfirst.go
authz						preloader.go
backoff						profiling
backoff.go					reflection
balancer					regenerate.sh
balancer_conn_wrappers.go			resolver
benchmark					resolver_conn_wrapper.go
binarylog					rpc_util.go
call.go						rpc_util_test.go
channelz					server.go
clientconn.go					server_test.go
clientconn_authority_test.go			service_config.go
clientconn_parsed_target_test.go		service_config_test.go
clientconn_state_transition_test.go		serviceconfig
clientconn_test.go				stats
codec.go					status
codec_test.go					stream.go
codegen.sh					stress
codes						tap
connectivity					test
credentials					testdata
default_dial_option_server_option_test.go	trace.go
dialoptions.go					trace_test.go
doc.go						version.go
encoding					vet.sh
go.mod						xds
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ ls grpc
grpc_test.go  grpclog/
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ ls grpc
ls: grpc: No such file or directory
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ ls grpc
grpc_test.go  grpclog/
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ pwd
/Users/vhatte/go/pkg/mod/google.golang.org/grpc@v1.48.0
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ pwd
/Users/vhatte/go/pkg/mod/google.golang.org/grpc@v1.48.0
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ pwd
/Users/vhatte/go/pkg/mod/google.golang.org/grpc@v1.48.0
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ ls /Users/vhatte/go-lang/source/grpc-example/chat
proto	server	src
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ cd
Veerabhushans-MacBook-Pro:~ vhatte$ cd go
go/      go-lang/
Veerabhushans-MacBook-Pro:~ vhatte$ cd go/
bin/ pkg/ src/
Veerabhushans-MacBook-Pro:~ vhatte$ cd go/bin/
dlv            goimports      goplay         gotests        protoc-gen-go
go-outline     gomodifytags   gopls          impl           staticcheck
Veerabhushans-MacBook-Pro:~ vhatte$ cd go/pkg/
darwin_amd64/ mod/          sumdb/
Veerabhushans-MacBook-Pro:~ vhatte$ go
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	work        workspace maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

	buildconstraint build constraints
export GOROOT=$HOME/go
	buildmode       build modes
	c               calling between Go and C
	cache           build and test caching
	environment     environment variables
	filetype        file types
	go.mod          the go.mod file
	gopath          GOPATH environment variable
	gopath-get      legacy GOPATH go get
	goproxy         module proxy protocol
	importpath      import path syntax
	modules         modules, module versions, and more
	module-get      module-aware go get
	module-auth     module authentication using go.sum
	packages        package lists and patterns
	private         configuration for downloading non-public code
	testflag        testing flags
	testfunc        testing functions
	vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.

Veerabhushans-MacBook-Pro:~ vhatte$ which go
/usr/local/go/bin/go
Veerabhushans-MacBook-Pro:~ vhatte$ pwd
/Users/vhatte
Veerabhushans-MacBook-Pro:~ vhatte$ cd go/bin/
Veerabhushans-MacBook-Pro:bin vhatte$ ls
dlv		goimports	goplay		gotests		protoc-gen-go
go-outline	gomodifytags	gopls		impl		staticcheck
Veerabhushans-MacBook-Pro:bin vhatte$ cd
Veerabhushans-MacBook-Pro:~ vhatte$ vi .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ source .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ export
export GOROOT=$HOME/go
declare -x COLORFGBG="7;0"
declare -x COLORTERM="truecolor"
declare -x DISPLAY="/private/tmp/com.apple.launchd.a0wnPkAjMQ/org.macosforge.xquartz:0"
declare -x GOROOT=":/Users/vhatte/go/go"
declare -x HOME="/Users/vhatte"
declare -x ITERM_PROFILE="local"
declare -x ITERM_SESSION_ID="w0t1p0:2850BF85-208E-4592-94E5-6B6ABDD44C44"
declare -x LANG="en_US.UTF-8"
declare -x LOGNAME="vhatte"
declare -x LaunchInstanceID="7688B80F-11DC-40BE-85DA-A1FCAC914A8E"
declare -x OLDPWD="/Users/vhatte/go/bin"
export GOROOT=$HOME/go
declare -x PATH="/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/usr/local/share/dotnet:/opt/X11/bin:~/.dotnet/tools:/Library/Frameworks/Mono.framework/Versions/Current/Commands:/Applications/Wireshark.app/Contents/MacOS:/Users/vhatte/go:/Users/vhatte/go/bin/::/Users/vhatte/go/go::/Users/vhatte/go/go/bin/"
declare -x PWD="/Users/vhatte"
declare -x SECURITYSESSIONID="186a6"
declare -x SHELL="/bin/bash"
declare -x SHLVL="1"
declare -x SSH_AUTH_SOCK="/private/tmp/com.apple.launchd.EAupGqqZNT/Listeners"
declare -x TERM="xterm-256color"
declare -x TERM_PROGRAM="iTerm.app"
declare -x TERM_PROGRAM_VERSION="3.2.8"
declare -x TERM_SESSION_ID="w0t1p0:2850BF85-208E-4592-94E5-6B6ABDD44C44"
declare -x TMPDIR="/var/folders/80/1v6c4ksd3ns1xbfx64lp110m0000gn/T/"
declare -x USER="vhatte"
declare -x XPC_FLAGS="0x0"
declare -x XPC_SERVICE_NAME="0"
declare -x __CF_USER_TEXT_ENCODING="0x0:0:0"
Veerabhushans-MacBook-Pro:~ vhatte$ export | grep GO
declare -x GOROOT=":/Users/vhatte/go/go"
Veerabhushans-MacBook-Pro:~ vhatte$ vi .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ source .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ export | grep GO
declare -x GOROOT="/Users/vhatte/go:/Users/vhatte/go/go"
Veerabhushans-MacBook-Pro:~ vhatte$ pw
-bash: pw: command not found
Veerabhushans-MacBook-Pro:~ vhatte$ pwd
/Users/vhatte
Veerabhushans-MacBook-Pro:~ vhatte$ cd go/
bin/ pkg/ src/
Veerabhushans-MacBook-Pro:~ vhatte$ source .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ vi .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ source .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ go
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	work        workspace maintenance
	run         compile and run Go program
	test        test packages
export GOROOT=$HOME/go
	tool        run specified go tool
	version     print Go version
export GOROOT=$GOROOT:$GOROOT/bin/
	vet         report likely mistakes in packages
export PATH=$PATH:$HOME/go:$HOME/go/bin

Use "go help <command>" for more information about a command.

Additional help topics:

	buildconstraint build constraints
	buildmode       build modes
	c               calling between Go and C
	cache           build and test caching
	environment     environment variables
	filetype        file types
	go.mod          the go.mod file
	gopath          GOPATH environment variable
	gopath-get      legacy GOPATH go get
	goproxy         module proxy protocol
	importpath      import path syntax
	modules         modules, module versions, and more
	module-get      module-aware go get
	module-auth     module authentication using go.sum
	packages        package lists and patterns
	private         configuration for downloading non-public code
	testflag        testing flags
	testfunc        testing functions
	vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.

Veerabhushans-MacBook-Pro:~ vhatte$ vi .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ source .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ vi .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$ vi .bash_profile
Veerabhushans-MacBook-Pro:~ vhatte$
Veerabhushans-MacBook-Pro:~ vhatte$
Veerabhushans-MacBook-Pro:~ vhatte$
Veerabhushans-MacBook-Pro:~ vhatte$ cd
Veerabhushans-MacBook-Pro:~ vhatte$
Veerabhushans-MacBook-Pro:~ vhatte$
Veerabhushans-MacBook-Pro:~ vhatte$ pwd
/Users/vhatte
Veerabhushans-MacBook-Pro:~ vhatte$ cd go
Veerabhushans-MacBook-Pro:go vhatte$ ls
bin	pkg	src
Veerabhushans-MacBook-Pro:go vhatte$ pwd
/Users/vhatte/go
Veerabhushans-MacBook-Pro:go vhatte$ ls
bin	pkg	src
Veerabhushans-MacBook-Pro:go vhatte$ cd pkg/
Veerabhushans-MacBook-Pro:pkg vhatte$ ls
darwin_amd64	mod		sumdb
Veerabhushans-MacBook-Pro:pkg vhatte$ cd mod/
cache/                                              google.golang.org/
github.com/                                         gopkg.in/
go.starlark.net@v0.0.0-20200821142938-949cc6f4b097/ honnef.co/
golang.org/                                         mvdan.cc/
Veerabhushans-MacBook-Pro:pkg vhatte$ ls
darwin_amd64	mod		sumdb
Veerabhushans-MacBook-Pro:pkg vhatte$ cd mod/
Veerabhushans-MacBook-Pro:mod vhatte$ cd google.golang.org/
Veerabhushans-MacBook-Pro:google.golang.org vhatte$ ls
genproto@v0.0.0-20200526211855-cb27e3aa2013	protobuf@v1.27.1
grpc@v1.48.0
Veerabhushans-MacBook-Pro:google.golang.org vhatte$ cd grpc\@v1.48.0/
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ ls
AUTHORS						go.sum
CODE-OF-CONDUCT.md				grpc_test.go
CONTRIBUTING.md					grpclog
Documentation					health
GOVERNANCE.md					interceptor.go
LICENSE						internal
MAINTAINERS.md					interop
Makefile					keepalive
NOTICE.txt					metadata
README.md					peer
SECURITY.md					picker_wrapper.go
admin						picker_wrapper_test.go
attributes					pickfirst.go
authz						preloader.go
backoff						profiling
backoff.go					reflection
balancer					regenerate.sh
balancer_conn_wrappers.go			resolver
benchmark					resolver_conn_wrapper.go
binarylog					rpc_util.go
call.go						rpc_util_test.go
channelz					server.go
clientconn.go					server_test.go
clientconn_authority_test.go			service_config.go
clientconn_parsed_target_test.go		service_config_test.go
clientconn_state_transition_test.go		serviceconfig
clientconn_test.go				stats
codec.go					status
codec_test.go					stream.go
codegen.sh					stress
codes						tap
connectivity					test
credentials					testdata
default_dial_option_server_option_test.go	trace.go
dialoptions.go					trace_test.go
doc.go						version.go
encoding					vet.sh
go.mod						xds
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ pwd
/Users/vhatte/go/pkg/mod/google.golang.org/grpc@v1.48.0
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ pwd
/Users/vhatte/go/pkg/mod/google.golang.org/grpc@v1.48.0
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ pwd
/Users/vhatte/go/pkg/mod/google.golang.org/grpc@v1.48.0
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ pwd
/Users/vhatte/go/pkg/mod/google.golang.org/grpc@v1.48.0
Veerabhushans-MacBook-Pro:grpc@v1.48.0 vhatte$ cd
Veerabhushans-MacBook-Pro:~ vhatte$ cd go
Veerabhushans-MacBook-Pro:go vhatte$ pwd
/Users/vhatte/go
Veerabhushans-MacBook-Pro:go vhatte$ mkdir chat
Veerabhushans-MacBook-Pro:go vhatte$ cd chat/
Veerabhushans-MacBook-Pro:chat vhatte$
Veerabhushans-MacBook-Pro:chat vhatte$
Veerabhushans-MacBook-Pro:chat vhatte$ ls
Veerabhushans-MacBook-Pro:chat vhatte$ pwd
/Users/vhatte/go/chat
Veerabhushans-MacBook-Pro:chat vhatte$ pwd
/Users/vhatte/go/chat
Veerabhushans-MacBook-Pro:chat vhatte$ cd ..
Veerabhushans-MacBook-Pro:go vhatte$ \rm -rf chat
Veerabhushans-MacBook-Pro:go vhatte$ cd src/
Veerabhushans-MacBook-Pro:src vhatte$ ls
github.com		golang.org		google.golang.org
Veerabhushans-MacBook-Pro:src vhatte$ mkdir chat
Veerabhushans-MacBook-Pro:src vhatte$ cd chat/
Veerabhushans-MacBook-Pro:chat vhatte$ pwd
/Users/vhatte/go/src/chat
Veerabhushans-MacBook-Pro:chat vhatte$ cd
Veerabhushans-MacBook-Pro:~ vhatte$
Veerabhushans-MacBook-Pro:~ vhatte$
Veerabhushans-MacBook-Pro:~ vhatte$ cd notes/
Veerabhushans-MacBook-Pro:notes vhatte$ vi go-build.notes
Veerabhushans-MacBook-Pro:notes vhatte$
package main
Veerabhushans-MacBook-Pro:notes vhatte$
Veerabhushans-MacBook-Pro:notes vhatte$
Veerabhushans-MacBook-Pro:notes vhatte$ pwd
/Users/vhatte/notes
Veerabhushans-MacBook-Pro:notes vhatte$ cd ..
Veerabhushans-MacBook-Pro:~ vhatte$ cd go
go/      go-lang/
Veerabhushans-MacBook-Pro:~ vhatte$ cd go-lang/
Veerabhushans-MacBook-Pro:go-lang vhatte$ ls
go-grpc			go-grpc-tutorial	greetings		hello			source
Veerabhushans-MacBook-Pro:go-lang vhatte$ cd go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ ls
chat		chat.proto	client.go	go.mod
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export | gre[ GO
>
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export | grep GO
declare -x GOROOT="/Users/vhatte/go:/Users/vhatte/go/bin/"
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export GOROOT=
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export | grep GO
declare -x GOROOT=""
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: import "chat" is a program, not an importable package
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi client.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi
chat/       chat.proto  client.go   go.mod      go.sum
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi go.mod
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi client.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: cannot find package "example.com/chat" in any of:
	/usr/local/go/src/example.com/chat (from $GOROOT)
	/Users/vhatte/go/src/example.com/chat (from $GOPATH)
package main
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export | grep GO
declare -x GOROOT=""
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export GOPATH=
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export GOPATH="/Users/vhatte/go-lang/"
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: cannot find package "example.com/chat" in any of:
	/usr/local/go/src/example.com/chat (from $GOROOT)
	/Users/vhatte/go-lang/src/example.com/chat (from $GOPATH)
client.go:6:5: cannot find package "golang.org/x/net/context" in any of:
	/usr/local/go/src/golang.org/x/net/context (from $GOROOT)
	/Users/vhatte/go-lang/src/golang.org/x/net/context (from $GOPATH)
client.go:7:5: cannot find package "google.golang.org/grpc" in any of:
	/usr/local/go/src/google.golang.org/grpc (from $GOROOT)
	/Users/vhatte/go-lang/src/google.golang.org/grpc (from $GOPATH)
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export GOPATH="/Users/vhatte/go-lang/go-grpc"
package main
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: cannot find package "example.com/chat" in any of:
	/usr/local/go/src/example.com/chat (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/example.com/chat (from $GOPATH)
client.go:6:5: cannot find package "golang.org/x/net/context" in any of:
	/usr/local/go/src/golang.org/x/net/context (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/golang.org/x/net/context (from $GOPATH)
client.go:7:5: cannot find package "google.golang.org/grpc" in any of:
	/usr/local/go/src/google.golang.org/grpc (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/google.golang.org/grpc (from $GOPATH)
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi client.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:6:5: cannot find package "golang.org/x/net/context" in any of:
	/usr/local/go/src/golang.org/x/net/context (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/golang.org/x/net/context (from $GOPATH)
chat/chat.pb.go:11:2: cannot find package "google.golang.org/grpc" in any of:
	/usr/local/go/src/google.golang.org/grpc (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/google.golang.org/grpc (from $GOPATH)
chat/chat.pb.go:12:2: cannot find package "google.golang.org/grpc/codes" in any of:
	/usr/local/go/src/google.golang.org/grpc/codes (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/google.golang.org/grpc/codes (from $GOPATH)
chat/chat.pb.go:13:2: cannot find package "google.golang.org/grpc/status" in any of:
	/usr/local/go/src/google.golang.org/grpc/status (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/google.golang.org/grpc/status (from $GOPATH)
chat/chat.pb.go:14:2: cannot find package "google.golang.org/protobuf/reflect/protoreflect" in any of:
	/usr/local/go/src/google.golang.org/protobuf/reflect/protoreflect (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/google.golang.org/protobuf/reflect/protoreflect (from $GOPATH)
chat/chat.pb.go:15:2: cannot find package "google.golang.org/protobuf/runtime/protoimpl" in any of:
	/usr/local/go/src/google.golang.org/protobuf/runtime/protoimpl (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/google.golang.org/protobuf/runtime/protoimpl (from $GOPATH)
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi client.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: cannot find package "chat" in any of:
	/usr/local/go/src/chat (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/chat (from $GOPATH)
client.go:6:5: cannot find package "golang.org/x/net/context" in any of:
	/usr/local/go/src/golang.org/x/net/context (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/golang.org/x/net/context (from $GOPATH)
client.go:7:5: cannot find package "google.golang.org/grpc" in any of:
	/usr/local/go/src/google.golang.org/grpc (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/google.golang.org/grpc (from $GOPATH)
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: cannot find package "chat" in any of:
	/usr/local/go/src/chat (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/chat (from $GOPATH)
client.go:6:5: cannot find package "golang.org/x/net/context" in any of:
	/usr/local/go/src/golang.org/x/net/context (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/golang.org/x/net/context (from $GOPATH)
client.go:7:5: cannot find package "google.golang.org/grpc" in any of:
	/usr/local/go/src/google.golang.org/grpc (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/google.golang.org/grpc (from $GOPATH)
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: cannot find package "chat" in any of:
	/usr/local/go/src/chat (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/chat (from $GOPATH)
client.go:6:5: cannot find package "golang.org/x/net/context" in any of:
	/usr/local/go/src/golang.org/x/net/context (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/golang.org/x/net/context (from $GOPATH)
client.go:7:5: cannot find package "google.golang.org/grpc" in any of:
	/usr/local/go/src/google.golang.org/grpc (from $GOROOT)
	/Users/vhatte/go-lang/go-grpc/src/google.golang.org/grpc (from $GOPATH)
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ cd chat
Veerabhushans-MacBook-Pro:chat vhatte$ ls
chat.pb.go	go.mod		go.sum
Veerabhushans-MacBook-Pro:chat vhatte$ pwd
/Users/vhatte/go-lang/go-grpc/chat
Veerabhushans-MacBook-Pro:chat vhatte$ vi chat.pb.go
Veerabhushans-MacBook-Pro:chat vhatte$ pwd
/Users/vhatte/go-lang/go-grpc/chat
Veerabhushans-MacBook-Pro:chat vhatte$ rm go.
go.mod  go.sum
Veerabhushans-MacBook-Pro:chat vhatte$ rm go.mod
Veerabhushans-MacBook-Pro:chat vhatte$ rm go.sum
Veerabhushans-MacBook-Pro:chat vhatte$
Veerabhushans-MacBook-Pro:chat vhatte$
Veerabhushans-MacBook-Pro:chat vhatte$
Veerabhushans-MacBook-Pro:chat vhatte$ pwd
/Users/vhatte/go-lang/go-grpc/chat
Veerabhushans-MacBook-Pro:chat vhatte$ ls
chat.pb.go
Veerabhushans-MacBook-Pro:chat vhatte$ go mod init
go: modules disabled by GO111MODULE=off; see 'go help modules'
Veerabhushans-MacBook-Pro:chat vhatte$ set GO111MODULE="on"
Veerabhushans-MacBook-Pro:chat vhatte$ go mod init
go: modules disabled by GO111MODULE=off; see 'go help modules'
Veerabhushans-MacBook-Pro:chat vhatte$ export GO111MODULE="on"
Veerabhushans-MacBook-Pro:chat vhatte$ go mod init
$GOPATH/go.mod exists but should not
Veerabhushans-MacBook-Pro:chat vhatte$ export | grep GO
declare -x GO111MODULE="on"
declare -x GOPATH="/Users/vhatte/go-lang/go-grpc"
declare -x GOROOT=""
Veerabhushans-MacBook-Pro:chat vhatte$ export GOPATH
Veerabhushans-MacBook-Pro:chat vhatte$ export | grep GO
declare -x GO111MODULE="on"
declare -x GOPATH="/Users/vhatte/go-lang/go-grpc"
declare -x GOROOT=""
Veerabhushans-MacBook-Pro:chat vhatte$ export GOPATH=
Veerabhushans-MacBook-Pro:chat vhatte$ export | grep GO
declare -x GO111MODULE="on"
declare -x GOPATH=""
declare -x GOROOT=""
Veerabhushans-MacBook-Pro:chat vhatte$ go mod init
go: cannot determine module path for source directory /Users/vhatte/go-lang/go-grpc/chat (outside GOPATH, module path must be specified)

Example usage:
	'go mod init example.com/m' to initialize a v0 or v1 module
	'go mod init example.com/m/v2' to initialize a v2 module

Run 'go help mod init' for more information.
Veerabhushans-MacBook-Pro:chat vhatte$ go mod init
go: cannot determine module path for source directory /Users/vhatte/go-lang/go-grpc/chat (outside GOPATH, module path must be specified)

Example usage:
	'go mod init example.com/m' to initialize a v0 or v1 module
	'go mod init example.com/m/v2' to initialize a v2 module

syntax = "proto3";
Run 'go help mod init' for more information.
Veerabhushans-MacBook-Pro:chat vhatte$ pwd
/Users/vhatte/go-lang/go-grpc/chat
syntax = "proto3";
Veerabhushans-MacBook-Pro:chat vhatte$ go mod init go-grpc/chat
go: creating new go.mod: module go-grpc/chat
go: to add module requirements and sums:
	go mod tidy
Veerabhushans-MacBook-Pro:chat vhatte$ go mod tidy
go: finding module for package google.golang.org/protobuf/runtime/protoimpl
go: finding module for package google.golang.org/protobuf/reflect/protoreflect
go: finding module for package google.golang.org/grpc/codes
go: finding module for package google.golang.org/grpc
syntax = "proto3";
go: finding module for package google.golang.org/grpc/status
go: found google.golang.org/grpc in google.golang.org/grpc v1.48.0
go: found google.golang.org/grpc/codes in google.golang.org/grpc v1.48.0
go: found google.golang.org/grpc/status in google.golang.org/grpc v1.48.0
go: found google.golang.org/protobuf/reflect/protoreflect in google.golang.org/protobuf v1.28.0
go: found google.golang.org/protobuf/runtime/protoimpl in google.golang.org/protobuf v1.28.0
Veerabhushans-MacBook-Pro:chat vhatte$ ls
syntax = "proto3";
chat.pb.go	go.mod		go.sum
Veerabhushans-MacBook-Pro:chat vhatte$ rm chat.pb.go
Veerabhushans-MacBook-Pro:chat vhatte$ cd ..
Veerabhushans-MacBook-Pro:go-grpc vhatte$ protoc --go_out=plugins=grpc:chat chat.prot
Could not make proto path relative: chat.prot: No such file or directory
Veerabhushans-MacBook-Pro:go-grpc vhatte$ protoc --go_out=plugins=grpc:chat chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat
chat/       chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat/
chat.pb.go  go.mod      go.sum
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat/chat.pb.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ protoc --go_out=plugins=grpc:chat chat.proto
chat.proto:4:21: Expected option value.
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ protoc --go_out=plugins=grpc:chat chat.proto
protoc-gen-go: invalid Go import path "chat" for "chat.proto"

The import path must contain at least one period ('.') or forward slash ('/') character.

See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

--go_out: protoc-gen-go: Plugin failed with status code 1.
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ protoc --go_out=plugins=grpc:chat chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat
syntax = "proto3";
chat/       chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat/
chat/       chat.pb.go  go.mod      go.sum
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat/chat.pb.go
// Code generated by protoc-gen-go. DO NOT EDIT.
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ protoc --go_out=plugins=grpc:chat chat.proto
protoc-gen-go: invalid Go import path "chat" for "chat.proto"

The import path must contain at least one period ('.') or forward slash ('/') character.

See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

--go_out: protoc-gen-go: Plugin failed with status code 1.
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
  1 package main
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat.proto

[No write since last change]
/Users/vhatte/go-lang/go-grpc

Press ENTER or type command to continue
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ protoc --go_out=plugins=grpc:chat chat.proto
protoc-gen-go: invalid Go import path "chat" for "chat.proto"

The import path must contain at least one period ('.') or forward slash ('/') character.

See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

--go_out: protoc-gen-go: Plugin failed with status code 1.
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ protoc --go_out=plugins=grpc:chat chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat/
chat/       chat.pb.go  go.mod      go.sum
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat/chat.pb.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ build client.go
-bash: build: command not found
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: package chat is not in GOROOT (/usr/local/go/src/chat)
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go mod init
go: /Users/vhatte/go-lang/go-grpc/go.mod already exists
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go mod tidy
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: package chat is not in GOROOT (/usr/local/go/src/chat)
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi client.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: no required module provides package github.com/tutorialedge/go-grpc-tutorial/chat; to add it:
	go get github.com/tutorialedge/go-grpc-tutorial/chat
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go mod init tutorialedge/go-grpc-tutorial/chat
go: /Users/vhatte/go-lang/go-grpc/go.mod already exists
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: no required module provides package github.com/tutorialedge/go-grpc-tutorial/chat; to add it:
	go get github.com/tutorialedge/go-grpc-tutorial/chat
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export
declare -x COLORFGBG="7;0"
declare -x COLORTERM="truecolor"
declare -x DISPLAY="/private/tmp/com.apple.launchd.a0wnPkAjMQ/org.macosforge.xquartz:0"
declare -x GO111MODULE="on"
declare -x GOPATH=""
declare -x GOROOT=""
declare -x HOME="/Users/vhatte"
declare -x ITERM_PROFILE="local"
declare -x ITERM_SESSION_ID="w0t1p0:2850BF85-208E-4592-94E5-6B6ABDD44C44"
package main
declare -x LANG="en_US.UTF-8"
declare -x LOGNAME="vhatte"
declare -x LaunchInstanceID="7688B80F-11DC-40BE-85DA-A1FCAC914A8E"
declare -x OLDPWD="/Users/vhatte/go-lang/go-grpc/chat"
declare -x PATH="/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/usr/local/share/dotnet:/opt/X11/bin:~/.dotnet/tools:/Library/Frameworks/Mono.framework/Versions/Current/Commands:/Applications/Wireshark.app/Contents/MacOS:/Users/vhatte/go:/Users/vhatte/go/bin/::/Users/vhatte/go/go::/Users/vhatte/go/go/bin/:/Users/vhatte/go:/Users/vhatte/go/go:/Users/vhatte/go:/Users/vhatte/go/go/bin/:/Users/vhatte/go:/Users/vhatte/go/go:/Users/vhatte/go:/Users/vhatte/go/go/bin/:/Users/vhatte/go:/Users/vhatte/go/bin/go:/Users/vhatte/go:/Users/vhatte/go/bin/go/bin/:/Users/vhatte/go:/Users/vhatte/go/bin/:/Users/vhatte/go:/Users/vhatte/go/bin//bin/"
declare -x PWD="/Users/vhatte/go-lang/go-grpc"
declare -x SECURITYSESSIONID="186a6"
declare -x SHELL="/bin/bash"
declare -x SHLVL="1"
declare -x SSH_AUTH_SOCK="/private/tmp/com.apple.launchd.EAupGqqZNT/Listeners"
declare -x TERM="xterm-256color"
syntax = "proto3";
declare -x TERM_PROGRAM="iTerm.app"
declare -x TERM_PROGRAM_VERSION="3.2.8"
declare -x TERM_SESSION_ID="w0t1p0:2850BF85-208E-4592-94E5-6B6ABDD44C44"
declare -x TMPDIR="/var/folders/80/1v6c4ksd3ns1xbfx64lp110m0000gn/T/"
declare -x USER="vhatte"
declare -x XPC_FLAGS="0x0"
declare -x XPC_SERVICE_NAME="0"
declare -x __CF_USER_TEXT_ENCODING="0x0:0:0"
Veerabhushans-MacBook-Pro:go-grpc vhatte$ export | grep GO
declare -x GO111MODULE="on"
declare -x GOPATH=""
declare -x GOROOT=""
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi client.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
client.go:9:5: no required module provides package github.com/vkhatte/go-lang-grpc-chat; to add it:
	go get github.com/vkhatte/go-lang-grpc-chat
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go get github.com/vkhatte/go-lang-grpc-chat
go: downloading github.com/vkhatte/go-lang-grpc-chat v0.0.0-20220724194448-79da4648f80d
go: added github.com/vkhatte/go-lang-grpc-chat v0.0.0-20220724194448-79da4648f80d
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
# command-line-arguments
./client.go:9:5: imported and not used: "github.com/vkhatte/go-lang-grpc-chat" as __
./client.go:20:10: undefined: chat
./client.go:22:16: undefined: chat
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat/chat.pb.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ protoc --go_out=plugins=grpc:chat chat.proto
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat/
chat/       chat.pb.go  github.com/ go.mod      go.sum
Veerabhushans-MacBook-Pro:go-grpc vhatte$ vi chat/chat.pb.go
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
# command-line-arguments
./client.go:9:5: imported and not used: "github.com/vkhatte/go-lang-grpc-chat" as __
./client.go:20:10: undefined: chat
./client.go:22:16: undefined: chat
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ pwd
/Users/vhatte/go-lang/go-grpc
Veerabhushans-MacBook-Pro:go-grpc vhatte$ cd chat
Veerabhushans-MacBook-Pro:chat vhatte$ ls
chat		chat.pb.go	github.com	go.mod		go.sum
Veerabhushans-MacBook-Pro:chat vhatte$ cd chat
chat/       chat.pb.go
Veerabhushans-MacBook-Pro:chat vhatte$ vi chat/chat.pb.go
Veerabhushans-MacBook-Pro:chat vhatte$ go build client.go
no required module provides package client.go; to add it:
	go get client.go
Veerabhushans-MacBook-Pro:chat vhatte$ ls
chat		chat.pb.go	github.com	go.mod		go.sum
Veerabhushans-MacBook-Pro:chat vhatte$ pwd
package main
/Users/vhatte/go-lang/go-grpc/chat
Veerabhushans-MacBook-Pro:chat vhatte$ cd ..
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
# command-line-arguments
./client.go:9:5: imported and not used: "github.com/vkhatte/go-lang-grpc-chat" as __
./client.go:20:10: undefined: chat
./client.go:22:16: undefined: chat
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go get github.com/vkhatte/go-lang-grpc-chat
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
# command-line-arguments
./client.go:9:5: imported and not used: "github.com/vkhatte/go-lang-grpc-chat" as __
./client.go:20:10: undefined: chat
package main
./client.go:22:16: undefined: chat
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go get github.com/vkhatte/go-lang-grpc-chat
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$
Veerabhushans-MacBook-Pro:go-grpc vhatte$ go build client.go
# command-line-arguments
./client.go:9:5: imported and not used: "github.com/vkhatte/go-lang-grpc-chat" as __
./client.go:20:10: undefined: chat
./client.go:22:16: undefined: chat
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
//  protoc-gen-go v1.28.0-devel
//  protoc        v3.21.3
// source: chat.proto

package chat

import (
    context "context"
    grpc "google.golang.org/grpc"
    codes "google.golang.org/grpc/codes"
    status "google.golang.org/grpc/status"
    protoreflect "google.golang.org/protobuf/reflect/protoreflect"
    protoimpl "google.golang.org/protobuf/runtime/protoimpl"
    reflect "reflect"
    sync "sync"
)

const (
    // Verify that this generated code is sufficiently up-to-date.
    _ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
    // Verify that runtime/protoimpl is sufficiently up-to-date.
    _ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
    state         protoimpl.MessageState
    sizeCache     protoimpl.SizeCache
    unknownFields protoimpl.UnknownFields

"chat/chat.pb.go" 230L, 7025C
