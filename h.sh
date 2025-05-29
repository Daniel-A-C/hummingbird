selection=$(hummingbird)

if [[ "$selection" == "x" ]]; then
    # Do nothing if 'x' is returned
    true
elif [[ -d "$selection" ]]; then
    # If it's a directory, change to it
    cd "$selection"
elif [[ -f "$selection" ]]; then
    # If it's a file, check if it's executable
    if [[ -x "$selection" ]]; then
        # If the file is executable, do nothing (don't open it with nvim)
        # You could optionally echo a message here, e.g.:
        # echo "Selected file '$selection' is executable, not opening."
        true
    else
        # If the file is not executable, open it with neovim
        source nv "$selection"
    fi
else
    # Handle cases where selection is not "x", not a directory, and not a regular file
    # This could happen if hummingbird returns an error string or an empty string.
    # For example, if hummingbird returned "error_getting_wd_at_exit"
    if [[ -n "$selection" ]]; then # Check if selection is not empty
        echo "Hummingbird returned: '$selection'. Not a recognized file or directory."
    fi
    # Script will then proceed to clear and ls
fi

clear
# ls -l will show the contents of the current directory.
# If cd occurred, it's the new directory. Otherwise, it's the original directory.
ls -l
