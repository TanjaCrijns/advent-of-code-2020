Day 18: Operation Order
---

<details>
  <summary>Puzzle explanation</summary>
  <br/>
As you look out the window and notice a heavily-forested continent slowly appear over the horizon, you are interrupted by the child sitting next to you. They're curious if you could help them with their math homework.

Unfortunately, it seems like this "math" follows different rules than you remember.

The homework (your puzzle input) consists of a series of expressions that consist of addition (+), multiplication (*), and parentheses ((...)). Just like normal math, parentheses indicate that the expression inside must be evaluated before it can be used by the surrounding expression. Addition still finds the sum of the numbers on both sides of the operator, and multiplication still finds the product.

However, the rules of operator precedence have changed. Rather than evaluating multiplication before addition, the operators have the same precedence, and are evaluated left-to-right regardless of the order in which they appear.

For example, the steps to evaluate the expression 1 + 2 * 3 + 4 * 5 + 6 are as follows:
```
1 + 2 * 3 + 4 * 5 + 6
  3   * 3 + 4 * 5 + 6
      9   + 4 * 5 + 6
         13   * 5 + 6
             65   + 6
                 71
```
Parentheses can override this order; for example, here is what happens if parentheses are added to form 1 + (2 * 3) + (4 * (5 + 6)):
```
1 + (2 * 3) + (4 * (5 + 6))
1 +    6    + (4 * (5 + 6))
     7      + (4 * (5 + 6))
     7      + (4 *   11   )
     7      +     44
            51
```
Here are a few more examples:
```
2 * 3 + (4 * 5) becomes 26.
5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 437.
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 12240.
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 13632.
```
Before you can help with the homework, you need to understand it yourself. Evaluate the expression on each line of the homework; what is the sum of the resulting values?
```
Your puzzle answer was 650217205854.
```
#### Part Two 
You manage to answer the child's questions and they finish part 1 of their homework, but get stuck when they reach the next section: advanced math.

Now, addition and multiplication have different precedence levels, but they're not the ones you're familiar with. Instead, addition is evaluated before multiplication.

For example, the steps to evaluate the expression 1 + 2 * 3 + 4 * 5 + 6 are now as follows:
```
1 + 2 * 3 + 4 * 5 + 6
  3   * 3 + 4 * 5 + 6
  3   *   7   * 5 + 6
  3   *   7   *  11
     21       *  11
         231
```
Here are the other examples from above:
```
1 + (2 * 3) + (4 * (5 + 6)) still becomes 51.
2 * 3 + (4 * 5) becomes 46.
5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 1445.
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 669060.
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 23340.
```
What do you get if you add up the results of evaluating the homework problems using these new rules?
```
Your puzzle answer was 20394514442037.
```
Both parts of this puzzle are complete! They provide two gold stars: **
</details>

<details>
  <summary>Comments about solution</summary>
  <br/>
Today was both fun and difficult, and it took me quite some time to come to a solution. I first was stuck trying to solve puzzle 1 using recursion, I had a solution that worked for all but one of the examples; the one that starts with double parentheses. I wasted a lot of time trying to solve it instead of trying to work with regexes, which in the end made my life a lot easier. I rebuilt everything with regexes, which made the overall code a lot shorter. I was able to reuse some code I wrote for evaluating strings as equations, as Golang unfortunately doesn't have anything to help you with that, but because we only needed to parse addition and multiplication, it was fine. I went through the same process for puzzle 2, I tried to rewrite my equation eval to change the operator precedence without regexes, but eventually it was a lot shorter and easier to just use regexes. 
  </details>
