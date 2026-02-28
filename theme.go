package cliout

// Theme defines the colors used for each output level.
type Theme struct {
	Name         string
	PrefixColor  Color
	InfoColor    Color
	DebugColor   Color
	TraceColor   Color
	WarnColor    Color
	ErrorColor   Color
	SuccessColor Color
}

// ThemeDefault is a simple theme using standard ANSI colors.
var ThemeDefault = Theme{
	Name:         "Default",
	PrefixColor:  ColorCyan,
	InfoColor:    ColorDefault,
	DebugColor:   ColorBrightBlack,
	TraceColor:   ColorBrightBlack,
	WarnColor:    ColorYellow,
	ErrorColor:   ColorRed,
	SuccessColor: ColorGreen,
}

// --- Ayu family ---

var ThemeAyu = Theme{
	Name:         "Ayu",
	PrefixColor:  Hex("FF8F40"), // orange accent
	InfoColor:    Hex("E6E1CF"), // foreground
	DebugColor:   Hex("5C6773"), // comment
	TraceColor:   Hex("5C6773"),
	WarnColor:    Hex("FFB454"), // yellow
	ErrorColor:   Hex("FF3333"), // red
	SuccessColor: Hex("B8CC52"), // green
}

var ThemeAyuLight = Theme{
	Name:         "Ayu Light",
	PrefixColor:  Hex("FF6A00"), // orange accent
	InfoColor:    Hex("575F66"), // foreground
	DebugColor:   Hex("ABB0B6"), // comment
	TraceColor:   Hex("ABB0B6"),
	WarnColor:    Hex("F2AE49"), // yellow
	ErrorColor:   Hex("F51818"), // red
	SuccessColor: Hex("86B300"), // green
}

var ThemeAyuMirage = Theme{
	Name:         "Ayu Mirage",
	PrefixColor:  Hex("FFAD66"), // orange accent
	InfoColor:    Hex("CCCAC2"), // foreground
	DebugColor:   Hex("B8CFE6"), // comment
	TraceColor:   Hex("B8CFE6"),
	WarnColor:    Hex("FFD173"), // yellow
	ErrorColor:   Hex("F28779"), // red
	SuccessColor: Hex("D5FF80"), // green
}

// --- Classic themes ---

var ThemeDracula = Theme{
	Name:         "Dracula",
	PrefixColor:  Hex("BD93F9"), // purple
	InfoColor:    Hex("F8F8F2"), // foreground
	DebugColor:   Hex("6272A4"), // comment
	TraceColor:   Hex("6272A4"),
	WarnColor:    Hex("F1FA8C"), // yellow
	ErrorColor:   Hex("FF5555"), // red
	SuccessColor: Hex("50FA7B"), // green
}

var ThemeOneDark = Theme{
	Name:         "One Dark",
	PrefixColor:  Hex("61AFEF"), // blue
	InfoColor:    Hex("ABB2BF"), // foreground
	DebugColor:   Hex("5C6370"), // comment
	TraceColor:   Hex("5C6370"),
	WarnColor:    Hex("E5C07B"), // yellow
	ErrorColor:   Hex("E06C75"), // red
	SuccessColor: Hex("98C379"), // green
}

// --- Solarized family ---

var ThemeSolarizedDark = Theme{
	Name:         "Solarized Dark",
	PrefixColor:  Hex("268BD2"), // blue
	InfoColor:    Hex("839496"), // base0
	DebugColor:   Hex("586E75"), // base01
	TraceColor:   Hex("586E75"),
	WarnColor:    Hex("B58900"), // yellow
	ErrorColor:   Hex("DC322F"), // red
	SuccessColor: Hex("859900"), // green
}

var ThemeSolarizedLight = Theme{
	Name:         "Solarized Light",
	PrefixColor:  Hex("268BD2"), // blue
	InfoColor:    Hex("657B83"), // base00
	DebugColor:   Hex("93A1A1"), // base1
	TraceColor:   Hex("93A1A1"),
	WarnColor:    Hex("B58900"), // yellow
	ErrorColor:   Hex("DC322F"), // red
	SuccessColor: Hex("859900"), // green
}

// --- Nord ---

var ThemeNord = Theme{
	Name:         "Nord",
	PrefixColor:  Hex("88C0D0"), // nord8 (frost)
	InfoColor:    Hex("D8DEE9"), // nord4 (snow storm)
	DebugColor:   Hex("616E88"), // comment
	TraceColor:   Hex("616E88"),
	WarnColor:    Hex("EBCB8B"), // nord13 (aurora yellow)
	ErrorColor:   Hex("BF616A"), // nord11 (aurora red)
	SuccessColor: Hex("A3BE8C"), // nord14 (aurora green)
}

