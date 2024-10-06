import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        int b = sc.nextInt();
        int sum = 0;
        for(int i = a; i <= b; i++){
            if(isPrime(i))
                sum += i;
        }
        System.out.println(sum);
    }
    static boolean isPrime(int n){
        if(n == 1 || n == 2) 
            return false; 
        for(int i = 2; i < n; i++){
            if(n % i == 0){
                return false;
            }
        }
        return true;
    }
}