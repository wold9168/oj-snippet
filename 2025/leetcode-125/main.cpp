#include<bits/stdc++.h>
using namespace std;
class Solution {
public:
    void blank(string s, string & res){
        for(const char to:s){
            if(isdigit(to)||isalpha(to)) res.append(1,to);
        }
    }
    bool equal(char a,char b){
        return (a==b) || ((a^32) == b);
    }
    bool isPalindrome(string s) {
        string res;
        blank(s,res);
        int l=res.length();
        for(int i=0;i<l/2;i++){
            if(equal(res[i],res[l-1-i])) continue;
            return false;
        }
        return true;
    }
};
int main(){
	Solution s;
	s.isPalindrome("race a car");
}