// --- Gruvbox family ---

var ThemeGruvboxDark = Theme{
	Name:         "Gruvbox Dark",
	PrefixColor:  Hex("83A598"), // aqua
	InfoColor:    Hex("EBDBB2"), // fg
	DebugColor:   Hex("928374"), // gray
	TraceColor:   Hex("928374"),
	WarnColor:    Hex("FABD2F"), // yellow
	ErrorColor:   Hex("FB4934"), // red
	SuccessColor: Hex("B8BB26"), // green
}

var ThemeGruvboxLight = Theme{
	Name:         "Gruvbox Light",
	PrefixColor:  Hex("427B58"), // aqua
	InfoColor:    Hex("3C3836"), // fg
	DebugColor:   Hex("928374"), // gray
	TraceColor:   Hex("928374"),
	WarnColor:    Hex("B57614"), // yellow
	ErrorColor:   Hex("9D0006"), // red
	SuccessColor: Hex("79740E"), // green
}

// --- Monokai ---

var ThemeMonokai = Theme{
	Name:         "Monokai",
	PrefixColor:  Hex("66D9EF"), // cyan
	InfoColor:    Hex("F8F8F2"), // foreground
	DebugColor:   Hex("75715E"), // comment
	TraceColor:   Hex("75715E"),
	WarnColor:    Hex("E6DB74"), // yellow
	ErrorColor:   Hex("F92672"), // red/magenta
	SuccessColor: Hex("A6E22E"), // green
}

// --- Monokai Pro family ---

var ThemeMonokaiPro = Theme{
	Name:         "Monokai Pro",
	PrefixColor:  Hex("78DCE8"), // accent5 (cyan)
	InfoColor:    Hex("FCFCFA"), // text
	DebugColor:   Hex("727072"), // dimmed3 (comment)
	TraceColor:   Hex("5B595C"), // dimmed4
	WarnColor:    Hex("FFD866"), // accent3 (yellow)
	ErrorColor:   Hex("FF6188"), // accent1 (red/pink)
	SuccessColor: Hex("A9DC76"), // accent4 (green)
}

var ThemeMonokaiProClassic = Theme{
	Name:         "Monokai Pro Classic",
	PrefixColor:  Hex("66D9EF"), // accent5 (cyan)
	InfoColor:    Hex("FDFFF1"), // text
	DebugColor:   Hex("6E7066"), // dimmed3 (comment)
	TraceColor:   Hex("57584F"), // dimmed4
	WarnColor:    Hex("E6DB74"), // accent3 (yellow)
	ErrorColor:   Hex("F92672"), // accent1 (red/pink)
	SuccessColor: Hex("A6E22E"), // accent4 (green)
}

var ThemeMonokaiProMachine = Theme{
	Name:         "Monokai Pro Machine",
	PrefixColor:  Hex("7CD5F1"), // accent5 (cyan)
	InfoColor:    Hex("F2FFFC"), // text
	DebugColor:   Hex("6B7678"), // dimmed3 (comment)
	TraceColor:   Hex("545F62"), // dimmed4
	WarnColor:    Hex("FFED72"), // accent3 (yellow)
	ErrorColor:   Hex("FF6D7E"), // accent1 (red/pink)
	SuccessColor: Hex("A2E57B"), // accent4 (green)
}

var ThemeMonokaiProOctagon = Theme{
	Name:         "Monokai Pro Octagon",
	PrefixColor:  Hex("9CD1BB"), // accent5 (cyan)
	InfoColor:    Hex("EAF2F1"), // text
	DebugColor:   Hex("696D77"), // dimmed3 (comment)
	TraceColor:   Hex("535763"), // dimmed4
	WarnColor:    Hex("FFD76D"), // accent3 (yellow)
	ErrorColor:   Hex("FF657A"), // accent1 (red/pink)
	SuccessColor: Hex("BAD761"), // accent4 (green)
}

var ThemeMonokaiProRistretto = Theme{
	Name:         "Monokai Pro Ristretto",
	PrefixColor:  Hex("85DACC"), // accent5 (cyan)
	InfoColor:    Hex("FFF1F3"), // text
	DebugColor:   Hex("72696A"), // dimmed3 (comment)
	TraceColor:   Hex("5B5353"), // dimmed4
	WarnColor:    Hex("F9CC6C"), // accent3 (yellow)
	ErrorColor:   Hex("FD6883"), // accent1 (red/pink)
	SuccessColor: Hex("ADDA78"), // accent4 (green)
}

