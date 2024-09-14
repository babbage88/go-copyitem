#compdef go-cp
compdef _go-cp go-cp

# Replace all occurrences of "go-cp" in this file with the actual name of your
# CLI go-cp. We recommend using Find+Replace feature of your editor. Let's say
# your CLI go-cp is called "acme", then replace like so:
# * go-cp => acme
# * _go-cp => _acme

_go-cp() {
	local -a opts
	local cur
	cur=${words[-1]}
	if [[ "$cur" == "-"* ]]; then
		opts=("${(@f)$(${words[@]:0:#words[@]-1} ${cur} --generate-shell-completion)}")
	else
		opts=("${(@f)$(${words[@]:0:#words[@]-1} --generate-shell-completion)}")
	fi

	if [[ "${opts[1]}" != "" ]]; then
		_describe 'values' opts
	else
		_files
	fi
}

# don't run the completion function when being source-ed or eval-ed
if [ "$funcstack[1]" = "_go-cp" ]; then
	_go-cp
fi