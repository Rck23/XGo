git add .
git commit -m "AWS Prueba6"
git push -u origin main

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -tags lambda.norpc -o bootstrap main.go
ulimx\go\bin\build-lambda-zip.exe -o main.zip bootstrap
del main.zip 
tar.exe -a -cf main.zip main
