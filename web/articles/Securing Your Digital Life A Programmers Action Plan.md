## Introduction
<blockquote>
    <p>What I'm about to tell you is top secret. There's a powerful group of people out there that are secretly running the world. I'm talking about the guys, no one knows about the guys who are invisible. The top 1% of the top 1%. The guys that play God without permission. And now I think they're following me.</p>
    <img src="https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1756472826.webp" alt="Elliot Alderson in a dark hoodie with a black background" />
    <h6>Elliot Alderson (Mr. Robot series)</h6>
</blockquote>

I've never started an article with a quote, but this dialogue from Elliot Alderson was the first thing that came to my head when I started writing it. **The surveillance and existential paranoia displayed in that great series perfectly capture the digital landscape we navigate today.**

Recent scandals involving British and other politicians pushing for widespread monitoring, data collection, and mandatory face scans to confirm your age have highlighted the escalating trend towards government and corporate surveillance. This raises serious concerns about privacy and individual freedom in the digital age.

![People are protesting. Holding banners with the sign: Stop watching all of us](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1756478455.jpg)

Even though this article will not be political, it's difficult to ignore the increasing tension between individual privacy and governmental/corporate agendas. **The greediness for more power over people and control over their digital footprints is evident**, particularly from whistleblowers like Edward Snowden, who revealed how much personal data governments the National Security Agency (NSA) collect.

