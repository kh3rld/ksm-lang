# String Documentation

## Overview
A powerful, flexible, and efficient system for string manipulation. The Str type is central to this system, offering various methods that simplify working with text. From basic operations like concatenation and length calculation to advanced features like substring extraction and string reversal, making it easy to work with strings.

## String Type and Creation

### Creating a String

The `Str` type is used to represent a string in KSM. You can create a new string using the `New()` method.

#### Example:

```go
// Create a new string
name := New("John Doe")

// Create an empty string
empty := New("")
```

## Basic Methods
Here’s a quick overview of the common string manipulation methods available:

```go
s := New("hello")         // Create a new string
s.Len()                   // Get string length
s.Concatenate(New(" world"))  // Concatenate another string
s.Sub(0, 2)               // Get substring
s.Up()                    // Convert to uppercase
s.Low()                   // Convert to lowercase
s.Has(New("ll"))          // Check if substring exists
s.Rev()                   // Reverse the string
s.Rep(3)                  // Repeat the string
s.Cut(New("-"))           // Split string by delimiter
s.Trim()                  // Remove whitespace
```

## Detailed Methods and Use Cases
`Len()`: **Get String Length**

The `Len()` method returns the number of characters in a string.

```go
greeting := New("hello")
length := greeting.Len()  // 5
```

`Concatenate()`: **Join Two Strings**

`Concatenate()` joins two strings together and returns the result.

Example:
```go
greeting := New("Hello")
fullGreeting := greeting.Concatenate(New(" World"))  // "Hello World"
```

This method can be used to combine names, messages, or any strings you need to join together.

`Sub()`: **Extract Substring**

`Sub()` extracts a part of the string, starting from the start index to the end index. It automatically adjusts out-of-bounds indices and handles negative values.

Example:
```go
text := New("hello")
part := text.Sub(0, 2)   // "he"
```
If start is greater than end, it returns an empty string. You can use this method for slicing strings in various scenarios, like retrieving portions of a sentence or processing tokens.

`Up()`: **Convert to Uppercase**

`Up()` converts all characters in a string to uppercase.

Example:
```go
text := New("hello")
uppercaseText := text.Up()  // "HELLO"
```

`Low()`: **Convert to Lowercase**
Similarly, `Low()` converts the string to lowercase.

Example:
```go
Copy code
text := New("HELLO")
lowercaseText := text.Low()  // "hello"
```

`Has()`: **Check Substring Existence**

`Has()` checks if a string contains a given substring.

Example:
```go
text := New("hello")
hasSubstring := text.Has(New("ll"))  // true
```

You can use this method to search for specific keywords or phrases in user input or other strings.

`Rev()`: **Reverse the String**

`Rev()` reverses the order of characters in the string.

Example:
```go
text := New("hello")
reversedText := text.Rev()  // "olleh"
```

This can be useful in algorithms that require palindromic checks or reversed sequences.

`Rep()`: **Repeat the String**

`Rep()` repeats the string n times and returns the result. It handles non-positive values by returning an empty string.

Example:
```go
text := New("ha")
repeatedText := text.Rep(3)  // "hahaha"
```

Great for generating repeated patterns, visual formatting, or constructing larger strings from smaller components.

`Cut()`: **Split the String by a Delimiter**

`Cut()` splits the string into parts based on the given delimiter and returns a slice of strings. Each part is represented as a Str.

Example:
```go
text := New("a-b-c")
parts := text.Cut(New("-"))  // ["a", "b", "c"]
```

Useful for processing data that’s delimited by a character, such as CSV files or user input containing separators.

`Trim()`: **Remove Whitespace from Both Ends**

`Trim()` removes leading and trailing spaces, tabs, and newlines from the string.

Example:
```go
text := New("  hello  ")
trimmedText := text.Trim()  // "hello"
```

This method is essential when sanitizing user input or working with data that might have inconsistent spacing.

## Chaining Methods
All methods on Str return a new Str object, making it easy to chain multiple operations together.

Example:
```go
result := New(" hello ")
    .Trim()
    .Up()
    .Rev()  // "OLLEH"
```

Here, we trim the string, convert it to uppercase, and then reverse it, all in a single expression.

## Error Handling
Here string methods are designed to handle errors gracefully:

**Empty Strings**: Operations like `Concatenate()`, `Sub()`, and `Has()` work seamlessly with empty strings without causing errors.

**Out-of-Bounds Indices**: `Sub()` adjusts invalid indices to valid ranges, ensuring safe slicing even if the indices are out of range.

**Negative or Zero Repeat Counts**: If you pass a non-positive value to `Rep()`, it returns an empty string without raising an error.

Example:
```go
text := New("hello")
invalidSub := text.Sub(-5, 10)  // "hello", corrected indices
emptyRepeat := text.Rep(0)      // ""
```

## Example Use Cases
1. **Constructing User Greetings**
You can use concatenation and string methods to dynamically generate messages:

```go
name := New("John")
greeting := New("Hello, ").Concatenate(name).Concatenate(New("!"))  // "Hello, John!"
```

2. **Validating Input**
Check if a user’s input contains a specific substring and handle it accordingly:

```go
input := New("reset password")
if input.Has(New("password")) {
    // Take some action
}
```

3. **Processing CSV-like Data**
With `Cut()`, you can split strings containing delimited data:

```go
csvLine := New("John,Doe,30")
fields := csvLine.Cut(New(","))  // ["John", "Doe", "30"]
```

4. **Formatting Data**
Use `Up()`, `Low()`, and `Trim()` to format data before storing or displaying:

```go
rawData := New("  alice  ")
formattedData := rawData.Trim().Up()  // "ALICE"
```
