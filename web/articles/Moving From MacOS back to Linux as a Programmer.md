# Moving From macOS back to Linux as a Programmer

## Introduction: The Journey to macOS

In May 2022, I bought my first Apple device, a **16-inch, 16 GB RAM, 512 GB SSD MacBook Pro 2021 with the M1 Pro chip**. I had always been interested in Apple products, seeing programmers and designers using them in their video vlogs and tutorials. Honestly, I was frustrated with my previous laptop, a **Dell Vostro 5590**, which was fast but had a terrible screen and poor build quality.

![Dell Vostro 5590](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1755759832.webp)

Dell Vostro 5590 that I had

You see those MacBook Pros on social media, and they look so sleek and nice, and I thought to myself, "This is the laptop I need to get my programming career to the next level." So, I took the plunge and bought one.

It felt nice to join the Apple ecosystem, but little did I know that this would lead to a series of frustrations that would ultimately drive me back to Linux.

Note. Throughout this article, every time I mention Linux, I mean [GNU](https://www.gnu.org/)/[Linux](https://www.linux.org/pages/download "Link to the GNU project"). Linux is just a kernel, and GNU is a collection of tools on top of it that makes it a complete operating system. Often times, people forget to give the GNU project the credit it deserves, so I want to stress that when I say Linux, I mean GNU/Linux.

![GNU/Linux Logo on a green background](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756374185.webp)

For the full context on why I made the switch to macOS in the first place, I documented it in my article [Moving From Linux to macOS as a Programmer](../../../posts/moving-from-linux-to-macos-as-a-programmer "Link to the article"). I highly recommend checking it out to understand the context of my journey. In this article, I will elaborate on the route I took to return to Linux, the challenges I faced, and the reasons why I believe Linux is a better choice for hackers. And I purposefully use the term "hacker" here; you will understand why later in the article.

## Why I Ditched macOS for Linux

I'm pretty sure there are people who enjoy using macOS, and that's great. However, for me, the experience was far from ideal. A part of it was that you start missing the freedom and flexibility that Linux offers. If Mac is the only thing you ever used, you might be happy with it for the rest of your life.

### Privacy Concerns

My biggest issue with Apple is privacy. Even though, as of 2025, Apple doesn't have the worst reputation when it comes to user privacy, I still don't trust them. I can't shake the feeling that Apple, as a massive US corporation, has ties and obligations to government entities that could compromise user privacy. If government wants to access your data, they will find a way to do it.

With Linux, I have full control over my data and who has access to it. I have the freedom to choose the flavor of Linux that suits my needs. If privacy is a concern for you, Linux is the way to go.

### Small frustrations Adding Up

Beyond the big philosophical issues, it was often the small, daily frustrations that chipped away at my enjoyment. They highlighted the lack of control and the "Apple knows best" mentality. There are apps that you can't get rid of, like Stickies as an example. Since it's a system app, you can't delete it. Why is Stickies a system app? It doesn't make any sense to me. I tried to use it, but it looks ugly, and I don't like it.

![Stickies app on macOS](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1755936541.webp)

Stickies app on macOS

If you want to close the Finder app, you can't do it. It's always there with an indicator in the dock. From users' perspective, it feels odd that you can't close the app you don't use. And the reason it's always running and you can't close it is that macOS is using this app to manage the desktop.

![Finder on macOS](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1755936457.webp)

Finder app on macOS

Why not separate the desktop management from the file manager? Apple didn't see it that way. They are too focused on earning money and locking users into their ecosystem.

There are many weird implementations on macOS that just don't make sense for me. Like a Time Machine app that backups your data to an external drive. Why does this app need the full disk to create a backup file? I don't know. Apple doesn't know either.

![Time Machine](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1755936607.webp)

I have an external **1 TB SSD** that I use for various things. I have a folder on it where I store my backups. If I want to back up my macOS, I need to give Time Machine the whole disk, not just a single directory. Sure, you can partition the drive and give it a partition, but why would I want to do that? I would like to use the drive as a whole, not split it into partitions.

On Linux, with Deja Dup Backup Tool on GNOME, I can choose any directory on any drive to store my backups. I don't need to give the app the entire disk.

![Deja Dup Backup Tool on Linux](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1755937018.webp)

This is how it should be. Shout out to [Michael Terry](https://mterry.name/ "Link to Michael Terry's website") for maintaining this beautiful app.

**I can sit here and list all the frustrations I had with macOS, like not being able to close apps by clicking on the red button, slow and glitchy AirPods connectivity, or inconsistent copy-paste experience across devices. Or buggy desktop (workflow) switcher.** But I think you get the point. Small frustrations add up, and eventually they become unbearable.

The purpose of this article is not to bash macOS but to share my journey as a programmer moving from Linux to macOS and back to Linux. If you're considering the switch to macOS, I would recommend you think twice. If you are already on macOS and feeling frustrated, know that you are not alone.

**When we all switch to Linux, the world will be a better place. Linux is the future, whether you like it or not. The power of people united is stronger than the power of corporations.** Maybe it sounds cringy, but that's just how I feel about it.

These frustrations made the decision final. The only thing left was to find hardware worthy of replacing my MacBook—a machine that could deliver a premium Linux experience. This led me to the most important step: choosing my new Linux machine.

## Choosing My Linux Machine: The Return to Freedom

Finding the right hardware for my return to Linux was mostly straightforward, though it required careful research. The first step is choosing the right machine. This is a big one. It needs to be something to match the MacBook's build quality and performance. You can technically go with almost any laptop that can run Windows, but I would recommend buying a laptop that has good Linux support.

There are a few options out there; here are some of them:

*   **Framework Laptop**: A modular laptop that you can easily upgrade and repair. Very famous in the Linux community. I would buy it if it wasn't US-only at that time.
*   **ThinkPad**: The ThinkPad series doesn't need an introduction. They are known for their durability and excellent Linux support.
*   **System76**: A company that sells laptops and desktops with Linux pre-installed. They have a good reputation in the Linux community and offer great hardware options.
*   **Dell XPS**: Dell's XPS series is also known for its good Linux support. Dell is one of the few manufacturers that officially sell laptops with Linux pre-installed.
*   **Tuxedo Laptop**: A lesser-known brand that offers various laptops with Linux pre-installed. They have a good selection of hardware options and are worth considering.

I went with the **ThinkPad T14s Gen 5**. It has a great keyboard, solid build quality, and excellent Linux support. Plus, it's pretty light and portable, making it a great choice for a programmer on the go.

![The photo of my Lenovo ThinkPad T14s Gen 5](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1755871852.webp)

The photo of my Lenovo ThinkPad T14s Gen 5

As for the specs, I went for this:

*   CPU: Intel i7 Ultra
*   SSD: 1 TB
*   RAM: 32 GB

I chose my laptop carefully, reading all the specs multiple times so that I could digest everything. If you are choosing a machine for yourself right now, I highly recommend you to be picky.

## Why Linux? Reclaiming Control as a Programmer

For many of us programmers, Linux is something you grow up with. It's the OS that not only powers the majority of servers in the world but also offers great desktop environments.

If you like animated and beautiful designs, [GNOME](https://www.gnome.org/ "Link to the GNOME official website") has you covered. If you want something more traditional, there's always [Cinnamon](https://en.wikipedia.org/wiki/Cinnamon_(desktop_environment) "Link to Wikipedia page"), [Budgie](https://ubuntubudgie.org/ "Link to Ubuntu Budgie website"), and [KDE](https://kde.org/plasma-desktop/ "Link to KDE website"). Thinking about something lightweight? [XFCE](https://xfce.org/ "Link to the official XFCE website") is a great choice.

I went as far as switching my mom from Windows to Linux Mint, and she's loving it.

![My mom's desktop look with Linux Mint](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1755948592.webp)

A screenshot of my mom's desktop look with Linux Mint

She noticed a few differences at first, but she adapted quickly and now appreciates Linux's simplicity. She even uses the terminal to install apps from [Flathub](https://flathub.org "Link to Flathub website") by copying commands from their site. If my mom, born in the '70s, can handle it, you and your parents can too. This accessibility is built on a foundation of incredible flexibility.

### The Freedom of Customization

If customization is your thing, Linux is the way to go. You can customize almost everything, from the desktop environment to the terminal emulator. You can choose your window manager, your file manager, or even distribution.

You can delete unwanted apps and customize your workflow to match your needs. For example, I deleted all the keyboard shortcuts that I don't use and added my own for quick workspace switching and window management.

> All the best people in life seem to like LINUX.
> 
> ![Steve Wozniak](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756212307.webp)
> 
> ###### Steve Wozniak

macOS is very restrictive and closed in this regard. It wasn't designed to be customized or modified. You are stuck with what Apple gives you. Want to change the fundamental way you navigate between windows or workspaces? Want to truly optimize your keyboard for coding efficiency? **On macOS, you're often fighting the OS. On Linux, you're collaborating with it.**

### Open-Source Power for Developers

Linux provides you with access to a vast array of open-source tools and libraries that can enhance your development workflow. You can easily install and manage packages using package managers. If you don't want to use the terminal, you can use GUI package managers like **GNOME Software, KDE Discover, or Software Manager on Linux Mint**.

![GNOME Software program showing that there are no updates available](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1755950040.webp)

If you are using macOS, do the math of how much money you spend on apps like Magnet, Alfred, and services like iCloud. On Linux, you can get similar functionality for free using open-source alternatives.

The trade-off is a different kind of support. Instead of a refund, you get a community. You can report issues, contribute fixes, or even fork the project—options that simply don't exist in the walled garden of macOS.

### Linux for Modern Development

Linux is the backbone of modern development. Containers, cloud computing, and DevOps practices are all built on Linux. If you want to stay relevant in the tech industry, learning Linux is a must.

I was heavily using containers in my development workflow, and you think the M1 Pro chip will handle Docker well? Think again. Docker on macOS is super slow compared to Linux.

![Docker Desktop app on macOS](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1755952640.webp)

The ARM architecture support is still not perfect. Docker Desktop for Mac is also known to be unstable and buggy. It uses a virtual machine to run Linux containers, which adds overhead and slows down performance.

As a bonus, you get a much better gaming experience on Linux compared to macOS. With the [Steam Deck](https://store.steampowered.com/steamdeck "Link to Steam Deck website") and [Proton](https://github.com/ValveSoftware/Proton "Link to Proton GitHub page"), you can play a lot more games on Linux than on macOS. Gaming on Linux is going to get even better in the future. [Valve](https://www.valvesoftware.com/ "Link to Valve website") is investing heavily in Linux gaming, and it's paying off. Thank you, [Gabe Newell,](https://en.wikipedia.org/wiki/Gabe_Newell "Link to Gabe Newell Wikipedia page") for this!

**All this beautiful ecosystem is what led me to choose Linux as my primary OS.** Let's talk a little bit about the distribution choice that I've made.

## Choosing Fedora: A Developer’s Ideal Linux

Choosing a Linux distribution and desktop environment is the most important step. I remember spending hours talking with AI discussing differences between distros and desktop environments. It was the hardest part of the switch because there are so many of them.

![Various Desktop Environments for Linux](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756211386.webp)

My ThinkPad hadn't even arrived, and I was already paralyzed by choice. First, I considered going back to Ubuntu, but then I was captivated by how awesome Linux Mint is, and I decided to try it instead. The next day, I was convinced Debian was the answer. Then, I was thinking about Arch Linux, but I didn't want something too bleeding edge.

Finally, after all that research and deliberation, I settled on Fedora.

### Why Fedora over Ubuntu?

Before buying myself a MacBook, I was using Ubuntu for around 5 years. While I used and mostly enjoyed Ubuntu before my MacBook, Canonical's aggressive push of Snaps was a primary reason I looked elsewhere. I dislike the forced integration and the performance overhead. **Fedora offers a refreshingly pure, upstream GNOME experience without such vendor lock-in, instead embracing open standards like Flatpak.**

![Desktop look of my Linux Fedora 42 setup](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756278763.webp)

My current look of the desktop. I put this wallpaper just for the sake of a screenshot

Fedora, on the other hand, is very similar to Ubuntu, since both use GNOME as the default desktop environment. While Fedora is backed by Red Hat and IBM, it has a strong community and is known for its commitment to open-source principles. Fedora tends to have more up-to-date packages, which is great for developers who want the latest tools and libraries.

The other thing that attracted me to Fedora is its attachment to tools like Podman for container management and [Flatpak](https://flathub.org/ "Link to Flathub repository") for application distribution. It happens that I love Podman and Flatpak, and Fedora embraces these technologies wholeheartedly.

### Why GNOME?

My choice of GNOME ultimately comes down to personal preference. I appreciate its simplicity and extensibility. By extensibility, I mean the ability to add GNOME extensions, which you can install from the [Extension Manager](https://flathub.org/apps/com.mattjakeman.ExtensionManager "Link to Flathub Extension Manager") or create your own with JavaScript.

![My GNOME Desktop look](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756277549.webp)

I'm thoroughly happy with my current GNOME setup. Its clean design gets out of my way, and the extensibility allows me to tailor it perfectly to my workflow—a level of control that simply wasn't possible on macOS.

It's easy to find help online and overcome any challenges you might face.

## Overcoming Challenges in the Switch

While the switch was ultimately straightforward thanks to my Linux background, I still faced a period of adjustment. It involved rewiring my brain for a new workflow and hunting down my essential tools on Flathub.

But **any obstacle can be overcome if you want the freedom badly enough.**

### Adjusting to Linux Workflows

The biggest and most time-consuming difficulty was the habit of using the old Mac keybinds. As you know, Apple keyboards have a "Command" key, which other keyboards lack.

![Apple Magic Keyboard](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756209549.webp)

Things like copying, pasting, and undoing changes are done through the "Command" key. Getting used to the same actions but with the "Ctrl" key was a 7-day journey. Of course, the more you do it, the easier it gets.

To make this transition easier, you can remap things to the "Alt" key to make it act as "Command," but I didn't do it. I needed to embrace the pain since I would rather not use Linux as macOS or with an Apple mindset. I wanted to learn Linux on its terms, not force it to be something it's not.

I wouldn't suggest that anyone remap keys to match their old Mac device, but if you really want to do it, Linux will not be on your way. But let's mention the elephant in the room when it comes to Linux vs. macOS and Windows: privacy and security.

## Privacy and Security

**The biggest and most significant factor influencing my decision was the enhanced privacy and robust security offered by Linux in general.** Trusting a company like Apple, which was integrated into the NSA's PRISM program and remains subject to the same secretive laws and national security letters, to me, like a crazy move. The framework for cooperation is still in place, making it impossible for me to trust that my data is truly private.

![People are walking on a protest with a poster 'HANDS OFF MY META-DATA'](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756293559.jpg)

It was clear in data leaks in 2013 by [Edward Snowden](https://en.wikipedia.org/wiki/Edward_Snowden "Link to Edward Snowden Wikipedia page") that Apple was working with the NSA to provide backdoor access to user data through the [PRISM program](https://en.wikipedia.org/wiki/PRISM "Link to a Wikipedia page").

![NSA Logo surrounded by other logos, Facebook, Google, Apple, Yahoo, Skype](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756293717.jpg)

Can you open a fresh document on macOS and start writing about something deeply personal? Like your health issues, your inner thoughts about life, intimate relationships with your partner, and conversations with your inner voice. Can you do it without the fear of being spied on? I can't. I don't trust Apple or any other big corporation with closed source code.

**Compare it to Linux, which is open-source and transparent. No spyware, no bloatware, no hidden backdoors and no telemetry.**

![Open-source Logo](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756293985.webp)

Note. You can find plenty of articles online about how to stay safe, but I haven't found any that align with my privacy standards. That's why my next article will be **Securing Your Digital Life: A Programmer's Action Plan**. It will be the most useful article you will read in your entire life, believe that. I'm already working on it while you are reading this one.

Linux gives you the tools to stay safe and private, but it's up to you to use them. It doesn't matter how protected your house is if you leave the door wide open. It's much better than living in a protected house with NSA agents inside.

![A person holding a sign at a protest against mass surveillance with text 'HANDS OFF MY META-DATA'](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756291185.webp)

I want to end this section with a quote from [Richard Stallman](https://en.wikipedia.org/wiki/Richard_Stallman "Link to Richard Stallman Wikipedia page"), the founder of the Free Software Foundation and the GNU Project:

> Free software is software that respects your freedom and the social solidarity of your community. So it's free as in freedom.
> 
> ![Richard Stallman](../../../storage/posts/content/10-ways-to-make-money-in-web-development-1664114567.jpg)
> 
> ###### Richard Stallman

## The Joy of Tinkering with Linux

Do you ever get that thrill from using your computer? That little rush of excitement when you open a Linux terminal and start typing commands? I know I can't be the only one who feels this way. There must be millions of us who love tinkering—not just with software but with hardware too.

Do you remember messing around with your old computers, routers, switches, ONTs, and other gadgets? I have such admiration for people who build this stuff from scratch. I remember when I was a kid, I used to take apart things to see how they worked and then put them back together or try to build something new.

![Disaccembling some electronic device with buttons](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756369164.webp)

Returning to Linux brought back the pure joy of hacking. It inspired me to go beyond the surface, to learn the meaning behind essential abbreviations like **FTTH, DSL, LAN, WAN, ADSL, VDSL, ONU, DNS, TCP, and UDP**, and finally understand how the surrounding technology actually works.

If you are a programmer and you still haven't switched to Linux, what are you waiting for? **Your computer is waiting to become a source of joy again.**

## Conclusion

Looking back at my experience with macOS, I realize that while it had its moments, it often rubbed me the wrong way. The lack of control, the privacy concerns, the giant corporation's influence on my digital life—all these factors pushed me to switch back to Linux.

**Not everything is perfect with Linux, just like with any other OS.** The longer you use it, the easier it becomes. It's a skill, and like all skills, practice makes perfect. Linux gives you the freedom to learn, tinker, and grow—not just as a programmer, but as a hacker.

And I don't mean "hacker" in the negative sense twisted by the media. A true hacker is a curious tinkerer, an explorer, and a creator. It's about understanding the system, bending it to your will, and valuing the freedoms of privacy and control. It's about being able to look under the hood and not just accept that something "just works"—but to know _why_.

![A man unpluging the wire from a server station](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756374622.webp)

### Looking Ahead

The trajectory is clear: Linux is only going to get better, and its adoption will continue to grow far beyond the server room. We are witnessing its rise on the desktop, in gaming, in business, and crucially, in the public sector around the world. This isn't a niche trend; it's a global shift towards sovereignty and security.

The evidence is everywhere:

*   Since 2008, France's military police (Gendarmerie) have run [GendBuntu](https://en.wikipedia.org/wiki/GendBuntu "The link to Wikipedia") on tens of thousands of PCs.
*   In 2024, the German state of Schleswig-Holstein announced a full migration of 30,000 government workers from Windows to Linux.
*   By 2025, Denmark’s Ministry of Digital Affairs will have fully transitioned to Linux.
*   Nations are building their digital infrastructure on open source: India with [Maya OS](https://en.wikipedia.org/wiki/Maya_(operating_system) "The link to Wikipedia"), Russia with [Astra Linux](https://en.wikipedia.org/wiki/Astra_Linux "The link to Wikipedia"), China with [Kylin](https://en.wikipedia.org/wiki/Kylin_(operating_system) "The link to Wikipedia"), and Turkey with [Pardus](https://en.wikipedia.org/wiki/Pardus_(operating_system) "The link to Wikipedia"). This is a powerful testament to the security and reliability these systems provide.

This movement isn't confined to governments. Private businesses are also making the switch, recognizing the long-term benefits. And this is where you come in. You can contribute to this movement directly by choosing Linux yourself. Every new user strengthens the community, drives development, and pushes the ecosystem forward.

Your journey doesn't just change your computer; it's a vote for a freer, more open digital future.

### Advice for Programmers Considering Linux

If you want to try Linux, start by running it in a virtual machine (VM) on your current OS. It's the safest way to get a feel for it. Just keep in mind that a VM won't give you the full experience—you'll miss out on native hardware compatibility and performance—but it's an excellent and risk-free first step.

Don't reach for advanced distros like Arch or Gentoo right away. Start with something user-friendly and stable like **Fedora**, **Pop!\_OS**, or **Linux Mint**. I, personally, don't recommend Ubuntu anymore due to its direction with Snap packages, but it remains a popular starting point and is still a far more programmable environment than Windows or macOS.

Start escaping proprietary ecosystems. Look at the services you're paying for and see if there's a free, open-source alternative. You'll be surprised by the powerful options available:

*   [GIMP](https://www.gimp.org/ "The link to GIMP official website") instead of Photoshop
*   [Blender](https://www.blender.org/download/ "The link to official blender website") instead of Autodesk Maya
*   [LibreOffice](https://www.libreoffice.org/ "The link to official LibreOffice website") instead of Microsoft Office
*   [Kdenlive](https://kdenlive.org/ "The link to official Kdenlive website") instead of Final Cut Pro

For nearly every proprietary app, a robust open-source alternative exists. The trade-off is often a different workflow for the sake of freedom and control. Package formats like Flatpak make these apps incredibly easy to find, install, and manage.

The visual customization is a massive part of the appeal. You can make your desktop look like anything you want.

![Garuda Linux desktop look with open file manager and a welcome window](../../../storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756380737.webp)

Garuda Linux demonstrates a modern, flashy desktop style.

If that's not your taste, you can choose something more traditional and Windows-like (Cinnamon) or immensely customizable ([KDE Plasma](https://kde.org/plasma-desktop/ "Link to KDE website")). That's the beauty of Linux—you own your system, not the other way around. You decide how it looks and behaves, not a giant corporation. You become the master of your own digital destiny.

Have you made the switch? Thinking about it? Leave a comment down below with your thoughts or questions.
