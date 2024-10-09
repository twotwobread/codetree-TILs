import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        String str = sc.next();
        
        System.out.println(isPalindrome(str)?"Yes":"No");
    }
    static boolean isPalindrome(String str){
        int len = str.length();
        for(int i = 0; i < len/2 ; i++){
            if(str.charAt(i) != str.charAt(len - 1 - i))
                return false;
        }
        return true;

    }
}