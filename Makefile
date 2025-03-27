OUT := bin

.PHONY: hello value clean

hello: $(OUT)/hello.exe
	./$(OUT)/hello.exe
	
$(OUT)/hello.exe: hello/main.go
	go build -o $(OUT)/hello.exe hello/main.go

value: $(OUT)/value.exe
	./$(OUT)/value.exe
	
$(OUT)/value.exe: value/main.go
	go build -o $(OUT)/value.exe value/main.go