var ThemeMonokaiProSpectrum = Theme{
	Name:         "Monokai Pro Spectrum",
	PrefixColor:  Hex("5AD4E6"), // accent5 (cyan)
	InfoColor:    Hex("F7F1FF"), // text
	DebugColor:   Hex("69676C"), // dimmed3 (comment)
	TraceColor:   Hex("525053"), // dimmed4
	WarnColor:    Hex("FCE566"), // accent3 (yellow)
	ErrorColor:   Hex("FC618D"), // accent1 (red/pink)
	SuccessColor: Hex("7BD88F"), // accent4 (green)
}

var ThemeMonokaiProLight = Theme{
	Name:         "Monokai Pro Light",
	PrefixColor:  Hex("1C8CA8"), // accent5 (cyan)
	InfoColor:    Hex("29242A"), // text
	DebugColor:   Hex("A59FA0"), // dimmed3 (comment)
	TraceColor:   Hex("BFB9BA"), // dimmed4
	WarnColor:    Hex("CC7A0A"), // accent3 (yellow)
	ErrorColor:   Hex("E14775"), // accent1 (red/pink)
	SuccessColor: Hex("269D69"), // accent4 (green)
}

// --- Material family ---

var ThemeMaterialDark = Theme{
	Name:         "Material Dark",
	PrefixColor:  Hex("82AAFF"), // blue
	InfoColor:    Hex("EEFFFF"), // foreground
	DebugColor:   Hex("546E7A"), // comment
	TraceColor:   Hex("546E7A"),
	WarnColor:    Hex("FFCB6B"), // yellow
	ErrorColor:   Hex("FF5370"), // red
	SuccessColor: Hex("C3E88D"), // green
}

var ThemeMaterialLight = Theme{
	Name:         "Material Light",
	PrefixColor:  Hex("6182B8"), // blue
	InfoColor:    Hex("90A4AE"), // foreground
	DebugColor:   Hex("90A4AE"), // comment
	TraceColor:   Hex("CCD7DA"),
	WarnColor:    Hex("FFB62C"), // yellow
	ErrorColor:   Hex("E53935"), // red
	SuccessColor: Hex("91B859"), // green
}

// --- Palenight ---

var ThemePalenight = Theme{
	Name:         "Palenight",
	PrefixColor:  Hex("82AAFF"), // blue
	InfoColor:    Hex("A6ACCD"), // foreground
	DebugColor:   Hex("676E95"), // comment
	TraceColor:   Hex("676E95"),
	WarnColor:    Hex("FFCB6B"), // yellow
	ErrorColor:   Hex("FF5370"), // red
	SuccessColor: Hex("C3E88D"), // green
}

// --- Catppuccino family ---

var ThemeCatppuccinoFrappe = Theme{
	Name:         "Catppuccino Frappe",
	PrefixColor:  Hex("8CAAEE"), // blue
	InfoColor:    Hex("C6D0F5"), // text
	DebugColor:   Hex("737994"), // overlay0
	TraceColor:   Hex("626880"), // surface2
	WarnColor:    Hex("E5C890"), // yellow
	ErrorColor:   Hex("E78284"), // red
	SuccessColor: Hex("A6D189"), // green
}

var ThemeCatppuccinoLatte = Theme{
	Name:         "Catppuccino Latte",
	PrefixColor:  Hex("1E66F5"), // blue
	InfoColor:    Hex("4C4F69"), // text
	DebugColor:   Hex("9CA0B0"), // overlay0
	TraceColor:   Hex("ACB0BE"), // surface2
	WarnColor:    Hex("DF8E1D"), // yellow
	ErrorColor:   Hex("D20F39"), // red
	SuccessColor: Hex("40A02B"), // green
}

var ThemeCatppuccinoMacchiato = Theme{
	Name:         "Catppuccino Macchiato",
	PrefixColor:  Hex("8AADF4"), // blue
	InfoColor:    Hex("CAD3F5"), // text
	DebugColor:   Hex("6E738D"), // overlay0
	TraceColor:   Hex("5B6078"), // surface2
	WarnColor:    Hex("EED49F"), // yellow
	ErrorColor:   Hex("ED8796"), // red
	SuccessColor: Hex("A6DA95"), // green
}

var ThemeCatppuccinoMocha = Theme{
	Name:         "Catppuccino Mocha",
	PrefixColor:  Hex("89B4FA"), // blue
	InfoColor:    Hex("CDD6F4"), // text
	DebugColor:   Hex("6C7086"), // overlay0
	TraceColor:   Hex("585B70"), // surface2
	WarnColor:    Hex("F9E2AF"), // yellow
	ErrorColor:   Hex("F38BA8"), // red
	SuccessColor: Hex("A6E3A1"), // green
}

// --- Rose Pine family ---

