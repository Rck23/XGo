git add .
git commit -m "AWS Prueba4"
git push -u origin main


set GOOS=linux

set GOARCH=amd64

go build -o bootstrap
go build main.go

del main.zip 
tar.exe -a -cf main.zip main
