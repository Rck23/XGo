git add .
git commit -m "AWS Prueba7"
git push -u origin main

set GOOS=linux

set GOARCH=amd64

go build -o bootstrap

del main.zip 
tar.exe -a -cf main.zip main