var ThemeRosePine = Theme{
	Name:         "Rose Pine",
	PrefixColor:  Hex("C4A7E7"), // iris
	InfoColor:    Hex("E0DEF4"), // text
	DebugColor:   Hex("6E6A86"), // muted
	TraceColor:   Hex("6E6A86"),
	WarnColor:    Hex("F6C177"), // gold
	ErrorColor:   Hex("EB6F92"), // love
	SuccessColor: Hex("31748F"), // pine
}

var ThemeRosePineDawn = Theme{
	Name:         "Rose Pine Dawn",
	PrefixColor:  Hex("907AA9"), // iris
	InfoColor:    Hex("575279"), // text
	DebugColor:   Hex("9893A5"), // muted
	TraceColor:   Hex("9893A5"),
	WarnColor:    Hex("EA9D34"), // gold
	ErrorColor:   Hex("B4637A"), // love
	SuccessColor: Hex("286983"), // pine
}

var ThemeRosePineMoon = Theme{
	Name:         "Rose Pine Moon",
	PrefixColor:  Hex("C4A7E7"), // iris
	InfoColor:    Hex("E0DEF4"), // text
	DebugColor:   Hex("6E6A86"), // muted
	TraceColor:   Hex("6E6A86"),
	WarnColor:    Hex("F6C177"), // gold
	ErrorColor:   Hex("EB6F92"), // love
	SuccessColor: Hex("3E8FB0"), // pine
}

// --- Tokyo Night family ---

var ThemeTokyoNightStorm = Theme{
	Name:         "Tokyo Night Storm",
	PrefixColor:  Hex("7AA2F7"), // blue
	InfoColor:    Hex("C0CAF5"), // foreground
	DebugColor:   Hex("565F89"), // comment
	TraceColor:   Hex("565F89"),
	WarnColor:    Hex("E0AF68"), // yellow
	ErrorColor:   Hex("F7768E"), // red
	SuccessColor: Hex("9ECE6A"), // green
}

var ThemeTokyoNightDay = Theme{
	Name:         "Tokyo Night Day",
	PrefixColor:  Hex("2E7DE9"), // blue
	InfoColor:    Hex("3760BF"), // foreground
	DebugColor:   Hex("848CB5"), // comment
	TraceColor:   Hex("848CB5"),
	WarnColor:    Hex("8C6C3E"), // yellow
	ErrorColor:   Hex("F52A65"), // red
	SuccessColor: Hex("587539"), // green
}

var ThemeTokyoNightNight = Theme{
	Name:         "Tokyo Night Night",
	PrefixColor:  Hex("7AA2F7"), // blue
	InfoColor:    Hex("A9B1D6"), // foreground
	DebugColor:   Hex("565F89"), // comment
	TraceColor:   Hex("565F89"),
	WarnColor:    Hex("E0AF68"), // yellow
	ErrorColor:   Hex("F7768E"), // red
	SuccessColor: Hex("9ECE6A"), // green
}

// --- Theme listing and lookup ---

// Themes returns a slice containing all built-in themes. The returned slice
// is a fresh copy each time, so callers are free to modify it.
func Themes() []Theme {
	return []Theme{
		ThemeDefault,
		ThemeAyu, ThemeAyuLight, ThemeAyuMirage,
		ThemeDracula,
		ThemeOneDark,
		ThemeSolarizedDark, ThemeSolarizedLight,
		ThemeNord,
		ThemeGruvboxDark, ThemeGruvboxLight,
		ThemeMonokai,
		ThemeMonokaiPro, ThemeMonokaiProClassic, ThemeMonokaiProMachine,
		ThemeMonokaiProOctagon, ThemeMonokaiProRistretto,
		ThemeMonokaiProSpectrum, ThemeMonokaiProLight,
		ThemeMaterialDark, ThemeMaterialLight,
		ThemePalenight,
		ThemeCatppuccinoFrappe, ThemeCatppuccinoLatte,
		ThemeCatppuccinoMacchiato, ThemeCatppuccinoMocha,
		ThemeRosePine, ThemeRosePineDawn, ThemeRosePineMoon,
		ThemeTokyoNightStorm, ThemeTokyoNightDay, ThemeTokyoNightNight,
	}
}

// ThemeByName returns the built-in theme with the given name and true if found,
// or an empty Theme and false if no match exists. The comparison is
// case-insensitive.
func ThemeByName(name string) (Theme, bool) {
	lower := toLower(name)
	for _, t := range Themes() {
		if toLower(t.Name) == lower {
			return t, true
		}
	}
	return Theme{}, false
}

// toLower returns s with ASCII A-Z mapped to a-z. This avoids importing
// strings just for a case-insensitive comparison.
func toLower(s string) string {
	b := make([]byte, len(s))
	for i := range s {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return string(b)
}
