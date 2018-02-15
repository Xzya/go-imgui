all:
	c-for-go imgui.yml

clean:
	rm -f imgui/cgo_helpers.go imgui/cgo_helpers.h imgui/cgo_helpers.c
	rm -f imgui/const.go imgui/doc.go imgui/types.go
	rm -f imgui/imgui.go

test:
	cd foobar && go build