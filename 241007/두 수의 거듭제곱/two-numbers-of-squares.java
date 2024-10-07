import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();
        int r = a;
        for(int i = 0; i < b - 1; i++){
            r *= a;
        }
        System.out.println(r);
    }
}