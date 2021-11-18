#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    vector<int> days = {31,28,31,30,31,30,31,31,30,31,30,31};
    int T;
    cin >> T;
    string input;

    for(int i=1; i<=T; i++) {
        cin >> input;
        int month = atoi(input.substr(4,2).c_str());
        int day = atoi(input.substr(6).c_str());

        cout << "#" << i << " ";

        if (1<=month && month<=12 && 1<=day && day<=days[month-1]) {
            cout << input.substr(0,4) << "/" << input.substr(4,2) << "/" << input.substr(6) << endl;
        } else {
            cout << -1 << endl;
        }
    }
}