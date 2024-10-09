import java.util.Scanner;
class WrappedInt{
    public int integer;
    WrappedInt(int n){
        this.integer = n;
    }

}
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        WrappedInt a = new WrappedInt(sc.nextInt());
        WrappedInt b = new WrappedInt(sc.nextInt());
        swap(a, b);

        System.out.println(a.integer +" "+ b.integer);

    }
    static void swap(WrappedInt a, WrappedInt b){
        int temp = a.integer;
        a.integer = b.integer;
        b.integer = temp;
    }
}