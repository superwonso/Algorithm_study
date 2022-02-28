#include <iostream>
using namespace std;

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);
    int t, tc;
    cin >> tc;
    for (t = 1; t <= tc; t++) {
        int h[2],m[2];
        cin >> h[0] >> m[0] >> h[1] >> m[1];
        int res[2] = {0,0};
        res[0]=h[0]+h[1];
        res[1]=m[0]+m[1];
        while(res[1]>=60) {
            res[0]++;
            res[1]-=60;
        }
        while(res[0]>12) {
            res[0]-=12;
        }
        cout << "#" << t << " " << res[0]<<" "<<res[1] << endl;
    }
}