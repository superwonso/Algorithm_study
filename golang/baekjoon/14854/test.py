import sys
input = sys.stdin.readline


def precompute(p, e):
    pe = p**e
    val = [1]
    for i in range(1, pe):
        if i % p:
            val.append(val[-1]*i)
        else:
            val.append(val[-1])
    return val


def modpow(c, n, d):
    r = 1
    while n:
        if n & 1:
            r = r*c % d
        c = c*c % d
        n >>= 1
    return r


def fact_n_pe(n, p, e):
    val = precompute(p, e)
    pe = p**e

    def do(n):
        if n < p:
            return (0, val[n])
        t = n//p
        nt, nm = do(t)
        tp = t + nt
        kp, rp = divmod(n, pe)
        m = nm * modpow(val[pe-1], kp, pe) % pe * val[rp] % pe
        return (tp, m)
    return do(n)


def comb(n, k, p):
    if not k:
        return 1
    if not n:
        return 0
    r, s, t = 1, 1, 1
    for i in range(2, n+1):
        r = r*i % p
        if i == k:
            s = modpow(r, p-2, p)
        if i == n-k:
            t = modpow(r, p-2, p)
    return r*s % p*t % p


def lucas(n, k, p):
    ni = []
    ki = []
    while n:
        ni.append(n % p)
        n //= p
    while k:
        ki.append(k % p)
        k //= p
    while len(ni) > len(ki):
        ki.append(0)

    ret = 1
    for i in range(len(ni)):
        ret = ret*comb(ni[i], ki[i], p) % p
    return ret


def comb_pe(n, k, p, e):
    if e == 1:
        return lucas(n, k, p)
    nt, nm = fact_n_pe(n, p, e)
    kt, km = fact_n_pe(k, p, e)
    nkt, nkm = fact_n_pe(n-k, p, e)
    pe = p**e
    t = nt - kt - nkt
    if t >= e:
        return 0

    for i in range(pe):
        if nkm * km * i % pe == 1:
            break

    return (p**t)*nm*i % pe


D = 142857
T = int(input())

result = []

for _ in range(T):
    n, r = map(int, input().split())
    a = comb_pe(n, r, 3, 3)
    b = comb_pe(n, r, 11, 1)
    c = comb_pe(n, r, 13, 1)
    d = comb_pe(n, r, 37, 1)
    result.append((a*137566 % D + b*103896 % D + c*109890 % D + d*77220 % D) % D)

for j in range (T):
    sys.stdout.write(str(result[j]) + '\n')
