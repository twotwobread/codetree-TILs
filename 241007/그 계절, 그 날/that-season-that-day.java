import java.util.Scanner;
public class Main {
    public static void main(String[] args) {
        // 여기에 코드를 작성해주세요.
        Scanner sc = new Scanner(System.in);
        int y = sc.nextInt();
        int m = sc.nextInt();
        int d = sc.nextInt();
        String season = getSeason(y, m, d);
        System.out.println(season);

    }
    static String getSeason(int y, int m, int d){
        if(y < 1 || y > 3000) return "-1";
        if(m < 1 || m > 12 ) return "-1";
        if(d < 1 || d > getLastDay(y, m)) return "-1";

        if     (m >= 3 && m <= 5)   return "Spring";
        else if(m >= 6 && m <= 8)   return "Summer";
        else if(m >= 9 && m <= 11)  return "Fall";
        else                        return "Winter";
    }
    static int getLastDay(int y, int m){
        if(m == 2){
            if(isLeapYear(y))   return 29;
            else                return 28;
        }
        if(m == 4 || m == 6 || m == 9 || m == 11)
            return 30;
        return 31;

    }
    static boolean isLeapYear(int y){
        if(y % 400 == 0) return true;
        if(y % 100 == 0) return false;
        return y % 4 == 0;
    }
}