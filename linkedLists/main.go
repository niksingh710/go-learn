package main

func main() {
	l := List{}
	l.add(5)
	l.add(6)
	l.add(16)
	l.add(9)
	l.add(6)
	print(l)

	dl := DList{}
	dl.add(5)
	dl.add(6)
	dl.add(7)
	dl.add(8)
	dl.add(9)
	print(dl)
}
