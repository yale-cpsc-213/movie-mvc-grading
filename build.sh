if [ -z "$1" ]
  then
    echo "You must supply a version"
    exit 1
fi


GOOS=windows GOARCH=386 go build -o ./movie-mvc-grading-windows-$1.exe
mv ./movie-mvc-grading-windows-$1.exe releases

GOOS=linux GOARCH=386 go build -o ./movie-mvc-grading-linux-$1
mv ./movie-mvc-grading-linux-$1 releases

go build -o movie-mvc-grading-mac-$1
mv ./movie-mvc-grading-mac-$1 releases
