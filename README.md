# compiler_design
This Repo is to learn about compiler design and learn its implementation.

First I am starting with desinging a interperter in go language. 

let five = 5;
let ten = 10;

let add = func(x,y){
    x + y;
};

let result add(five, ten);

numbers are integers -> this is a type.
veriable names are a type
idenifiers are a type. (idenifiers are actually part of the language and they are called keywords.)

Defining the token data structure. 
    1. We need a type. (to distinguish between "integer" and "the right bracket")
    2. We also need a field to hold the literal value of the token.
    3. we are taking illegal and EOF. These two we are taking as extra to help our parser
    4. Usually token contians line number and file name to easily identify the error.

Lexer
    We are gonna write our own lexer. This wil take code as input and it is gonna give tokens as output.
    This gonna contain only one function nextToken() which will help in getting the token. In lexer we have read input and get current position while reading.
    A function is required to read character. This function will give the next charater and advance our position in the input.
    Lexer only supports the ASCII characters instead of full unicode range.

    Extending our token set and lexer
        a. We have added a function to read words and check if it is a idenifier or is it a keyword. 
        b. We have added a function to read integers
        c. We have added more operators other than add
        d. We are adding peekChar to the lexer, just in case we want to peek in the character ahead


Note: Difficulty of parsing different language often comes down to how far can you peek ahead and backwards
Start of a REPL
    a. Read Eval Print loop: reads input, sends it to interpreter for evaluation, prints the result/ouput of the interpreter.

Now we need to parse the token
    a. It takes input and builds a data structure - often some kind of parse tree, abstract syntax tree
    b. We need the data structure to present the input
    c. In most programming, the data structure that is used is "abstract syntax tree"
    d. This process of parsing is also known as syntactic analysis
    e. There are pre build parser generators

AST: There are no parenthesis, command, semicolons and other things it is pretty abstract in nature.
Context free grammer: The set of rules to describe how to form correct sentences in a language.
    ECMAScript(javascript) this is described using Backus - Nuar form (a notation used to describe the syntax)

We are gonna write our own parser. Usually are there are different types of parsers top down, bottom up, recursive parser and others


In our language, we are using let for binding. let <identifier> = <expression>
Remember expression create values not the statements. let x  = 5; does not create a value but 5 does.
    a. So out AST will have an identifier and expresssion, and a Node which is gonna contain tokenLiteral()
    b. So our AST is just list of statements. 
    c. For our variable binding we need a name of the variable, and we also need a field that points to the expression on the right side of the variable and the expression can be anything.

There are statements and then there are expression. In this language there are only two statements and rest are just expressisons and we do not need to do anything for the expressions.
Writing a parser(We are creating parser for let,return and other things one by one)
    a. We are using 2 pointer kind of approach, we have current token and next token, we requrire nextToken in order to identify if after 5 (INT) there is semicolon or there is a arithematic expression
    b. Parser contains three things one is lexer(in order to get the tokens), we need current token and we need next token
    c. We have a parse program, we are using recursive descent parser
        c.1: Idea behind the parse program is to create a node and we need to identify the token and based on the token we are going to proceed, like if there is let token or if there is IF token or else if there is return token. 
        c.2: Based on the current token we need to identify if the next token is correct or not.
    d. We are creating test case for the parser. 
        d.1: Idea is pretty much the same which is to get input 
        d.2: Create ast from the input and check if the ast is correct or not
            d.2.i: We are creating ast for let statement, return statement and other things
        d.3: Creating a parse program based on the first token we will create a switch statement for "let", "return" etc
        d.4: Create a function for parse let statement and create a parse function for return statement
        d.5: Based on the parse/return function we are making test to chech weather it is correct or not.
        d.6: We are creating a expectedToken() function where we are going to check the validity of the next function
        d.7: We also have error string in our parser which is going to help us with the debugging
        d.8: We need to write parser for the expression which is not that straight forward. 
            d.8.i: We need to take care of the operator precedence
            d.8.ii: Same type can appear in mulitple positions. Validity of a token's position now depends on the context. It heavily depends on the language and what operators and what syntax that you have allowed for the programming language.
            d.8.iii: We are write a String function in our ast so that it will become easy to debug

To do list
    1. Our lexer only supports ASCII characters instead of full Unicode range. 
        a. In order to suport Unicode and UTF - 8 we need to change lexer character(l.ch) from byte to rune (int32)
    2. Right now we are only reading integers in number need to extend support for float as well as double and other things. Atleast do float.
