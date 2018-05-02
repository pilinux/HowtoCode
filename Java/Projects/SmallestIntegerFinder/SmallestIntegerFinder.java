package pilinux;

public class SmallestIntegerFinder {
    public static int findSmallestInt(int[] args) {

        int smallestNumber = 0;

        for(int i = 0; i < args.length-1; i++) {
            if(args[i] < args[i+1]) {
                smallestNumber = args[i];
                args[i+1] = args[i];
            }
            else
                smallestNumber = args[i+1];
        }

        return smallestNumber;
    }
}
