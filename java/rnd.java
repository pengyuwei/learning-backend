import java.util.Random;

public class RanTest {

    public static void main(String[] args) {

        for(int i=190901;i<190931;i++){
            java.util.Random r = new java.util.Random(i);
            int rnd = r.nextInt(100);
            System.out.printf("%d:%d %s\n", i, rnd, rnd<=19?"T":"F");
        }

    }
}
