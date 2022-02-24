#include <vector>
#include <algorithm>
#include <iostream>
#include <string>

using namespace std;

int main() {
    string inp;
        cin>>inp;
    for (int i=0; i<inp.size(); i++){
        printf("%d ",inp[i]-64);
}
return 0;
}