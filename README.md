<h3 align="center"> Easy Regex </h3>
<p align="center">
  <a href="#usage-examples">Usage Examples</a> •
  <a href="#usage-parameters">Usage Parameters</a> •
  <a href="#installation">Install</a> •
  <a href="https://t.me/ImAyrix">Contact me</a>
</p>

---

"Easy Regex" is a valuable and straightforward tool that takes your regular expression and showcases the matched parts of the text in the output. Give it a try and witness its usefulness in action.

## Usage Examples
Suppose we possess a file that encompasses a plethora of information, comprising ASN numbers. Our objective is to extract solely the ASN numbers. Let us witness how our tool accomplishes this task for us... 

```
▶ cat things.txt
The foremost ASN, AS11111, reigns supreme while the second in line, AS2222, follows closely behind.
Additionally, we possess another noteworthy ASN number - AS3333.

▶ cat things.txt | er -r "AS\d+"
AS11111
AS2222
AS3333

▶ cat things.txt | er -r "AS(\d+)" -g 1
11111
2222
3333
```

## Usage Parameters
```
  -r string
        Your regex [e.g: ^\w+]
  -g int
        Group number
  -h    Help, show usage parameters

```

## Installation
```
go install github.com/ImAyrix/er@latest
```
