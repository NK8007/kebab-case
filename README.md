# kebab-case

## Description

This repository contains a utility for converting strings to kebab-case format. Kebab-case is a lowercase string format where words are separated by hyphens (`-`). It is commonly used in URLs, file names, and CSS class names.

## Usage

To use this utility, simply import the function and pass the string you want to convert:

```javascript
import { toKebabCase } from 'kebab-case';

const result = toKebabCase('Hello World');
console.log(result); // Output: hello-world
```

## API

### `toKebabCase(input: string): string`

Converts the given string to kebab-case.

#### Parameters

- `input` (string): The string to convert.

#### Returns

- (string): The kebab-case formatted string.    
- Example:

```javascript
toKebabCase('Hello World'); // Output: 'hello-world'
``` 

