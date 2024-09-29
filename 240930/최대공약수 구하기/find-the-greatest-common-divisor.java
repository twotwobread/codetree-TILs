import java.io.*;
import java.util.*;
public class Main {
    public static void main(String[] args) throws IOException {
        // 여기에 코드를 작성해주세요.\
        Scanner sc = new Scanner(System.in);

        int n = sc.nextInt();
        int m = sc.nextInt();

        int r = n;
        while(m % r != 0){
            int temp = r;
            r = m % n;
            m = temp;
        }
        System.out.print(r);

        //18 % 12 = 6
        //12 % 6 = 0

    }
}