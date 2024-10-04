import java.util.*;
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        System.out.println(isMagicNumber(sc.nextInt())?"Yes":"No");
    }
    static boolean isMagicNumber(int a){
        int sum = 0;
        while(a > 0){
            sum += a % 10;
            a /= 10;
        }
        if(a % 2 == 0 && sum % 5 == 0)
            return true;
        return false;
    }
}