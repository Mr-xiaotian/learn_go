OUT := bin
TARGETS := hello value variable constants for if_else switch

.PHONY: all $(TARGETS)

# 默认目标：构建所有
all: $(TARGETS)

# 每个目标构建规则
$(TARGETS): %: $(OUT)/%.exe
	./$<

# 编译规则
$(OUT)/%.exe: %/main.go
	go build -o $@ $<
