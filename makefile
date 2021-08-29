compile:
	echo "Compiling for every Arch"
	rm -rf bin/*
	GOOS=linux GOARCH=386 go build -o bin/linux-386/twist main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/twist main.go
	GOOS=linux GOARCH=arm go build -o bin/linux-arm/twist main.go
	GOOS=linux GOARCH=arm64 go build -o bin/linux-arm64/twist main.go
	GOOS=linux GOARCH=ppc64 go build -o bin/linux-ppc64/twist main.go
	GOOS=linux GOARCH=ppc64le go build -o bin/linux-ppc64le/twist main.go
	GOOS=linux GOARCH=mips go build -o bin/linux-mips/twist main.go
	GOOS=windows GOARCH=386 go build -o bin/windows-386/twist main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64/twist main.go
