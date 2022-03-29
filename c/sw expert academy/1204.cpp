#include <iostream>
#include <vector>
#include <algorithm>
#include <memory.h>
using namespace std;

int main() {
    ios_base::sync_with_stdio(false);
    int t,tc;
    int check[100];
    check[0]=0;
    cin>>tc;
    for (t=1; t<=tc; t++) {
        int tcn;
        cin>>tcn;
        memset(check, 0, sizeof(check));
        for(int i=0; i<1000; i++) {
            int a;
            cin>>a;
           check[a]++;
        }
        int now = 0, score = 0;
        for (int i = 0; i <= 100; i++) {
            if (now <= check[i]) score = i, now = check[i];
        }
        cout << "#" << tcn << " " << score << '\n';
    }
return 0;
}