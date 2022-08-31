import sys

def pow_mod(a,b,m):
    result = 1
    while b > 0:
        if b % 2 != 0:
            result = (result * a) % m
        b //= 2
        a = (a * a) % m
    return result

def millar_rabin(n,a):
    r = 0
    d = n-1
    while (d%2 == 0):
        r += 1
        d = d // 2
    x = pow_mod(a,d,n)
    if x == 1 or x == n-1:
        return True
    for i in range(0,r-1):
        x = pow_mod(x,2,n)
        if x == n-1:
            return True
    return False

def check_prime(k):
    check = 0
    if k <= 71:
        if k in [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71]:
            return True
        else:
            return False
    else:
        for i in [ 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37]:
            if millar_rabin(k,i) == False:
                break
            else:
                check += 1
        if check == 12:
            return True
        else:
            return False

n = int(sys.stdin.readline())
original_n = n

def gcd(a, b):
    if b > a:
        k = a
        a = b
        b = k
    while b != 0:
        r = a % b
        a = b
        b = r
    return a

def g(x,n):
    return ((x*x) + 1) % n


def pollard_roh(n,x):
    p = x
    if check_prime(n):
        return n
    else:
        for i in [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71]:
            if n % i == 0:
                return i
        y = x
        d = 1
        while d == 1:
            x = g(x,n)
            y = g(g(y,n),n)
            d = gcd(abs(x-y),n)
        if d == n:
            return pollard_roh(n,p+1)
        else:
            if check_prime(d):
                return d
            else:
                return pollard_roh(d,2)
ans = []
while n!= 1 :
    k = pollard_roh(n,2)
    ans.append(int(k))
    n//=k

from itertools import combinations

check_sum = original_n
ans = list(set(ans))
for i in range(1,len(ans)+1):
    for t in combinations(ans,i):
        k = 1
        for p in t :
            k*=p
        if i%2 == 1 :
            check_sum -=original_n//k
        else :
            check_sum +=original_n//k

if original_n == 1:
    print(1)
else :
    print(check_sum)
