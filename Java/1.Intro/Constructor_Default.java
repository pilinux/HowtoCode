/*
* Default Constructor
* */

//Java program
public class Main {

    //main method
    public static void main(String[] args) {
        //calling a constructor
        Food food = new Food();

        food.displayOnTerminal();
    }

}

class Food {
    //constructor
    int ID;
    String name;
    
    //Code block A - user-defined constructor
    /*
    Food() {
        ID = 1;
        name = "Mango";
    }
    */

    void displayOnTerminal() {
        System.out.println("Food ID is " + ID + " and name is " + name);
    }
}


/*
* Output:
* >>> Food ID is 0 and name is null
* Note: Here compiler adds a default constructor Food() {}
* */

/*
* Output (when code block A is included):
* >>> Food ID is 1 and name is Mango
* Note: Here user-defined constructor Food() { ID = 1; name = "Mango"; } is used.
* */
