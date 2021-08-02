#include <cstdio>
int a, b, c, n;
int main() {
    scanf("%d %d %d", &a, &b, &n);
    while (n--) a %= b, a *= 10, c = a / b;
    printf("%d", c);
    return 0;
}