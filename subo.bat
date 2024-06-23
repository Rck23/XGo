 git add .
 git commit -m "Login realizado"
 git push -u origin main

@REM set GOOS=linux

@REM set GOARCH=amd64

@REM go build main.go
@REM del main.zip 
@REM tar.exe -a -cf main.zip main

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -tags lambda.norpc -o bootstrap main.go
%USERPROFILE%\go\bin\build-lambda-zip.exe -o main.zip bootstrap
