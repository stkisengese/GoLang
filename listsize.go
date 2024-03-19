package piscine

func ListSize(l *List) int {
  	size := 0
  	newHead := l.Head
  	for newHead != nil {
    	    size++
    		newHead = newHead.Next 
  	}
  	return size
}
