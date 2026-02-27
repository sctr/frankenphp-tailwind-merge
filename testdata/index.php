<?php

// Test: extension is loaded
echo "extension_loaded: " . (extension_loaded('tailwind_merge') ? 'true' : 'false') . "\n";

// Test: shorthand overrides longhand
echo "shorthand: " . tailwind_merge(['px-2 py-1', 'p-3']) . "\n";

// Test: last conflicting wins
echo "last_wins: " . tailwind_merge(['text-red-500', 'text-blue-500']) . "\n";

// Test: modifier scoping
echo "modifiers: " . tailwind_merge(['hover:bg-red-500', 'hover:bg-blue-500', 'bg-green-500']) . "\n";

// Test: arbitrary values
echo "arbitrary: " . tailwind_merge(['bg-red-500', 'bg-[#1da1f2]']) . "\n";

// Test: non-TW preserved
echo "non_tw: " . tailwind_merge(['my-custom-class px-2', 'px-4']) . "\n";

// Test: important modifier
echo "important: " . tailwind_merge(['!font-bold', '!font-thin']) . "\n";

// Test: empty input
echo "empty: " . tailwind_merge([]) . "\n";

// Test: hero example
echo "hero: " . tailwind_merge(['px-2 py-1 bg-red hover:bg-dark-red', 'p-3 bg-[#B91C1C]']) . "\n";
