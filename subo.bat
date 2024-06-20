git add .
git commit -m "AWS Prueba2"
git push -u origin main


set GOARCH=amd64 
set GOOS=linux 

go build -o bootstrap main.go
del main.zip 
tar.exe -a -cf main.zip main
