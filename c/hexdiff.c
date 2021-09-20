#include <ctype.h>
#include <stdio.h>
#include <string.h>

const char cReset[] = "\033[0m";
const char cRed[] = "\033[31m";

typedef struct {
  const char* data;
  size_t size;
} hexstr_t;

static inline int isCharEq(char l, char r) {
  const char magic = 'A' - 'a';
  if (isupper(l)) l -= magic;
  if (isupper(r)) r -= magic;
  return (l == r);
}

static inline int isHexDigitEq(const char* l, const char* r) {
  return (isCharEq(l[0], r[0]) && isCharEq(l[1], r[1]));
}

static inline void hexstr_process(const hexstr_t self, const hexstr_t other) {
  const size_t cmnSize = (self.size < other.size) ? self.size : other.size;
  for (size_t i = 0; i < cmnSize; i += 2) {
    if (isHexDigitEq(self.data + i, other.data + i)) {
      printf("%c%c", self.data[i], self.data[i + 1]);
    } else {
      printf("%s%c%c%s", cRed, self.data[i], self.data[i + 1], cReset);
    }
  }
  if (self.size > other.size) {
    printf("%s%s%s", cRed, self.data + other.size, cReset);
  }
  printf("\n");
}

int main(int argc, char* argv[]) {
  if (argc < 2) {
    fprintf(stderr, "missing first hex string\n");
    return 1;
  }
  const hexstr_t first = {argv[1], strlen(argv[1])};
  if (argc < 3) {
    fprintf(stderr, "missing second hex string\n");
    return 1;
  }
  const hexstr_t second = {argv[2], strlen(argv[2])};

  if (first.size % 2 != 0) {
    fprintf(stderr, "first arg is not a hex string(len is not even)\n");
    return 1;
  }
  if (second.size % 2 != 0) {
    fprintf(stderr, "second arg is not a hex string(len is not even)\n");
    return 1;
  }

  hexstr_process(first, second);
  hexstr_process(second, first);

  return 0;
}
