/*
* Copy values without java constructor:
* - assign one object values to another object
* */

//Java program
public class Main {

    //main method
    public static void main(String[] args) {
        //creating an object f1 by calling the constructor Food
        Food f1 = new Food(1, "Maracuja");
        
        
        Food f2 = new Food();
        Food f3 = new Food();

        f1.displayOnTerminal();
        /*
        * Output:
        * >>> Food ID is 1 and name is Maracuja
        * */

        f2.displayOnTerminal();
        /*
         * Output:
         * >>> Food ID is 0 and name is null
         * */

        //assigning values of f1 to f3
        f3.ID = f1.ID;
        f3.name = f1.name;
        f3.displayOnTerminal();
        /*
         * Output:
         * >>> Food ID is 1 and name is Maracuja
         * */
    }

}

class Food {
    int ID;
    String name;

    //constructor to initialize int and String
    Food(int i, String n) {
        ID = i;
        name = n;
    }

    //empty constructor
    Food() {
    }

    void displayOnTerminal() {
        System.out.println("Food ID is " + ID + " and name is " + name);
    }
}
