# compiler_design
This Repo is to learn about compiler design and learn its implementation.

First I am starting with desinging a interperter in go language. 

To do list
    1. Our lexer only supports ASCII characters instead of full Unicode range. 
        a. In order to suport Unicode and UTF - 8 we need to change lexer character(l.ch) from byte to rune (int32)
    2. Right now we are only reading integers in number need to extend support for float as well as double and other things. Atleast do float.
