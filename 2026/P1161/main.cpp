#include<iostream>

using namespace std;

const int N = 2e7+5;
bool lamp[N];

int main(){
    int n;
    cin>>n;
    for(int i=0;i<n;i++){
        double x;
        int m;
        cin>>x>>m;
        double op=0;
        for(int i=0;i<m;i++){
            op+=x;
            int ix=static_cast<int>(op);
            lamp[ix]=!lamp[ix];
        }
    }
    for(int i=1;i<N;i++){
        if(lamp[i]==true){
            cout<<i;
            break;
        }
    }
}
