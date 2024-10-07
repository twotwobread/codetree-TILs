import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.]
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();
        int count = 0;
        for(int i = a; i <= b; i++){
            if(isValid(i)) count++;
        }
        System.out.println(count);
    }
    static boolean isValid(int n){
        if(isPrime(n) && getDigitSum(n) % 2 == 0)
            return true;
        return false;
    }
    static int getDigitSum(int n){
        int sum = 0;
        while(n > 0){
            sum += n % 10;
            n /= 10;
        }
        return sum;
    }
    static boolean isPrime(int n){
        if(n < 1)
            return false;
        if(n == 1)
            return true;
        for(int i = 2; i < n; i++){
            if(n % i == 0)
                return false;
        }
        return true;
    }
}