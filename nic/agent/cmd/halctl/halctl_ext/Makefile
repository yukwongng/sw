default: halctl-ext

halctl-ext:
	go install
	@# Generates files for:
	@#    - Bash completion
	@#    - Documentation:
	@#      - Markdown files
	@#      - man pages. Do this for man pages to work: cp *.3 /usr/local/share/man/man3
	mkdir -p bash_completion
	mkdir -p markdown
	mkdir -p man_pages
	halctl_ext
	@# For bash completion
	#cp bash_completion/halctl_bash_complete.sh /etc/bash_completion.d/halctl_bash_complete.sh
	# Do this for halctl bash completion: source /etc/bash_completion.d/halctl_bash_complete.sh
	#source /etc/bash_completion.d/halctl_bash_complete.sh