![NSA logo surrounded by logos of Facebook, Yahoo, Skype, Apple, and Google](https://serhiicho.com/storage/posts/content/moving-from-macos-back-to-linux-as-a-programmer-1756293717.jpg)

This ongoing war on privacy has led millions to investigate how to reclaim their digital freedom. **DuckDuckGo search engine went from 1.64 million daily queries in 2013 (Snowden revelation) to over 100 million in 2024. It's a mind-blowing 6000% increase in daily searches.**

![Advertizing banner with the DuckDuckGo ad: Google tracks you. We don't](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1756478715.jpg)

<a href="https://proton.me/mail" title="Link to ProtonMail website" target="_blank">ProtonMail</a> founded their company after Snowden's leaks in 2013, and in 2025 they have over **70 million users**, illustrating a clear and growing public mandate for privacy-focused digital services.

Switching from Gmail to ProtonMail is a good first move, but it's like putting window bars on a single window when you have 10 of them. I don't consider myself a guru of privacy, but I have been pushing hard on privacy since January 2019, and since then, I've changed the digital footprints of my some of relatives and myself.

And yes, **as tech professionals, this responsibility extends to teaching and protecting our relatives who are strangers to all that tech stuff.** Because while you understand what DNS is, for your parents, this abbreviation doesn't ring a bell.

<blockquote>
    <p>You don't know how much you appreciate your privacy until you don't have it.</p>
    <img src="https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1756479356.webp" alt="Photo of Morgan Freeman, actor and producer" />
    <h6>Morgan Freeman (actor, producer)</h6>
</blockquote>

Anyway, I'm not your privacy advisor; I'm just a tech who cares about this stuff. So, take my advice with a grain of salt and do your own research if you want to delve deeper. But I'm going to walk you through a comprehensive action plan for securing your digital life, offering practical steps to enhance your digital life and those people you care about. And we'll start with the first and most important step, your operating system.

## Use Linux
The most fundamental aspect of digital security begins with the underlying free and open-source, fully functional operating system. It’s like a foundation upon which all the other security measures are built. If you skip this step and keep using Windows or macOS then all of the following steps in this article are pretty much don't make sense.

In truth, the field of fully open-source operating systems is narrow. As of 2025, we have only FreeBSD and GNU/Linux.

Tip. GNU/Linux is the full, proper name for the operating system most people refer to simply as "Linux." The term GNU/Linux is used to give credit to both projects and acknowledge their movement. However, throughout this article, I'm going to refer to GNU/Linux as just Linux, for simplicity.

FreeBSD is a good choice; I have nothing against it; however, I’ve never used it myself and thus cannot recommend it or comment on it. I just know it has a robust, long-standing reputation for stability and security in server environments.

**Linux, on the other hand, is a great choice** because it has a vast and active community, frequent security updates, and lots of different flavors. This diversity allows users to select a distribution tailored to their specific security needs and technical proficiency, ranging from user-friendly options to highly configurable systems like <a href="https://archlinux.org/" title="Link to Arch Linux website" target="_blank">Arch Linux</a>.

Keep in mind that Linux is not a silver bullet. While it's a great alternative to Windows and macOS, it's not perfect. A common downside of some Linux desktop environments is that they can be less polished and more buggy than the tightly integrated macOS.

![A Steam Deck with Linux logo on the screen](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1757157853.jpg)

You can just start by running Linux on a virtual machine if you've never done it before. This approach allows for risk-free experimentation with various distributions and desktop environments, familiarizing oneself with the nuances of a Linux system without altering the primary operating system. Keep in mind that the experience with Linux inside a virtual machine is not as performant as running it natively, but it's a good starting point.

## Safe ISO File
It doesn't matter how safe your house is if your back door is not locked. That's why it's important to follow basic security practices when it comes to downloading installation files like ISO (often called an ISO image).

**ISO image is a complete snapshot of a disc, commonly used for operating system installations.** All you need to know about this file format is that it encapsulates a bootable environment. It means that this environment can be tinkered with. That's why we have such thing as checksums.

To check the checksum of a file, you typically use command-line utility like `sha256sum` on Linux. This cryptographic hash function generates a unique string of characters for a given file, and if you change at least 1 byte in that target file, the resulting checksum will be drastically different. This allows you to know if the downloaded file has been modified or not.

You run `sha256sum your-file` and compare the output to the published checksum provided by the official website.

```bash
# Navigate to your Downloads folder first, then run:
sha256sum fedora-workstation-42.x86_64.iso
```
If they match, then you can be confident that your downloaded ISO file is authentic and untampered with.

## No Proprietary Software
**Proprietary software, by its nature, often operates as a black box, making it impossible for users to inspect its inner workings for vulnerabilities or malicious functionalities.** Processes that are constantly running on your machine can collect sensitive data and transmit it without explicit user consent. I do not consider myself a paranoid individual, but it is not a secret that companies like Microsoft, Apple, Facebook, and Google routinely collect vast amounts of user data. Especially Google, since their whole business model is to sell you targeted ads.

![chrome, watsup, skype and google logos are spying behind the girl who is sitting with her laptop on a sofa in a cartoon style](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1757248167.webp)

If you installed a fresh Linux installation and are thinking of getting your favorite Google Chrome browser, be mindful that you are reintroducing the same issue that you had before. Running Google Chrome on Linux is like installing a surveillance camera with the default password `admin` into your secured house. The rule is simple: if it's not open source, it doesn't get root-level trust.

And it applies to browser extensions, which can introduce huge security vulnerabilities. **Before installing a browser extension check if it's open source and keep close attention to its requested permissions.** Unfortunately, there is no perfect way to check the checksum of the extension itself and compare it to the original version, we can only rely on Mozilla and Google to ensure the integrity of their extension marketplaces.

### Use Alternatives
For every proprietary software application, there is typically a robust open-source alternative that offers similar or even the same functionality. There are exceptions, of course, but most of the time you can find pretty much anything, including flatpak applications.

Here is a list of recommendations if you are looking for free alternatives:

- **Email:** ProtonMail or Tutanota offer end-to-end encrypted email services with a bunch of features
- **Cloud Storage:** ProtonDrive and Nextcloud provide secure storage. Nextcloud is a self-hosted solution
- **Daily Browser:** Firefox, Brave, Vivaldi and tons of others. If you like Chrome, just use Brave and you will not see any differences in terms of user interface
- **Search Engine:** DuckDuckGo, Brave Search and Startpage are the best privacy-focused solutions
- **Graphic Design:** GIMP is the best choice if you want to switch from Photoshop
- **Office Suite:** LibreOffice is a comprehensive alternative to Microsoft Office, offering word processing, spreadsheets, and presentation tools that are compatible with common file formats.
- **Code Editor:** Vim, Neovim, and Emacs are true open-source options for developers, but if you prefer a more modern UI, VS Code, while partially proprietary, still provides significant open-source components and a vast ecosystem of extensions. If you can avoid VS Code (due to its Microsoft telemetry and proprietary elements), that’s a huge advantage. We will talk more about this later in the article.

I can continue this list for a long time, but the core principle remains: prioritize open-source solutions to enhance security, transparency, and user control over your digital environment. You can always search for alternatives and you will find help online.

## Secure Browsing
Browsing is something we as tech professionals do constantly, whether for research, development, debugging, or just personal use. The personal use part is particularly important and requires extra vigilance.

When I type a query in a search bar, I don't want anyone to attach this search query to my identity, nor do I want my browsing habits or visited websites tracked and correlated with my personal data. I want to be able to save my browser tabs in private, encrypted storage so that even if someone gets access to my computer, they cannot see those saved tabs. That’s where browser extensions in combination with open-source philosophy shine.

### Browser Extensions
There are at least 3 browser extensions that you must have by default. These are:
- <a href="https://addons.mozilla.org/en-US/firefox/addon/duckduckgo-for-firefox/" title="Link to Mozilla DuckDuckGo Privacy Essentials page" target="_blank">DuckDuckGo Privacy Essentials</a>: Only essential if you use DuckDuckGo as your default search engine. It enhances that specific experience.
- <a href="https://addons.mozilla.org/en-US/firefox/addon/ublock-origin/" title="Link to Mozilla uBlock Origin page" target="_blank">uBlock Origin</a>: Essential for everyone. The best content blocker period.
- <a href="https://addons.mozilla.org/en-US/firefox/addon/privacy-badger17/" title="Link to Mozilla Privacy Badger page" target="_blank">Privacy Badger</a>: A great supplement to block sneaky trackers that slip through.

Note. DuckDuckGo Privacy Essentials extension only makes sense if you use DuckDuckGo as your primary search engine.

I'm giving you Mozilla links because I recommend Firefox-based browsers for their stronger privacy stance. Plus, uBlock Origin's <a href="https://github.com/gorhill/uBlock/wiki/uBlock-Origin-works-best-on-Firefox" title="Link to the Wiki" target="_blank">maintainers also recommend</a> using it on Firefox. But if you need it for Chromium, just get the extension from the official <a href="https://ublockorigin.com/" title="Link to official uBlock Origin extension website" target="_blank">ublockorigin.com</a> website instead of Chrome Store. If you use the Brave browser, you may not need uBlock Origin as it has a built-in blocker, though many still use it for finer control.

![uBlock Origin additional information for different browsers](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1757592362.webp)

Before installing any new extension, make sure that it's open source and doesn't have a permission of accessing your data for all websites. Of course, those 3 extensions that I've listed above are the exeption from this rule because they cannot work without getting access to all the data for all the websites.

### Tab Guardian
This desire for private tab management is exactly why I developed the <a href="https://addons.mozilla.org/en-US/firefox/addon/tab-guardian/" title="Link to Mozilla Tab Guardian extension page" target="_blank">Tab Guardian</a> extension that is available for Chromium and Firefox-based browsers.

![Tab Guardian extension banner with a sign 'See all of your saved tab groups at a glance on the main screen'](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1757587311.png)

I'm a bit biased towards this extension because I developed it specifically to address the critical need for encrypted and private tab management, ensuring that sensitive browsing data remains inaccessible even in the event of system compromise.

It allows you to store and retrieve your browser tabs in a few clicks, but **the coolest feature of Tab Guardian is “Private groups”**, which allows you to create encrypted groups with a password to prevent anyone from accessing them.

### Use Tor for Maximum Privacy
For the ultimate level of anonymity in your browsing, the Tor network is the gold standard. But keep in mind that it's a separate tool for specific, high-privacy needs. It doesn’t replace your daily browsers since Tor is significantly slower due to its multi-layered encryption and relay system.

![Tor Browser home page that says 'Explore. Privately'](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1757591016.webp)

The other reason is that companies like Google and Meta actively discourage its use because if they don't know who you are, they cannot monetize your online activity. But honestly, **you shouldn’t use their services anyway if you care about your privacy.**

## The Security Fundamentals
Beyond your browser, true digital security is built on fundamentals. These tools don't just make you safer—they make your life easier. You might already use some, like a password manager, while others remain powerful, underutilized secrets.

Let’s start with the most basic yet critical components that you need to have in your arsenal: **a robust password manager**.

### Use a Password Manager
**A password manager is a software application that helps users create, store, and manage their passwords securely.** It typically requires a master password to access a vault where all other passwords are stored, making it easier to use unique and complex passwords for different accounts without needing to remember each one.

Instead of trying to come up with a strong password that you can memorize, a password manager generates passwords like `xDc@P#SXk6&skP3y%zpv` which is virtually impossible to guess or brute-force.

Stop putting it off. If you've never used a password manager, your life is about to get infinitely easier. The principle of open source is also applicable here, since password managers are just applications. If there is nothing to hide, they should be open source.

I'm not sponsored by any password manager, but I recommend <a href="https://proton.me/pass" title="Link to Proton Pass password manager" target="_blank">Proton Pass</a> and <a href="https://bitwarden.com/" title="Link to Bitwarden website" target="_blank">Bitwarden</a>. I’d love to recommend more, but I haven’t used others extensively. You can also explore offline options available on <a href="https://flathub.org" title="Link to Flathub" target="_blank">Flathub</a>—just search for “password manager” there. These are typically self-hosted or local-only options (like <a href="https://flathub.org/en/apps/org.keepassxc.KeePassXC" title="Link to KeePassXC on Flathub" target="_blank">KeePassXC</a>) that don't sync to a vendor's cloud, which is the trade-off for being offline.

![Metal lock with Bitwarden logo on it guarding a door](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1757665744.webp)

Disable password managers in your browser, in your phone, and remove any saved credentials from your physical notes. There has to be only a single place for your passwords, not a scattered collection across multiple locations.

Lastly, go through your accounts and replace weak passwords with strong ones generated by your password manager. Replace your passwords like `i-love-sydney-sweeney` with randomly generated alphanumeric strings with extra characters like `zQp!7$bW@v5F#kL%R9`.

Once your passwords are secure, the next layer of defense is ensuring your data is unreadable even if your device is stolen—through encryption.

### Use Encryption
Imagine losing your laptop. Without encryption, every file, message, and credential becomes an open book. Encryption is your digital shield—it scrambles your data into unreadable code, ensuring even physical theft doesn’t compromise your privacy.

While a password manager encrypts your secrets, the next step is to encrypt everything. **Encryption is the process of transforming information into a secure code to make this information unreadable.** For example, if I encrypt a string `Emma Stone is the greatest actress` with `aes-256-cbc` algorithm like this:

```php
<?php

$text = 'Emma Stone is the greatest actress';
$algo = 'aes-256-cbc';
$passphrase = 'my-passphrase';

$encrypted = openssl_encrypt($text, $algo, $passphrase);

// Result: string(64) "j+7Si0X86xhS96RCZCSXYVQUyv6/irmmSTPMJYkep9You1+rj3bT4XX3FTR8xGKX"
```

In result, I'll get an encrypted string, which is complete gibberish to anyone without the passphrase `my-passphrase`. But when I decrypt this string back like this:

```php
<?php

$algo = 'aes-256-cbc';
$passphrase = 'my-passphrase';
$encrypted = 'j+7Si0X86xhS96RCZCSXYVQUyv6/irmmSTPMJYkep9You1+rj3bT4XX3FTR8xGKX';

$text = openssl_decrypt($encrypted, $algo, $passphrase);

var_dump($text);

// Result: string(39) "Emma Stone is the greatest actress"
```
Now, full-disk encryption is very similar, except that instead of encrypting file by file, it encrypts your whole disk partition. One of the most widely used Linux encryption tools is <a href="https://en.wikipedia.org/wiki/Linux_Unified_Key_Setup" title="Link to Wikipedia page about LUKS" target="_blank">LUKS</a> (Linux Unified Key Setup). **A partition is essentially a massive sequence of 0s and 1s on your drive.** LUKS grabs that entire chunk, encrypts it as-is with a strong cipher, and your passphrase unlocks the master key to make it readable again. Your passphrase is the master key to unlock it all at boot time.

![Scheme that shows how file gets encrypted and decrypted](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1757678153.webp)

**Most Linux distros have encryption built-in into the installer; just make sure to enable it during installation.** If you already installed Linux and do not use encryption, you need to backup your user data and reinstall your OS with full-disk encryption.

This is crucial, because your login password doesn’t matter and doesn’t protect your SSD; your login password is just a gatekeeper to your user session, not your underlying data. Without encryption, a thief can simply plug your SSD into another machine and read all your data directly. Encryption ensures they’ll only see gibberish.

Encryption isn’t just a feature—it's your right. Enable it, and turn your hardware into a fortress.

### Use 2FA
A strong password is your first line of defense, but it can still be stolen in a data breach or through phishing. This is where **2FA (Two-Factor Authentication)** becomes essential. It adds a critical second layer of security by requiring two proofs of identity: something you know (your password) and something you have (a code from your phone or a physical key).

**Avoid SMS-based codes when possible,** as they can be intercepted through SIM-swapping attacks. Instead, use one of these more secure methods:

- **Authentication Apps:** Generate codes offline on your device.
- **Hardware Security Keys:** The gold standard. A physical device like a **YubiKey** that you plug into your USB port. This offers the strongest protection, though it requires carrying the key with you.

![An authentication key lying on a laptop](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1757691755.webp)

Note. I don't recommend apps like  **Authy** as an authentication app because they don't have a way to export your authentication accounts. If you decide to switch apps in the future, it will be very difficult to do. Use those which have import/export functionality like **Bitwarden Authenticator**, **Proton Auth**, **Authman**, etc.

**Your action is simple: enable 2FA on every account that offers it,** especially your email, password manager, and social media accounts. This single step will block the vast majority of automated account takeover attempts.

## Secure Software Management
For programmers, **security isn’t just a best practice—it's a core responsibility**, especially if you work on critical infrastructure or financial systems. If your company decides which operating system and programs you are going to use, you are out of luck. Typically, you will just have to use what is provided, but you can still make sure your own personal environment is safe and secure.

We’ve already talked about Linux and safe browsing, but I also want to talk about secure software management for Linux.

### 1. Use Official Repos
Relying on official distribution repositories is a crucial step in preventing software supply chain attacks, which can compromise systems through malicious code injected into legitimate-looking packages. This approach is great but not always sufficient, since we might need to install a program from a third-party source. A good example for this is something like <a href="https://tableplus.com/download/linux" title="Link to Table Plus download page" target="_blank">Table Plus</a> that is not available on Flathub or on the official Fedora repositories.

![TablePlus installation page for Linux](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1758019838.webp)

The good thing is that there are suitable alternatives like <a href="https://flathub.org/en/apps/it.fabiodistasio.AntaresSQL" title="Link to Antares SQL app on Flathub" target="_blank">Antares SQL</a> on Flathub, even though alternatives might not always offer the same feature set or user experience.

### 2. Use Flatpak Packages
Flatpak offers a sandboxed environment for applications, isolating them from the core operating system and mitigating potential security risks associated with third-party software.

Package formats like Debian and RPM are traditional packaging systems that integrate software directly into the system, often leading to complex dependency management and a broader attack surface. Flatpak helps to mitigate these risks by sandboxing applications with control groups and namespaces, thereby limiting their access to system resources.

Flatpak is getting a lot of traction in the Linux community, and when you start using it, you will appreciate it as well. Especially with a tool like <a href="https://flathub.org/apps/io.github.flattool.Warehouse" title="Link to Flathub store" target="_blank">Warehouse</a> app to manage them.

![Warehouse app interface for managing Flapak packages](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1758024928.webp)

### 3. Use the Flatseal App
Another crucial tool for managing Flatpak applications’ security is <a href="https://flathub.org/apps/com.github.tchx84.Flatseal" title="Link to Flathub store" target="_blank">Flatseal</a>. This graphical utility provides fine-grained control over Flatpak permissions, allowing users to modify access to various system resources and enhance the overall security.

![Flatseal user interface](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1758025376.webp)

If you don't want Obsidian to have access to the internet, you can just revoke the “Share → Network” permission using Flatseal. If you want to give your app access only to a specific directory on your system without giving the whole access, you can tweak it in the “Filesystem” section.

![Filesystem setting in the Flatseal app](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1758025603.webp)

Tip. Always look at your installed Flatpak app in Flatseal before even opening it. Because if an application has access to your filesystem, it could potentially compromise sensitive data. Restricting permissions before you open the app can prevent unauthorized data access or malicious activities.

"Filesystem," "Network," and "System Bus" are the most critical permissions to control, as they represent the primary avenues for data leakage and system intrusion.

### 4. AppImages
**AppImages are also another great option for distributing Linux applications, providing a portable and self-contained format that suits various Linux distributions.** However, as for writing this article, AppImages still require manual updates, and features like seamless desktop integration (e.g., launching AppImages as native applications) are still a work in progress.

But for me, it's still better to have an AppImage than having to install TablePlus with 30+ dependencies on my system. I like the encapsulation that AppImages and Flatpak provide. It gives me peace of mind, and I'm willing to tolerate a bit of manual management for the peace of mind that comes with dependency-free, encapsulated software.

## Development Workflow
A guide on securing your digital life would be incomplete without addressing your development environment. The tools you use to write code can become unexpected vectors for data leakage, making your editor and repository choices a critical part of your security posture.

For developers working with sensitive or proprietary information, tools provided by large corporations like Microsoft present a unique consideration. Their history of participating in programs like <a href="https://en.wikipedia.org/wiki/PRISM" title="Link to Wikipedia page" target="_blank">PRISM</a> is a matter of public record, making it prudent to understand and limit their data collection.

This isn't to say you must abandon GitHub or VS Code—but you should use them with awareness and configure them for maximum privacy.

### Configuring Your Tools
Here is my personal advice on how to keep using software development tools more secure:

**For VS Code:**
The VS Code editor collects significant telemetry by default. To disable it, add these lines to your `settings.json` if you don't have them yet:

```json
{
    "telemetry.feedback.enabled": false,
    "telemetry.editStats.enabled": false,
    "telemetry.telemetryLevel": "off",
}
```

**For GitHub:**
Use SSH keys instead of password and disable Copilot on private projects to prevent your code snippets from being sent to and stored on Microsoft's servers for model training. It's not a big deal if you work on open-source, but if you work on a private project you might disable copilot to avoid any code leakage.

**JetBrains IDEs:**
JetBrains doesn't have a history of controversial data collection practices since their income stream relies on product sales rather than data monetization. From all the big tech companies out there, JetBrains has the most amount of respect from the developer community for their ethical approach to user data and privacy.

If you use one of their products, you can disable telemetry in your settings if you like, but this step is largely optional due to their transparent data handling policies and clean record. **The only JetBrains-related IDE I recommend to disable telemetry on is Android Studio, since it's proprietary software managed by Google instead of JetBrains**, posing a similar privacy concern as other Google products.

### Use Podman
Since software development is heavily dependent on container technologies, we should understand their inherent security implications. Running Docker as root on your system exposes a significant attack surface.

The “Docker is secure” or “Docker is safe” mentality is marketing crap. Docker containers are isolated in theory, but root access means if anything's hacked, you're done. Escaping container and getting access to your system is is a well-documented risk. One wrong `docker run -v /:/host` and boom, the attacker is root.

**Podman addresses this by never running a root daemon.** Containers stay in your namespace; there's no way to jump out unless you literally give them privileges by running `podman run —privileged <image>`. Don't do that! Don't give your container a privileged access. It isn't bulletproof, but it's way more secure than Docker. That's why big corporations like Accenture, Red Hat, DHL and thousands more are using Podman.

![Podman logo with 3 seal heads sticking out from the water](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1758107516.webp)

### Check your Scripts
As programmers, we often copy and paste code snippets from various online sources. If you have no clue what the code does, I highly recommend pasting it into a trusted AI assistant and asking it to break it down and explain.

This practice helps mitigate the risks of executing unfamiliar code. It's especially important given that studies show developers using AI assistants are more prone to introducing vulnerabilities, often due to overconfidence in the generated code's security.

If you don't believe me, here are some articles talking about this issue:

- <a href="https://thehackernews.com/expert-insights/2025/08/ais-hidden-security-debt.html" title="Link to the article" target="_blank">AI's Hidden Security Debt</a>
- <a href="https://www.itpro.com/development/programming/369763/developers-introduce-security-vulnerabilities-ai-assistants" title="Link to the article" target="_blank">Developers more likely to introduce security vulnerabilities in code when using AI assistants</a>
- <a href="https://www.helpnetsecurity.com/2025/06/19/silviu-asandei-sonar-ai-code-assistants-security/" title="Link to the article" target="_blank">Why AI code assistants need a security reality check</a>
- <a href="https://www.securecodewarrior.com/article/deep-dive-navigating-vulnerabilities-generated-by-ai-coding-assistants" title="Link to the article" target="_blank">Deep-dive: Navigating vulnerabilities generated by AI coding assistants</a>

These are just a few; there are a lot more you can find. This is not to say AI assistants are inherently bad; rather, their output requires scrutiny and validation.

Another critical moment is to avoid piping `curl` directly to `bash` and always review scripts before executing them in your terminal. Here is a bad example that you should never do:

```bash
# NEVER PIPE CURL INTO BASH
curl https://example.com/some-script.sh | bash
```

Instead, do this:
```bash
# Download and save the script into a file
curl https://example.com/some-script.sh -sS -o script.sh

# Open with the editor of your choice
vim script.sh

# Check for sketchy stuff like rm, wget, or weird network calls

# Run only if it looks safe
bash script.sh
```

Tip. Use `curl -sSf` for best practices. `-s` to hide the download noise, `-S` to show errors if anything goes wrong, and `-f` to totally fail if the server returns an error. This prevents running a broken or malicious script from a failed download.

This meticulous review process is particularly crucial for shell scripts, as they often execute with elevated privileges and can access sensitive system resources.

## Network Security
Network security is another important area that demands extra attention, especially if you work remotely. Remote work is often done in coffee shops, libraries, or other public spaces, which can expose an individual to significant security risk.

Using things like a VPN, firewall, and custom DNS can help mitigate these risks by encrypting traffic and securing connections over untrusted networks. Let’s start with VPNs and discuss their benefits.

### Virtual Private Network
Back then, dial-up was slow, but it felt safe, even though your ISP (Internet Service Provider) could monitor all your traffic since HTTPS wasn’t widely adopted yet. Then Wi-Fi rolled into cafés and airports, and suddenly a stranger next to you could sniff everything: your URLs, your passwords, your credit card information. Most sites still used plain HTTP, no encryption. Spying was free. VPNs (Virtual Private Networks) weren't fancy; they were the only wall between you and anyone who cared to look.

<blockquote>
    <p>If you go to a coffee shop or at the airport, and you're using open wireless, I would use a VPN service that you could subscribe for 10 bucks a month. Everything is encrypted in an encryption tunnel, so a hacker cannot tamper with your connection.</p>
    <img src="https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1758197358.webp" alt="Kevin Mitnick in glasses, hacker and a computer security consultant" />
    <h6>Kevin Mitnick (hacker, computer security consultant)</h6>
</blockquote>

Today, VPNs remain a critical component of privacy, especially in combination with TLS (Transport Layer Security). Always enable VPN when you are connected to a public network to obscure your traffic even when you use HTTPS. This is because TLS doesn’t encrypt all metadata while a VPN encrypts all traffic between your device and the VPN server so that the attacker only sees that you are communicating with a VPN server.

### Domain Name System
The DNS (Domain Name System) is a foundational component of Internet communication, translating human-readable domain names into numerical IP addresses necessary for establishing connections. However, if you use the default (automatic) DNS settings, it means you use the DNS of your Internet provider, and it can monitor your browsing activities and potentially intercept your traffic, which is mostly the case.

It can be clearly demonstrated with websites that are blocked in your country. Moreover, ISP DNS is trash for privacy and security because it never blocks malicious domains such as phishing sites, making users vulnerable to various cyber threats.

Note. Phishing sites are fake websites that look real and are designed to trick you into sharing passwords, credit card information, or personal data. They often mimic banks, social media, or email logins, using similar URLs or designs. Clicking links from shady emails or texts can lead you there. This is where secure DNS providers like Cloudflare and Quad9 come in—they automatically block known phishing and malicious domains, protecting you from these threats.

If you still rely on your ISP's default DNS, consider changing your DNS settings to point to a different free provider like Quad9 or Cloudflare. It’s super easy to do; you don’t need any special software. Depending on your OS and desktop environment, it will look different, but the concept is still the same.

You just need to go into your WiFi settings and configure your current WiFi connection by disabling the automatic DNS setting. On my GNOME desktop environment, I find my WiFi connection `WWWmi` and changing automatic DNS to 2 IP addresses `9.9.9.9, 149.112.112.112` which you can find on the <a href="https://quad9.net/" title="Link to Quad9 official website" target="_blank">Quad9</a> website if you choose it as your DNS provider. For <a href="https://developers.cloudflare.com/1.1.1.1/ip-addresses/" title="Link to the official Cloudflare website" target="_blank">Cloudflare</a>, it will be `1.1.1.1, 1.0.0.1`.

![Configuration window for WiFi connection for GNOME desktop environment](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1758207136.webp)

There are many other DNS providers that offer enhanced privacy features and malware blocking capabilities, but if you want the maximum level of privacy, you should consider using your own DNS server. However, it might be overkill for most of us. If you are into the hacking side of things, I recommend buying a Raspberry Pi and setting up a DNS server with firewall capabilities.

![Raspberry Pi 5 in a case connected with USB cables connected](https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1758210608.webp)

I can write a series of articles just about DNS because there are so many things I would love to talk about. However, I’m going to stop here and move to the firewall section.

### Firewall
A firewall acts as a critical barrier, controlling incoming and outgoing network traffic based on predetermined security rules, thereby preventing unauthorized access and mitigating various cyber threats. By default, a firewall is enabled on Windows 11, macOS, and many Linux distributions like Fedora Workstation, Ubuntu, Debian, etc.

If you use Debian-based distros like Ubuntu or Linux Mint, run this to check if your firewall is enabled:

```bash
# ufw stands for Uncomplicated FireWall
sudo ufw status
```

If it's not running, make sure you enable it with the enable command instead of status.

On a Fedora Workstation, you can check it by running:

```bash
sudo firewall-cmd --state
# or, as an alternative, you can run this:
sudo systemctl status firewalld
```

To enable it, you just run:

```bash
# The '--now' flag starts the service immediately
sudo systemctl enable --now firewalld
```

For other distros, you need to find it yourself since each distribution often implements its firewall management tools and configurations. For most users, the default firewall settings provide robust protection without impacting usability, making it a critical and effortless layer of defense.

## Conclusion
Securing your digital life is an ongoing process. Start with the fundamentals (strong passwords, encryption), then gradually replace tools and adopt more advanced practices (using Tor, avoiding big tech).

Try Linux if you are still not using it and make sure the ISO file has the correct checksum. Start moving to a password manager and always enable full-disk encryption to protect your data.

**Start using a VPN. Modern VPNs have minimal impact on performance, often with only a negligible speed reduction.** This minimalized performance overhead is a small price to pay for the significant enhancement in security and privacy, especially when traversing untrusted networks. You can disable it when you watch a movie or access content that might block the country your VPN is connected to.

Changing DNS is not even a technical challenge; it's free and it takes a maximum of 7 minutes to set up. If you are not sure which DNS to use, just use Quad9 or Cloudflare. Avoid Google DNS and your ISP’s default DNS, as these can compromise your privacy.

<blockquote>
    <p>Privacy matters; privacy is what allows us to determine who we are and who we want to be.</p>
    <img src="https://serhiicho.com/storage/posts/content/securing-your-digital-life-a-programmers-action-plan-1758213689.webp" alt="A photo of Edward Snowden" />
    <h6>Edward Snowden</h6>
</blockquote>

If I could leave you with only a single piece of advice, it would be to **assume that every program you install is a potential threat to your system**. It's not being paranoid; it's being stoical and pragmatic. This mindset aligns with the <a href="https://en.wikipedia.org/wiki/Zero_trust_architecture" title="Link to the Wikipedia article about Zero Trust" target="_blank">Zero Trust</a> security model, which advocates for strict identity verification for every user and device.

Remember, the goal is to make it as difficult as possible for anyone to track you or access your data. Take care and leave a comment down below with anything you want; just let me know that you were here 😉.
