OUT := bin

.PHONY: hello value clean

hello: $(OUT)/hello.exe
	./$(OUT)/hello.exe
	
$(OUT)/hello.exe: hello/main.go
	go build -o $(OUT)/hello.exe hello/main.go
