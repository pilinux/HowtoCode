// Generate random number

#include <stdio.h>
#include <stdlib.h> // For rand() function
#include <time.h>

int randomNumber(int min, int max) {
	int num = (rand() % (max - min + 1)) + min;
	return num;
}

int main() {
	// Use current time as a seed for the random generator
	srand(time(0));
	// Random number between 10 and 90
	int number = randomNumber(10, 90);

	printf("Random number: %d\n", number);

	return 0;
}
