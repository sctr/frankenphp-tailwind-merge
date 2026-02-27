# tailwind-merge for PHP

Utility function to efficiently merge Tailwind CSS classes in PHP without style conflicts.

A [FrankenPHP](https://frankenphp.dev/) extension written in Go. Port of the Tailwind v4 class resolution logic from [tailwind_merge](https://github.com/gjtorikian/tailwind_merge) (Ruby), inspired by [tailwind-merge](https://github.com/dcastil/tailwind-merge) (JS).

- Supports Tailwind v4
- Works with FrankenPHP worker mode
- In-memory LRU cache powered by Go — shared across requests
- Zero PHP dependencies

```php
tailwind_merge(['px-2 py-1 bg-red hover:bg-dark-red', 'p-3 bg-[#B91C1C]']);
// → 'hover:bg-dark-red p-3 bg-[#B91C1C]'
```

## Get started

- [What is it for?](#what-is-it-for)
- [Installation](#installation)
- [Usage](#usage)
- [API reference](#api-reference)
- [How it works](#how-it-works)

## What is it for?

If you use Tailwind with a component-based templating approach (Blade, Twig, Livewire, etc.), you probably run into class conflicts when composing components. When two conflicting classes exist in the same string, CSS specificity — not source order — determines which one applies. This makes overriding styles unreliable.

`tailwind_merge()` solves this by intelligently resolving conflicts so the last class always wins:

```php
// Without tailwind_merge — conflicting padding, unpredictable result
'px-2 py-1' . ' ' . 'p-3'
// → "px-2 py-1 p-3" ← browser may apply px-2/py-1 OR p-3 depending on stylesheet order

// With tailwind_merge — p-3 overrides px-2 and py-1
tailwind_merge(['px-2 py-1', 'p-3']);
// → "p-3"
```

The primary use case is component override patterns:

```php
function button(string $class = ''): string
{
    return tailwind_merge([
        'inline-flex items-center px-4 py-2 bg-blue-600 text-white font-medium rounded-md',
        $class,
    ]);
}

button('bg-red-600 py-3');
// → "inline-flex items-center px-4 py-3 bg-red-600 text-white font-medium rounded-md"
```

## Installation

### Building with FrankenPHP

Add the module when building FrankenPHP:

```bash
CGO_ENABLED=1 \
XCADDY_GO_BUILD_FLAGS="-ldflags='-w -s'" \
CGO_CFLAGS=$(php-config --includes) \
CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" \
xcaddy build \
    --output frankenphp \
    --with github.com/sctr/frankenphp-tailwind-merge
```

### Requirements

- [FrankenPHP](https://frankenphp.dev/) with extension support
- Go 1.22+
- PHP 8.2+

### Verify

```php
var_dump(extension_loaded('tailwind_merge'));
// bool(true)
```

## Usage

### Basic merging

```php
tailwind_merge(['px-2 py-1 bg-red-500', 'p-3 bg-blue-500']);
// → "p-3 bg-blue-500"
```

### Conflict resolution

```php
// Shorthand overrides longhand
tailwind_merge(['px-2 py-1', 'p-3']);
// → "p-3"

// Last conflicting class wins
tailwind_merge(['text-red-500', 'text-blue-500']);
// → "text-blue-500"

// Modifiers are scoped independently
tailwind_merge(['hover:bg-red-500', 'hover:bg-blue-500', 'bg-green-500']);
// → "hover:bg-blue-500 bg-green-500"

// Arbitrary values
tailwind_merge(['bg-red-500', 'bg-[#1da1f2]']);
// → "bg-[#1da1f2]"

// Non-Tailwind classes are preserved
tailwind_merge(['my-custom-class px-2', 'px-4']);
// → "my-custom-class px-4"
```

### Laravel Blade

```php
// resources/views/components/button.blade.php
<button {{ $attributes->merge(['class' => tailwind_merge([
    'inline-flex items-center px-4 py-2 bg-blue-600 text-white rounded-md',
    $attributes->get('class', ''),
])]) }}>
    {{ $slot }}
</button>
```

```html
<x-button class="bg-red-600 py-3">Delete</x-button>
<!-- → "inline-flex items-center px-4 py-3 bg-red-600 text-white rounded-md" -->
```

### Features

| Feature | Example | Result |
|---------|---------|--------|
| Last class wins | `['text-sm', 'text-lg']` | `text-lg` |
| Shorthand overrides longhand | `['px-2 py-1', 'p-4']` | `p-4` |
| Modifier-aware | `['hover:text-sm', 'hover:text-lg']` | `hover:text-lg` |
| Arbitrary values | `['bg-red-500', 'bg-[#B91C1C]']` | `bg-[#B91C1C]` |
| Important modifier | `['!font-bold', '!font-thin']` | `!font-thin` |
| Postfix modifiers | `['text-lg/7', 'text-lg/8']` | `text-lg/8` |
| Non-TW classes preserved | `['custom px-2', 'px-4']` | `custom px-4` |

## API reference

### `tailwind_merge(array $classes): string`

Merges an array of Tailwind CSS class strings, resolving conflicts by keeping the last conflicting class.

| Parameter | Type | Description |
|-----------|------|-------------|
| `$classes` | `array<string>` | Array of class strings to merge |

**Returns:** `string` — The merged class string with conflicts resolved.

## How it works

1. **Cache lookup** — A Go-powered LRU cache checks if this exact input was seen before. On hit, returns instantly.
2. **Parse** — Each class string is split into individual classes, then parsed into groups, modifiers, and values.
3. **Resolve** — Conflicting classes are identified using Tailwind's class group hierarchy. Last conflicting class wins.
4. **Cache & return** — The result is stored in the LRU cache and returned to PHP.

The cache lives in Go memory and persists across PHP requests in FrankenPHP worker mode. This means:

- **No serialization overhead** — No PHP serialize/unserialize, no Redis, no filesystem
- **Shared across requests** — All PHP workers share the same cache
- **Bounded memory** — LRU eviction keeps memory usage predictable

## Acknowledgments

- [tailwind-merge](https://github.com/dcastil/tailwind-merge) (JS) by Dany Castillo — the original implementation and API inspiration
- [tailwind_merge](https://github.com/gjtorikian/tailwind_merge) (Ruby) by Garen Torikian — Tailwind v4 class resolution logic this port is based on
- [FrankenPHP](https://frankenphp.dev/) by Kevin Dunglas — Go-powered PHP extensions

## License

[MIT](LICENSE)
