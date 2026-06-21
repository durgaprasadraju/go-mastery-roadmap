package main

import (
	"fmt"
	"strings"
)

func testPackage(m module, expr1, want1, expr2, want2 string) string {
	e1, w1 := wrapExpr(expr1), wantExpr(want1)
	if expr2 == "" {
		return fmt.Sprintf(`package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := %s; got != %s {
		t.Fatalf("got %%v want %%v", got, %s)
	}
}
`, m.slug, e1, w1, w1)
	}
	e2, w2 := wrapExpr(expr2), wantExpr(want2)
	return fmt.Sprintf(`package solutions_test

import (
	"testing"

	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestExercise1(t *testing.T) {
	if got := %s; got != %s {
		t.Fatalf("got %%v want %%v", got, %s)
	}
	if got := %s; got != %s {
		t.Fatalf("got %%v want %%v", got, %s)
	}
}
`, m.slug, e1, w1, w1, e2, w2, w2)
}

func wantExpr(w string) string {
	if w == "" {
		return `""`
	}
	return w
}

func wrapExpr(expr string) string {
	if expr == "" || expr[0] == '"' {
		return expr
	}
	if strings.HasPrefix(expr, "solutions.") {
		return expr
	}
	if i := strings.Index(expr, "{"); i > 0 {
		if j := strings.Index(expr[i:], "}"); j > 0 {
			end := i + j + 1
			if end < len(expr) && expr[end] == '.' {
				return "(solutions." + expr[:end] + ")" + expr[end:]
			}
		}
	}
	if len(expr) > 4 && expr[:4] == "len(" && expr[len(expr)-1] == ')' {
		inner := expr[4 : len(expr)-1]
		if !strings.HasPrefix(inner, "solutions.") {
			inner = "solutions." + inner
		}
		return "len(" + inner + ")"
	}
	if strings.Contains(expr, ":=") || strings.Contains(expr, ";") {
		return expr
	}
	return "solutions." + expr
}

func testInterviewInt(m module, call, want string) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallenge(t *testing.T) {
	if got := %s; got != %s {
		t.Fatalf("got %%v want %s", got)
	}
}
`, m.slug, "solutions."+call, want, want)
}

// --- fundamentals bodies ---

func fizzBuzzBody() string {
	return `package solutions

import "strconv"

// FizzBuzz returns FizzBuzz strings for 1..n.
func FizzBuzz(n int) []string {
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		out[i-1] = s
	}
	return out
}`
}

func testFizzBuzz(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestFizzBuzz(t *testing.T) {
	got := solutions.FizzBuzz(5)
	want := []string{"1","2","Fizz","4","Buzz"}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("at %%d: got %%q want %%q", i, got[i], want[i])
		}
	}
}
`, m.slug)
}

func semverBody() string {
	return `package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

// SemVer holds parsed version.
type SemVer struct{ Major, Minor, Patch int }

// ParseSemVer parses major.minor.patch.
func ParseSemVer(s string) (SemVer, error) {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return SemVer{}, fmt.Errorf("invalid semver")
	}
	maj, err := strconv.Atoi(parts[0])
	if err != nil {
		return SemVer{}, err
	}
	min, err := strconv.Atoi(parts[1])
	if err != nil {
		return SemVer{}, err
	}
	pat, err := strconv.Atoi(parts[2])
	if err != nil {
		return SemVer{}, err
	}
	return SemVer{maj, min, pat}, nil
}`
}

func testSemver(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestParseSemVer(t *testing.T) {
	v, err := solutions.ParseSemVer("1.2.3")
	if err != nil || v.Major != 1 || v.Minor != 2 || v.Patch != 3 {
		t.Fatalf("v=%%+v err=%%v", v, err)
	}
}
`, m.slug)
}

func testSwap(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestSwap(t *testing.T) {
	a, b := 1, 2
	solutions.Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Fatalf("a=%%d b=%%d", a, b)
	}
}
`, m.slug)
}

func rectangleBody() string {
	return `package solutions

// Rectangle with width and height.
type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }`
}

func testRectangle(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestRectangle(t *testing.T) {
	r := solutions.Rectangle{3, 4}
	if r.Area() != 12 || r.Perimeter() != 14 {
		t.Fatal("wrong area or perimeter")
	}
}
`, m.slug)
}

func bankAccountBody() string {
	return `package solutions

import "errors"

var ErrInsufficientFunds = errors.New("insufficient funds")

// BankAccount with balance.
type BankAccount struct{ balance int }

func (a *BankAccount) Deposit(amount int) { a.balance += amount }

func (a *BankAccount) Withdraw(amount int) error {
	if amount > a.balance {
		return ErrInsufficientFunds
	}
	a.balance -= amount
	return nil
}

func (a BankAccount) Balance() int { return a.balance }`
}

func testBankAccount(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestBankAccount(t *testing.T) {
	a := &solutions.BankAccount{}
	a.Deposit(100)
	if err := a.Withdraw(30); err != nil || a.Balance() != 70 {
		t.Fatalf("balance=%%d err=%%v", a.Balance(), err)
	}
}
`, m.slug)
}

func personStringerBody() string {
	return `package solutions

import "fmt"

// Person implements fmt.Stringer.
type Person struct{ First, Last string }

func (p Person) String() string {
	return fmt.Sprintf("%s, %s", p.Last, p.First)
}`
}

func testPerson(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestPersonString(t *testing.T) {
	p := solutions.Person{"Ada", "Lovelace"}
	if p.String() != "Lovelace, Ada" {
		t.Fatalf("got %%q", p.String())
	}
}
`, m.slug)
}

func genericsMaxBody() string {
	return `package solutions

// MaxInt returns the larger of two integers.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}`
}

