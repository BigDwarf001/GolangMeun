package main

import "fmt"

type data interface{}

type LinkTableNode struct {
	pValue data
	pNext  *LinkTableNode
}

type LinkTable struct {
	pHead  *LinkTableNode
	pTail  *LinkTableNode
	Length int
}

func CreateLinkTable() *LinkTable {
	var pLinkTable *LinkTable = new(LinkTable)
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.Length = 0
	return pLinkTable
}

func IsEmpty(pLinkTable *LinkTable) bool {
	return pLinkTable.Length == 0
}

func DeleteLinkTable(pLinkTable *LinkTable) bool {
	if pLinkTable == nil {
		return false
	}
	for !IsEmpty(pLinkTable) {
		var pTemp *LinkTableNode = pLinkTable.pHead
		pLinkTable.pHead = pLinkTable.pHead.pNext
		pTemp.pNext = nil
		pTemp.pValue = nil
		pLinkTable.Length--
	}
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	return true
}

func ADDLinkTableNode(pLinkTableNode *LinkTableNode, pLinkTable *LinkTable) bool {
	if pLinkTable == nil || pLinkTableNode == nil {
		return false
	}
	if IsEmpty(pLinkTable) {
		pLinkTable.pHead = pLinkTableNode
		pLinkTable.pTail = pLinkTableNode
		pLinkTable.Length++
	} else {
		pLinkTable.pTail.pNext = pLinkTableNode
		pLinkTable.Length++
		pLinkTable.pTail = pLinkTableNode
	}
	return true
}

func DeleteLinkTableNode(pLinkTableNode *LinkTableNode, pLinkTable *LinkTable) bool {
	if pLinkTable == nil || pLinkTableNode == nil {
		return false
	}

	if pLinkTableNode == pLinkTable.pHead {
		if pLinkTable.Length == 1 {
			pLinkTable.pHead = nil
			pLinkTable.pTail = nil
		} else {
			pLinkTable.pHead = pLinkTable.pHead.pNext
		}
		pLinkTable.Length--
		return true
	}

	var p *LinkTableNode = pLinkTable.pHead
	for p.pNext != pLinkTable.pTail {
		if p.pNext == pLinkTableNode {
			p = p.pNext.pNext
			pLinkTable.Length--
			return true
		}
		p = p.pNext
	}

	if p.pNext == pLinkTableNode {
		pLinkTable.pTail = p
		pLinkTable.Length--
		return true
	}
	return false
}

func SelectLinkTableNode(pLinkTableNode *LinkTableNode, pLinkTable *LinkTable) bool {
	if pLinkTable == nil || pLinkTableNode == nil {
		return false
	}
	var p *LinkTableNode = pLinkTable.pHead
	for p != nil {
		if p == pLinkTableNode {
			return true
		}
		p = p.pNext
	}
	return false
}

func main() {
	var pLinkTable *LinkTable = CreateLinkTable()
	var pLinkTableNode1 *LinkTableNode = new(LinkTableNode)
	pLinkTableNode1.pValue = "LinkTableNodeExample"
	ADDLinkTableNode(pLinkTableNode1, pLinkTable)

	var pLinkTableNode2 *LinkTableNode = new(LinkTableNode)
	pLinkTableNode2.pValue = 2
	ADDLinkTableNode(pLinkTableNode2, pLinkTable)

	fmt.Println("Length:", pLinkTable.Length)
	fmt.Println("Head:", pLinkTable.pHead.pValue)
	fmt.Println("Tail:", pLinkTable.pTail.pValue)
}
