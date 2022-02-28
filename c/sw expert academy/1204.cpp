#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    ios_base::sync_with_stdio(false);
    int t,tc;
    cin>>tc;
    int check[100];
    vector<int> v;
    vector<int> res;
    for (t=1; t<=tc; t++) {
        int tcn;
        cin>>tcn;
        for(int i=0; i<1000; i++) {
            int a;
            cin>>a;
            v.push_back(a);
        }
        for(int i=0; i<1000; i++) {
            check[v[i]]++;
        }
    res[t]=*max_element(v.begin(),v.end());
    cout<<"#"<<tcn<<" "<<res[t]<<endl;
    v.clear();
    }
return 0;
}