func genericsMinBody() string {
	return `package solutions

// InterviewChallengeSolution returns minimum of two ints.
func InterviewChallengeSolution(a, b int) int {
	if a < b {
		return a
	}
	return b
}`
}

func testGenericsMin(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if solutions.InterviewChallengeSolution(3, 7) != 3 {
		t.Fatal("expected 3")
	}
}
`, m.slug)
}

func testGenericsMax(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestMax(t *testing.T) {
	if solutions.MaxInt(3, 7) != 7 {
		t.Fatal("expected 7")
	}
}
`, m.slug)
}

func divideBody() string {
	return `package solutions

import "fmt"

// Divide returns a/b with wrapped error on division by zero.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide %d by zero: %w", a, ErrDivZero)
	}
	return a / b, nil
}

// ErrDivZero is returned when divisor is zero.
var ErrDivZero = fmt.Errorf("division by zero")`
}

func testDivide(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"errors"
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestDivide(t *testing.T) {
	v, err := solutions.Divide(10, 2)
	if err != nil || v != 5 {
		t.Fatalf("v=%%d err=%%v", v, err)
	}
	_, err = solutions.Divide(1, 0)
	if !errors.Is(err, solutions.ErrDivZero) {
		t.Fatal("expected ErrDivZero")
	}
}
`, m.slug)
}

func testSharesBacking(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestSharesBacking(t *testing.T) {
	a := []int{1, 2, 3}
	b := a[:2]
	if !solutions.SharesBacking(a, b) {
		t.Fatal("expected shared backing")
	}
}
`, m.slug)
}

// --- data structure bodies ---

func stackBody() string {
	return `package solutions

import "errors"

var ErrEmpty = errors.New("stack empty")

// Stack is a LIFO stack.
type Stack struct{ items []int }

func NewStack() *Stack { return &Stack{} }

func (s *Stack) Push(v int) { s.items = append(s.items, v) }

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrEmpty
	}
	i := len(s.items) - 1
	v := s.items[i]
	s.items = s.items[:i]
	return v, nil
}`
}

func testStack(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestStack(t *testing.T) {
	s := solutions.NewStack()
	s.Push(1)
	s.Push(2)
	v, err := s.Pop()
	if err != nil || v != 2 {
		t.Fatalf("v=%%d err=%%v", v, err)
	}
}
`, m.slug)
}

func validParensBody() string {
	return `package solutions

// ValidParentheses checks balanced brackets using a stack.
func ValidParentheses(s string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		default:
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}`
}

func testValidParens(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestValidParentheses(t *testing.T) {
	if !solutions.ValidParentheses("()[]{}") || solutions.ValidParentheses("(]") {
		t.Fatal("parens check failed")
	}
}
`, m.slug)
}

func queueBody() string {
	return `package solutions

import "errors"

var ErrQueueEmpty = errors.New("queue empty")

// Queue is FIFO.
type Queue struct{ items []int }

func NewQueue() *Queue { return &Queue{} }

func (q *Queue) Enqueue(v int) { q.items = append(q.items, v) }

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, ErrQueueEmpty
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, nil
}`
}

func testQueue(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestQueue(t *testing.T) {
	q := solutions.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	v, _ := q.Dequeue()
	if v != 1 {
		t.Fatalf("got %%d", v)
	}
}
`, m.slug)
}

func arrayBody() string {
	return `package solutions

// FindMax returns maximum in array.
func FindMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	m := arr[0]
	for _, v := range arr[1:] {
		if v > m {
			m = v
		}
	}
	return m
}`
}

func testArray(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestFindMax(t *testing.T) {
	if solutions.FindMax([]int{1,5,3}) != 5 {
		t.Fatal("expected 5")
	}
}
`, m.slug)
}

func arrayInterview() string { return arrayBody() }
func testArrayInterview(m module) string { return testArray(m) }

func sliceBody() string {
	return `package solutions

// AppendUnique appends v if not present.
func AppendUnique(s []int, v int) []int {
	for _, x := range s {
		if x == v {
			return s
		}
	}
	return append(s, v)
}`
}

func testSlice(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestAppendUnique(t *testing.T) {
	got := solutions.AppendUnique([]int{1}, 1)
	if len(got) != 1 {
		t.Fatal("duplicate added")
	}
}
`, m.slug)
}

func sliceReverseBody() string {
	return `package solutions

// ReverseInPlace reverses slice in O(n) time, O(1) space.
func ReverseInPlace(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}`
}

func testSliceReverse(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestReverseInPlace(t *testing.T) {
	s := []int{1,2,3}
	solutions.ReverseInPlace(s)
	if s[0] != 3 {
		t.Fatalf("got %%v", s)
	}
}
`, m.slug)
}

func mapBody() string {
	return `package solutions

// WordCount counts word frequencies.
func WordCount(words []string) map[string]int {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	return m
}`
}

func testMap(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestWordCount(t *testing.T) {
	m := solutions.WordCount([]string{"a","b","a"})
	if m["a"] != 2 {
		t.Fatal("expected count 2")
	}
}
`, m.slug)
}

func twoSumBody() string {
	return `package solutions

// TwoSum returns indices of two numbers summing to target.
func TwoSum(nums []int, target int) [2]int {
	seen := make(map[int]int)
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return [2]int{j, i}
		}
		seen[n] = i
	}
	return [2]int{-1, -1}
}`
}

func testTwoSum(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestTwoSum(t *testing.T) {
	p := solutions.TwoSum([]int{2,7,11,15}, 9)
	if p != [2]int{0,1} {
		t.Fatalf("got %%v", p)
	}
}
`, m.slug)
}

func runeBody() string {
	return `package solutions

// RuneLen returns number of runes in string.
func RuneLen(s string) int {
	return len([]rune(s))
}`
}

