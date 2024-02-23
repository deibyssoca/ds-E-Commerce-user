git add .
git commit -m "Last Commit"
git push
set GOOS=linux
set GOARCH=amd64
@REM go build main.go
-tags lambda.norpc -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip bootstrap