package glap

// Arg TODO
type Arg struct {
	short    string
	long     string
	aliases  []string
	help     string
	longHelp string
	required bool
}

// Short TODO
func (a *Arg) Short(short string) *Arg {
	a.short = short
	return a
}

// Long TODO
func (a *Arg) Long(long string) *Arg {
	a.long = long
	return a
}

// Alias TODO
func (a *Arg) Alias(alias string) *Arg {
	a.aliases = append(a.aliases, alias)
	return a
}

// Aliases TODO
func (a *Arg) Aliases(aliases []string) *Arg {
	a.aliases = append(a.aliases, aliases...)
	return a
}

// Help TODO
func (a *Arg) Help(help string) *Arg {
	a.help = help
	return a
}

// LongHelp TODO
func (a *Arg) LongHelp(longHelp string) *Arg {
	a.longHelp = longHelp
	return a
}

// Required TODO
func (a *Arg) Required(required bool) *Arg {
	a.required = required
	return a
}