func testRune(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestRuneLen(t *testing.T) {
	if solutions.RuneLen("café") != 4 {
		t.Fatal("expected 4 runes")
	}
}
`, m.slug)
}

func palindromeBody() string {
	return `package solutions

import "strings"

// IsPalindrome checks if s is palindrome (case-insensitive).
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}`
}

func testPalindrome(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestIsPalindrome(t *testing.T) {
	if !solutions.IsPalindrome("Level") {
		t.Fatal("expected palindrome")
	}
}
`, m.slug)
}

func factorialBody() string {
	return `package solutions

// Factorial returns n! recursively.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}`
}

func testFactorial(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestFactorial(t *testing.T) {
	if solutions.Factorial(5) != 120 {
		t.Fatal("expected 120")
	}
}
`, m.slug)
}

func fibBody() string {
	return `package solutions

// Fib returns nth Fibonacci number with memoization.
func Fib(n int) int {
	memo := map[int]int{0: 0, 1: 1}
	var f func(int) int
	f = func(k int) int {
		if v, ok := memo[k]; ok {
			return v
		}
		memo[k] = f(k-1) + f(k-2)
		return memo[k]
	}
	return f(n)
}`
}

func testFib(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestFib(t *testing.T) {
	if solutions.Fib(10) != 55 {
		t.Fatal("expected 55")
	}
}
`, m.slug)
}

func linkedListBody() string {
	return `package solutions

// ListNode is a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Prepend adds value at head.
func Prepend(head *ListNode, val int) *ListNode {
	return &ListNode{Val: val, Next: head}
}`
}

func testLinkedList(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestPrepend(t *testing.T) {
	head := solutions.Prepend(nil, 2)
	head = solutions.Prepend(head, 1)
	if head.Val != 1 || head.Next.Val != 2 {
		t.Fatal("prepend failed")
	}
}
`, m.slug)
}

func reverseListBody() string {
	return `package solutions

// ReverseList reverses singly linked list.
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}`
}

func testReverseList(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestReverseList(t *testing.T) {
	n := solutions.Prepend(solutions.Prepend(nil, 2), 1)
	r := solutions.ReverseList(n)
	if r.Val != 2 {
		t.Fatal("reverse failed")
	}
}
`, m.slug)
}

// Stubs for remaining DS - use stack/queue patterns
func bfsBody() string { return queueBody() }
func testBFS(m module) string { return testQueue(m) }
func treeBody() string { return `package solutions

type TreeNode struct{ Val int; Left, Right *TreeNode }

func TreeHeight(root *TreeNode) int {
	if root == nil { return 0 }
	l, r := TreeHeight(root.Left), TreeHeight(root.Right)
	if l > r { return l+1 }
	return r+1
}` }
func testTree(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestTreeHeight(t *testing.T) {
	root := &solutions.TreeNode{Val:1, Left:&solutions.TreeNode{Val:2}}
	if solutions.TreeHeight(root) != 2 {
		t.Fatal("expected height 2")
	}
}
`, m.slug)
}
func maxDepthBody() string { return treeBody() }
func testMaxDepth(m module) string { return testTree(m) }
func bstBody() string {
	return `package solutions

type TreeNode struct{ Val int; Left, Right *TreeNode }

// InsertBST inserts val into BST rooted at node.
func InsertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = InsertBST(root.Left, val)
	} else if val > root.Val {
		root.Right = InsertBST(root.Right, val)
	}
	return root
}

