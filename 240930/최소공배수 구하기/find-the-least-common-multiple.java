import java.io.*;
import java.util.*;
public class Main {
    public static void main(String[] args) throws IOException {
        // 여기에 코드를 작성해주세요.\
        Scanner sc = new Scanner(System.in);

        int n = sc.nextInt();
        int m = sc.nextInt();
        if(n > m){
            int temp = m;
            m = n;
            n = temp;
        }
        int lcm = lcm(n, m);
        System.out.print(lcm);
    }
    static int gcd(int a, int b){
        int r = a;
        while(b % r != 0){
            int temp = r;
            r = b % r;
            b = temp;
        }
        return r;
    }
    static int lcm(int a, int b){
        int gcd = gcd(a, b);
        int lcm = a * b / gcd;
        return lcm;
    }
}