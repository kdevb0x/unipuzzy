env := $(shell env)
modoff = GO111MODULE=off
modon = GO111MODULE=on


PKG_CONFIG_PATH := /usr/local/lib64/pkgconfig
LD_LIBRARY_PATH := /usr/local/lib64
CGO_CPPFLAGS = -I/usr/local/include
CGO_CXXFLAGS = --std=c++1z
CGO_LDFLAGS := -L/usr/local/lib -lopencv_core -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d -lopencv_plot -lopencv_tracking

.PHONY: run

run:
	$(env) $(modoff) go run main.go


build:
	$(env) $(modoff) go build -ldflags=$(CGO_LDFLAGS)

dockerbuild:
	go dockerbuildgen.go

dockerruntime:
	go generate dockerruntimegen.go

