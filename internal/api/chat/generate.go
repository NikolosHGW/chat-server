package chat

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../../bin/minimock -i Service -o ./mocks/ -s "_minimock.go"
