package main

func swap(ptr_a, ptr_b *int) {
	ori_a := *ptr_a
	*ptr_a = *ptr_b
	*ptr_b = ori_a
}

func initValue(ptr *int) {
	*ptr = 100
}

func safeAssign(ptr_target *int, ptr_value *int) {
	if ptr_value == nil {
		return
	}
	*ptr_target = *ptr_value
}

func swapDoublePointer(ptr_ptr_a, ptr_ptr_b **int) {
	ori_a := **ptr_ptr_a
	**ptr_ptr_a = **ptr_ptr_b
	**ptr_ptr_b = ori_a
}

func createPtr() *int {
	x := 100
	return &x
}
