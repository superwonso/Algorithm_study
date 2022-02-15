#include <iostream>

using namespace std;

long long mod = 1e9 + 7;
long long n, m, k;
long long ans[4000010];

long long mod_func(long long a, long long b) {
    if (!b) return 1;
    if (b & 1) return (a*mod_func(a, b - 1) % mod) % mod;
    return mod_func((a*a) % mod, b / 2) % mod;
}

int main() {
    scanf(" %lld", &m);
    ans[0] = 1;

    for (long long i = 1; i <= 4000000; i++) {
        ans[i] = i*ans[i - 1];
        ans[i] %= mod;
    }

    while (m--) {
        long long a, b, c;
        scanf(" %lld %lld", &n, &k);
        a = ans[n];
        b = ans[k] * ans[n - k];
        b %= mod;
        c = mod_func(b, mod - 2);
        a *= c;
        a %= mod;
        printf("%lld\n", a);
    }
}

