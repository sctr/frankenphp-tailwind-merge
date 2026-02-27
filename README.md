# FrankenPHP Extension for Tailwind Merge

A high-performance [PHP](https://php.net) utility for merging [Tailwind CSS](https://tailwindcss.com) classes without style conflicts, designed to work with [FrankenPHP](https://frankenphp.dev).
It leverages Go to perform intelligent class resolution based on Tailwind v4's utility hierarchy, ported from [tailwind_merge](https://github.com/gjtorikian/tailwind_merge) (Ruby) and inspired by [tailwind-merge](https://github.com/dcastil/tailwind-merge) (JS).

This extension provides a single `tailwind_merge()` function that can be shared across all requests and worker script instances,
backed by an in-memory LRU cache powered by Go — ensuring efficient resource usage and high performance with zero PHP dependencies.

```php
tailwind_merge(['px-2 py-1 bg-red hover:bg-dark-red', 'p-3 bg-[#B91C1C]']);
// → 'hover:bg-dark-red p-3 bg-[#B91C1C]'
```

## Installation

First, [install FrankenPHP](https://frankenphp.dev/docs/) and its dependencies including a ZTS (Zend Thread Safety) build of libphp and [xcaddy](https://github.com/caddyserver/xcaddy).

Then, compile FrankenPHP with the extension:

```console
CGO_ENABLED=1 \
XCADDY_GO_BUILD_FLAGS="-ldflags='-w -s'" \
CGO_CFLAGS=$(php-config --includes) \
CGO_LDFLAGS="$(php-config --ldflags) $(php-config --libs)" \
xcaddy build \
    --output frankenphp \
    --with github.com/dunglas/frankenphp/caddy \
    --with github.com/sctr/frankenphp-tailwind-merge
    # Add extra Caddy modules and FrankenPHP extensions here
```

That's it! Your custom FrankenPHP build now contains the `tailwind_merge` extension.

Verify it's loaded:

```php
var_dump(extension_loaded('tailwind_merge'));
// bool(true)
```

## Usage

If you use Tailwind with a component-based templating approach (Blade, Twig, Livewire, etc.), you'll run into class conflicts when composing components. When two conflicting classes exist in the same string, CSS specificity — not source order — determines which one applies, making overrides unreliable.

`tailwind_merge()` solves this by intelligently resolving conflicts so the last class always wins:

```php
// Without tailwind_merge — conflicting padding, unpredictable result
'px-2 py-1' . ' ' . 'p-3'
// → "px-2 py-1 p-3" ← browser may apply px-2/py-1 OR p-3 depending on stylesheet order

// With tailwind_merge — p-3 overrides px-2 and py-1
tailwind_merge(['px-2 py-1', 'p-3']);
// → "p-3"

// Last conflicting class wins
tailwind_merge(['text-red-500', 'text-blue-500']);
// → "text-blue-500"

// Modifiers are scoped independently
tailwind_merge(['hover:bg-red-500', 'hover:bg-blue-500', 'bg-green-500']);
// → "hover:bg-blue-500 bg-green-500"

// Arbitrary values work too
tailwind_merge(['bg-red-500', 'bg-[#1da1f2]']);
// → "bg-[#1da1f2]"

// Non-Tailwind classes are always preserved
tailwind_merge(['my-custom-class px-2', 'px-4']);
// → "my-custom-class px-4"
```

The primary use case is component override patterns. Here is an example using a plain PHP function:

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

And here is an equivalent Laravel Blade component:

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
<!-- renders with: "inline-flex items-center px-4 py-3 bg-red-600 text-white rounded-md" -->
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

### How it works

1. **Cache lookup** — A Go-powered LRU cache checks if this exact input was seen before. On hit, returns instantly.
2. **Parse** — Each class string is split into individual classes, then parsed into groups, modifiers, and values.
3. **Resolve** — Conflicting classes are identified using Tailwind's class group hierarchy. The last conflicting class wins.
4. **Cache & return** — The result is stored in the LRU cache and returned to PHP.

The cache lives in Go memory and persists across PHP requests in FrankenPHP worker mode. All PHP workers share the same cache with no serialization overhead and bounded memory via LRU eviction.

## Credits

- [tailwind-merge](https://github.com/dcastil/tailwind-merge) (JS) by Dany Castillo — the original implementation and API inspiration
- [tailwind_merge](https://github.com/gjtorikian/tailwind_merge) (Ruby) by Garen Torikian — Tailwind v4 class resolution logic this port is based on
- [FrankenPHP](https://frankenphp.dev/) by Kévin Dunglas — Go-powered PHP extensions

## License

[MIT](LICENSE)
