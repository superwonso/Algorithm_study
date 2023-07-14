#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;
bool sieve[4000000] = {1, 1};

int main() {
    int T;
    long long n,k,m;
    m = 142857;
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    cin >> T;

    
    for (int i=0; i<T; i++) {
        
    cin >> n >> k;
    vector<long long> prime;
    
    for(int i=2; i*i<=n; ++i) {
        if (sieve[i]) continue;
        for (int j=i*i; j<=n; j+=i) {
            sieve[j] = 1;
        }
    }

    for(int i=2; i<=n; ++i) {
        if(!sieve[i]) prime.push_back(i);
    }

    vector<long long> cnt(prime.size());

    for(int i=0; i<prime.size(); ++i) {
        for(long long j= prime[i]; j<=n; j*=prime[i]) {
            cnt[i]+= (n/j)-(k/j)-((n-k)/j); 
        }
    }
    
    long long res = 1;
    
    for(int i=0; i<cnt.size(); ++i) {
        for(int j=0; j<cnt[i]; ++j){
        res*=prime[i];
        res%=m;
        if (res==0) {
            cout<<"0"<<endl; 
            return 0;
            }
        }
    }

    cout<<res<<endl;
    free(&prime);
    free(&cnt);

    }
    }
