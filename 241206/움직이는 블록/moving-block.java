import java.util.*;

public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.

        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int[] arr = new int[n];
        arr = Arrays.stream(arr)
                .map(i->sc.nextInt())
                .toArray();

        int count = 0;
        int avg = Arrays.stream(arr).sum() / n;
        for (int i = 0; i < n; i++) {
            if(avg < arr[i]){
                count += (arr[i] - avg);
                arr[i] = avg;
            }
        }
        System.out.println(count);
    }
}
