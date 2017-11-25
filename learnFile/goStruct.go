package Learn

import("fmt")

/// 申明一个结构
type Teacher struct{

	Name string 
	Age int
	Object string
}

/// 为Teacher绑定一个方法
func (t *Teacher)Teach(){

	fmt.Println(t.Object)
}

/// 外部可用
func Play(){

	fmt.Println("out play")
	play()
}

/// 内部可用
func play(){

	fmt.Println("in play")
}