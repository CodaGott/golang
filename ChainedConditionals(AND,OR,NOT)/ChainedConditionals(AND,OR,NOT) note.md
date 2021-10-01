# Chained Conditionals (AND, OR, NOT)

GOLANG like many programming languages has 3 chained conditional operators.

These Operators are written like this in a code.
    && = "AND"
    || = "OR"
    ! = "NOT"

    When using "||" both values has to be false for the result to be false every other value is true.

    When using "&&" both has to be true for the result to be true every other value is false.

Chained Condition Precedence.
    The order of precedence is: logical complements (!) are performed first, logical conjunctions (&&) are performed next, and logical disjunctions (||) are performed at the end.