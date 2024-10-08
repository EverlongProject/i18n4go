package common

type Options struct {
	CommandFlag string

	HelpFlag     bool
	LongHelpFlag bool

	VerboseFlag bool
	DryRunFlag  bool
	PoFlag      bool
	MetaFlag    bool

	SourceLanguageFlag        string
	LanguagesFlag             string
	GoogleTranslateApiKeyFlag string

	OutputDirFlag          string
	OutputMatchImportFlag  bool
	OutputMatchPackageFlag bool
	OutputFlatFlag         bool
	OutputFormatFlatFlag   bool

	ExcludedFilenameFlag  string
	SubstringFilenameFlag string
	FilenameFlag          string
	DirnameFlag           string

	RecurseFlag bool

	IgnoreRegexpFlag string

	LanguageFilesFlag string

	I18nStringsFilenameFlag string
	I18nStringsDirnameFlag  string

	RootPathFlag string

	InitCodeSnippetFilenameFlag string

	QualifierFlag string
}

type I18nStringInfo struct {
	ID          string `json:"id"`
	Translation string `json:"translation"`
	Modified    bool   `json:"modified"`
}

type StringInfo struct {
	Filename string `json:"filename"`
	Value    string `json:"value"`
	Offset   int    `json:"offset"`
	Line     int    `json:"line"`
	Column   int    `json:"column"`

	// optional, empty means "all locales"
	Locales []string `json:"locales,omitempty"`
}

type ExcludedStrings struct {
	ExcludedStrings     []string `json:"excludedStrings"`
	ExcludedLines       []string `json:"excludedLines"`
	ExcludedRegexps     []string `json:"excludedRegexps"`
	ExcludedFileRegexps []string `json:"excludedFileRegexps"`
	EnforcedFuncs       []string `json:"enforcedFuncs"`
}

type PrinterInterface interface {
	Println(a ...interface{}) (int, error)
	Printf(msg string, a ...interface{}) (int, error)
}

var BLANKS = []string{", ", "\t", "\n", "\n\t", "\t\n"}
