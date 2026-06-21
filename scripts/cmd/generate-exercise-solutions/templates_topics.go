package main

import (
	"fmt"
	"strings"
)

func fundamentalsSpec(m module) topicSpec {
	var body, test, bench, interviewSol, interviewTest, problem string
	ep := errPrefix(m.slug)

	switch m.slug {
	case "02-types-system":
		problem = "Implement TypeName(v any) string returning the dynamic type name"
		body = `package solutions

import "fmt"

// TypeName returns the dynamic type name of v.
func TypeName(v any) string {
	if v == nil {
		return "nil"
	}
	return fmt.Sprintf("%T", v)
}

// IsNumeric reports whether v is an integer or float type.
func IsNumeric(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}`
		test = testPackage(m, "TypeName(42)", `"int"`, "IsNumeric(42)", "true")
		bench = "TypeName(42)"
		interviewSol = `package solutions

// InterviewChallengeSolution returns zero value for int.
func InterviewChallengeSolution() int {
	var z int
	return z
}`
		interviewTest = testInterviewInt(m, "InterviewChallengeSolution()", "0")

	case "03-control-flow":
		problem = "FizzBuzz(n) returning slice of strings 1..n with Fizz/Buzz rules"
		body = fizzBuzzBody()
		test = testFizzBuzz(m)
		bench = `FizzBuzz(100)`

	case "04-functions":
		problem = "Compose two functions: Compose(f,g)(x) = f(g(x))"
		body = `package solutions

// Compose chains g then f.
func Compose(f, g func(int) int) func(int) int {
	return func(x int) int { return f(g(x)) }
}

// MapInts applies fn to each element.
func MapInts(xs []int, fn func(int) int) []int {
	out := make([]int, len(xs))
	for i, v := range xs {
		out[i] = fn(v)
	}
	return out
}`
		test = testPackage(m, "Compose(func(x int) int { return x * 2 }, func(x int) int { return x + 1 })(3)", "8", "len(MapInts([]int{1,2}, func(x int) int { return x * 2 }))", "2")
		bench = `Compose(func(x int) int { return x }, func(x int) int { return x })(42)`

	case "05-packages-modules":
		problem = "ParseSemVer parses major.minor.patch from string"
		body = semverBody()
		test = testSemver(m)
		bench = `ParseSemVer("1.2.3")`

	case "06-pointers":
		problem = "Swap swaps two integers via pointers"
		body = `package solutions

// Swap exchanges values at a and b.
func Swap(a, b *int) {
	*a, *b = *b, *a
}

// IntPtr returns a pointer to v.
func IntPtr(v int) *int { return &v }`
		test = testSwap(m)
		bench = `a,b:=1,2; Swap(&a,&b)`

	case "07-structs":
		problem = "Rectangle with Area() and Perimeter() methods"
		body = rectangleBody()
		test = testRectangle(m)
		bench = `Rectangle{3,4}.Area()`

	case "08-methods":
		problem = "BankAccount with Deposit, Withdraw, Balance (value vs pointer receivers)"
		body = bankAccountBody()
		test = testBankAccount(m)
		bench = `a:=BankAccount{}; a.Deposit(10)`

	case "09-interfaces":
		problem = "Implement Stringer for Person struct"
		body = personStringerBody()
		test = testPerson(m)
		bench = `Person{"A","B"}.String()`

	case "10-generics":
		problem = "Generic Min function for ordered types"
		body = genericsMaxBody()
		test = testGenericsMax(m)
		bench = `MaxInt(3, 7)`
		interviewSol = genericsMinBody()
		interviewTest = testGenericsMin(m)

	case "11-error-handling":
		problem = "Divide with wrapped errors using fmt.Errorf %w"
		body = divideBody()
		test = testDivide(m)
		bench = `Divide(10, 2)`

	case "12-memory-management":
		problem = "Detect if slice reuses same backing array after append"
		body = `package solutions

// SharesBacking reports whether a and b share the same backing array.
func SharesBacking(a, b []int) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	return &a[0] == &b[0]
}`
		test = testSharesBacking(m)
		bench = `SharesBacking([]int{1}, []int{1,2})`

	default:
		body = fmt.Sprintf(`package solutions

// Greet returns a greeting for %s.
func Greet(name string) string {
	if name == "" {
		return "Hello, Go!"
	}
	return "Hello, " + name + "!"
}`, m.title)
		test = testPackage(m, `Greet("Go")`, `"Hello, Go!"`, `Greet("")`, `"Hello, Go!"`)
		bench = `Greet("benchmark")`
		problem = "Implement variable shadowing detector or greeting function"
	}

	if interviewSol == "" {
		interviewSol = genericInterviewSolution(m)
	}
	if interviewTest == "" {
		interviewTest = genericInterviewTest(m)
	}

	_ = ep
	return topicSpec{
		summary:           fmt.Sprintf("Hands-on exercises for **%s** — variables, types, idioms, and Go fundamentals.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  problem,
		interviewSolution: interviewSol,
		interviewTest:     interviewTest,
		benchCall:         bench,
	}
}

func dataStructureSpec(m module) topicSpec {
	slug := m.slug
	switch slug {
	case "13-arrays":
		return dsSpec(m, arrayBody(), testArray(m), "FindDuplicates in fixed array", arrayInterview(), testArrayInterview(m))
	case "14-slices":
		return dsSpec(m, sliceBody(), testSlice(m), "Reverse slice in-place O(1) space", sliceReverseBody(), testSliceReverse(m))
	case "15-maps":
		return dsSpec(m, mapBody(), testMap(m), "Two-sum using hash map O(n)", twoSumBody(), testTwoSum(m))
	case "16-strings-runes-bytes":
		return dsSpec(m, runeBody(), testRune(m), "IsPalindrome ignoring case", palindromeBody(), testPalindrome(m))
	case "17-recursion":
		return dsSpec(m, factorialBody(), testFactorial(m), "Fibonacci with memoization", fibBody(), testFib(m))
	case "18-linked-lists":
		return dsSpec(m, linkedListBody(), testLinkedList(m), "Reverse singly linked list", reverseListBody(), testReverseList(m))
	case "19-stacks":
		return dsSpec(m, stackBody(), testStack(m), "Valid parentheses using stack", validParensBody(), testValidParens(m))
	case "20-queues":
		return dsSpec(m, queueBody(), testQueue(m), "BFS level order using queue", bfsBody(), testBFS(m))
	case "21-trees":
		return dsSpec(m, treeBody(), testTree(m), "Max depth of binary tree", maxDepthBody(), testMaxDepth(m))
	case "22-binary-search-tree":
		return dsSpec(m, bstBody(), testBST(m), "Validate BST", validateBSTBody(), testValidateBST(m))
	case "23-heaps":
		return dsSpec(m, heapBody(), testHeap(m), "Kth largest element using heap", kthLargestBody(), testKthLargest(m))
	case "24-priority-queue":
		return dsSpec(m, priorityQueueBody(), testPQ(m), "Merge k sorted lists", mergeKBody(), testMergeK(m))
	case "25-hash-tables":
		return dsSpec(m, hashMapBody(), testHashMap(m), "Group anagrams", groupAnagramsBody(), testGroupAnagrams(m))
	case "26-trie":
		return dsSpec(m, trieBody(), testTrie(m), "Autocomplete prefix search", trieSearchBody(), testTrieSearch(m))
	case "27-graphs":
		return dsSpec(m, graphBody(), testGraph(m), "Number of islands (DFS)", islandsBody(), testIslands(m))
	case "28-union-find":
		return dsSpec(m, unionFindBody(), testUF(m), "Count connected components", countComponentsBody(), testCountComponents(m))
	case "29-segment-tree":
		return dsSpec(m, segTreeBody(), testSegTree(m), "Range sum query mutable", segQueryBody(), testSegQuery(m))
	case "30-fenwick-tree":
		return dsSpec(m, fenwickBody(), testFenwick(m), "Range sum with point updates", fenwickRangeBody(), testFenwickRange(m))
	}
	return dsSpec(m, stackBody(), testStack(m), "Implement core data structure", stackBody(), testStack(m))
}

func dsSpec(m module, body, test, problem, interview, interviewTest string) topicSpec {
	bench := defaultBenchCall(firstExportedFunc(body))
	if bench == `solutions.()` {
		bench = `s := solutions.NewStack(); s.Push(1)`
	}
	switch m.slug {
	case "15-maps":
		bench = `solutions.WordCount([]string{"a","b","a"})`
	case "18-linked-lists":
		bench = `solutions.Prepend(nil, 1)`
	case "25-hash-tables":
		bench = `solutions.WordCount([]string{"a","b"})`
	case "26-trie":
		bench = `tr := solutions.NewTrie(); tr.Insert("go")`
	case "21-trees":
		bench = `root := &solutions.TreeNode{Val:1, Left:&solutions.TreeNode{Val:2}}; solutions.TreeHeight(root)`
	case "22-binary-search-tree":
		bench = `root := solutions.InsertBST(nil, 5); solutions.SearchBST(root, 5)`
	case "27-graphs":
		bench = `g := solutions.Graph{1: {2}}; g.BFS(1)`
	default:
		if !strings.Contains(body, "NewStack") {
			bench = `_ = solutions.NewStack()`
		}
	}
	return topicSpec{
		summary:           fmt.Sprintf("Implement **%s** from scratch with tests and complexity analysis.", m.title),
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  problem,
		interviewSolution: interview,
		interviewTest:     interviewTest,
		benchCall:         bench,
	}
}

func algorithmSpec(m module) topicSpec {
	cases := map[string]topicSpec{
		"31-sorting":           algo("MergeSort", mergeSortBody(), testMergeSort(m), "Sort colors (Dutch flag)", dutchFlagBody(), testDutchFlag(m)),
		"32-searching":         algo("BinarySearch", binarySearchBody(), testBinarySearch(m), "Search rotated sorted array", searchRotatedBody(), testSearchRotated(m)),
		"33-sliding-window":    algo("MaxSubarrayLenK", slidingWindowBody(), testSliding(m), "Longest substring without repeat", longestSubstrBody(), testLongestSubstr(m)),
		"34-two-pointers":      algo("TwoSumSorted", twoPointerBody(), testTwoPointer(m), "Container with most water", containerWaterBody(), testContainerWater(m)),
		"35-fast-slow-pointers": algo("HasCycle", cycleBody(), testCycle(m), "Find duplicate (Floyd)", findDuplicateBody(), testFindDuplicate(m)),
		"36-monotonic-stack":   algo("DailyTemperatures", monoStackBody(), testMonoStack(m), "Largest rectangle in histogram", largestRectBody(), testLargestRect(m)),
		"37-monotonic-queue":   algo("SlidingWindowMax", monoQueueBody(), testMonoQueue(m), "Jump game VI", jumpGameBody(), testJumpGame(m)),
		"38-backtracking":      algo("Permutations", permuteBody(), testPermute(m), "N-Queens count solutions", nQueensBody(), testNQueens(m)),
		"39-divide-conquer":    algo("MergeSortDC", mergeSortBody(), testMergeSort(m), "Count inversions", countInversionsBody(), testInversions(m)),
		"40-greedy":            algo("ActivitySelection", greedyBody(), testGreedy(m), "Jump game minimum jumps", minJumpsBody(), testMinJumps(m)),
		"41-dynamic-programming": algo("CoinChange", coinChangeBody(), testCoinChange(m), "Longest increasing subsequence", lisBody(), testLIS(m)),
		"42-bit-manipulation":  algo("SingleNumber", singleNumberBody(), testSingleNumber(m), "Count set bits", countBitsBody(), testCountBits(m)),
		"43-math-algorithms":   algo("GCD", gcdBody(), testGCD(m), "Sieve of Eratosthenes", sieveBody(), testSieve(m)),
		"44-string-algorithms": algo("KMP", kmpBody(), testKMP(m), "Longest palindromic substring", longestPalindromeBody(), testLongestPalindrome(m)),
		"45-graph-algorithms":  algo("Dijkstra", dijkstraBody(), testDijkstra(m), "Course schedule (topological sort)", courseScheduleBody(), testCourseSchedule(m)),
		"46-tree-algorithms":   algo("InorderTraversal", inorderBody(), testInorder(m), "Lowest common ancestor", lcaBody(), testLCA(m)),
	}
	if s, ok := cases[m.slug]; ok {
		s.summary = fmt.Sprintf("**%s** — implement, analyze complexity, and benchmark.", m.title)
		return s
	}
	return algo("Algo", mergeSortBody(), testMergeSort(m), "Classic algorithm", mergeSortBody(), testMergeSort(m))
}

func algo(name, body, test, problem, interview, interviewTest string) topicSpec {
	fn := firstExportedFunc(body)
	if fn == "" {
		fn = name
	}
	return topicSpec{
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  problem,
		interviewSolution: interview,
		interviewTest:     interviewTest,
		benchCall:         defaultBenchCall(fn),
	}
}

func complexitySpec(m module) topicSpec {
	body := `package solutions

// CountOperations returns op count for linear scan of n elements.
func CountOperations(n int) int { return n }

// CountNested returns op count for naive O(n^2) nested loop.
func CountNested(n int) int { return n * n }

// ClassifyGrowth returns big-O label for common growth rates.
func ClassifyGrowth(n int, kind string) string {
	switch kind {
	case "constant":
		return "O(1)"
	case "linear":
		return "O(n)"
	case "quadratic":
		return "O(n^2)"
	case "log":
		return "O(log n)"
	default:
		return "unknown"
	}
}`
	test := testPackage(m, `CountOperations(10)`, "10", `ClassifyGrowth(100,"linear")`, `"O(n)"`)
	return topicSpec{
		summary:           "Analyze time/space complexity of algorithms.",
		exercise1:         body,
		exercise1Test:     test,
		interviewProblem:  "Compare O(n), O(n log n), O(n^2) for given input sizes",
		interviewSolution: body,
		interviewTest:     test,
		benchCall:         `CountOperations(1000)`,
	}
}
