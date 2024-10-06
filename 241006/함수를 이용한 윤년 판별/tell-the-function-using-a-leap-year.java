import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        boolean b = isLeapYear(sc.nextInt());
        System.out.println(b);
    }
    public static boolean isLeapYear(int y) {
        if(y % 100 == 0 && y % 400 != 0)
            return false;
        if(y % 4 == 0)
            return true;
        return false;

    }
}