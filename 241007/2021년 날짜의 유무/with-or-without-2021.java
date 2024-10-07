import java.util.Scanner;
public class Main{
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        int m = sc.nextInt();
        int d = sc.nextInt();

        System.out.println(isValidDay(m, d)?"Yes":"No");
    }
    static boolean isValidDay(int m, int d){
        if(m < 1 || m > 12) return false;
        if(d < 0 || d > 31) return false;
        if(m == 2 && d > 28) return false;
        if(d == 28 && m != 2) return false;
        if(d == 31 && m <= 7 && m % 2 != 1) return false;
        if(d == 30 && m > 7 &&  m % 2 != 0) return false;
        return true;
    }
    
}