import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();
        int count = 0;
        for(int i = a; i <= b; i++){
            if(isOnjeonsu(i)) count++;
        }
        System.out.println(count);
    }
    static boolean isOnjeonsu(int n){
        // 아래조건을 만족하지 않아야 한다
        //2로 나누어 떨어지는 경우
        //일의 자리가 5인 경우
        //3으로 나누어 떨어지면서 9로는 나누어 떨어지지 않는 수
        if(n % 2 == 0)
            return false;
        if(n % 10 == 5)
            return false;
        if(n % 3 == 0 && n % 9 != 0)
            return false;

        return true;
    }
}