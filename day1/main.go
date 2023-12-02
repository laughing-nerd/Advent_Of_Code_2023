package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
    file, err := os.Open("./input");
    if err != nil{
        log.Fatal("Input file could not be opened");        
    }
    lines := bufio.NewScanner(file);
    sum := 0;

    for lines.Scan(){
        line := lines.Text();

        // 2 pointers
        i := 0;
        j := len(line)-1;

        digit_left := 0;
        digit_right := 0;

        for i<=j{
            if line[i]>48 && line[i]<58{
                digit_left = int(line[i])-48;
            }
           if line[j]>48 && line[j]<58{ 
                digit_right = int(line[j])-48;
            }
            if(digit_left>0 && digit_right>0){
                break;
            } else if digit_left>0 && digit_right == 0{
                j -= 1;
            } else if digit_left == 0 && digit_right > 0{
                i += 1;
            } else {
                i +=1;
                j -= 1;
            }
        }
        sum = sum + (digit_left*10 + digit_right);
    }
    fmt.Println(sum)
}
