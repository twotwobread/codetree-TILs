import java.io.*;
import java.util.*;
public class Main {
    public static void main(String[] args) throws IOException {
        // 여기에 코드를 작성해주세요.
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
        StringTokenizer st = new StringTokenizer(br.readLine());

        int n = Integer.parseInt(st.nextToken());
        int m = Integer.parseInt(st.nextToken());

        int r = n;
        while(m % r != 0){
            int temp = r;
            r = m % n;
            m = temp;
        }
        bw.write(r + "\n");
        bw.flush();
        bw.close();
        br.close();
        //18 % 12 = 6
        //12 % 6 = 0

    }
}