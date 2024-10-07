import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        int n1 = sc.nextInt();
        int n2 = sc.nextInt();
        int[] arr1 = new int[n1];
        int[] arr2 = new int[n2];

        for(int i = 0; i < n1; i++){
            arr1[i] = sc.nextInt();
        }
        for(int i = 0; i < n2; i++){
            arr2[i] = sc.nextInt();
        }

        for(int i = 0; i <= n1 - n2; i++){
            if(arr1[i] == arr2[0]){
                if(isSubString(arr1, arr2, i, 0 )){
                    System.out.println("Yes");
                    return;
                }
            }
        }

        System.out.println("No");

    }
    static boolean isSubString(int[] arr1, int[] arr2, int idx1, int idx2){
        if(arr1.length - idx1 < idx2)
            return false;
        for(int i = 0; i < arr2.length; i++){
            if(arr1[idx1 + i] != arr2[i])
                return false;
        }
        return true;
    }
}