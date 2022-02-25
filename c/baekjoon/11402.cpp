#include <iostream>

using namespace std;

int Combination(int n,int k, int m);

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);

    long long n, k;
    int m;
    
    cin >> n >> k >> m;

    int result = 1;
    while (n || k){
        result = result*Combination(n%m, k%m, m) %m;
        n /= m, k /= m;
    }
    cout << result << endl;
    return 0;
}

int Combination(int n, int k, int m){
    int a = 1; 
    int b = 1;
    for (int i = n; i > n-k; i--)
        a = a*i %m;
	
    for (int i = k; i >= 2; i--)
        b = b*i %m;

    n = 1, k = m-2;
    while (k){
        if (k & 1) n = n*b %m;
        k >>= 1;
        b = b*b %m;
    }
    return a*n %m;
}