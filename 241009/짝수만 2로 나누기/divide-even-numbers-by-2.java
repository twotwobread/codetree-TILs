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
        int n = sc.nextInt();
        int[] arr = new int[n];
        for(int i = 0; i < n; i++){
            arr[i] = sc.nextInt();
        }

        func(arr);
        for(int num : arr){
            System.out.print(num + " ");
        }
        System.out.println();
        

    }
    static void func(int[] arr){
        for(int i = 0; i < arr.length; i++){
            if(arr[i] % 2 == 0){
                arr[i] /= 2;
            }
        }
    }
}