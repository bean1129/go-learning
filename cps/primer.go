/* 我们先是调用GenerateNatural()生成最原始的从2开始的自然数序列。然后开始一个100次迭代的循环，希望生成100个素数。
在每次循环迭代开始的时候，管道中的第一个数必定是素数，我们先读取并打印这个素数。然后基于管道中剩余的数列，
并以当前取出的素数为筛子过滤后面的素数。不同的素数筛子对应的管道是串联在一起的。 */
package main

import (
	"fmt"
)

var flag int32

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(in <-chan int, primer int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in
			if i%primer != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}

/*c语言求素数
#include<stdio.h>
int Prime(int n)      //判断n是不是素数，0代表不是，1代表是
{
    int i;
    for(i=2;i*i<=n;i++)
    {
        if(n%i==0)
            return 0;
    }
    return 1;
}
int main()
{
    int n,count=0;     //count代表已经找到了几个素数
    scanf("%d",&n);
    int i=2;
    while(count<n)
    {
        if(Prime(i)==1)
            count++;
        i++;
    }
    printf("第%d个素数是%d\n",n,i-1);
    return 0;
}*/
