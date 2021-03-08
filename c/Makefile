CFLAGS = -Wall -Wextra -O3
BINNAME = hexdiff

.PHONY: build install clean

build:
	$(CC) -Wall -Wextra -O3 -o $(BINNAME) main.c

install: build
	strip $(BINNAME)
	install $(BINNAME) /usr/local/bin

clean:
	rm -f $(BINNAME)
