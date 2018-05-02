package pilinux;

import java.lang.Character;

public class VowelCount {
    public static int vowelCount(char[] args){

        int NoOfVowels = 0;
        char Letter;



        for(int i=0; i<args.length; i++){

            Letter = Character.toLowerCase(args[i]);

            switch (Letter){
                case 'a':
                    NoOfVowels++;
                    break;
                case 'e':
                    NoOfVowels++;
                    break;
                case 'i':
                    NoOfVowels++;
                    break;
                case 'o':
                    NoOfVowels++;
                    break;
                case 'u':
                    NoOfVowels++;
                    break;
            }
        }

        return NoOfVowels;
    }
}
