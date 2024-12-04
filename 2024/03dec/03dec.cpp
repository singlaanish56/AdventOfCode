#include<iostream>
#include<string>
#include<sstream>
#include<vector>
#include<regex>
using namespace std;







void part1(string& input){
    std::regex pattern(R"(mul\([0-9]{1,3},[0-9]{1,3}\))");

    vector<string> matches;
    auto it = sregex_iterator(input.begin(), input.end(), pattern);
    auto end = sregex_iterator();
    for(;it!=end;it++){
        auto match = *it;
        auto element = match[1].str();
        matches.push_back(element);
        cout<<element<<"\n";
    }

}

int main(){

    freopen("input2.txt", "r", stdin);

    string x;

    string input;
    while(getline(cin, x)){
        input+=x;
    }

    part1(input);
    //part2(input);
    return 0;
}