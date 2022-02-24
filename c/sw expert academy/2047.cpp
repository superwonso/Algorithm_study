#include <vector>
#include <algorithm>
#include <iostream>
#include <string>

using namespace std;

int main() {
    string inp;
        cin>>inp;
    for (int i=0; i<inp.size(); i++){
        inp[i]=toupper(inp[i]);
}
cout<<inp<<endl;
return 0;
}