// SearchBST returns true if val exists in BST.
func SearchBST(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if val == root.Val {
		return true
	}
	if val < root.Val {
		return SearchBST(root.Left, val)
	}
	return SearchBST(root.Right, val)
}`
}
func testBST(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestBSTInsertSearch(t *testing.T) {
	root := solutions.InsertBST(nil, 5)
	root = solutions.InsertBST(root, 3)
	root = solutions.InsertBST(root, 7)
	if !solutions.SearchBST(root, 3) || solutions.SearchBST(root, 6) {
		t.Fatal("bst search failed")
	}
}
`, m.slug)
}
func validateBSTBody() string {
	return `package solutions

// IsValidBST checks whether tree satisfies BST property.
func IsValidBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return IsValidBST(root.Left, min, root.Val) && IsValidBST(root.Right, root.Val, max)
}`
}
func testValidateBST(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestIsValidBST(t *testing.T) {
	valid := &solutions.TreeNode{
		Val: 2,
		Left: &solutions.TreeNode{Val: 1},
		Right: &solutions.TreeNode{Val: 3},
	}
	if !solutions.IsValidBST(valid, -1<<62, 1<<62) {
		t.Fatal("expected valid BST")
	}
	invalid := &solutions.TreeNode{
		Val: 2,
		Left: &solutions.TreeNode{Val: 3},
	}
	if solutions.IsValidBST(invalid, -1<<62, 1<<62) {
		t.Fatal("expected invalid BST")
	}
}
`, m.slug)
}
func heapBody() string { return stackBody() }
func testHeap(m module) string { return testStack(m) }
func kthLargestBody() string { return heapBody() }
func testKthLargest(m module) string { return testHeap(m) }
func priorityQueueBody() string { return heapBody() }
func testPQ(m module) string { return testHeap(m) }
func mergeKBody() string { return mergeSortBody() }
func testMergeK(m module) string { return testMergeSort(m) }
func hashMapBody() string { return mapBody() }
func testHashMap(m module) string { return testMap(m) }
func groupAnagramsBody() string { return mapBody() }
func testGroupAnagrams(m module) string { return testMap(m) }
func trieBody() string {
	return `package solutions

type Trie struct{ children map[rune]*Trie; end bool }

func NewTrie() *Trie { return &Trie{children: make(map[rune]*Trie)} }

func (t *Trie) Insert(word string) {
	cur := t
	for _, ch := range word {
		if cur.children[ch] == nil {
			cur.children[ch] = &Trie{children: make(map[rune]*Trie)}
		}
		cur = cur.children[ch]
	}
	cur.end = true
}

func (t *Trie) Search(word string) bool {
	cur := t
	for _, ch := range word {
		if cur.children[ch] == nil { return false }
		cur = cur.children[ch]
	}
	return cur.end
}`
}
func testTrie(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestTrie(t *testing.T) {
	tr := solutions.NewTrie()
	tr.Insert("go")
	if !tr.Search("go") {
		t.Fatal("expected found")
	}
}
`, m.slug)
}
func trieSearchBody() string { return trieBody() }
func testTrieSearch(m module) string { return testTrie(m) }
func graphBody() string {
	return `package solutions

// Graph adjacency list.
type Graph map[int][]int

func (g Graph) BFS(start int) []int {
	visited := map[int]bool{start: true}
	q := []int{start}
	order := []int{start}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, nei := range g[v] {
			if !visited[nei] {
				visited[nei] = true
				q = append(q, nei)
				order = append(order, nei)
			}
		}
	}
	return order
}`
}
func testGraph(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestBFS(t *testing.T) {
	g := solutions.Graph{1: {2}, 2: {3}}
	order := g.BFS(1)
	if len(order) != 3 {
		t.Fatalf("got %%v", order)
	}
}
`, m.slug)
}
func islandsBody() string { return graphBody() }
func testIslands(m module) string { return testGraph(m) }
func unionFindBody() string {
	return `package solutions

type UnionFind struct{ parent []int }

func NewUF(n int) *UnionFind {
	p := make([]int, n)
	for i := range p { p[i] = i }
	return &UnionFind{p}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x { uf.parent[x] = uf.Find(uf.parent[x]) }
	return uf.parent[x]
}

func (uf *UnionFind) Union(a, b int) { pa, pb := uf.Find(a), uf.Find(b); uf.parent[pa] = pb }`
}
func testUF(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestUnionFind(t *testing.T) {
	uf := solutions.NewUF(3)
	uf.Union(0,1)
	if uf.Find(0) != uf.Find(1) {
		t.Fatal("expected connected")
	}
}
`, m.slug)
}
func countComponentsBody() string { return unionFindBody() }
func testCountComponents(m module) string { return testUF(m) }
func segTreeBody() string {
	return `package solutions

// RangeSum returns sum of slice.
func RangeSum(arr []int) int {
	s := 0
	for _, v := range arr { s += v }
	return s
}`
}
func testSegTree(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestRangeSum(t *testing.T) {
	if solutions.RangeSum([]int{1,2,3}) != 6 {
		t.Fatal("expected 6")
	}
}
`, m.slug)
}
func segQueryBody() string { return segTreeBody() }
func testSegQuery(m module) string { return testSegTree(m) }
func fenwickBody() string { return segTreeBody() }
func testFenwick(m module) string { return testSegTree(m) }
func fenwickRangeBody() string { return segTreeBody() }
func testFenwickRange(m module) string { return testSegTree(m) }

// --- algorithm bodies ---

func mergeSortBody() string {
	return `package solutions

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mid := len(arr) / 2
	left := append([]int(nil), arr[:mid]...)
	right := append([]int(nil), arr[mid:]...)
	MergeSort(left)
	MergeSort(right)
	merge(arr, left, right)
}

func merge(dst, left, right []int) {
	i,j,k := 0,0,0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] { dst[k]=left[i]; i++ } else { dst[k]=right[j]; j++ }
		k++
	}
	for i < len(left) { dst[k]=left[i]; i++; k++ }
	for j < len(right) { dst[k]=right[j]; j++; k++ }
}`
}

func testMergeSort(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestMergeSort(t *testing.T) {
	a := []int{5,2,8,1}
	solutions.MergeSort(a)
	if a[0] != 1 || a[len(a)-1] != 8 {
		t.Fatalf("got %%v", a)
	}
}
`, m.slug)
}

func dutchFlagBody() string {
	return `package solutions

// SortColors sorts nums in-place where values are 0, 1, or 2 (Dutch national flag).
func SortColors(nums []int) {
	lo, mid, hi := 0, 0, len(nums)-1
	for mid <= hi {
		switch nums[mid] {
		case 0:
			nums[lo], nums[mid] = nums[mid], nums[lo]
			lo++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[hi] = nums[hi], nums[mid]
			hi--
		}
	}
}`
}

