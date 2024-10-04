import java.util.*;
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();
        int cnt = 0;

        for(int i = a; i <= b; i++){
            if(isMagicNumber(i)) cnt++;
        }

        System.out.println(cnt);
    }
    static boolean isMagicNumber(int a){
        if(a % 3 == 0 || isThreeSixNine(a)){
            return true;
        }
        return false;
    }
    static boolean isThreeSixNine(int a){
        int n;
        while(a > 0){
            n = a % 10;
            a /= 10;
            if(n != 0 && n % 3 == 0){
                return true;
            }
        }
        return false;

    }
}