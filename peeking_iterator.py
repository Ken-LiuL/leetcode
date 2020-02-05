# Below is the interface for Iterator, which is already defined for you.
#
class Iterator:
    def __init__(self, nums):
        """
        Initializes an iterator object to the beginning of a list.
        :type nums: List[int]
        """
        self.nums = nums

    def hasNext(self):
        """
        Returns true if the iteration has more elements.
        :rtype: bool
        """
        return len(self.nums) > 0

    def next(self):
        """
        Returns the next element in the iteration.
        :rtype: int
        """
        tmp =  self.nums[-1]
        self.nums = self.nums[:-1]
        return tmp

class PeekingIterator:
    def __init__(self, iterator):
        """
        Initialize your data structure here.
        :type iterator: Iterator
        """
        self.iterator = iterator
        if self.iterator.hasNext():
            self.top =  self.iterator.next()
        else:
            self.top = None
        

    def peek(self):
        """
        Returns the next element in the iteration without advancing the iterator.
        :rtype: int
        """
        return self.top
        

    def next(self):
        """
        :rtype: int
        """
        cur = self.top
        self.top = self.iterator.hasNext() and self.iterator.next() or None
        return cur
        

    def hasNext(self):
        """
        :rtype: bool
        """
        return self.iterator.hasNext() or self.top  != None
        

# Your PeekingIterator object will be instantiated and called as such:
iter = PeekingIterator(Iterator([1,2,3]))
while iter.hasNext():
    val = iter.peek()   # Get the next element but not advance the iterator.
    iter.next()         # Should return the same value as [val].