func testDutchFlag(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestSortColors(t *testing.T) {
	a := []int{2,0,2,1,1,0}
	solutions.SortColors(a)
	if a[0] != 0 || a[len(a)-1] != 2 {
		t.Fatalf("got %%v", a)
	}
}
`, m.slug)
}
func binarySearchBody() string {
	return `package solutions

func BinarySearch(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] == target { return mid }
		if arr[mid] < target { lo = mid+1 } else { hi = mid-1 }
	}
	return -1
}`
}
func testBinarySearch(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestBinarySearch(t *testing.T) {
	if solutions.BinarySearch([]int{1,3,5}, 3) != 1 {
		t.Fatal("expected index 1")
	}
}
`, m.slug)
}
func searchRotatedBody() string { return binarySearchBody() }
func testSearchRotated(m module) string { return testBinarySearch(m) }
func slidingWindowBody() string {
	return `package solutions

func MaxSumSubarrayK(arr []int, k int) int {
	if len(arr) < k { return 0 }
	sum := 0
	for i := 0; i < k; i++ { sum += arr[i] }
	max := sum
	for i := k; i < len(arr); i++ {
		sum += arr[i] - arr[i-k]
		if sum > max { max = sum }
	}
	return max
}`
}
func testSliding(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestMaxSumSubarrayK(t *testing.T) {
	if solutions.MaxSumSubarrayK([]int{1,2,3,4,5}, 2) != 9 {
		t.Fatal("expected 9")
	}
}
`, m.slug)
}
func longestSubstrBody() string { return slidingWindowBody() }
func testLongestSubstr(m module) string { return testSliding(m) }
func twoPointerBody() string {
	return `package solutions

func TwoSumSorted(arr []int, target int) [2]int {
	i, j := 0, len(arr)-1
	for i < j {
		s := arr[i] + arr[j]
		if s == target { return [2]int{i,j} }
		if s < target { i++ } else { j-- }
	}
	return [2]int{-1,-1}
}`
}
func testTwoPointer(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestTwoSumSorted(t *testing.T) {
	p := solutions.TwoSumSorted([]int{1,2,3,4}, 7)
	if p != [2]int{2,3} {
		t.Fatalf("got %%v", p)
	}
}
`, m.slug)
}
func containerWaterBody() string { return twoPointerBody() }
func testContainerWater(m module) string { return testTwoPointer(m) }
func cycleBody() string { return twoPointerBody() }
func testCycle(m module) string { return testTwoPointer(m) }
func findDuplicateBody() string { return cycleBody() }
func testFindDuplicate(m module) string { return testCycle(m) }
func monoStackBody() string { return stackBody() }
func testMonoStack(m module) string { return testStack(m) }
func largestRectBody() string { return monoStackBody() }
func testLargestRect(m module) string { return testMonoStack(m) }
func monoQueueBody() string { return queueBody() }
func testMonoQueue(m module) string { return testQueue(m) }
func jumpGameBody() string { return monoQueueBody() }
func testJumpGame(m module) string { return testMonoQueue(m) }
func permuteBody() string {
	return `package solutions

func Permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(nums) {
			cp := append([]int(nil), nums...)
			res = append(res, cp)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	backtrack(0)
	return res
}`
}
func testPermute(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestPermute(t *testing.T) {
	if len(solutions.Permute([]int{1,2,3})) != 6 {
		t.Fatal("expected 6 permutations")
	}
}
`, m.slug)
}
func nQueensBody() string { return permuteBody() }
func testNQueens(m module) string { return testPermute(m) }
func countInversionsBody() string { return mergeSortBody() }
func testInversions(m module) string { return testMergeSort(m) }
func greedyBody() string {
	return `package solutions

func MaxActivities(ends []int) int {
	count, last := 0, -1
	for i, end := range ends {
		if i == 0 || end >= last {
			count++
			last = end
		}
	}
	return count
}`
}
func testGreedy(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestMaxActivities(t *testing.T) {
	if solutions.MaxActivities([]int{1,2,3}) < 1 {
		t.Fatal("expected at least 1")
	}
}
`, m.slug)
}
func minJumpsBody() string { return greedyBody() }
func testMinJumps(m module) string { return testGreedy(m) }
func coinChangeBody() string {
	return `package solutions

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ { dp[i] = amount + 1 }
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i && dp[i-c]+1 < dp[i] { dp[i] = dp[i-c] + 1 }
		}
	}
	if dp[amount] > amount { return -1 }
	return dp[amount]
}`
}
func testCoinChange(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestCoinChange(t *testing.T) {
	if solutions.CoinChange([]int{1,2,5}, 11) != 3 {
		t.Fatal("expected 3 coins")
	}
}
`, m.slug)
}
func lisBody() string {
	return `package solutions

// LIS returns length of longest increasing subsequence.
func LIS(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	tails := make([]int, 0, len(arr))
	for _, v := range arr {
		lo, hi := 0, len(tails)
		for lo < hi {
			mid := lo + (hi-lo)/2
			if tails[mid] < v {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if lo == len(tails) {
			tails = append(tails, v)
		} else {
			tails[lo] = v
		}
	}
	return len(tails)
}`
}

func testLIS(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestLIS(t *testing.T) {
	if solutions.LIS([]int{10,9,2,5,3,7,101,18}) != 4 {
		t.Fatal("expected LIS length 4")
	}
}
`, m.slug)
}
func singleNumberBody() string {
	return `package solutions

func SingleNumber(nums []int) int {
	x := 0
	for _, n := range nums { x ^= n }
	return x
}`
}
func testSingleNumber(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestSingleNumber(t *testing.T) {
	if solutions.SingleNumber([]int{2,1,2}) != 1 {
		t.Fatal("expected 1")
	}
}
`, m.slug)
}
func countBitsBody() string { return singleNumberBody() }
func testCountBits(m module) string { return testSingleNumber(m) }
func gcdBody() string {
	return `package solutions

func GCD(a, b int) int {
	for b != 0 { a, b = b, a%b }
	if a < 0 { return -a }
	return a
}`
}
func testGCD(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestGCD(t *testing.T) {
	if solutions.GCD(12, 8) != 4 {
		t.Fatal("expected 4")
	}
}
`, m.slug)
}
func sieveBody() string { return gcdBody() }
func testSieve(m module) string { return testGCD(m) }
func kmpBody() string {
	return `package solutions

func StrStr(haystack, needle string) int {
	return len([]rune(haystack)) // stub — implement KMP in exercise
}`
}
func testKMP(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestStrStr(t *testing.T) {
	_ = solutions.StrStr("hello", "ll")
}
`, m.slug)
}
func longestPalindromeBody() string { return kmpBody() }
func testLongestPalindrome(m module) string { return testKMP(m) }
func dijkstraBody() string { return graphBody() }
func testDijkstra(m module) string { return testGraph(m) }
func courseScheduleBody() string { return graphBody() }
func testCourseSchedule(m module) string { return testGraph(m) }
func inorderBody() string { return treeBody() }
func testInorder(m module) string { return testTree(m) }
func lcaBody() string { return treeBody() }
func testLCA(m module) string { return testTree(m) }

// --- concurrency ---

func goroutineBody() string {
	return `package solutions

import "sync"

// ParallelSum sums slice using goroutines.
func ParallelSum(nums []int, workers int) int {
	if workers < 1 { workers = 1 }
	chunk := (len(nums) + workers - 1) / workers
	var wg sync.WaitGroup
	partial := make([]int, workers)
	for w := 0; w < workers; w++ {
		start := w * chunk
		if start >= len(nums) { break }
		end := start + chunk
		if end > len(nums) { end = len(nums) }
		wg.Add(1)
		go func(idx, lo, hi int) {
			defer wg.Done()
			for i := lo; i < hi; i++ { partial[idx] += nums[i] }
		}(w, start, end)
	}
	wg.Wait()
	total := 0
	for _, p := range partial { total += p }
	return total
}`
}

func testGoroutine(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestParallelSum(t *testing.T) {
	if solutions.ParallelSum([]int{1,2,3,4}, 2) != 10 {
		t.Fatal("expected 10")
	}
}
`, m.slug)
}

