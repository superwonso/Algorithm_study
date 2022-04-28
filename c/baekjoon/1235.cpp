#include <iostream>
#include <algorithm>
#include <vector>
#include <string>
#include <map>

#define ll long long
using namespace std;
int main(){
    ios::sync_with_stdio(0);
    cin.tie(0);
    ll N, ans=1;
    string arr[1005];
    cin >> N;
    for (int i=0;i<N;i++){
        cin >> arr[i];
        reverse(arr[i].begin(),arr[i].end());
    }
    while (1){
        map<string,ll> m;
        ll chk=1;
        for (int i=0;i<N;i++){
            if (m.find(arr[i].substr(0,ans))!=m.end()){
                chk=0;
                break;
            }
            m[arr[i].substr(0,ans)]=1;
        }
        if (chk)
            break;
        ans++;
    }
    cout << ans;
    return 0;
}