package pilinux;

public class Main {

    public static void main(String[] args) {
        // Variables
        int totalTests = 0;
        int passed = 0;
        int failed = 0;

        VowelCount vowelCount = new VowelCount();

        //Test 1
        char[] chars = new char[]{'A', 'i', 'B', 'z', 'p', 'o', 'O'};
        if( vowelCount.vowelCount(chars) == 4)
            passed++;
        else
            failed++;
        totalTests++;

        //Test 2
        chars = new char[]{'1', 'i', 'Q', 'z', 'p', '0', 'O', 's', 'a', 'y'};
        if( vowelCount.vowelCount(chars) == 3)
            passed++;
        else
            failed++;
        totalTests++;

        //Test 3
        chars = new char[]{'A', 'E', 'O', 'I', 'A', 'o', 'O', 'U'};
        if( vowelCount.vowelCount(chars) == 8)
            passed++;
        else
            failed++;
        totalTests++;

        System.out.println("Total tests: " + totalTests + " | Passed: " + passed + " | Failed: " + failed);
    }
}


/*
* Output:
* Total tests: 3 | Passed: 3 | Failed: 0
* Process finished with exit code 0
*/