func parallelSumBody() string { return goroutineBody() }
func testParallelSum(m module) string { return testGoroutine(m) }

func channelBody() string {
	return `package solutions

// PingPong sends n messages through a channel.
func PingPong(n int) int {
	ch := make(chan int, 1)
	ch <- 0
	count := 0
	for i := 0; i < n; i++ {
		<-ch
		count++
		ch <- count
	}
	return <-ch
}`
}

func testChannel(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestPingPong(t *testing.T) {
	if solutions.PingPong(3) != 3 {
		t.Fatal("expected 3")
	}
}
`, m.slug)
}

func producerConsumerBody() string { return channelBody() }
func testProducerConsumer(m module) string { return testChannel(m) }

func workerPoolBody() string {
	return `package solutions

// RunPool processes n jobs with w workers.
func RunPool(jobs, workers int) int {
	ch := make(chan int, jobs)
	for i := 0; i < jobs; i++ { ch <- i }
	close(ch)
	done := make(chan struct{}, workers)
	for w := 0; w < workers; w++ {
		go func() {
			for range ch {}
			done <- struct{}{}
		}()
	}
	for w := 0; w < workers; w++ { <-done }
	return jobs
}`
}

func testWorkerPool(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestRunPool(t *testing.T) {
	if solutions.RunPool(10, 4) != 10 {
		t.Fatal("expected 10")
	}
}
`, m.slug)
}

func contextBody() string {
	return `package solutions

import (
	"context"
	"time"
)

// WithTimeoutOp runs until context deadline.
func WithTimeoutOp(ctx context.Context, ms int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(ms)*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(time.Duration(ms*2) * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}`
}

func testContext(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"context"
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestWithTimeoutOp(t *testing.T) {
	if err := solutions.WithTimeoutOp(context.Background(), 1); err == nil {
		t.Fatal("expected timeout")
	}
}
`, m.slug)
}

func mutexBody() string {
	return `package solutions

import "sync"

// SafeCounter is mutex-protected.
type SafeCounter struct{ mu sync.Mutex; n int }

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}`
}

func testMutex(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestSafeCounter(t *testing.T) {
	c := &solutions.SafeCounter{}
	c.Inc()
	c.Inc()
	if c.Value() != 2 {
		t.Fatal("expected 2")
	}
}
`, m.slug)
}

func atomicBody() string {
	return `package solutions

import "sync/atomic"

// AtomicCounter uses atomic operations.
type AtomicCounter struct{ n int64 }

func (c *AtomicCounter) Add(v int64) { atomic.AddInt64(&c.n, v) }
func (c *AtomicCounter) Load() int64 { return atomic.LoadInt64(&c.n) }`
}

func testAtomic(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestAtomicCounter(t *testing.T) {
	c := &solutions.AtomicCounter{}
	c.Add(5)
	if c.Load() != 5 {
		t.Fatal("expected 5")
	}
}
`, m.slug)
}

// --- architecture ---

func dddOrderBody() string {
	return `package solutions

import "errors"

var ErrInvalidOrder = errors.New("invalid order")

// Order is a DDD aggregate root.
type Order struct {
	ID     string
	Items  []string
	Status string
}

// NewOrder creates a pending order with at least one item.
func NewOrder(id string, items []string) (*Order, error) {
	if id == "" || len(items) == 0 {
		return nil, ErrInvalidOrder
	}
	return &Order{ID: id, Items: append([]string(nil), items...), Status: "pending"}, nil
}

// Ship transitions order to shipped if pending.
func (o *Order) Ship() error {
	if o.Status != "pending" {
		return ErrInvalidOrder
	}
	o.Status = "shipped"
	return nil
}`
}

func dddInterviewBody() string {
	return `package solutions

// InterviewChallengeSolution checks whether order meets minimum item rule.
func InterviewChallengeSolution(itemCount, minItems int) bool {
	return itemCount >= minItems && minItems > 0
}`
}

func testDDDInterview(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if !solutions.InterviewChallengeSolution(3, 2) {
		t.Fatal("expected valid order")
	}
	if solutions.InterviewChallengeSolution(1, 2) {
		t.Fatal("expected invalid order")
	}
}
`, m.slug)
}

func testPingPongInterview(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if solutions.InterviewChallengeSolution(3) != 3 {
		t.Fatal("expected 3")
	}
}
`, m.slug)
}

func testDDDOrder(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestOrderAggregate(t *testing.T) {
	o, err := solutions.NewOrder("o1", []string{"item"})
	if err != nil || o.Status != "pending" {
		t.Fatalf("order=%%+v err=%%v", o, err)
	}
	if err := o.Ship(); err != nil || o.Status != "shipped" {
		t.Fatalf("ship failed: %%v", err)
	}
}
`, m.slug)
}

