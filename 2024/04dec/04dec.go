package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func part2(store []string, m, n int){
	ans:=0

	for i:=0;i<m;i++{
		str := store[i]
		for j:=0;j<n;j++{
			c := str[j]

			if c=='A'{
				mcount:=0
				scount:=0

				if(i-1>=0 && j-1>=0 && i+1<m && j+1<n){
					if(store[i-1][j-1]=='M' && store[i+1][j+1]!='M'){
						mcount++
					}else if(store[i-1][j-1]=='S'  && store[i+1][j+1]!='S'){
						scount++
					}
					if(store[i-1][j+1]=='M'  && store[i+1][j-1]!='M'){
						mcount++
					}else if(store[i-1][j+1]=='S'  && store[i+1][j-1]!='S'){
						scount++
					}
					if(store[i+1][j-1]=='M'  && store[i-1][j+1]!='M'){
						mcount++
					}else if(store[i+1][j-1]=='S'  && store[i-1][j+1]!='S'){
						scount++
					}
					if(store[i+1][j+1]=='M'  && store[i-1][j-1]!='M'){
						mcount++
					}else if(store[i+1][j+1]=='S'  && store[i-1][j-1]!='S'){
						scount++
					}
				
					if mcount==2 && scount==2{
						ans++
					}
				}
			}
		}
	}

	fmt.Println(ans)
}
func part1(store []string , m,n int){

	ans:=0;

	for i:=0;i<m;i++{
		str := store[i]
		for j:=0;j<n;j++{
			c := str[j]

			if c=='X'{
				//up
				if(i-3>=0 && (store[i-1][j]=='M' && store[i-2][j]=='A' && store[i-3][j]=='S')){
					ans++
				}
				//down
				if(i+3<m && (store[i+1][j]=='M' && store[i+2][j]=='A' && store[i+3][j]=='S')){
					ans++;
				}
				//left
				if(j-3>=0 && (store[i][j-1]=='M' && store[i][j-2]=='A' && store[i][j-3]=='S')){
					ans++;
				}
				//right
				if(j+3<n && (store[i][j+1]=='M' && store[i][j+2]=='A' && store[i][j+3]=='S')){
					ans++;
				}
				//up left diagnal
				if(i-3>=0 && j-3>=0 && (store[i-1][j-1]=='M' && store[i-2][j-2]=='A' && store[i-3][j-3]=='S')){
					ans++;
				}				
				//up right diagnal
				if(i-3>=0 && j+3<n && (store[i-1][j+1]=='M' && store[i-2][j+2]=='A' && store[i-3][j+3]=='S')){
					ans++;
				}	
				//bottom left diagnal
				if(i+3<m && j-3>=0 && (store[i+1][j-1]=='M' && store[i+2][j-2]=='A' && store[i+3][j-3]=='S')){
					ans++;
				}	
				//bottom  right diagnal
				if(i+3<m && j+3<n && (store[i+1][j+1]=='M' && store[i+2][j+2]=='A' && store[i+3][j+3]=='S')){
					ans++;
				}	
			}
		}
	}

	fmt.Println(ans)
}

func main() {

	start := time.Now()
	if len(os.Args)<1 {
		fmt.Println("Error while gettin the file")
		return 
	}

	file, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("error while opening the file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var store []string
	for scanner.Scan(){
		str := scanner.Text()

		store=append(store, str)
	}

	m := len(store)
	n := len(store[0])

	part1(store, m, n)
	part2(store, m, n)

	end := time.Now()

	fmt.Println(end.Sub(start))

}