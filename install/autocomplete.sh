#compdef cpgo
compdef _cpgo cpgo

# Replace all occurrences of "cpgo" in this file with the actual name of your
# CLI cpgo. We recommend using Find+Replace feature of your editor. Let's say
# your CLI cpgo is called "acme", then replace like so:
# * cpgo => acme
# * _cpgo => _acme

_cpgo() {
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
if [ "$funcstack[1]" = "_cpgo" ]; then
	_cpgo
fi