Discover how modern PHP embraces strong typing, from core types to advanced concepts. Master type safety, avoid type juggling pitfalls, and boost code quality with static analysis and more.

## Introduction
PHP has significantly evolved, with one of the most notable changes being its shift to strong typing. It was always a loosely typed language like Python and Ruby, allowing developers to write code without explicitly defining types. However, with new versions of PHP, the language has introduced several features to support strong typing.

In this comprehensive guide, we'll explore the PHP type system in depth, covering core types, type juggling, and static analysis tools for type safety. I'm sure you'll find something new, even if you are an experienced PHP developer.

By the end of this guide, you'll have a solid understanding of PHP's type system and how to leverage it to write more robust and maintainable code. If you are still someone who is not convinced about the importance of types in PHP and still prefer the dynamic typing, this guide will help you understand the significance of types in modern PHP development and teach you how to work with them effectively.

By mastering PHP's type system, you'll write more robust and maintainable code and strengthen your skills in crafting reliable software. Let's dive in!

### Types are Optional in PHP
Types (or type hints) in PHP are not mandatory. You can write PHP code without using types, and it will work just fine. We all have written PHP code without types at some point, if you have been working with PHP for a while.

There is nothing wrong with not using types in PHP, it all depends on the context and the project you are working on. Don't feel that you are doing something wrong if you are not using types in PHP. Everything discussed in this article is to help you understand the importance of types in PHP and how to work with them effectively.

If you would rather not use types and static analysis, that's perfectly fine. But if you are working on a large project with a team of developers, using types and static analysis tools can help you catch bugs early, improve code quality, and make your codebase more maintainable.

<p class="tip">
If you want to lean more about PHP, refer to the <a href="https://laracasts.com/referral/SerhiiCho" target="_blank">Laracasts</a>, the best place to learn PHP, Laravel, and modern web development. I've been using it for years, and it's the best investment I've made in my career.
</p>

### The Shift to Strict Typing in PHP
PHP has come a long way since its first release in June 1995, and now, it's the most used server-side language on the web.

One of the most significant changes in PHP's evolution is the shift to strong typing. With the release of PHP 7 in December 2015, PHP introduced scalar type declarations and return type declarations. This was a significant step towards making PHP a more robust and reliable language.

Since PHP doesn't have a compile step, so types are checked at runtime. In most languages like Python, Java, and C#, types are checked at compile time, which means you can catch type errors before running your code. But in PHP, you can only catch type errors when you run your code. To fix that, we need to add static analysis tools to our workflow. We'll talk about that later in this guide in the <a href="#static-analysis-for-type-safety">Static Analysis for Type Safety</a> section.

#### Coercive Typing
There are 2 ways you can go about using types in PHP: “coercive typing” and “strict typing”. Coercive typing is the default behavior in PHP, where the type of variable is automatically converted to the expected type. How it works is, if you pass a string to a function that expects an integer, PHP will automatically convert the string to an integer (which is not always a good thing).

```php
<?php

function add(int $a, int $b): int
{
    return $a + $b;
}

echo add('1', '2'); // Output: 3
```

It's a default behavior in PHP not because it's a preferred way of working with types, but because it's backward compatible with older versions of PHP. PHP's Core team doesn't want to break millions of existing PHP applications by changing the default behavior of type coercion.

#### Strict Typing
Strict typing, on the other hand, is a way to enforce types in PHP. When you enable strict typing, PHP will throw a `TypeError` if you pass the wrong type to a function or return an invalid type from a function. This is a more reliable way of working with types in PHP.

All you need to do to enable strict typing in PHP is to add the following line at the beginning of your PHP file:

```php
<?php

declare(strict_types=1);
```

And yes, you need to add this line at the beginning of every PHP file where you want to enable strict typing. It's not a global setting that you can enable in your php.ini file.

If we take the example from above and enable strict typing, PHP will throw a `TypeError` because we are passing a string to a function that expects an integer:

```php
<?php

declare(strict_types=1);

function add(int $a, int $b): int
{
    return $a + $b;
}

echo add('1', '2'); // Fatal error: Uncaught TypeError: add(): Argument #1 ($a) must be of type int, string given...
```

As a beginner, you might find strict typing a bit overwhelming, but trust me, it's a good practice to enable strict typing in your PHP projects. It will help you catch bugs early and improve code quality.

Don't mistake this for strict mode in JavaScript. They are not the same thing. In JavaScript, you enable strict mode like this:

```javascript
"use strict"
age = 3.14 // This will cause an error because 'age' is not declared
```

I personally use strict types in all my PHP projects from the day I've discovered it. But, I must admin, it's not easy to force strict types in the existing codebase. Unless you have a good test coverage.

## Core PHP Types
As for a non-compiled language, PHP has a rich set of core types that you can use as parameters, return, property, constant, and enum types. Let's explore each of them to understand how they work and when to use them.

Here is the list of core types in PHP:

- `int`: Represents an integer value like `42` or `-42`.
- `float`: Represents a floating-point number like `3.14` or `-3.14`.
- `string`: Represents a sequence of characters like `'Hello, World!'`.
- `bool`: Represents a boolean value like `true` or `false`.
- `mixed`: Represents a value of any type.
- `never`: Represents a value that never occurs.
- `array`: Represents an array of values.
- `object`: Represents an instance of a class.
- `callable`: Represents a callable function or method.
- `iterable`: Represents an iterable value like an array or an object that implements the `Traversable` interface.
- `resource`: Represents a resource like a file handle or a database connection.
- `null`: Represents a null value.
- `void`: Represents a function that doesn't return a value.
- `self`: Represents the current class.
- `static`: Represents the class where the method is called.
- `parent`: Represents the parent class.

### Understanding Parameter, Return, and Property Typing

## Type Errors in PHP
— Passing the wrong data type to a function with strict typing enabled.
— Returning an invalid type from a function with a declared return type.
— Assigning an incompatible value to a typed property.

## Type Juggling

### Demystifying Type Juggling: Definition and Impact
### Best Practices for Avoiding Common Type Conversion Pitfalls

## Advanced Types: Mixed and Never

### Exploring the Role of Mixed and Never in PHP
### Practical Use Cases for Mixed and Never Types

## Static Analysis for Type Safety

### Enhancing Type Safety with PHPStan and Psalm

Psalm is another popular static analysis tool for PHP that focuses on type safety. As they describe it, Psalm is a free and open-source static analysis tool that helps you identify problems in your code, so you can sleep a little better.

From what I've seen, it's very similar to PHPStan, almost identical in terms of features and functionality. Although, Psalm has a few extra features like SQL injection detection and other security-related checks that PHPStan doesn't have.

But, PHPStan was recommended by <a href="https://en.wikipedia.org/wiki/Rasmus_Lerdorf" target="_blank">Rasmus Lerdorf</a> himself in one of his talks, so I went for PHPStan from the beginning.

### Leveraging Static Analysis Tools to Improve Code Quality

<p class="tip">
If you want to help the PHP community, one way of doing it is by supporting it financially. You can sponsor PHP on <a href="https://opencollective.com/phpfoundation" target="_blank">Open Collective</a>. Your money will be used to fund the development of the PHP language.
</p>

## Conclusion
// todo: Mastering the PHP Type System