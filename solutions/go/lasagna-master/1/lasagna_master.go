package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePreL int) int {
   
    if timePreL==0 {
       timePreL = 2 
    }
    return len(layers)*timePreL
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    countN :=0
    countS := 0
   for _, layer := range layers {
       if layer == "noodles" {
          countN++ 
       } else if layer == "sauce" {
          countS++ 
       } 
   }
    return countN*50, float64(countS)*0.2
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
   myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, points int) []float64 {
    if points ==2 {
        return quantities
    }
    newq := make([]float64, len(quantities))
    for i, val := range quantities {
      newq[i] = float64(val)/2.0*float64(points) 
    }
    return newq
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
