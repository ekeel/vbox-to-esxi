ci: tools libs build_all

tools:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	go get github.com/mitchellh/gox

libs:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	go get github.com/hashicorp/packer && \
	go get github.com/mitchellh/packer

build_all:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

build:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="$(os_arch)" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

deploy:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="$(os_arch)" -output="$(output)/packer-post-processor-vbox-to-esxi" vbox-to-esxi

darwinamd64:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="darwin/amd64" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

linux_386:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="linux/386" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

darwin_386:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="darwin/386" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

windows_amd64:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="windows/amd64" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

freebsd_amd64:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="freebsd/amd64" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

linux_arm:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="linux/arm" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

netbsd_386:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="netbsd/386" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

freebsd_arm:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="freebsd/arm" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

netbsd_arm:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="netbsd/arm" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

linux_amd64:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="linux/amd64" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

freebsd_386:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="freebsd/386" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

openbsd_386:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="openbsd/386" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

linux_mips64:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="linux/mips64" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi


linux_mipsle:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="linux/mipsle" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

linux_s390x:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="linux/s390x" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

windows_386:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="windows/386" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

openbsd_amd64:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="openbsd/amd64" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

netbsd_amd64:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="netbsd/amd64" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

linux_mips:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="linux/mips" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi

linux_mips64le:
	@export GOPATH=$$(pwd) && \
	export PATH=$$PATH:$$(pwd)/bin && \
	gox -osarch="linux/mips64le" -output="bin/$(version)/packer-post-processor-vbox-to-esxi_{{.OS}}_{{.Arch}}_$(version)" vbox-to-esxi
