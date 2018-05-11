package pilinux;

public class Main {

    public static void main(String[] args) {

        HourMinSec hourMinSec = new HourMinSec();

        //Test 1
        String time = hourMinSec.toClockTime(9);
        System.out.println("Test 1:  " + time);

        //Test 2
        time = hourMinSec.toClockTime(59);
        System.out.println("Test 2:  " + time);

        //Test 3
        time = hourMinSec.toClockTime(60);
        System.out.println("Test 3:  " + time);

        //Test 4
        time = hourMinSec.toClockTime(61);
        System.out.println("Test 4:  " + time);

        //Test 5
        time = hourMinSec.toClockTime(600);
        System.out.println("Test 5:  " + time);

        //Test 6
        time = hourMinSec.toClockTime(659);
        System.out.println("Test 6:  " + time);

        //Test 7
        time = hourMinSec.toClockTime(3659);
        System.out.println("Test 7:  " + time);

        //Test 8
        time = hourMinSec.toClockTime(3659, 200);
        System.out.println("Test 8:  " + time);

        //Test 9
        time = hourMinSec.toClockTime(3659, 500, 400);
        System.out.println("Test 9:  " + time);

        //Test 10
        time = hourMinSec.toClockTime(36, 59, 61);
        System.out.println("Test 10: " + time);

        //Test 11
        time = hourMinSec.toClockTime(7659);
        System.out.println("Test 11: " + time);

        //Test 12
        time = hourMinSec.toClockTime(9832, 8746);
        System.out.println("Test 12: " + time);
    }
}


/*
* Output:
* Test 1:  00h:00m:09s
* Test 2:  00h:00m:59s
* Test 3:  00h:01m:00s
* Test 4:  00h:01m:01s
* Test 5:  00h:10m:00s
* Test 6:  00h:10m:59s
* Test 7:  01h:00m:59s
* Test 8:  61h:02m:20s
* Test 9:  3667h:26m:40s
* Test 10: 37h:00m:01s
* Test 11: 02h:07m:39s
* Test 12: 166h:17m:46s
* 
* Process finished with exit code 0
* */
