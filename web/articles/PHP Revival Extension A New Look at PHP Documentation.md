# PHP Revival Extension: A New Look at PHP Documentation

## Introduction

As a software engineer focused on solving problems and building tools to make others' lives easier, I noticed that while PHP continues to evolve with new features, its official documentation site remains outdated. This disconnect inspired me to take action.

Even developers who have worked with PHP for years wish to see an update to improve the overall user experience. **Why hasn't anyone tried to improve the documentation website** and make a contribution back to the PHP community? That's the question I was asking myself in July 2019.

It seemed odd that there are so many talented developers who use PHP daily, but none of them have tackled the issue of the outdated documentation website, despite the importance of the language.

What if they did? Perhaps they tried, but that's where my story begins.

Note. Why user experience is important? User experience (UX) is crucial because it directly impacts how users interact with your website or application. A positive user experience can lead to increased user engagement, higher conversion rates, and improved customer satisfaction. By focusing on improving the user experience, we can create a more enjoyable and effective product that meets the needs of your users.

## Initial Attempt: Documentation Website Redesign

In July 2019, I took it upon myself to start working on [pull request #289](https://github.com/php/web-php/pull/289 "Link to the pull request"), which would improve the look of the [www.php.net](https://www.php.net "Link to PHP documentation") website. I decided to start small, focusing on the outdated navigation bar and, in my opinion, the poor color scheme.

After several iterations, I came up with a new design that would improve the overall user experience. Below, is an image of the proposed changes:

![https://serhiicho.com/storage/posts/content/php-revival-extension-a-new-look-at-php-documentation-1728463497.png](../../../storage/posts/content/php-revival-extension-a-new-look-at-php-documentation-1728463497.png)

After 40 days of waiting for the PR to be reviewed, I received feedback. It stated that my style changes to [www.php.net](http://www.php.net "Link to PHP documentation") would affect sites like doc.php.net, news.php.net, people.php.net, etc. because they share the same CSS files.

Realizing that modifying the style for [www.php.net](https://www.php.net) while maintaining compatibility with all the other sites wasn't feasible, I sought an alternative solution.

## Customizing the Documentation Experience

At that point, I decided to create a browser [extension](https://php-revival.github.io/ "Link to the PHP Revival landing page") to let users experience an improved version of `www.php.net` without impacting other PHP websites on the same domain. An additional goal was to provide users with features such as code copying, code execution, and embedded PHP-related videos on the homepage.

![Maximize your PHP learning with our browser extension that enhances your experience on www.php.net](../../../storage/posts/content/php-revival-extension-a-new-look-at-php-documentation-1728832143.png)

The extension was getting attention, the feedback was overwhelmingly positive, and it was motivation to continue working on it.

PHP Revival is a browser extension that enhances the official PHP documentation website. It improves the syntax highlighting to code examples, enables code copying, allows PHP code evaluation, and provides quick access to PHP-related links and videos.

Here are some features of PHP Revival:

*   Improves site styling
*   Applies a dark theme with syntax highlighting to code examples
*   Enables copying of code examples
*   Allows evaluation of PHP code examples
*   Adds useful PHP-related links to the homepage
*   Includes a section of PHP-related videos on the homepage
*   Provides quick access to [www.php.net](http://www.php.net "Link to PHP documentation") with a single click on the extension icon

Looking ahead, I plan to introduce more features and enhance the existing ones. I welcome any suggestions and feedback from the community to help shape the extension's development.

> This works great, really improves the documentation by being able to run all the examples. Thank you!
>
> ###### J Post
>
> <img src="https://php-revival.github.io/assets/J%20Post-8hISC9Z6.jpeg" alt="J Post" />

## Behind the Code: Technical Challenges and Solutions

The PHP Revival extension is built using modern web technologies such as SASS, TypeScript, and Webpack. The extension is compatible with Chromium-based browsers like Brave, Edge and Chrome, and also supports Firefox.

![Some browsers that support the PHP Revival extension](../../../storage/posts/content/php-revival-extension-a-new-look-at-php-documentation-1728832153.png)

For everyone who is interested, I want to share some of the technical challenges I faced while developing the extension and how I overcame them.

### Challenge 1: Override CSS Styles

**One of the technical challenges I faced was the ability to override some CSS styles on the official PHP documentation with the `!important` rule.** If you ever watched a CSS tutorial, you know that using `!important` is considered bad practice, but on www.php.net there are plenty of style rules that use the `!important` flag.

```
 .example {
    color: red !important;
}
```

There are several ways to override CSS styles with the `!important` flag, but the most common are either using JavaScript to inject styles directly into the HTML or applying the `!important` flag in CSS. Since browser extension stylesheets have higher specificity than the website stylesheets, I decided to go with the second option because it was easier to implement and maintain.

Tip. Higher specificity in CSS means that the styles in your extension's stylesheet will take precedence over the website's existing styles, allowing you to effectively override them without needing complex workarounds. This is particularly beneficial in the context of browser extensions because it ensures that your custom styles will be applied consistently across the site without being unintentionally overridden by the original site's CSS.

### Challenge 2: Modify DOM Elements

**Another challenge was modifying DOM elements on specific pages without the user noticing.** Since my JavaScript executes after the DOM has fully loaded, any manipulation would briefly appear to the user, causing a flicker effect due to visible changes happening in real time.

I'm still working on fully resolving this issue, but the best solution I've found so far is to hide the elements with CSS, making the opacity `0`, and then show them when the modifications are done. This way, the user won't notice any flickering or changes on the page.

Tip. If you're working on a similar project, try using CSS to set elements' opacity to `0` while modifying them. This approach minimizes flickering and ensures the changes aren't visible to the user.

```
.element {
    opacity: 0;
}
```

```
const el = document.querySelector('.element')

if (!el) {
    return
}

// Update the element

el.style.opacity = '1'
```

### Challenge 3: Handling Inconsistent HTML Structure

`www.php.net` has seen many changes over the years, and as a result, the HTML structure of the pages can be inconsistent. This inconsistency occurs because different parts of the pages were created at various times by different contributors, each using their own rules and approaches.

Trying to modify the DOM elements on such pages can be challenging because you can't rely on a consistent structure. Plus, I have to account for potential future updates to the site's HTML structure, which could break the extension if not handled carefully.

To overcome this challenge, I target and parse specific sections of the page, remove them from the DOM, and reinsert them after making modifications. This way, I can ensure that the extension works consistently across all pages, regardless of the HTML structure. While this requires extra work, it ensures the extension functions reliably across all pages, regardless of the underlying HTML inconsistencies.

I'm applying this technique to the right sidebar on the homepage because:

![www.php.net home page screenshot](../../../storage/posts/content/php-revival-extension-a-new-look-at-php-documentation-1728830269.png)

It contains many useful links that I want to keep, but I also would like to rebuild it with a new design to get something like this:

![www.php.net home page screenshot with the PHP Revival extension](../../../storage/posts/content/php-revival-extension-a-new-look-at-php-documentation-1728830430.png)

## Should You Use the PHP Revival Extension?

If you're a PHP developer who uses the official PHP documentation regularly, the PHP Revival extension is a must-have. You can at least try it out and see if it's something you'd like to use. If not, you can always disable or remove the extension.

If you like the extension but feel that something is missing or could be improved, please let me know in the comments below. I'm always open to suggestions and feedback. I'll do my best to implement the most requested features and improve the extension, as long as it benefits the PHP developers.

> This looks AWESOME! Thanks!
>
> ###### Alex Labanino
>
> <img src="https://php-revival.github.io/assets/Alex%20Labanino-Cte-ydIb.png" alt="Alex Labanino" />

## Conclusion

Looking back at 2019, I couldn't have imagined that a small contribution to the community would lead to such an unexpected change in direction. I'm grateful for the support and feedback from the community, and I'm excited to continue working on the [PHP Revival](https://php-revival.github.io/ "Link to the PHP Revival landing page") extension as long as it benefits everybody.

I hope this article has inspired you to make a small contribution to any other community you're part of. Remember, every bit helps, and you never know where it might lead you.

If you want to install the PHP Revival extension, just head over to the [official website](https://php-revival.github.io/ "Link to the PHP Revival landing page"), and you can install the extension from the Chrome Web Store or Firefox Add-ons.

> AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME AWESOME
>
> ###### for prog
>
> <img src="https://php-revival.github.io/assets/for%20prog-BFfddD2_.jpeg" alt="for prog" />