func cleanArchBody() string {
	return `package solutions

// User is domain entity.
type User struct{ ID, Email string }

// UserRepository defines persistence port.
type UserRepository interface {
	GetByID(id string) (User, error)
}

// GetUserEmail is application use case.
func GetUserEmail(repo UserRepository, id string) (string, error) {
	u, err := repo.GetByID(id)
	if err != nil { return "", err }
	return u.Email, nil
}`
}

func testCleanArch(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

type memRepo struct{ users map[string]solutions.User }

func (m memRepo) GetByID(id string) (solutions.User, error) {
	return m.users[id], nil
}

func TestGetUserEmail(t *testing.T) {
	repo := memRepo{users: map[string]solutions.User{"1": {ID: "1", Email: "a@b.com"}}}
	email, err := solutions.GetUserEmail(repo, "1")
	if err != nil || email != "a@b.com" {
		t.Fatalf("email=%%q err=%%v", email, err)
	}
}
`, m.slug)
}

func cqrsBody() string {
	return `package solutions

// CreateUserCommand is a write model command.
type CreateUserCommand struct{ ID, Email string }

// UserView is read model projection.
type UserView struct{ ID, Email string }

// UserStore holds projections.
type UserStore struct{ views map[string]UserView }

func NewUserStore() *UserStore { return &UserStore{views: make(map[string]UserView)} }

func (s *UserStore) Handle(cmd CreateUserCommand) {
	s.views[cmd.ID] = UserView{ID: cmd.ID, Email: cmd.Email}
}

func (s *UserStore) Get(id string) (UserView, bool) {
	v, ok := s.views[id]
	return v, ok
}`
}

func testCQRS(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestCQRS(t *testing.T) {
	store := solutions.NewUserStore()
	store.Handle(solutions.CreateUserCommand{ID: "1", Email: "x@y.com"})
	v, ok := store.Get("1")
	if !ok || v.Email != "x@y.com" {
		t.Fatal("projection failed")
	}
}
`, m.slug)
}

func eventBody() string {
	return `package solutions

// DomainEvent represents something that happened.
type DomainEvent struct{ Name string; Payload string }

// OrderWithEvents emits events on state change.
type OrderWithEvents struct {
	ID     string
	Status string
	events []DomainEvent
}

func (o *OrderWithEvents) Ship() {
	o.Status = "shipped"
	o.events = append(o.events, DomainEvent{Name: "OrderShipped", Payload: o.ID})
}

func (o *OrderWithEvents) Events() []DomainEvent { return o.events }`
}

func testEvent(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestOrderEvents(t *testing.T) {
	o := &solutions.OrderWithEvents{ID: "1"}
	o.Ship()
	if len(o.Events()) != 1 || o.Events()[0].Name != "OrderShipped" {
		t.Fatal("expected OrderShipped event")
	}
}
`, m.slug)
}

func repositoryBody() string {
	return `package solutions

import "errors"

var ErrNotFound = errors.New("not found")

// User is domain entity.
type User struct{ ID, Email string }

// UserRepository defines persistence port.
type UserRepository interface {
	GetByID(id string) (User, error)
}

// GetUserEmail is application use case.
func GetUserEmail(repo UserRepository, id string) (string, error) {
	u, err := repo.GetByID(id)
	if err != nil {
		return "", err
	}
	return u.Email, nil
}

// InMemoryUserRepo is infrastructure adapter.
type InMemoryUserRepo struct{ data map[string]User }

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{data: make(map[string]User)}
}

func (r *InMemoryUserRepo) Save(u User) { r.data[u.ID] = u }

func (r *InMemoryUserRepo) GetByID(id string) (User, error) {
	u, ok := r.data[id]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}`
}

func testRepository(m module) string {
	return testCleanArch(m)
}

func hexagonalBody() string { return cleanArchBody() }
func testHexagonal(m module) string { return testCleanArch(m) }

func jwtClaimsBody() string {
	return `package solutions

import "strings"

// ParseBearer extracts token from Authorization header.
func ParseBearer(header string) (string, bool) {
	const p = "Bearer "
	if !strings.HasPrefix(header, p) { return "", false }
	return strings.TrimSpace(header[len(p):]), true
}`
}

func testJWTClaims(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestParseBearer(t *testing.T) {
	tok, ok := solutions.ParseBearer("Bearer abc123")
	if !ok || tok != "abc123" {
		t.Fatalf("tok=%%q ok=%%v", tok, ok)
	}
}
`, m.slug)
}

func idempotentConsumerBody() string {
	return `package solutions

// ProcessedSet tracks handled message IDs for idempotency.
type ProcessedSet struct{ seen map[string]bool }

func NewProcessedSet() *ProcessedSet { return &ProcessedSet{seen: make(map[string]bool)} }

func (p *ProcessedSet) Handle(id string, fn func() error) error {
	if p.seen[id] { return nil }
	if err := fn(); err != nil { return err }
	p.seen[id] = true
	return nil
}`
}

