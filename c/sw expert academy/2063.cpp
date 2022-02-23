#include <iostream>
#include <cstdio>
#include <malloc.h>
#include <algorithm>
#include <vector>

using namespace std;

int main() {
int N;
cin >> N;
vector<int> inp;
for(int i=0; i<N; i++) {
    int temp;
    cin >> temp;
    inp.push_back(temp);
}
sort(inp.begin(), inp.end());
printf("%d",inp[((inp.size())/2)]);
}