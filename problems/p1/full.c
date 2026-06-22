#include<stdio.h>

int main(void){

	for(int i = 0;i < 3;i++){
		for(int ii = 0;ii < 3-i-1;ii++){
			printf(" ");
		}
		for(int ii = 0;ii <= i;ii++){
			printf("*");
		}
		printf("\n");
	}

	return 0;
}
