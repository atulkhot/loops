package main

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	sentinel *Cell
}

type Cell struct {
	data string
	next *Cell
}

func make_linked_list() LinkedList {
	cell := Cell{}
	list := LinkedList{sentinel: &cell}
	return list
}

func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	me.next = after
}

func (list *LinkedList) add_range(values []string) {
	last_cell := list.sentinel
	for ; last_cell.next != nil; last_cell = last_cell.next {
	}

	for _, v := range values {
		cell := Cell{data: v}
		last_cell.add_after(&cell)
		last_cell = &cell
	}
}

func (list *LinkedList) to_string(separator string) string {
	sbuf := make([]string, 0)
	for p := list.sentinel.next; p != nil; p = p.next {
		sbuf = append(sbuf, p.data)
	}
	return strings.Join(sbuf[:], ",")
}

func (list *LinkedList) to_string_max(separator string, max int) string {
	sbuf := make([]string, 0)
	for p := list.sentinel.next; p != nil && max > 0; p = p.next {
		sbuf = append(sbuf, p.data)
		max--
	}
	return strings.Join(sbuf[:], ",")
}

func (list *LinkedList) has_loop() bool {
	fast := list.sentinel.next
	if fast != nil {
		fast = fast.next
	}
	slow := list.sentinel.next

	for fast != nil && slow != nil && fast != slow {
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}
		slow = slow.next
	}
	return fast == slow && fast != nil
}

func main() {
	// Make a list from a slice of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := make_linked_list()
	list.add_range(values)

	fmt.Println(list.to_string(" "))
	if list.has_loop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.to_string_max(" ", 10))
	if list.has_loop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	list.sentinel.next.next.next.next.next = list.sentinel.next.next

	fmt.Println(list.to_string_max(" ", 10))
	if list.has_loop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
}
