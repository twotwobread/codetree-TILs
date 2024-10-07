import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        int a = sc.nextInt();
        char operator = sc.next().charAt(0);
        int b = sc.nextInt();

        int r = 0;
        switch(operator){
            case'+':
                r = plus(a, b);
                break;
            case'-':
                r = minus(a, b);
                break;
            case'*' :
                r = mutilply(a, b);
                break;
            case'/':
                r = divide(a, b);
                break;
            default:
                System.out.println("False");
                return;
        };

        System.out.printf("%d %c %d = %d%n\n", a , operator, b , r);
    }

    static int plus (int a, int b){
        return a + b;
    }

    static int minus (int a, int b){
        return a - b;
    }

    static int mutilply (int a, int b){
        return a * b;
    }

    static int divide (int a, int b){
        return (int)a / b;
    }
}