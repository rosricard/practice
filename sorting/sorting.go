package sort

func Selection_sort(arr []int) []int {
	// todo:
	//for i in 0 to n-1:  			//c8*n
	// 	minvalue = a[i]   			// c1*n
	//	minindex = i				// c2*n
	// for smol in i+1 to n-1:
	// 	if a[smol] < minvalue    // c3 		// (n(n-1))/2 *c9
	// 		minvalue = a[smol]   // c4
	// 		minindex = smol      // c5
	// swap a[i], a[minindex]    // c6

	for i := 0; i < len(arr); i++ {
		// Find the index of the smallest element in the unsorted part of the array
		min := i
		for inner := i; inner < len(arr); inner++ {
			if arr[inner] < arr[min] {
				min = inner
			}
		}

		// Swap arr[i] and arr[min]
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}

// how do we analyze this algorithm?
// 1. We have to argue that it's CORRECT and solves the computational problem that is given
// 2. Quantify its EFFICIENCY? ie: time (processor)/ space (memory) complexity

// Note you could prove correctness of the above algorithm with induction

// t(n) = c8*n + c1*n + c2*n + (n(n-1))/2 *c9 + c6*n + c7
// t(n) = an^2 + bn + c
// simplify to t(n) = theta(n^2)
