CFLAGS = -Wall -Wextra -O3
BINNAME = hexdiff

.PHONY: build install clean

build: $(BINNAME)

install: build
	strip $(BINNAME)
	install $(BINNAME) /usr/local/bin

clean:
	rm -f $(BINNAME)

$(BINNAME): main.c
	$(CC) -Wall -Wextra -O3 -o $(BINNAME) main.c
