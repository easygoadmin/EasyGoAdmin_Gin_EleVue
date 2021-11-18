/**
 *
 * @author 半城风雨
 * @since 2021/11/16
 * @File : list
 */
package common

type ArrayList struct {
	elements []interface{}
	size     int
}

func New(values ...interface{}) *ArrayList {
	list := &ArrayList{}
	list.elements = make([]interface{}, 10)
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *ArrayList) Add(values ...interface{}) {
	if list.size+len(values) >= len(list.elements)-1 {
		newElements := make([]interface{}, list.size+len(values)+1)
		copy(newElements, list.elements)
		list.elements = newElements
	}

	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}

}

func (list *ArrayList) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		return nil
	}

	curEle := list.elements[index]
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	return curEle
}

func (list *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= list.size {
		return nil
	}
	return list.elements[index]
}

func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

func (list *ArrayList) Size() int {
	return list.size
}
func (list *ArrayList) Contains(value interface{}) bool {
	for _, curValue := range list.elements {
		if curValue == value {
			return true
		}
	}
	return false
}
