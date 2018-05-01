package pilinux;

public class Main {

    public static void main(String[] args) {
        // Variables
        int totalTests = 0;
        int passed = 0;
        int failed = 0;

        TenMinWalk tenMinWalk = new TenMinWalk();

        //Test 1
        char[] charArray = new char[]{'e', 'w', 'n', 's', 'n', 'e', 'w', 's', 'e', 'w'};
        if( tenMinWalk.isValid(charArray) == true)
            passed++;
        else
            failed++;
        totalTests++;

        //Test 2
        charArray[9] = 's';
        if( tenMinWalk.isValid(charArray) == false)
            passed++;
        else
            failed++;
        totalTests++;

        //Test 3
        charArray = new char[]{'e', 'w', 'n', 's', 'n', 'e', 'w', 's', 'e', 's', 'w'};
        if( tenMinWalk.isValid(charArray) == false)
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
