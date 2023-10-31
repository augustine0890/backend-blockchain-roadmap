import LinkedListNode from "./LinkedListNode";

class LinkedList {
    constructor() {
        this.head = null;
        this.tail = null;
    }

    prepend(value) {
        const newNode = new LinkedListNode(value, this.head);
        this.head = newNode;
        if (!this.tail) {
            this.tail = newNode;
        }
        return this;
    }

    append(value) {
        const newNode = new LinkedListNode(value);
        if (!this.head) {
            this.head = newNode;
            this.tail = newNode;
            return this;
        }

        const currentTail = this.tail;
        currentTail.next = newNode;
        this.tail = newNode;

        return this;
    }

    delete(value) {
        if (!this.head) {
            return null;
        }

        let deletedNode = null;

        // If the head must be deleted then make next node that is different
        // from the head to be a new head.
        while (this.head && this.head.value === value) {
            deletedNode = this.head;
            this.head = this.head.next;
        }

        let currentNode = this.head;

        if (currentNode !== null) {
            while (currentNode.next) {
                if (currentNode.next.value === value) {
                    // If next node must be deleted then make next node to be a next next one.
                    deletedNode = currentNode.next;
                    currentNode.next = currentNode.next.next;
                } else {
                    currentNode = currentNode.next;
                }
            }
        }

        // Check if tail must be deleted.
        if (this.tail.value === value) {
            this.tail = currentNode;
        }

        return deletedNode;
    }

}