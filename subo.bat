git add .
git commit -m "Fix: Token Bearer Authorization"
git push -u origin main

@REM set GOOS=linux

@REM set GOARCH=amd64

@REM go build main.go
@REM del main.zip 
@REM tar.exe -a -cf main.zip main

@REM set GOOS=linux
@REM set GOARCH=amd64
@REM set CGO_ENABLED=0
@REM del main.zip bootstrap
@REM go build -tags lambda.norpc -o bootstrap main.go
@REM %USERPROFILE%\go\bin\build-lambda-zip.exe -o main.zip bootstrap
