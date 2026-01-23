## Introduction
**Are you currently working on a new programming language? Wondering how to make it more accessible to others? Wonder what to do next?** I hope you wrote documentation for it because without the documentation, not many people can use your language. Not only that, but users expect this documentation to be easily searchable and accessible. One `README.md` file on GitHub is not sufficient. Thankfully, tools like [VitePress](https://vitepress.dev/ "Link to the VitePress website") and [Docusaurus](https://docusaurus.io/ "Link to the Docusaurus website") make it relatively easy to write documentation websites.

![VitePress and Docusaurus](https://serhiicho.com/storage/posts/content/your-language-works-now-the-real-work-begins-1753513711.png)

Language and docs are great, but now you need at least syntax highlighting to start. What should we use for that, TextMate grammar? Tree-sitter? It depends on the editor you want to support. VSCode because of its popularity and maybe something like Neovim or Emacs for 10x developers.

For VSCode we can just write regex for the TextMate grammar, but for Neovim, which uses Tree-sitter, we need to write a Tree-sitter parser and query files that look something like S-expressions. You’ll likely want an LSP for advanced features like autocompletion.

My point is that **after creating a language, more effort is required to build the surrounding ecosystem**. If the language you are writing is something small and insignificant, it might not matter, but if your creation is something complex, it requires an ecosystem around it.

Imagine you've just created a revolutionary new programming language, but without the right tools and documentation, it might as well be a secret code known only to you. The ecosystem improves usability, enhances developer experience, and can attract more users to your language.

The ecosystem is precisely what I’ve been spending the past 6 months building, which includes:

* [VSCode extension](https://github.com/textwire/vscode-textwire "Link to Textwire VSCode extension on GitHub")
* [Neovim plugin](https://github.com/textwire/textwire.nvim "Link to Textwire.nvim plugin on GitHub")
* [Tree-sitter parser](https://github.com/textwire/tree-sitter-textwire "Link to Textwire Tree-sitter parser on GitHub")
* [Language Server Protocol](https://github.com/textwire/lsp "Link to Textwire LSP on GitHub")
* [A documentation website](https://github.com/textwire/textwire.github.io "Link to Textwire docs on GitHub")
* [The language itself](https://github.com/textwire/textwire "Link to Textwire repo on GitHub")

I'll go through some challenges and learnings I experienced after building [Textwire](https://github.com/textwire "Link to Textwire GitHub"), a template language for writing HTML templates in Go. Here is what I've learned...

## This Article is For
Planning to build a programming language or a program with custom syntax? If that's the case, then you are at the right place. This article will be good information for your next project because I was exactly in your shoes. **I'll share practical advice on building an ecosystem around your language, including pitfalls I've encountered.**

Keep in mind that this article doesn't require you to have an advanced understanding of programming, but if you are planning to build your own programming language, advanced skills are necessary for building a language. Not for reading this material.

I'm not going into small details on how to build your programming language; this article is focused more on the ecosystem itself rather than a particular program. **It covers tools and resources for syntax highlighting, editor support, documentation, and other essential components.**

Tip. I have an article, [Writing Your Own Lexer with Simple Steps](/posts/writing-your-own-lexer-with-simple-steps "Link to the article"), which is very popular on this platform. Check it out if you are interested in that.

## Why Ecosystems Are Important
**No language survives on syntax alone—it thrives through its ecosystem.** It’s not just about writing code; it’s about creating a community and tools that make your language more accessible and usable. Without an ecosystem, your language is just a set of syntax rules—challenging to learn and use.

A good example is PHP’s success, which took off thanks to a strong ecosystem built around it. Rasmus Lerdorf, the creator of PHP, said in his [talk at phpCE in 2018](https://youtu.be/SvEGwtgLtjA?si=2_TtH_uJp6LF3LoL&t=699 "Link to a YouTube video"):

> I and we, later on when the project became bigger, we spent a lot of time on ecosystem and making sure that everything worked together and that we could scale up and also more importantly scale down and make it really really accessible to new users coming to the language.
> ###### Rasmus Lerdorf (the creator of PHP)
> ![Rasmus Lerdorf](https://serhiicho.com/storage/posts/content/how-to-make-your-software-developer-job-more-enjoyable-7.png)

As he later emphasized in another talk, "The ecosystem is what made PHP successful at the time". Modern examples like Swift and .NET are more than just languages or runtimes, they are ecosystems with libraries, frameworks, and tools that make them more usable.

**Technically revolutionary languages like ALGOL 68 still failed** because:

- **No standard libraries** (required OS-specific implementations)
- **Compiler disagreements** (split the tiny community)
- **Almost no tooling** (no debuggers, package managers, or IDE support)

Despite innovations like parallel processing, ALGOL 68’s fragmented tooling and lack of standard libraries made it impractical compared to rivals like C or Fortran.

While we're focusing on the minimal viable ecosystem here, these lessons scale - whether you're building a niche Domain-Specific Language (DSL) or the next Rust.

## Writing a Language Is Tough
Designing a programming language is inherently complex. There is a reason why people who work on programming languages are usually programmers with years of experience. **Building a programming language is an arduous and time-consuming task.** One overlooked challenge: error messages with proper line numbers. A bad parser error can frustrate users for years.

Tutorials make it seem easy—because they’ve done the hard work for you. Don't listen to those pseudo gurus (myself included) on YouTube and Twitch saying that you just need to read a special book that will teach you everything.

Books are a starting point, but mastery requires grinding through your mistakes. I’m no expert—just someone who’s wrestled with these problems firsthand. And that’s the only way to truly learn. So, I've started my journey with writing and designing a language.

## 1. Writing a Language
You define syntax and decide on semantics. Doing thorough preliminary work will make your life easier later. Depending on your specifics, you might want to write a tokenizer, a parser, a compiler or interpreter.

* **A tokenizer (lexer) converts source text into structured tokens like keywords, identifiers, and literals.** Usually, tokens are just constants. For example, for, `x = 5 + 3`, the tokenizer outputs something like:
    ```javascript
    [IDENT(x), EQUALS, NUMBER(5), PLUS, NUMBER(3)]
    ```
* **A parser receives a stream of tokens and outputs a recursively generated abstract syntax tree (AST for short).** The AST is a hierarchical representation of your program’s structure, where each node describes an operation or statement. In the example `x = 5 + 3`, our AST node can be:
    ```javascript
    AssignStatement(
        target: 'x',
        value: BinaryExpression(
            operator: '+',
            left: 5,
            right: 3,
        ),
    )
    ```
    Program has statements; one of those statements is the assign statement, which contains all the information about the statement.
* **An evaluator is the component that executes the code.** The evaluator either walks the AST to execute code directly (interpreter) or lowers it to bytecode/machine code (compiler). It's also responsible for handling the semantics of your language, such as variable scope, function calls, and control flow.

Note. Sometimes developers use the word "interpreter" interchangeably with "evaluator", though technically the interpreter is the whole system (lexer + parser + evaluator). The evaluator is the part that executes the code.

These components are fundamental in transforming human-readable code into machine-executable instructions. **The tokenizer breaks down the source code into a stream of tokens, which are then parsed to form an AST. The AST will be fed into the evaluator, which executes the code or generates low-level code.**

If your language (like [Textwire](https://github.com/textwire "Link to Textwire GitHub")) is a combination of multiple syntaxes, it will be more complex. In my case, I had Textwire embedded into HTML, which meant I could just treat all HTML code as a string and parse it as such.

80% of all the complexity that I had to deal with was the fact that Textwire is tied to HTML. Handling ambiguous syntax like `<span @if>Text</span>` required custom parsing rules. The code example below is just HTML, there is no Textwire syntax in it:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <title>@myTitle</title>
</head>
<body>
    @content
</body>
</html>
```

Notice how `@myTitle` and `@content` could be either Textwire directives or just part of HTML text - this ambiguity forces complex parsing rules.

My parser needs to be able to differentiate between things like `@content` and valid Textwire directives like `@if`, `@for`, etc. This is where the complexity arises, especially when your language is elegant and easy to read. The simpler the syntax, the more complex the parser becomes.

Parsing `@if(true) { @html('<span>Text</span>') }` is so much easier than parsing `@if(true)<span>Text</span>@end`, which is what Textwire does. The main motto of Textwire is "Easy for you, hard for the parser".

If you have something similar, you will have to deal with the same issues because there are not many resources online about this topic. But no matter how elegant your language design, it's worthless without documentation that makes it accessible. Let's discuss how to create docs that users will love.

## 2. Writing Documentation
**I recommend writing documentation alongside your language—not after.** I find it to be a lot easier because you document programming structures first and implement them afterward. Writing docs first forces you to clarify designs upfront—giving you a reference to guide implementation.

**Documentation is crucial for adoption. It’s not just about explaining syntax; it’s about guiding users through your language's features, error handling, best practices, and common pitfalls.**

For example, documenting how `@if` handles edge cases (like `@if(nil)`, `@if("")`, `@if([])`) before implementing it prevents messy surprises later.

> Documentation is a love letter that you write to your future self.
> ###### Damian Conway (a computer scientist)
> ![Damian Conway](https://serhiicho.com/storage/posts/content/your-language-works-now-the-real-work-begins-1753606371.jpg)

**Effective documentation transcends syntax.** It provides:  
- Best practices and idioms
- Clear error handling guidance
- Common pitfalls to avoid
- A reference for the future self

The [Textwire docs](https://textwire.github.io/ "Link to the Textwire documentation") have proven invaluable for both users and myself, serving as both reference and design compass during development.

I recommend using documentation generators like [VitePress](https://vitepress.dev/ "Link to the VitePress website") and [Docusaurus](https://docusaurus.io/ "Link to the Docusaurus website"). Typically I reach for VitePress but Docusaurus is better for complex docs.

![VitePress and Docusaurus](https://serhiicho.com/storage/posts/content/your-language-works-now-the-real-work-begins-1753614151.webp)

Both generate static sites from Markdown, which is easy to read and write. Markdown keeps docs version-controllable and collaborative, allowing you to write in a familiar format. You can use any text editor to write Markdown, and it’s easy to convert to HTML.

**You can then upload your docs to GitHub Pages, which will give you a free subdomain.** For personal projects, GitHub Pages gives you `torvalds.github.io/linux`, but an organization account offers cleaner URLs like `linux.github.io`. As an option, you can attach your domain to it, which is not difficult to do, and there are plenty of resources online that explain how to do it.

Note. Even basic docs put you ahead of most language projects because most projects have only a `README.md` file in their repo. For small projects (like [Tiny Logger](https://github.com/tiny-logger/tiny-logger "Link to GitHub repo tiny-logger")) this is enough, but for larger projects you should have a dedicated documentation website.

But adoption requires more—starting with editor support. Let’s dive into syntax highlighting.

## 3. Syntax Highlighting
**Syntax highlighting isn’t just cosmetic—it’s essential for code comprehension and productivity.** It helps users quickly understand code structure, making it easier to spot errors and navigate complex files. Without syntax highlighting, your language feels like a wall of text—intimidating and difficult to read.

You can still use any language without syntax highlighting; in fact, developers used many languages without syntax highlighting for years. Nowadays, **users expect syntax highlighting to be available for any language they use**. So how do you implement it? Let’s compare TextMate grammars (regex-based) versus Tree-sitter (AST-based) without going deep into the details.

### TextMate Grammars
**TextMate grammars use regular expressions to identify syntax patterns and are defined in a single JSON file.** They're supported across editors like VSCode and Sublime. It's pretty straightforward to write a TextMate grammar if you know regular expression.

Tip. What I found useful for me was to see how other languages implement TextMate and Tree-sitter grammars. You can find many examples on GitHub. I suggest looking on languages that are similar to yours, as they will give you a good starting point. In my case, I looked at how PHP highlighting is implemented because it also has HTML embedded into it.

Here is a simple example of a `tmLanguage.json` file that you use for TextMate grammar, which supports syntax highlighting only for integers and floats:

```json
{
    "$schema": "https://raw.githubusercontent.com/martinring/tmlanguage/master/tmlanguage.json",
    "scopeName": "text.html.basic.myLang",
    "name": "MyLang",
    "fileTypes": ["mylang"],
    "patterns": [
        {
            "name": "constant.numeric.integer",
            "match": "\\b\\d+\\b"
        },
        {
            "name": "constant.numeric.float",
            "match": "\\b\\d+\\.\\d+\\b"
        }
    ]
}
```

The only good reference I found for TextMate grammars is [Syntax Highlight Guide](https://code.visualstudio.com/api/language-extensions/syntax-highlight-guide "Link to Tree-sitter documentation") for VSCode. That is what I was using to learn how to write TextMate grammars for Textwire.

### Tree-sitter Parsers
**Tree-sitter uses a parser to build an AST, which is then used for syntax highlighting.** It’s more powerful and accurate, but sometimes requires writing a parser in C language. While Tree-sitter parsers compile to C, you’ll usually write them in a JavaScript DSL.

Tip. **DSL stands for Domain Specific Language, which is a language designed to be used for a specific task.** In the context of Tree-sitter, the DSL is used to define the grammar rules for the parser. DSL is a general term that basically refers to a mini-language tailored for specific tasks. Like SQL for databases or Regex for text processing. Fun fact: The Tree-sitter CLI converts this DSL to optimized C code—you’ll see `grammar.c` in your output!

Here is an example of `grammar.js` that you would probably write:

```javascript
// Enables Tree-sitter’s DSL for rule definitions
/// <reference types="tree-sitter-cli/dsl" />

module.exports = grammar({
  name: 'MyLang',

  rules: {
    // The root rule - matches one or more numbers
    program: $ => repeat1(
        choice(
            $.integer_literal,
            $.float_literal,
        ),
    ),

    integer_literal: _ => /\d+/,
    float_literal: _ => /\d+\.\d+/
  }
})
```

The example above is a simple Tree-sitter grammar that matches integers and floats. As you can see, our program consists of one or more literals, which can be either integers or floats. The `integer_literal` and `float_literal` rules define how to match those literals using regular expressions.

Later, this JavaScript code will be compiled into a parser that can be used by Tree-sitter. It might seem simple at first, but it can get complex quickly, especially if your language has a lot of weird syntax rules. You can read more about Tree-sitter in the [Tree-sitter documentation](https://tree-sitter.github.io/tree-sitter/ "Link to the Tree-sitter documentation").

But what about the editor support? Tree-sitter is supported by editors like Neovim, Atom, and Textadept. It's more complex to set up, but it provides better performance and accuracy.

My use case was very different from what the usual language syntax is; thus, I spent a lot of time writing a Tree-sitter parser for Textwire. I had to write a custom parser that could handle HTML and Textwire syntax together, which was not easy. The parser needs to be able to differentiate between HTML tags and Textwire code, which is not a trivial task.

In theory, Tree-sitter is better for complex languages, but it requires more effort to set up. If you are writing a simple language, TextMate grammars are easier to implement and maintain. But you are not dictating the rules here; the editor you want to support is the one that decides which syntax highlighting method you should use. VSCode uses TextMate grammars, while Neovim uses Tree-sitter.

In your case, you might want to support other editors like Emacs or JetBrains IDEs, which also use TextMate grammars. While Tree-sitter's editor support is growing thanks to plugins and extensions for different editors, TextMate remains the safest bet for broad compatibility.

With syntax highlighting handled, we face the next hurdle: making your language smart with autocompletion and diagnostics (via LSP).

## 4. Language Server Protocol
We already talked about writing your language, documentation, and syntax highlighting for it, but now we need to make it smart. LSP stands for Language Server Protocol, which is a standardized way to provide language features like autocompletion, go-to-definition, and diagnostics (error checking) across different editors.

LSP separates language intelligence (your server) from editor UI (client), so you write features once and support VSCode, Neovim, etc., simultaneously.

I'm not going to go deep into details about LSP, because the best description of LSP is on this official [Wikipedia Page](https://en.wikipedia.org/wiki/Language_Server_Protocol "Link to Wikipedia") that explains it better than me.

What I do want to say is that if you don't have experience in building LSP (at least one), then it's going to be a journey. At the time I was building an LSP for my Textwire language, there were not many resources online that I could find to fit my needs. The best resource I found was the [YouTube video of TJ DeVries](https://youtu.be/YsdlcQoHqPY?si=yQoS2_692DkokH1k "Link to YouTube video of TJ DeVries").

The hardest part? I don't think it was writing the LSP itself, but rather understanding how to implement it correctly. I had to read the [LSP specification](https://microsoft.github.io//language-server-protocol/specifications/specification-current "Link to LSP specification") and figure out how to implement it in Go. But overall, it wasn't as hard as I thought it would be.

There are different tools that can help you with LSP, but I didn't use any; I just wrote everything from scratch, and I think you should too.

The most enjoyable thing about writing your LSP is that you can use any programming language that can be compiled to a binary. Many programmers choose C, C++, Zig, Rust, and Go.

While Go’s simplicity got me started, C/C++ offer finer control over memory and performance—critical for large codebases. Moreover, **I respect the community around C and C++ and Dennis Ritchie.** His impact on the programming world is enormous. We can conclude this section with a quote from him that fits perfectly here:

> The only way to learn a new programming language is by writing programs in it.
> ###### Dennis Ritchie (a legendary computer scientist)
> ![Dennis Ritchie, creator of C, at a computer](https://serhiicho.com/storage/posts/content/your-language-works-now-the-real-work-begins-1753806109.webp)

Next, we’ll explore how these pieces fit into an editor plugin and what challenges you might face.

## 5. Editor's Support
**Great editor support transforms your language from a curiosity to a tool people choose to use daily.** Users expect a seamless experience with features like syntax highlighting, autocompletion, and error checking. Without these, your language feels incomplete and unpolished.

![The list of various VSCode Extensions](https://serhiicho.com/storage/posts/content/your-language-works-now-the-real-work-begins-1753868143.webp)

For Textwire, adding autocompletion and hover information for directives saves user time by giving them detailed documentation right in the editor.

Note. For small projects, basic support (like a VSCode syntax extension) may suffice. But ambitious languages demand full editor integration to shine. In my case, the language is indeed small, but I wanted to go all the way and build an ecosystem around it for the sake of learning and experience.

In my case, I went for 2 editors and built the [VSCode extension](https://github.com/textwire/vscode-textwire "Link to Textwire VSCode extension on GitHub") and the [Neovim plugin](https://github.com/textwire/textwire.nvim "Link to Textwire.nvim plugin on GitHub").

You might think that if you build a VSCode extension then a Neovim plugin will be easier, but it's not. Those projects are entirely different; Neovim uses Lua for plugins, while VSCode uses TypeScript. The only thing that is similar is the LSP, which is used by both editors.

Here is a small Lua code snippet that shows how to register a Tree-sitter parser in Neovim:

```lua
local parser_config = require("nvim-treesitter.parsers")
    .get_parser_configs()

parser_config.textwire = {
    install_info = {
        url = "https://github.com/textwire/tree-sitter-textwire",
        files = { "src/parser.c", "src/scanner.c" },
        branch = "master",
        -- if stand-alone parser without npm dependencies
        generate_requires_npm = false,
        -- if directory contains pre-generated src/parser.c
        requires_generate_from_grammar = true,
    },
}
```

The most difficult part of writing plugins was understanding the API of each editor and how to implement the features correctly. For VSCode, there are many resources available, but for Neovim, I had to use a “trial and error” approach in some cases.

**For Neovim, I wasted many hours debugging why syntax highlighting doesn't work.** That was the most difficult thing to understand because there are not many materials online that go into details of how Tree-sitter actually works in Neovim. That's a big thumbs down for Neovim.

Even right now, while I'm writing this article, I've noticed my syntax highlighting disappeared in Textwire files after moving my Neovim config from macOS to Linux Fedora. The whole Neovim workflow for me was a challenge; I spent too much time on it.

VSCode, surprisingly, was a lot easier to implement. The VSCode extension API is well-documented, and there are many examples available online. I was able to implement syntax highlighting, autocompletion, and hover information without much trouble.

Anyway, let's look back at everything we've discussed and summarize the key points. Plus, I want to give you my conclusion on this topic.

## Conclusion
What started as a language design challenge became **a lesson in ecosystem building**—where documentation, tooling, and editor support consumed far more time than the core interpreter.

**What I've discovered throughout that time is that building a programming language is not just about writing code; it’s about creating an ecosystem that makes your language more usable and accessible.**

If you build a community around your language before building all the surrounding tools, you will have a better chance of success. The community will help you with feedback, bug reports, feature requests, spreading the word about your language, and attracting more users.

If your language is great and people like it, they will help you with building the surrounding ecosystem. But if your language is for a small niche, you might not need a community at all. In that case, you can just build the tools yourself and use them for your projects.

That's what I did with Textwire. I built the tools myself and used them for my projects. I didn't have a community, but I still managed to build a small ecosystem around my language. Likewise, I hope this article will help you to build your ecosystem around your language and make it more usable and accessible.

This is your invitation. Whether you’re building for thousands or just your toolbox, remember: **the ecosystem is the product**. Now go make something that, at least, you will enjoy using.
