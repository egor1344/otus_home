package doublelinkedlist

import (
	"testing"
	"log"
	)

func TestDuobleLinkedItem(t *testing.T) {
	testItem := Item{}
	if result := testItem.GetValue(); (result != nil) {
		t.Error("Wrong result ", result)
	}
	testItem.SetValue(3)
	if result := testItem.GetValue(); (result != 3) {
		t.Error("Wrong result ", result)
	}
	if result := testItem.GetValue(); (result != 3) {
		t.Error("Wrong result ", result)
	}
	i1 := Item{}
	i1.SetValue(1)
	i2 := Item{}
	i2.SetValue(2)
	testItem.next = &i1
	testItem.prev = &i2
	if testItem.next == nil && testItem.prev == nil {		
		t.Error("Not next or prev item")
	}
	// log.Println(i2, testItem, i1)
	testItem.Remove()
	// log.Println(i2, testItem, i1)
}

func TestDuobleLinkedList(t *testing.T) {
	ddlTest := DoubleLinkedList{}
	i1 := Item{}
	i1.SetValue(3)
	ddlTest.PushFront(i1)
	log.Println(ddlTest)
	if ddlTest.first == nil {
		t.Error("Erorr PushFront")
	}
	i2 := Item{}
	i2.SetValue(2)
	ddlTest.PushBack(i2)
	log.Println(ddlTest)
	if  counter:= ddlTest.Len(); (counter != 2) {
		t.Error("Erorr len", counter)
	}
}