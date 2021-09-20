CFLAGS += \
	-O2 \
	-Wall \
	-Wextra \
	-Werror \
	-march=native \
	-fno-omit-frame-pointer \
	-fstack-protector-strong \
	-fstack-clash-protection \
	-g \
	-D_FORTIFY_SOURCE=2

CFLAGS_DEBUG = $(CFLAGS) \
	-fsanitize=address \
	-fsanitize=pointer-compare \
	-fsanitize=pointer-subtract \
	-fsanitize=leak \
	-fsanitize=undefined \
	-fsanitize=float-divide-by-zero \
	-fsanitize=float-cast-overflow

BINNAME = hexdiff
BINNAME_DEBUG = $(BINNAME).debug

.PHONY: all release install clean debug

all: debug release

release: $(BINNAME_DEBUG)

debug: $(BINNAME)

install: build
	install $(BINNAME) /usr/local/bin

clean:
	rm -f $(BINNAME) $(BINNAME_DEBUG)

$(BINNAME): hexdiff.c
	$(CC) $(CFLAGS) -o $(BINNAME) hexdiff.c

$(BINNAME_DEBUG): hexdiff.c
	$(CC) $(CFLAGS_DEBUG) -o $(BINNAME_DEBUG) hexdiff.c
