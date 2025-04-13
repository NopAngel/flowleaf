# FlowLeaf

FlowLeaf is a programming language (superset) of HTML, which allows you to create large applications with just a few lines of code. FlowLeaf corrects errors in HTML itself.

## Characteristics

1. **Optimization**: Allows you to write code in a clean, fast way.
2. **Convenience**: FlowLeaf is optimized and comfortable to write in, allowing you to create pages in seconds!
3. **Server**: FlowLeaf comes with its own server, to run FlowLeaf

It's probably going to give some errors, for now it's in development :)

## Install

### How to install?

To install FlowLeaf you have to follow these steps:

1. **Install**: Go to the link [here](https://github.com/NopAngel/flowleaf/releases/download/flowleaf-version-1.0/FlowLeaf.exe), and you will download FlowLeaf, you can also go to GitHub Releases and download it
2. After installing it you will need to create a file with the `flowleaf` extension.

And that's it, you can now create FlowLeaf code.

### Documentation

**FlowLeaf** is based on the syntax of a framework called **[Astro.js](https://astro.build)**, for its simplicity when writing code.

To create tags, you have to create the **HTML** structure, as is done in **FlowLeaf**? Well, you simply have to write: `Body:`

And everything that follows the Body will be what's in the body, and from there you continue the code. If you want to write a `section, h1, p`, or any other tag, just put it after the Body, and write it simply like this:

```html
tag: child element:
```

To close the tag, you'll have to repeat it, but at the beginning, put a **"|"**, e.g.:

```
Body:
    h1:
        Hi from FlowLeaf!!
    |h1
```

The code doesn't necessarily have to be **indented**, but if you want to write **comfortably**, it's best if it is.

At first it may seem uncomfortable or strange, but then you'll get used to it :)

To create a class for an element, simply put the element (tag) followed by a "." and after that point, the name of the class, e.g.:

```
parent_element.class:

    child_item.class_child
```

## These are the tags that are in FlowLeaf:

**H1 (up to H6)** <br />
**UL** <br />
**P** <br />
**_SPAN_** <br />
**_U_** <br />
**_HEADER_** <br />
**_FOOTER_** <br />
**_BR_** <br />
**_I_** <br />
**_TABLE_** <br />
**_TD_** <br />
**_TR_** <br />
**_TH_** <br />
**_TBODY_** <br />
**_CAPTION_** <br />
**_TFOOT_** <br />
**_THEAD_** <br />
**_MAIN_** <br />
**_ARTICLE_** <br />
**_NAV_**<br />
**_FORM_**<br />
**_OPTION_**<br />
**_TEXTAREA_**<br />
**_LABEL_**<br />
**_LI_**<br />
