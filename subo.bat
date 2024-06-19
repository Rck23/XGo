git add .
git commit -m "AWS Prueba1"
git push -u origin main

go build main.go
del main.zip 
tar -a -cf main.zip main
