package pilinux;

public class HourMinSec {

    public static final String INVALID_MESSAGE = "Invalid value!"; //define a constant with 'final' keyword

    //Method
    public static String toClockTime(long seconds) {
        if(seconds < 0)
            return INVALID_MESSAGE;

        long hours = 0;
        long minutes = 0;


        if (seconds >= 60) {
            minutes = seconds / 60;
            seconds = seconds % 60;
        }

        if (minutes >= 60) {
            hours = minutes / 60;
            minutes = minutes % 60;
        }



        if (hours < 10 && minutes < 10 && seconds < 10)
            return "0" + hours + "h:0" + minutes + "m:0" + seconds + "s";

        if (hours < 10 && minutes < 10 && seconds >= 10)
            return "0" + hours + "h:0" + minutes + "m:" + seconds + "s";

        if (hours < 10 && minutes >= 10 && seconds < 10)
            return "0" + hours + "h:" + minutes + "m:0" + seconds + "s";

        if (hours < 10 && minutes >= 10 && seconds >= 10)
            return "0" + hours + "h:" + minutes + "m:" + seconds + "s";

        if (hours >= 10 && minutes < 10 && seconds < 10)
            return hours + "h:0" + minutes + "m:0" + seconds + "s";

        if (hours >= 10 && minutes < 10 && seconds >= 10)
            return hours + "h:0" + minutes + "m:" + seconds + "s";

        if (hours >= 10 && minutes >= 10 && seconds < 10)
            return hours + "h:" + minutes + "m:0" + seconds + "s";

        else
            return hours + "h:" + minutes + "m:" + seconds + "s";
    }

    //Method Overloading
    public static String toClockTime(long minutes, long seconds) {
        seconds = seconds + minutes * 60;
        return toClockTime(seconds);
    }

    //Method Overloading
    public static String toClockTime(long hours, long minutes, long seconds) {
        seconds = seconds + minutes * 60 + hours * 60 * 60;
        return toClockTime(seconds);
    }
}
