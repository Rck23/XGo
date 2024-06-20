git add .
git commit -m "AWS Prueba1"
git push -u origin main

set GOOS=linux
set GOARCH=amd64

go build main.go
del main.zip 
zip -r main.zip main
