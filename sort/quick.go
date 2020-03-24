//
// Revision History:
//     Initial: 2019-02-19 20:44    Jon Snow

package sort

func quick(arr []int, start, stop int) {
	if start >= stop {
		return
	}

	i, j := start, stop
	for i < j {
		for i < j {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				break
			}
			i++
		}
		for i < j {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				break
			}
			j--
		}
	}
	if start < i {
		quick(arr, start, i-1)
	}
	if j < stop {
		quick(arr, j+1, stop)
	}
}
