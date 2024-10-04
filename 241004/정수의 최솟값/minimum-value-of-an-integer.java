import java.util.*;
public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int min = Integer.MAX_VALUE;
        while(sc.hasNext()){
            min = Math.min(min, sc.nextInt());
        }
        System.out.println(min);
    }
}