package twmerge

// GetDefaultConfig returns the default tailwind-merge configuration
// supporting Tailwind CSS v4.
func GetDefaultConfig() *Config {
	// Theme getters
	themeColor := FromTheme("color")
	themeFont := FromTheme("font")
	themeText := FromTheme("text")
	themeFontWeight := FromTheme("font-weight")
	themeTracking := FromTheme("tracking")
	themeLeading := FromTheme("leading")
	themeBreakpoint := FromTheme("breakpoint")
	themeContainer := FromTheme("container")
	themeSpacing := FromTheme("spacing")
	themeRadius := FromTheme("radius")
	themeShadow := FromTheme("shadow")
	themeInsetShadow := FromTheme("inset-shadow")
	themeTextShadow := FromTheme("text-shadow")
	themeDropShadow := FromTheme("drop-shadow")
	themeBlur := FromTheme("blur")
	themePerspective := FromTheme("perspective")
	themeAspect := FromTheme("aspect")
	themeEase := FromTheme("ease")
	themeAnimate := FromTheme("animate")

	// Scale helpers
	scaleBreak := func() []ClassDefinition {
		return d("auto", "avoid", "all", "avoid-page", "page", "left", "right", "column")
	}

	scalePosition := func() []ClassDefinition {
		return d("center", "top", "bottom", "left", "right",
			"top-left", "left-top",
			"top-right", "right-top",
			"bottom-right", "right-bottom",
			"bottom-left", "left-bottom")
	}

	scalePositionWithArbitrary := func() []ClassDefinition {
		return append(scalePosition(), IsArbitraryVariable, IsArbitraryValue)
	}

	scaleOverflow := func() []ClassDefinition {
		return d("auto", "hidden", "clip", "visible", "scroll")
	}

	scaleOverscroll := func() []ClassDefinition {
		return d("auto", "contain", "none")
	}

	scaleUnambiguousSpacing := func() []ClassDefinition {
		return []ClassDefinition{IsArbitraryVariable, IsArbitraryValue, themeSpacing}
	}

	scaleInset := func() []ClassDefinition {
		return append([]ClassDefinition{IsFraction, "full", "auto"}, scaleUnambiguousSpacing()...)
	}

	scaleGridTemplateColsRows := func() []ClassDefinition {
		return []ClassDefinition{IsInteger, "none", "subgrid", IsArbitraryVariable, IsArbitraryValue}
	}

	scaleGridColRowStartAndEnd := func() []ClassDefinition {
		return []ClassDefinition{
			"auto",
			map[string][]ClassDefinition{
				"span": {"full", IsInteger, IsArbitraryVariable, IsArbitraryValue},
			},
			IsInteger, IsArbitraryVariable, IsArbitraryValue,
		}
	}

	scaleGridColRowStartOrEnd := func() []ClassDefinition {
		return []ClassDefinition{IsInteger, "auto", IsArbitraryVariable, IsArbitraryValue}
	}

	scaleGridAutoColsRows := func() []ClassDefinition {
		return []ClassDefinition{"auto", "min", "max", "fr", IsArbitraryVariable, IsArbitraryValue}
	}

	scaleAlignPrimaryAxis := func() []ClassDefinition {
		return d("start", "end", "center", "between", "around", "evenly",
			"stretch", "baseline", "center-safe", "end-safe")
	}

	scaleAlignSecondaryAxis := func() []ClassDefinition {
		return d("start", "end", "center", "stretch", "center-safe", "end-safe")
	}

	scaleMargin := func() []ClassDefinition {
		return append([]ClassDefinition{"auto"}, scaleUnambiguousSpacing()...)
	}

	scaleSizing := func() []ClassDefinition {
		return append([]ClassDefinition{
			IsFraction, "auto", "full", "dvw", "dvh", "lvw", "lvh", "svw", "svh",
			"min", "max", "fit",
		}, scaleUnambiguousSpacing()...)
	}

	scaleSizingInline := func() []ClassDefinition {
		return append([]ClassDefinition{
			IsFraction, "screen", "full", "dvw", "lvw", "svw",
			"min", "max", "fit",
		}, scaleUnambiguousSpacing()...)
	}

	scaleSizingBlock := func() []ClassDefinition {
		return append([]ClassDefinition{
			IsFraction, "screen", "full", "lh", "dvh", "lvh", "svh",
			"min", "max", "fit",
		}, scaleUnambiguousSpacing()...)
	}

	scaleColor := func() []ClassDefinition {
		return []ClassDefinition{themeColor, IsArbitraryVariable, IsArbitraryValue}
	}

	scaleBgPosition := func() []ClassDefinition {
		return append(scalePosition(),
			IsArbitraryVariablePosition, IsArbitraryPosition,
			map[string][]ClassDefinition{
				"position": {IsArbitraryVariable, IsArbitraryValue},
			},
		)
	}

	scaleBgRepeat := func() []ClassDefinition {
		return []ClassDefinition{
			"no-repeat",
			map[string][]ClassDefinition{
				"repeat": {"", "x", "y", "space", "round"},
			},
		}
	}

	scaleBgSize := func() []ClassDefinition {
		return []ClassDefinition{
			"auto", "cover", "contain",
			IsArbitraryVariableSize, IsArbitrarySize,
			map[string][]ClassDefinition{
				"size": {IsArbitraryVariable, IsArbitraryValue},
			},
		}
	}

	scaleGradientStopPosition := func() []ClassDefinition {
		return []ClassDefinition{IsPercent, IsArbitraryVariableLength, IsArbitraryLength}
	}

	scaleRadius := func() []ClassDefinition {
		return []ClassDefinition{"", "none", "full", themeRadius, IsArbitraryVariable, IsArbitraryValue}
	}

	scaleBorderWidth := func() []ClassDefinition {
		return []ClassDefinition{"", IsNumber, IsArbitraryVariableLength, IsArbitraryLength}
	}

	scaleBlendMode := func() []ClassDefinition {
		return d("normal", "multiply", "screen", "overlay", "darken", "lighten",
			"color-dodge", "color-burn", "hard-light", "soft-light", "difference",
			"exclusion", "hue", "saturation", "color", "luminosity")
	}

	scaleMaskImagePosition := func() []ClassDefinition {
		return []ClassDefinition{IsNumber, IsPercent, IsArbitraryVariablePosition, IsArbitraryPosition}
	}

	scaleRotate := func() []ClassDefinition {
		return []ClassDefinition{"none", IsNumber, IsArbitraryVariable, IsArbitraryValue}
	}

	scaleScale := func() []ClassDefinition {
		return []ClassDefinition{"none", IsNumber, IsArbitraryVariable, IsArbitraryValue}
	}

	scaleTranslate := func() []ClassDefinition {
		return append([]ClassDefinition{IsFraction, "full"}, scaleUnambiguousSpacing()...)
	}

	return &Config{
		CacheSize: 500,

		Theme: map[string][]ClassDefinition{
			"animate":     {"spin", "ping", "pulse", "bounce"},
			"aspect":      {"video"},
			"blur":        {IsTshirtSize},
			"breakpoint":  {IsTshirtSize},
			"color":       {IsAny},
			"container":   {IsTshirtSize},
			"drop-shadow": {IsTshirtSize},
			"ease":        {"in", "out", "in-out"},
			"font":        {IsAnyNonArbitrary},
			"font-weight": {"thin", "extralight", "light", "normal", "medium", "semibold", "bold", "extrabold", "black"},
			"inset-shadow": {IsTshirtSize},
			"leading":     {"none", "tight", "snug", "normal", "relaxed", "loose"},
			"perspective":  {"dramatic", "near", "normal", "midrange", "distant", "none"},
			"radius":      {IsTshirtSize},
			"shadow":      {IsTshirtSize},
			"spacing":     {"px", IsNumber},
			"text":        {IsTshirtSize},
			"text-shadow": {IsTshirtSize},
			"tracking":    {"tighter", "tight", "normal", "wide", "wider", "widest"},
		},

		ClassGroups: map[string][]ClassDefinition{
			// Layout
			"aspect":         {m("aspect", d("auto", "square", IsFraction, IsArbitraryValue, IsArbitraryVariable, themeAspect)...)},
			"container":      {"container"},
			"columns":        {m("columns", IsNumber, IsArbitraryValue, IsArbitraryVariable, themeContainer)},
			"break-after":    {m("break-after", scaleBreak()...)},
			"break-before":   {m("break-before", scaleBreak()...)},
			"break-inside":   {m("break-inside", d("auto", "avoid", "avoid-page", "avoid-column")...)},
			"box-decoration": {m("box-decoration", d("slice", "clone")...)},
			"box":            {m("box", d("border", "content")...)},
			"display":        d("block", "inline-block", "inline", "flex", "inline-flex", "table", "inline-table", "table-caption", "table-cell", "table-column", "table-column-group", "table-footer-group", "table-header-group", "table-row-group", "table-row", "flow-root", "grid", "inline-grid", "contents", "list-item", "hidden"),
			"sr":             {"sr-only", "not-sr-only"},
			"float":          {m("float", d("right", "left", "none", "start", "end")...)},
			"clear":          {m("clear", d("left", "right", "both", "none", "start", "end")...)},
			"isolation":      {"isolate", "isolation-auto"},
			"object-fit":     {m("object", d("contain", "cover", "fill", "none", "scale-down")...)},
			"object-position": {m("object", scalePositionWithArbitrary()...)},
			"overflow":       {m("overflow", scaleOverflow()...)},
			"overflow-x":     {m("overflow-x", scaleOverflow()...)},
			"overflow-y":     {m("overflow-y", scaleOverflow()...)},
			"overscroll":     {m("overscroll", scaleOverscroll()...)},
			"overscroll-x":   {m("overscroll-x", scaleOverscroll()...)},
			"overscroll-y":   {m("overscroll-y", scaleOverscroll()...)},
			"position":       d("static", "fixed", "absolute", "relative", "sticky"),
			"inset":          {m("inset", scaleInset()...)},
			"inset-x":        {m("inset-x", scaleInset()...)},
			"inset-y":        {m("inset-y", scaleInset()...)},
			"start":          {m("inset-s", scaleInset()...), m("start", scaleInset()...)},
			"end":            {m("inset-e", scaleInset()...), m("end", scaleInset()...)},
			"inset-bs":       {m("inset-bs", scaleInset()...)},
			"inset-be":       {m("inset-be", scaleInset()...)},
			"top":            {m("top", scaleInset()...)},
			"right":          {m("right", scaleInset()...)},
			"bottom":         {m("bottom", scaleInset()...)},
			"left":           {m("left", scaleInset()...)},
			"visibility":     d("visible", "invisible", "collapse"),
			"z":              {m("z", IsInteger, "auto", IsArbitraryVariable, IsArbitraryValue)},

			// Flexbox and Grid
			"basis":          {m("basis", IsFraction, "full", "auto", themeContainer, IsArbitraryVariable, IsArbitraryValue, themeSpacing)},
			"flex-direction":  {m("flex", d("row", "row-reverse", "col", "col-reverse")...)},
			"flex-wrap":      {m("flex", d("nowrap", "wrap", "wrap-reverse")...)},
			"flex":           {m("flex", IsNumber, IsFraction, "auto", "initial", "none", IsArbitraryValue)},
			"grow":           {m("grow", d("", IsNumber, IsArbitraryVariable, IsArbitraryValue)...)},
			"shrink":         {m("shrink", d("", IsNumber, IsArbitraryVariable, IsArbitraryValue)...)},
			"order":          {m("order", IsInteger, "first", "last", "none", IsArbitraryVariable, IsArbitraryValue)},
			"grid-cols":      {m("grid-cols", scaleGridTemplateColsRows()...)},
			"col-start-end":  {m("col", scaleGridColRowStartAndEnd()...)},
			"col-start":      {m("col-start", scaleGridColRowStartOrEnd()...)},
			"col-end":        {m("col-end", scaleGridColRowStartOrEnd()...)},
			"grid-rows":      {m("grid-rows", scaleGridTemplateColsRows()...)},
			"row-start-end":  {m("row", scaleGridColRowStartAndEnd()...)},
			"row-start":      {m("row-start", scaleGridColRowStartOrEnd()...)},
			"row-end":        {m("row-end", scaleGridColRowStartOrEnd()...)},
			"grid-flow":      {m("grid-flow", d("row", "col", "dense", "row-dense", "col-dense")...)},
			"auto-cols":      {m("auto-cols", scaleGridAutoColsRows()...)},
			"auto-rows":      {m("auto-rows", scaleGridAutoColsRows()...)},
			"gap":            {m("gap", scaleUnambiguousSpacing()...)},
			"gap-x":          {m("gap-x", scaleUnambiguousSpacing()...)},
			"gap-y":          {m("gap-y", scaleUnambiguousSpacing()...)},
			"justify-content": {m("justify", append(scaleAlignPrimaryAxis(), "normal")...)},
			"justify-items":  {m("justify-items", append(scaleAlignSecondaryAxis(), "normal")...)},
			"justify-self":   {m("justify-self", d("auto", "start", "end", "center", "stretch", "center-safe", "end-safe")...)},
			"align-content":  {m("content", d("normal", "start", "end", "center", "between", "around", "evenly", "stretch", "baseline", "center-safe", "end-safe")...)},
			"align-items": {m("items", d("start", "end", "center", "stretch", "center-safe", "end-safe",
				map[string][]ClassDefinition{"baseline": {"", "last"}})...)},
			"align-self": {m("self", d("auto", "start", "end", "center", "stretch", "center-safe", "end-safe",
				map[string][]ClassDefinition{"baseline": {"", "last"}})...)},
			"place-content": {m("place-content", scaleAlignPrimaryAxis()...)},
			"place-items":   {m("place-items", d("start", "end", "center", "stretch", "center-safe", "end-safe", "baseline")...)},
			"place-self":    {m("place-self", d("auto", "start", "end", "center", "stretch", "center-safe", "end-safe")...)},

			// Spacing
			"p":   {m("p", scaleUnambiguousSpacing()...)},
			"px":  {m("px", scaleUnambiguousSpacing()...)},
			"py":  {m("py", scaleUnambiguousSpacing()...)},
			"ps":  {m("ps", scaleUnambiguousSpacing()...)},
			"pe":  {m("pe", scaleUnambiguousSpacing()...)},
			"pbs": {m("pbs", scaleUnambiguousSpacing()...)},
			"pbe": {m("pbe", scaleUnambiguousSpacing()...)},
			"pt":  {m("pt", scaleUnambiguousSpacing()...)},
			"pr":  {m("pr", scaleUnambiguousSpacing()...)},
			"pb":  {m("pb", scaleUnambiguousSpacing()...)},
			"pl":  {m("pl", scaleUnambiguousSpacing()...)},
			"m":   {m("m", scaleMargin()...)},
			"mx":  {m("mx", scaleMargin()...)},
			"my":  {m("my", scaleMargin()...)},
			"ms":  {m("ms", scaleMargin()...)},
			"me":  {m("me", scaleMargin()...)},
			"mbs": {m("mbs", scaleMargin()...)},
			"mbe": {m("mbe", scaleMargin()...)},
			"mt":  {m("mt", scaleMargin()...)},
			"mr":  {m("mr", scaleMargin()...)},
			"mb":  {m("mb", scaleMargin()...)},
			"ml":  {m("ml", scaleMargin()...)},
			"space-x":         {m("space-x", scaleUnambiguousSpacing()...)},
			"space-x-reverse": {"space-x-reverse"},
			"space-y":         {m("space-y", scaleUnambiguousSpacing()...)},
			"space-y-reverse": {"space-y-reverse"},

			// Sizing
			"size":            {m("size", scaleSizing()...)},
			"inline-size":     {m("inline", append(d("auto"), scaleSizingInline()...)...)},
			"min-inline-size": {m("min-inline", append(d("auto"), scaleSizingInline()...)...)},
			"max-inline-size": {m("max-inline", append(d("none"), scaleSizingInline()...)...)},
			"block-size":      {m("block", append(d("auto"), scaleSizingBlock()...)...)},
			"min-block-size":  {m("min-block", append(d("auto"), scaleSizingBlock()...)...)},
			"max-block-size":  {m("max-block", append(d("none"), scaleSizingBlock()...)...)},
			"w":               {m("w", append([]ClassDefinition{themeContainer, "screen"}, scaleSizing()...)...)},
			"min-w":           {m("min-w", append([]ClassDefinition{themeContainer, "screen", "none"}, scaleSizing()...)...)},
			"max-w": {m("max-w", append([]ClassDefinition{themeContainer, "screen", "none", "prose",
				map[string][]ClassDefinition{"screen": {themeBreakpoint}}}, scaleSizing()...)...)},
			"h":     {m("h", append(d("screen", "lh"), scaleSizing()...)...)},
			"min-h":  {m("min-h", append(d("screen", "lh", "none"), scaleSizing()...)...)},
			"max-h":  {m("max-h", append(d("screen", "lh"), scaleSizing()...)...)},

			// Typography
			"font-size":      {m("text", d("base", themeText, IsArbitraryVariableLength, IsArbitraryLength)...)},
			"font-smoothing": d("antialiased", "subpixel-antialiased"),
			"font-style":     d("italic", "not-italic"),
			"font-weight":    {m("font", d(themeFontWeight, IsArbitraryVariableWeight, IsArbitraryWeight)...)},
			"font-stretch":   {m("font-stretch", d("ultra-condensed", "extra-condensed", "condensed", "semi-condensed", "normal", "semi-expanded", "expanded", "extra-expanded", "ultra-expanded", IsPercent, IsArbitraryValue)...)},
			"font-family":    {m("font", d(IsArbitraryVariableFamilyName, IsArbitraryFamilyName, themeFont)...)},
			"font-features":  {m("font-features", d(IsArbitraryValue)...)},
			"fvn-normal":     {"normal-nums"},
			"fvn-ordinal":    {"ordinal"},
			"fvn-slashed-zero": {"slashed-zero"},
			"fvn-figure":     d("lining-nums", "oldstyle-nums"),
			"fvn-spacing":    d("proportional-nums", "tabular-nums"),
			"fvn-fraction":   d("diagonal-fractions", "stacked-fractions"),
			"tracking":       {m("tracking", d(themeTracking, IsArbitraryVariable, IsArbitraryValue)...)},
			"line-clamp":     {m("line-clamp", IsNumber, "none", IsArbitraryVariable, IsArbitraryNumber)},
			"leading":        {m("leading", d(themeLeading, IsArbitraryVariable, IsArbitraryValue, themeSpacing)...)},
			"list-image":     {m("list-image", d("none", IsArbitraryVariable, IsArbitraryValue)...)},
			"list-style-position": {m("list", d("inside", "outside")...)},
			"list-style-type":     {m("list", d("disc", "decimal", "none", IsArbitraryVariable, IsArbitraryValue)...)},
			"text-alignment":      {m("text", d("left", "center", "right", "justify", "start", "end")...)},
			"placeholder-color":   {m("placeholder", scaleColor()...)},
			"text-color":          {m("text", scaleColor()...)},
			"text-decoration":     d("underline", "overline", "line-through", "no-underline"),
			"text-decoration-style": {m("decoration", d("solid", "dashed", "dotted", "double", "wavy")...)},
			"text-decoration-thickness": {m("decoration", IsNumber, "from-font", "auto", IsArbitraryVariable, IsArbitraryLength)},
			"text-decoration-color":     {m("decoration", scaleColor()...)},
			"underline-offset":   {m("underline-offset", IsNumber, "auto", IsArbitraryVariable, IsArbitraryValue)},
			"text-transform":     d("uppercase", "lowercase", "capitalize", "normal-case"),
			"text-overflow":      d("truncate", "text-ellipsis", "text-clip"),
			"text-wrap":          {m("text", d("wrap", "nowrap", "balance", "pretty")...)},
			"indent":             {m("indent", scaleUnambiguousSpacing()...)},
			"vertical-align":     {m("align", d("baseline", "top", "middle", "bottom", "text-top", "text-bottom", "sub", "super", IsArbitraryVariable, IsArbitraryValue)...)},
			"whitespace":         {m("whitespace", d("normal", "nowrap", "pre", "pre-line", "pre-wrap", "break-spaces")...)},
			"break":              {m("break", d("normal", "words", "all", "keep")...)},
			"wrap":               {m("wrap", d("break-word", "anywhere", "normal")...)},
			"hyphens":            {m("hyphens", d("none", "manual", "auto")...)},
			"content":            {m("content", d("none", IsArbitraryVariable, IsArbitraryValue)...)},

			// Backgrounds
			"bg-attachment": {m("bg", d("fixed", "local", "scroll")...)},
			"bg-clip":       {m("bg-clip", d("border", "padding", "content", "text")...)},
			"bg-origin":     {m("bg-origin", d("border", "padding", "content")...)},
			"bg-position":   {m("bg", scaleBgPosition()...)},
			"bg-repeat":     {m("bg", scaleBgRepeat()...)},
			"bg-size":       {m("bg", scaleBgSize()...)},
			"bg-image": {m("bg", d("none",
				map[string][]ClassDefinition{
					"linear": {
						map[string][]ClassDefinition{
							"to": {"t", "tr", "r", "br", "b", "bl", "l", "tl"},
						},
						IsInteger, IsArbitraryVariable, IsArbitraryValue,
					},
					"radial": {"", IsArbitraryVariable, IsArbitraryValue},
					"conic":  {IsInteger, IsArbitraryVariable, IsArbitraryValue},
				},
				IsArbitraryVariableImage, IsArbitraryImage)...),
			},
			"bg-color":           {m("bg", scaleColor()...)},
			"gradient-from-pos":  {m("from", scaleGradientStopPosition()...)},
			"gradient-via-pos":   {m("via", scaleGradientStopPosition()...)},
			"gradient-to-pos":    {m("to", scaleGradientStopPosition()...)},
			"gradient-from":      {m("from", scaleColor()...)},
			"gradient-via":       {m("via", scaleColor()...)},
			"gradient-to":        {m("to", scaleColor()...)},

			// Borders
			"rounded":    {m("rounded", scaleRadius()...)},
			"rounded-s":  {m("rounded-s", scaleRadius()...)},
			"rounded-e":  {m("rounded-e", scaleRadius()...)},
			"rounded-t":  {m("rounded-t", scaleRadius()...)},
			"rounded-r":  {m("rounded-r", scaleRadius()...)},
			"rounded-b":  {m("rounded-b", scaleRadius()...)},
			"rounded-l":  {m("rounded-l", scaleRadius()...)},
			"rounded-ss": {m("rounded-ss", scaleRadius()...)},
			"rounded-se": {m("rounded-se", scaleRadius()...)},
			"rounded-ee": {m("rounded-ee", scaleRadius()...)},
			"rounded-es": {m("rounded-es", scaleRadius()...)},
			"rounded-tl": {m("rounded-tl", scaleRadius()...)},
			"rounded-tr": {m("rounded-tr", scaleRadius()...)},
			"rounded-br": {m("rounded-br", scaleRadius()...)},
			"rounded-bl": {m("rounded-bl", scaleRadius()...)},
			"border-w":   {m("border", scaleBorderWidth()...)},
			"border-w-x": {m("border-x", scaleBorderWidth()...)},
			"border-w-y": {m("border-y", scaleBorderWidth()...)},
			"border-w-s": {m("border-s", scaleBorderWidth()...)},
			"border-w-e": {m("border-e", scaleBorderWidth()...)},
			"border-w-bs": {m("border-bs", scaleBorderWidth()...)},
			"border-w-be": {m("border-be", scaleBorderWidth()...)},
			"border-w-t": {m("border-t", scaleBorderWidth()...)},
			"border-w-r": {m("border-r", scaleBorderWidth()...)},
			"border-w-b": {m("border-b", scaleBorderWidth()...)},
			"border-w-l": {m("border-l", scaleBorderWidth()...)},
			"divide-x":   {m("divide-x", scaleBorderWidth()...)},
			"divide-x-reverse": {"divide-x-reverse"},
			"divide-y":         {m("divide-y", scaleBorderWidth()...)},
			"divide-y-reverse": {"divide-y-reverse"},
			"border-style":     {m("border", d("solid", "dashed", "dotted", "double", "hidden", "none")...)},
			"divide-style":     {m("divide", d("solid", "dashed", "dotted", "double", "hidden", "none")...)},
			"border-color":     {m("border", scaleColor()...)},
			"border-color-x":   {m("border-x", scaleColor()...)},
			"border-color-y":   {m("border-y", scaleColor()...)},
			"border-color-s":   {m("border-s", scaleColor()...)},
			"border-color-e":   {m("border-e", scaleColor()...)},
			"border-color-bs":  {m("border-bs", scaleColor()...)},
			"border-color-be":  {m("border-be", scaleColor()...)},
			"border-color-t":   {m("border-t", scaleColor()...)},
			"border-color-r":   {m("border-r", scaleColor()...)},
			"border-color-b":   {m("border-b", scaleColor()...)},
			"border-color-l":   {m("border-l", scaleColor()...)},
			"divide-color":     {m("divide", scaleColor()...)},
			"outline-style":    {m("outline", d("solid", "dashed", "dotted", "double", "none", "hidden")...)},
			"outline-offset":   {m("outline-offset", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"outline-w":        {m("outline", d("", IsNumber, IsArbitraryVariableLength, IsArbitraryLength)...)},
			"outline-color":    {m("outline", scaleColor()...)},

			// Effects
			"shadow":             {m("shadow", d("", "none", themeShadow, IsArbitraryVariableShadow, IsArbitraryShadow)...)},
			"shadow-color":       {m("shadow", scaleColor()...)},
			"inset-shadow":       {m("inset-shadow", d("none", themeInsetShadow, IsArbitraryVariableShadow, IsArbitraryShadow)...)},
			"inset-shadow-color": {m("inset-shadow", scaleColor()...)},
			"ring-w":             {m("ring", d("", IsNumber, IsArbitraryVariableLength, IsArbitraryLength)...)},
			"ring-w-inset":       {"ring-inset"},
			"ring-color":         {m("ring", scaleColor()...)},
			"ring-offset-w":      {m("ring-offset", IsNumber, IsArbitraryLength)},
			"ring-offset-color":  {m("ring-offset", scaleColor()...)},
			"inset-ring-w":       {m("inset-ring", scaleBorderWidth()...)},
			"inset-ring-color":   {m("inset-ring", scaleColor()...)},
			"text-shadow":        {m("text-shadow", d("none", themeTextShadow, IsArbitraryVariableShadow, IsArbitraryShadow)...)},
			"text-shadow-color":  {m("text-shadow", scaleColor()...)},
			"opacity":            {m("opacity", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"mix-blend":          {m("mix-blend", append(scaleBlendMode(), "plus-darker", "plus-lighter")...)},
			"bg-blend":           {m("bg-blend", scaleBlendMode()...)},

			// Masks
			"mask-clip":      {m("mask-clip", d("border", "padding", "content", "fill", "stroke", "view")...), "mask-no-clip"},
			"mask-composite": {m("mask", d("add", "subtract", "intersect", "exclude")...)},
			"mask-image-linear-pos":         {m("mask-linear", d(IsNumber)...)},
			"mask-image-linear-from-pos":    {m("mask-linear-from", scaleMaskImagePosition()...)},
			"mask-image-linear-to-pos":      {m("mask-linear-to", scaleMaskImagePosition()...)},
			"mask-image-linear-from-color":  {m("mask-linear-from", scaleColor()...)},
			"mask-image-linear-to-color":    {m("mask-linear-to", scaleColor()...)},
			"mask-image-t-from-pos":         {m("mask-t-from", scaleMaskImagePosition()...)},
			"mask-image-t-to-pos":           {m("mask-t-to", scaleMaskImagePosition()...)},
			"mask-image-t-from-color":       {m("mask-t-from", scaleColor()...)},
			"mask-image-t-to-color":         {m("mask-t-to", scaleColor()...)},
			"mask-image-r-from-pos":         {m("mask-r-from", scaleMaskImagePosition()...)},
			"mask-image-r-to-pos":           {m("mask-r-to", scaleMaskImagePosition()...)},
			"mask-image-r-from-color":       {m("mask-r-from", scaleColor()...)},
			"mask-image-r-to-color":         {m("mask-r-to", scaleColor()...)},
			"mask-image-b-from-pos":         {m("mask-b-from", scaleMaskImagePosition()...)},
			"mask-image-b-to-pos":           {m("mask-b-to", scaleMaskImagePosition()...)},
			"mask-image-b-from-color":       {m("mask-b-from", scaleColor()...)},
			"mask-image-b-to-color":         {m("mask-b-to", scaleColor()...)},
			"mask-image-l-from-pos":         {m("mask-l-from", scaleMaskImagePosition()...)},
			"mask-image-l-to-pos":           {m("mask-l-to", scaleMaskImagePosition()...)},
			"mask-image-l-from-color":       {m("mask-l-from", scaleColor()...)},
			"mask-image-l-to-color":         {m("mask-l-to", scaleColor()...)},
			"mask-image-x-from-pos":         {m("mask-x-from", scaleMaskImagePosition()...)},
			"mask-image-x-to-pos":           {m("mask-x-to", scaleMaskImagePosition()...)},
			"mask-image-x-from-color":       {m("mask-x-from", scaleColor()...)},
			"mask-image-x-to-color":         {m("mask-x-to", scaleColor()...)},
			"mask-image-y-from-pos":         {m("mask-y-from", scaleMaskImagePosition()...)},
			"mask-image-y-to-pos":           {m("mask-y-to", scaleMaskImagePosition()...)},
			"mask-image-y-from-color":       {m("mask-y-from", scaleColor()...)},
			"mask-image-y-to-color":         {m("mask-y-to", scaleColor()...)},
			"mask-image-radial":             {m("mask-radial", d(IsArbitraryVariable, IsArbitraryValue)...)},
			"mask-image-radial-from-pos":    {m("mask-radial-from", scaleMaskImagePosition()...)},
			"mask-image-radial-to-pos":      {m("mask-radial-to", scaleMaskImagePosition()...)},
			"mask-image-radial-from-color":  {m("mask-radial-from", scaleColor()...)},
			"mask-image-radial-to-color":    {m("mask-radial-to", scaleColor()...)},
			"mask-image-radial-shape":       {m("mask-radial", d("circle", "ellipse")...)},
			"mask-image-radial-size": {m("mask-radial", map[string][]ClassDefinition{
				"closest":  {"side", "corner"},
				"farthest": {"side", "corner"},
			})},
			"mask-image-radial-pos":         {m("mask-radial-at", scalePosition()...)},
			"mask-image-conic-pos":          {m("mask-conic", d(IsNumber)...)},
			"mask-image-conic-from-pos":     {m("mask-conic-from", scaleMaskImagePosition()...)},
			"mask-image-conic-to-pos":       {m("mask-conic-to", scaleMaskImagePosition()...)},
			"mask-image-conic-from-color":   {m("mask-conic-from", scaleColor()...)},
			"mask-image-conic-to-color":     {m("mask-conic-to", scaleColor()...)},
			"mask-mode":                     {m("mask", d("alpha", "luminance", "match")...)},
			"mask-origin":                   {m("mask-origin", d("border", "padding", "content", "fill", "stroke", "view")...)},
			"mask-position":                 {m("mask", scaleBgPosition()...)},
			"mask-repeat":                   {m("mask", scaleBgRepeat()...)},
			"mask-size":                     {m("mask", scaleBgSize()...)},
			"mask-type":                     {m("mask-type", d("alpha", "luminance")...)},
			"mask-image":                    {m("mask", d("none", IsArbitraryVariable, IsArbitraryValue)...)},

			// Filters
			"filter":            {m("filter", d("", "none", IsArbitraryVariable, IsArbitraryValue)...)},
			"blur":              {m("blur", d("", "none", themeBlur, IsArbitraryVariable, IsArbitraryValue)...)},
			"brightness":        {m("brightness", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"contrast":          {m("contrast", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"drop-shadow":       {m("drop-shadow", d("", "none", themeDropShadow, IsArbitraryVariableShadow, IsArbitraryShadow)...)},
			"drop-shadow-color": {m("drop-shadow", scaleColor()...)},
			"grayscale":         {m("grayscale", d("", IsNumber, IsArbitraryVariable, IsArbitraryValue)...)},
			"hue-rotate":        {m("hue-rotate", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"invert":            {m("invert", d("", IsNumber, IsArbitraryVariable, IsArbitraryValue)...)},
			"saturate":          {m("saturate", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"sepia":             {m("sepia", d("", IsNumber, IsArbitraryVariable, IsArbitraryValue)...)},
			"backdrop-filter":     {m("backdrop-filter", d("", "none", IsArbitraryVariable, IsArbitraryValue)...)},
			"backdrop-blur":       {m("backdrop-blur", d("", "none", themeBlur, IsArbitraryVariable, IsArbitraryValue)...)},
			"backdrop-brightness": {m("backdrop-brightness", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"backdrop-contrast":   {m("backdrop-contrast", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"backdrop-grayscale":  {m("backdrop-grayscale", d("", IsNumber, IsArbitraryVariable, IsArbitraryValue)...)},
			"backdrop-hue-rotate": {m("backdrop-hue-rotate", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"backdrop-invert":     {m("backdrop-invert", d("", IsNumber, IsArbitraryVariable, IsArbitraryValue)...)},
			"backdrop-opacity":    {m("backdrop-opacity", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"backdrop-saturate":   {m("backdrop-saturate", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"backdrop-sepia":      {m("backdrop-sepia", d("", IsNumber, IsArbitraryVariable, IsArbitraryValue)...)},

			// Tables
			"border-collapse":  {m("border", d("collapse", "separate")...)},
			"border-spacing":   {m("border-spacing", scaleUnambiguousSpacing()...)},
			"border-spacing-x": {m("border-spacing-x", scaleUnambiguousSpacing()...)},
			"border-spacing-y": {m("border-spacing-y", scaleUnambiguousSpacing()...)},
			"table-layout":     {m("table", d("auto", "fixed")...)},
			"caption":          {m("caption", d("top", "bottom")...)},

			// Transitions and Animation
			"transition":          {m("transition", d("", "all", "colors", "opacity", "shadow", "transform", "none", IsArbitraryVariable, IsArbitraryValue)...)},
			"transition-behavior": {m("transition", d("normal", "discrete")...)},
			"duration":            {m("duration", IsNumber, "initial", IsArbitraryVariable, IsArbitraryValue)},
			"ease":                {m("ease", d("linear", "initial", themeEase, IsArbitraryVariable, IsArbitraryValue)...)},
			"delay":               {m("delay", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"animate":             {m("animate", d("none", themeAnimate, IsArbitraryVariable, IsArbitraryValue)...)},

			// Transforms
			"backface":          {m("backface", d("hidden", "visible")...)},
			"perspective":       {m("perspective", d(themePerspective, IsArbitraryVariable, IsArbitraryValue)...)},
			"perspective-origin": {m("perspective-origin", scalePositionWithArbitrary()...)},
			"rotate":            {m("rotate", scaleRotate()...)},
			"rotate-x":          {m("rotate-x", scaleRotate()...)},
			"rotate-y":          {m("rotate-y", scaleRotate()...)},
			"rotate-z":          {m("rotate-z", scaleRotate()...)},
			"scale":             {m("scale", scaleScale()...)},
			"scale-x":           {m("scale-x", scaleScale()...)},
			"scale-y":           {m("scale-y", scaleScale()...)},
			"scale-z":           {m("scale-z", scaleScale()...)},
			"scale-3d":          {"scale-3d"},
			"skew":              {m("skew", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"skew-x":            {m("skew-x", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"skew-y":            {m("skew-y", IsNumber, IsArbitraryVariable, IsArbitraryValue)},
			"transform":         {m("transform", d(IsArbitraryVariable, IsArbitraryValue, "", "none", "gpu", "cpu")...)},
			"transform-origin":  {m("origin", scalePositionWithArbitrary()...)},
			"transform-style":   {m("transform", d("3d", "flat")...)},
			"translate":         {m("translate", IsFraction, "full", IsArbitraryVariable, IsArbitraryValue, themeSpacing)},
			"translate-x":       {m("translate-x", scaleTranslate()...)},
			"translate-y":       {m("translate-y", scaleTranslate()...)},
			"translate-z":       {m("translate-z", scaleTranslate()...)},
			"translate-none":    {"translate-none"},

			// Interactivity
			"accent":       {m("accent", scaleColor()...)},
			"appearance":   {m("appearance", d("none", "auto")...)},
			"caret-color":  {m("caret", scaleColor()...)},
			"color-scheme": {m("scheme", d("normal", "dark", "light", "light-dark", "only-dark", "only-light")...)},
			"cursor":       {m("cursor", d("auto", "default", "pointer", "wait", "text", "move", "help", "not-allowed", "none", "context-menu", "progress", "cell", "crosshair", "vertical-text", "alias", "copy", "no-drop", "grab", "grabbing", "all-scroll", "col-resize", "row-resize", "n-resize", "e-resize", "s-resize", "w-resize", "ne-resize", "nw-resize", "se-resize", "sw-resize", "ew-resize", "ns-resize", "nesw-resize", "nwse-resize", "zoom-in", "zoom-out", IsArbitraryVariable, IsArbitraryValue)...)},
			"field-sizing":   {m("field-sizing", d("fixed", "content")...)},
			"pointer-events": {m("pointer-events", d("auto", "none")...)},
			"resize":         {m("resize", d("none", "", "y", "x")...)},
			"scroll-behavior": {m("scroll", d("auto", "smooth")...)},
			"scroll-m":       {m("scroll-m", scaleUnambiguousSpacing()...)},
			"scroll-mx":      {m("scroll-mx", scaleUnambiguousSpacing()...)},
			"scroll-my":      {m("scroll-my", scaleUnambiguousSpacing()...)},
			"scroll-ms":      {m("scroll-ms", scaleUnambiguousSpacing()...)},
			"scroll-me":      {m("scroll-me", scaleUnambiguousSpacing()...)},
			"scroll-mbs":     {m("scroll-mbs", scaleUnambiguousSpacing()...)},
			"scroll-mbe":     {m("scroll-mbe", scaleUnambiguousSpacing()...)},
			"scroll-mt":      {m("scroll-mt", scaleUnambiguousSpacing()...)},
			"scroll-mr":      {m("scroll-mr", scaleUnambiguousSpacing()...)},
			"scroll-mb":      {m("scroll-mb", scaleUnambiguousSpacing()...)},
			"scroll-ml":      {m("scroll-ml", scaleUnambiguousSpacing()...)},
			"scroll-p":       {m("scroll-p", scaleUnambiguousSpacing()...)},
			"scroll-px":      {m("scroll-px", scaleUnambiguousSpacing()...)},
			"scroll-py":      {m("scroll-py", scaleUnambiguousSpacing()...)},
			"scroll-ps":      {m("scroll-ps", scaleUnambiguousSpacing()...)},
			"scroll-pe":      {m("scroll-pe", scaleUnambiguousSpacing()...)},
			"scroll-pbs":     {m("scroll-pbs", scaleUnambiguousSpacing()...)},
			"scroll-pbe":     {m("scroll-pbe", scaleUnambiguousSpacing()...)},
			"scroll-pt":      {m("scroll-pt", scaleUnambiguousSpacing()...)},
			"scroll-pr":      {m("scroll-pr", scaleUnambiguousSpacing()...)},
			"scroll-pb":      {m("scroll-pb", scaleUnambiguousSpacing()...)},
			"scroll-pl":      {m("scroll-pl", scaleUnambiguousSpacing()...)},
			"snap-align":     {m("snap", d("start", "end", "center", "align-none")...)},
			"snap-stop":      {m("snap", d("normal", "always")...)},
			"snap-type":      {m("snap", d("none", "x", "y", "both")...)},
			"snap-strictness": {m("snap", d("mandatory", "proximity")...)},
			"touch":          {m("touch", d("auto", "none", "manipulation")...)},
			"touch-x":        {m("touch-pan", d("x", "left", "right")...)},
			"touch-y":        {m("touch-pan", d("y", "up", "down")...)},
			"touch-pz":       {"touch-pinch-zoom"},
			"select":         {m("select", d("none", "text", "all", "auto")...)},
			"will-change":    {m("will-change", d("auto", "scroll", "contents", "transform", IsArbitraryVariable, IsArbitraryValue)...)},

			// SVG
			"fill":     {m("fill", d("none", themeColor, IsArbitraryVariable, IsArbitraryValue)...)},
			"stroke-w": {m("stroke", IsNumber, IsArbitraryVariableLength, IsArbitraryLength, IsArbitraryNumber)},
			"stroke":   {m("stroke", d("none", themeColor, IsArbitraryVariable, IsArbitraryValue)...)},

			// Accessibility
			"forced-color-adjust": {m("forced-color-adjust", d("auto", "none")...)},
		},

		ConflictingClassGroups: map[string][]string{
			"overflow":   {"overflow-x", "overflow-y"},
			"overscroll": {"overscroll-x", "overscroll-y"},
			"inset":      {"inset-x", "inset-y", "inset-bs", "inset-be", "start", "end", "top", "right", "bottom", "left"},
			"inset-x":    {"right", "left"},
			"inset-y":    {"top", "bottom"},
			"flex":       {"basis", "grow", "shrink"},
			"gap":        {"gap-x", "gap-y"},
			"p":          {"px", "py", "ps", "pe", "pbs", "pbe", "pt", "pr", "pb", "pl"},
			"px":         {"pr", "pl"},
			"py":         {"pt", "pb"},
			"m":          {"mx", "my", "ms", "me", "mbs", "mbe", "mt", "mr", "mb", "ml"},
			"mx":         {"mr", "ml"},
			"my":         {"mt", "mb"},
			"size":       {"w", "h"},
			"font-size":  {"leading"},
			"fvn-normal":       {"fvn-ordinal", "fvn-slashed-zero", "fvn-figure", "fvn-spacing", "fvn-fraction"},
			"fvn-ordinal":      {"fvn-normal"},
			"fvn-slashed-zero": {"fvn-normal"},
			"fvn-figure":       {"fvn-normal"},
			"fvn-spacing":      {"fvn-normal"},
			"fvn-fraction":     {"fvn-normal"},
			"line-clamp":       {"display", "overflow"},
			"rounded": {
				"rounded-s", "rounded-e", "rounded-t", "rounded-r", "rounded-b", "rounded-l",
				"rounded-ss", "rounded-se", "rounded-ee", "rounded-es",
				"rounded-tl", "rounded-tr", "rounded-br", "rounded-bl",
			},
			"rounded-s": {"rounded-ss", "rounded-es"},
			"rounded-e": {"rounded-se", "rounded-ee"},
			"rounded-t": {"rounded-tl", "rounded-tr"},
			"rounded-r": {"rounded-tr", "rounded-br"},
			"rounded-b": {"rounded-br", "rounded-bl"},
			"rounded-l": {"rounded-tl", "rounded-bl"},
			"border-spacing": {"border-spacing-x", "border-spacing-y"},
			"border-w": {
				"border-w-x", "border-w-y", "border-w-s", "border-w-e",
				"border-w-bs", "border-w-be", "border-w-t", "border-w-r", "border-w-b", "border-w-l",
			},
			"border-w-x": {"border-w-r", "border-w-l"},
			"border-w-y": {"border-w-t", "border-w-b"},
			"border-color": {
				"border-color-x", "border-color-y", "border-color-s", "border-color-e",
				"border-color-bs", "border-color-be", "border-color-t", "border-color-r",
				"border-color-b", "border-color-l",
			},
			"border-color-x":  {"border-color-r", "border-color-l"},
			"border-color-y":  {"border-color-t", "border-color-b"},
			"translate":       {"translate-x", "translate-y", "translate-none"},
			"translate-none":  {"translate", "translate-x", "translate-y", "translate-z"},
			"scroll-m": {
				"scroll-mx", "scroll-my", "scroll-ms", "scroll-me",
				"scroll-mbs", "scroll-mbe", "scroll-mt", "scroll-mr", "scroll-mb", "scroll-ml",
			},
			"scroll-mx": {"scroll-mr", "scroll-ml"},
			"scroll-my": {"scroll-mt", "scroll-mb"},
			"scroll-p": {
				"scroll-px", "scroll-py", "scroll-ps", "scroll-pe",
				"scroll-pbs", "scroll-pbe", "scroll-pt", "scroll-pr", "scroll-pb", "scroll-pl",
			},
			"scroll-px": {"scroll-pr", "scroll-pl"},
			"scroll-py": {"scroll-pt", "scroll-pb"},
			"touch":     {"touch-x", "touch-y", "touch-pz"},
			"touch-x":   {"touch"},
			"touch-y":   {"touch"},
			"touch-pz":  {"touch"},
		},

		ConflictingClassGroupModifiers: map[string][]string{
			"font-size": {"leading"},
		},

		OrderSensitiveModifiers: []string{
			"*", "**", "after", "backdrop", "before", "details-content",
			"file", "first-letter", "first-line", "marker", "placeholder", "selection",
		},
	}
}

// d is a helper to convert variadic ClassDefinition values to a slice.
func d(defs ...ClassDefinition) []ClassDefinition {
	return defs
}

// m is a helper that creates a map[string][]ClassDefinition with a single key.
// This is the most common pattern in class group definitions.
func m(key string, defs ...ClassDefinition) map[string][]ClassDefinition {
	return map[string][]ClassDefinition{key: defs}
}
