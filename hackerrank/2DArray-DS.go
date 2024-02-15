func hourglassSum(arr [][]int32) int32 {
    var maxSum int32 = -999999;
    
    for i:=0;i<=3;i++ {
       for j:=0; j<=3;j++{
           sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + 
                            arr[i+1][j+1] +  
                  arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
           
           if (sum > maxSum) {
            maxSum = sum   
           }
       } 
    }
    
    return maxSum;
}

/*
Qual a complexidade de tempo e espaço desse algoritmo?
Espaço O(1) - Não manipulamos nada, é o msm array
Tempo O(4x4) * O(1)
 porem ao pé da letra e um for dentro de um for (nxn)
Tempo O(n^2)
*/

