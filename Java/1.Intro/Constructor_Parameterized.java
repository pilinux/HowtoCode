/*
* Parameterized Constructor:
* - has a specific number of parameters
* - can have any number of parameters
* */

//Java program
public class Main {

    //main method
    public static void main(String[] args) {
        //calling a constructor
        Food f2 = new Food(2, "Apple");
        Food f3 = new Food(3, "Bread");

        f2.displayOnTerminal();
        f3.displayOnTerminal();
    }

}

class Food {
    //constructor
    int ID;
    String name;

    Food(int i, String n) {
        ID = i;
        name = n;
    }

    void displayOnTerminal() {
        System.out.println("Food ID is " + ID + " and name is " + name);
    }
}


/*
* Output:
* >>> Food ID is 2 and name is Apple
* >>> Food ID is 3 and name is Bread
* */
