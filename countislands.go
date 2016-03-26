package main

func CountIslands(grid [][]int) int {
	// TODO: add your code here!
    var i, j int
    var count int
   count = 0;

  
   for  i = 0; i < len(grid); i++ {
      for j = 0; j < len(grid[0]); j++ {
         if(grid[i][j] == 1){
          count++
          merge(grid,i,j)

         }
      }
     
   }
	return count
}

func merge(a [][]int, i, j int) {
  if(i<0 || j<0 || i>len(a)-1 || j>len(a[0])-1){
        return
 }
 
    if(a[i][j] != 1) {return}
    
    a[i][j] = 0
    
     merge(a, i-1, j)
     merge(a, i+1, j)
     merge(a, i, j-1)
     merge(a, i, j+1)
}
