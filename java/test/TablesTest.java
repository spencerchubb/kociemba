import kociemba.Tables;

public class TablesTest {

    private static void print2D(int[][] arr2D) {
        for (int i = 0; i < arr2D.length; i++) {
            for (int j = 0; j < arr2D[i].length; j++) {
                System.out.print(arr2D[i][j] + " ");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        Tables.init(false);
        print2D(Tables.UDSliceMove);
    }
}
