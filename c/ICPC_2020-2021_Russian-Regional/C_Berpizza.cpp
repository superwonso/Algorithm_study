#include <cstdio>
const long long int max=5*(100000);

int main(){
int q; // number of queries
int m[max]; // customer
int c[max]; // cost 
bool surved[max]; // signal
int result[q];

for (int i=0; i<q; i++) {
    if(m[i]==EOF) break;
    surved[i]=true;
    for (int j=i+1; i<q; j++) {
    if(c[j]<c[j+1]) surved[j+1]=true;
    else if(c[j]==c[j+1]) surved[j]=true; 
    }
}
for(int i=0; i<q; i++){
    for(int j=0; j<q; j++){
        if(surved[i]==true) result[q]++; 
    }
}

}