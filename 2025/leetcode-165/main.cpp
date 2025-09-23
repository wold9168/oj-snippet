#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    std::vector<size_t> findAllIndexes(const std::string& str, char ch) {
        std::vector<size_t> indexes;
        size_t pos = 0;

        while ((pos = str.find(ch, pos)) != std::string::npos) {
            indexes.push_back(pos);
            pos++;
        }

        return indexes;
    }
    int fetch(const std::string&str,const std::vector<size_t>& dotlist, int cnt){
        assert(cnt>=0);
        if (dotlist.size()==0) return std::stoi(str);
        else if(cnt ==0) return std::stoi(str.substr(0,dotlist[0]));
        else if(cnt < dotlist.size()) return std::stoi(str.substr(dotlist[cnt-1]+1,dotlist[cnt]-dotlist[cnt-1]-1));
        else return std::stoi(str.substr(dotlist[cnt-1]+1));
    }
    int compareVersion(const string& version1, const string& version2) {
        auto dot1 = findAllIndexes(version1, '.');
        auto dot2 = findAllIndexes(version2, '.');
        if(dot1.size()>dot2.size()) return -1*compareVersion(version2,version1);//grant dot1 < dot2
        int ver1=fetch(version1,dot1,0);
        int ver2=fetch(version2,dot2,0);
        int cnt = 0;
        while(ver1==ver2){
            if(cnt<=dot1.size()){
                ver1=fetch(version1,dot1,cnt);
                ver2=fetch(version2,dot2,cnt);
            }else if(dot1.size()<dot2.size()){
                ver1=0;
                ver2=0;
                for(cnt;cnt<=dot2.size();cnt++){
                    ver2=fetch(version2,dot2,cnt);
                    if(ver1!=ver2) break;
                }
                break;
            }else break;
            cnt++;
        }
        if (ver1==ver2) return 0;
        return (ver1>ver2)?1:-1;

    }
};

int main(){
	Solution s;
	std::cout<<s.compareVersion("1","1.1");
	return 0;
}