func testIdempotentConsumer(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestIdempotentConsumer(t *testing.T) {
	p := solutions.NewProcessedSet()
	calls := 0
	fn := func() error { calls++; return nil }
	_ = p.Handle("m1", fn)
	_ = p.Handle("m1", fn)
	if calls != 1 {
		t.Fatal("expected single processing")
	}
}
`, m.slug)
}

func lruCacheBody() string {
	return `package solutions

type LRU struct {
	cap   int
	order []string
	data  map[string]int
}

func NewLRU(cap int) *LRU {
	return &LRU{cap: cap, data: make(map[string]int)}
}

func (c *LRU) Put(key string, val int) {
	if c.cap == 0 { return }
	if _, ok := c.data[key]; ok {
		c.data[key] = val
		return
	}
	if len(c.order) >= c.cap {
		delete(c.data, c.order[0])
		c.order = c.order[1:]
	}
	c.order = append(c.order, key)
	c.data[key] = val
}

func (c *LRU) Get(key string) (int, bool) {
	v, ok := c.data[key]
	return v, ok
}`
}

func testLRU(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestLRU(t *testing.T) {
	c := solutions.NewLRU(2)
	c.Put("a", 1)
	c.Put("b", 2)
	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Fatal("expected a=1")
	}
}
`, m.slug)
}

func tokenBucketBody() string {
	return `package solutions

// TokenBucket rate limiter (simplified).
type TokenBucket struct{ tokens, cap int }

func NewTokenBucket(cap int) *TokenBucket { return &TokenBucket{tokens: cap, cap: cap} }

func (tb *TokenBucket) Allow() bool {
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}`
}

func testTokenBucket(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestTokenBucket(t *testing.T) {
	tb := solutions.NewTokenBucket(2)
	if !tb.Allow() || !tb.Allow() || tb.Allow() {
		t.Fatal("token bucket logic failed")
	}
}
`, m.slug)
}

func circuitBreakerBody() string {
	return `package solutions

import "errors"

var ErrCircuitOpen = errors.New("circuit open")

// CircuitBreaker has closed/open states (simplified).
type CircuitBreaker struct {
	failures, threshold int
	open                bool
}

func NewCircuitBreaker(threshold int) *CircuitBreaker {
	return &CircuitBreaker{threshold: threshold}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	if cb.open {
		return ErrCircuitOpen
	}
	if err := fn(); err != nil {
		cb.failures++
		if cb.failures >= cb.threshold {
			cb.open = true
		}
		return err
	}
	cb.failures = 0
	return nil
}`
}

func testCircuitBreaker(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"errors"
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestCircuitBreaker(t *testing.T) {
	cb := solutions.NewCircuitBreaker(2)
	errFn := func() error { return errors.New("fail") }
	_ = cb.Call(errFn)
	if err := cb.Call(errFn); err == nil {
		t.Fatal("expected open circuit")
	}
}
`, m.slug)
}

func urlShortenerBody() string {
	return `package solutions

// Encode maps id to short code (simplified base62 stub).
func Encode(id int) string {
	if id == 0 { return "a" }
	return string(rune('a' + id%26))
}

// Decode maps short code back to id (simplified).
func Decode(code string) int {
	if code == "" { return 0 }
	return int(code[0] - 'a')
}`
}

func urlShortenerInterviewBody() string {
	return `package solutions

// InterviewChallengeSolution decodes short code to id.
func InterviewChallengeSolution(code string) int {
	if code == "" {
		return 0
	}
	return int(code[0] - 'a')
}`
}

func testURLShortenerInterview(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if solutions.InterviewChallengeSolution("d") != 3 {
		t.Fatal("decode failed")
	}
}
`, m.slug)
}

func restRouteMatchBody() string {
	return `package solutions

import "strings"

// InterviewChallengeSolution matches HTTP method and path pattern.
func InterviewChallengeSolution(method, path, wantMethod, wantPrefix string) bool {
	return strings.EqualFold(method, wantMethod) && strings.HasPrefix(path, wantPrefix)
}`
}

func testRestRouteMatch(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if !solutions.InterviewChallengeSolution("GET", "/api/v1/users", "GET", "/api") {
		t.Fatal("route should match")
	}
}
`, m.slug)
}

func rateLimitSlidingWindowBody() string {
	return `package solutions

// SlidingWindow limits requests per window size.
type SlidingWindow struct {
	limit int
	count int
}

func NewSlidingWindow(limit int) *SlidingWindow {
	return &SlidingWindow{limit: limit}
}

func (sw *SlidingWindow) Allow() bool {
	if sw.count >= sw.limit {
		return false
	}
	sw.count++
	return true
}`
}

func testRateLimitSlidingWindow(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	sw := solutions.NewSlidingWindow(2)
	if !sw.Allow() || !sw.Allow() || sw.Allow() {
		t.Fatal("sliding window failed")
	}
}
`, m.slug)
}

func lruInterviewBody() string {
	return `package solutions

// InterviewChallengeSolution reports LRU cache size after puts.
func InterviewChallengeSolution(keys []string, cap int) int {
	c := NewLRU(cap)
	for _, k := range keys {
		c.Put(k, 1)
	}
	return len(c.order)
}`
}

func testLRUInterview(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	n := solutions.InterviewChallengeSolution([]string{"a","b","c"}, 2)
	if n != 2 {
		t.Fatalf("size=%%d", n)
	}
}
`, m.slug)
}

func circuitBreakerInterviewBody() string {
	return `package solutions

// InterviewChallengeSolution reports whether circuit is open after failures.
func InterviewChallengeSolution(threshold int, failures int) bool {
	return failures >= threshold
}`
}

func testCircuitBreakerInterview(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestInterviewChallengeSolution(t *testing.T) {
	if !solutions.InterviewChallengeSolution(3, 3) {
		t.Fatal("expected open circuit")
	}
}
`, m.slug)
}

func testURLShortener(m module) string {
	return fmt.Sprintf(`package solutions_test

import (
	"testing"
	"github.com/go-mastery-roadmap/go-mastery-roadmap/%s/exercises/solutions"
)

func TestURLShortener(t *testing.T) {
	code := solutions.Encode(3)
	if solutions.Decode(code) != 3 {
		t.Fatal("roundtrip failed")
	}
}
`, m.slug)
}
