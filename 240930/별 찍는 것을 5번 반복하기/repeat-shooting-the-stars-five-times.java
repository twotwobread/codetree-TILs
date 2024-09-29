public class Main {
    public static void main(String[] args) {
        printStars();
    }
    static void printStars(){
        for(int i = 0; i < 5; i++){
            for(int j = 0; j < 10; j++){
                System.out.print("*");
            }
            System.out.println();
        }
    }
}