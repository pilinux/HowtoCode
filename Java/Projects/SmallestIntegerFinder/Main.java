package pilinux;

public class Main {

    public static void main(String[] args) {
        // Variables
        int totalTests = 0;
        int passed = 0;
        int failed = 0;

        SmallestIntegerFinder smallestIntegerFinder = new SmallestIntegerFinder();

        //Test 1
        int[] integers = new int[]{1, 4, -215, 314, -315, 279, -179};
        if( smallestIntegerFinder.findSmallestInt(integers) == -315)
            passed++;
        else
            failed++;
        totalTests++;

        //Test 2
        integers = new int[]{500, 400, -215, 3140, 315, -279, -1790, 670, 21478905};
        if( smallestIntegerFinder.findSmallestInt(integers) == -1790)
            passed++;
        else
            failed++;
        totalTests++;

        //Test 3
        integers = new int[]{56476, 65498, 138, 15840, 168461, 3548946, 216846, 35186413};
        if( smallestIntegerFinder.findSmallestInt(integers) == 138)
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
