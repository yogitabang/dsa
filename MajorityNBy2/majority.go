//Find element which is present more than length of array times


//Naive algo.
//Uses O(n) space
//Uses O(n) time
func majorityElement(nums []int) (int,err) {
    major := len(nums)/2
    var freqmap = make(map[int]int)
    for _,val := range nums {
        freqmap[val] = freqmap[val]+1
    }
    
    for key,value:=range freqmap {
        if value > major {
            return key,nil
        }
    }
	//Assume no majority element present
    return 0,errors.New("No majority element present")
}

//Optimal algo - moore's voting algo
//Uses O(n) time
//Uses O(1) space
func findCandidate(nums []int) (int, err) {
	majIndex := 0
	count := 1
	for _,val := range nums {
		if nums[maj_index] == val {
			count++
		} else {
			count--
		}
		if count == 0 {
			majIndex = i
			count = 1
		}
	}
	
	//nums[majIndex] is the candidate
	count = 0
	for _,val := range nums {
		if val == nums[majIndex]{
			count++
		}
	}
	
	if count > len(nums)/2{
		return nums[majIndex],nil
	}else {
		return 0, errors.New("No majority element present")
	}
}
